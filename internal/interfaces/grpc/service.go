package grpcinterface

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"path/filepath"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/tdex-network/tdex-daemon/internal/core/ports"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	log "github.com/sirupsen/logrus"
	"github.com/tdex-network/tdex-daemon/internal/core/application"
	interfaces "github.com/tdex-network/tdex-daemon/internal/interfaces"
	grpchandler "github.com/tdex-network/tdex-daemon/internal/interfaces/grpc/handler"
	"github.com/tdex-network/tdex-daemon/internal/interfaces/grpc/interceptor"
	"github.com/tdex-network/tdex-daemon/internal/interfaces/grpc/permissions"
	"github.com/tdex-network/tdex-daemon/pkg/macaroons"
	"google.golang.org/grpc"
	"gopkg.in/macaroon-bakery.v2/bakery"

	daemonv2 "github.com/tdex-network/tdex-daemon/api-spec/protobuf/gen/tdex-daemon/v2"
	tdexv1 "github.com/tdex-network/tdex-daemon/api-spec/protobuf/gen/tdex/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

const (
	// OperatorTLSKeyFile is the name of the TLS key file for the Operator
	// interface.
	OperatorTLSKeyFile = "key.pem"
	// OperatorTLSCertFile is the name of the TLS certificate file for the
	// Operator interface.
	OperatorTLSCertFile = "cert.pem"
	// Location is used as the macaroon's location hint. This is not verified as
	// part of the macaroons itself. Check the doc for more info:
	// https://github.com/go-macaroon/macaroon#func-macaroon-location.
	Location = "tdexd"
	// DbFile is the name of the macaroon database file.
	DBFile = "macaroons.db"
	// AdminMacaroonFile is the name of the admin macaroon.
	AdminMacaroonFile = "admin.macaroon"
	// ReadOnlyMacaroonFile is the name of the read-only macaroon.
	ReadOnlyMacaroonFile = "readonly.macaroon"
	// MarketMacaroonFile is the name of the macaroon allowing to open, close and
	// update the strategy of a market.
	MarketMacaroonFile = "market.macaroon"
	// PriceMacaroonFile is the name of the macaroon allowing to update only the
	// prices of markets.
	PriceMacaroonFile = "price.macaroon"
	// WalletMacaroonFile is the name of the macaroon allowing to manage the
	// so called "Wallet" subaccount of the daemon's wallet.
	WalletMacaroonFile = "wallet.macaroon"
	// WebhookMacaroonFile is the name of the macaroon allowing to add, remove or
	// list webhooks.
	WebhookMacaroonFile = "webhook.macaroon"
)

var (
	serialNumberLimit = new(big.Int).Lsh(big.NewInt(1), 128)

	Macaroons = map[string][]bakery.Op{
		AdminMacaroonFile:    permissions.AdminPermissions(),
		ReadOnlyMacaroonFile: permissions.ReadOnlyPermissions(),
		MarketMacaroonFile:   permissions.MarketPermissions(),
		PriceMacaroonFile:    permissions.PricePermissions(),
		WebhookMacaroonFile:  permissions.WebhookPermissions(),
		WalletMacaroonFile:   permissions.WalletPermissions(),
	}
)

type service struct {
	opts        ServiceOpts
	macaroonSvc *macaroons.Service

	operatorServer *http.Server
	tradeServer    *http.Server
	password       string
}

type ServiceOpts struct {
	NoMacaroons bool

	Datadir                  string
	DBLocation               string
	TLSLocation              string
	MacaroonsLocation        string
	OperatorExtraIPs         []string
	OperatorExtraDomains     []string
	WalletUnlockPasswordFile string

	OperatorPort int
	TradePort    int
	TradeTLSKey  string
	TradeTLSCert string

	AppConfig *application.Config
	BuildData ports.BuildData

	NoOperatorTls bool
	ConnectAddr   string
	ConnectProto  string
}

func (o ServiceOpts) validate() error {
	if !pathExists(o.Datadir) {
		return fmt.Errorf("%s: datadir must be an existing directory", o.Datadir)
	}

	if !o.NoMacaroons {
		macDir := o.macaroonsDatadir()
		adminMacExists := pathExists(filepath.Join(macDir, AdminMacaroonFile))
		roMacExists := pathExists(filepath.Join(macDir, ReadOnlyMacaroonFile))
		marketMacExists := pathExists(filepath.Join(macDir, MarketMacaroonFile))
		priceMacExists := pathExists(filepath.Join(macDir, PriceMacaroonFile))

		if adminMacExists != roMacExists ||
			adminMacExists != marketMacExists ||
			adminMacExists != priceMacExists {
			return fmt.Errorf(
				"all macaroons must be either existing or not in path %s", macDir,
			)
		}
	}

	if !o.NoOperatorTls {
		tlsDir := o.tlsDatadir()
		tlsKeyExists := pathExists(filepath.Join(tlsDir, OperatorTLSKeyFile))
		tlsCertExists := pathExists(filepath.Join(tlsDir, OperatorTLSCertFile))
		if !tlsKeyExists && tlsCertExists {
			return fmt.Errorf(
				"found %s file but %s is missing. Please delete %s to make the "+
					"daemon recreating both files in path %s",
				OperatorTLSCertFile, OperatorTLSKeyFile, OperatorTLSCertFile, tlsDir,
			)
		}

		if len(o.OperatorExtraIPs) > 0 {
			for _, ip := range o.OperatorExtraIPs {
				if net.ParseIP(ip) == nil {
					return fmt.Errorf("invalid operator extra ip %s", ip)
				}
			}
		}
	}

	if ok := isValidPort(o.OperatorPort); !ok {
		return fmt.Errorf("operator port must be in range [%d, %d]", minPort, maxPort)
	}
	if ok := isValidPort(o.TradePort); !ok {
		return fmt.Errorf("trade port must be in range [%d, %d]", minPort, maxPort)
	}

	tradeTLSKeyExist := pathExists(o.TradeTLSKey)
	tradeTLSCertExist := pathExists(o.TradeTLSCert)
	if tradeTLSKeyExist != tradeTLSCertExist {
		return fmt.Errorf(
			"TLS key and certificate for Trade interface must be either existing " +
				"or undefined",
		)
	}

	if o.WalletUnlockPasswordFile != "" {
		if !pathExists(o.WalletUnlockPasswordFile) {
			return fmt.Errorf("wallet unlock password file not found")
		}
	}

	if o.AppConfig == nil {
		return fmt.Errorf("missing app config")
	}
	if err := o.AppConfig.Validate(); err != nil {
		return fmt.Errorf("invalid app config: %s", err)
	}

	if o.BuildData == nil {
		return fmt.Errorf("missing build data")
	}

	return nil
}

func (o ServiceOpts) dbDatadir() string {
	return filepath.Join(o.Datadir, o.DBLocation)
}

func (o ServiceOpts) macaroonsDatadir() string {
	return filepath.Join(o.Datadir, o.MacaroonsLocation)
}

func (o ServiceOpts) tlsDatadir() string {
	return filepath.Join(o.Datadir, o.TLSLocation)
}

func (o ServiceOpts) operatorTLSKey() string {
	if o.NoOperatorTls {
		return ""
	}
	return filepath.Join(o.tlsDatadir(), OperatorTLSKeyFile)
}

func (o ServiceOpts) operatorTLSCert() string {
	if o.NoOperatorTls {
		return ""
	}
	return filepath.Join(o.tlsDatadir(), OperatorTLSCertFile)
}

func (o ServiceOpts) operatorTLSConfig() (*tls.Config, error) {
	if o.NoOperatorTls {
		return nil, nil
	}
	return getTlsConfig(o.operatorTLSKey(), o.operatorTLSCert())
}

func (o ServiceOpts) tradeTLSConfig() (*tls.Config, error) {
	if o.TradeTLSCert == "" {
		return nil, nil
	}
	return getTlsConfig(o.TradeTLSKey, o.TradeTLSCert)
}

func (o ServiceOpts) operatorServerAddr() string {
	return fmt.Sprintf(":%d", o.OperatorPort)
}

func (o ServiceOpts) tradeServerAddr() string {
	return fmt.Sprintf(":%d", o.TradePort)
}

func (o ServiceOpts) tradeClientAddr() string {
	return fmt.Sprintf("localhost:%d", o.TradePort)
}

func NewService(opts ServiceOpts) (interfaces.Service, error) {
	if err := opts.validate(); err != nil {
		return nil, fmt.Errorf("invalid opts: %s", err)
	}

	var macaroonSvc *macaroons.Service
	if !opts.NoMacaroons {
		macaroonSvc, _ = macaroons.NewService(
			opts.dbDatadir(), Location, DBFile, false, macaroons.IPLockChecker,
		)
		if err := permissions.Validate(); err != nil {
			return nil, err
		}
	}

	if !opts.NoOperatorTls {
		if err := generateOperatorTLSKeyCert(
			opts.tlsDatadir(), opts.OperatorExtraIPs, opts.OperatorExtraDomains,
		); err != nil {
			return nil, err
		}
	}

	var password string
	if opts.WalletUnlockPasswordFile != "" {
		passwordBytes, err := ioutil.ReadFile(opts.WalletUnlockPasswordFile)
		if err != nil {
			return nil, err
		}

		trimmedPass := bytes.TrimFunc(passwordBytes, func(r rune) bool {
			return r == 10 || r == 32
		})

		password = string(trimmedPass)
	}

	return &service{
		opts:        opts,
		macaroonSvc: macaroonSvc,
		password:    password,
	}, nil
}

func (s *service) Start() error {
	withWalletOnly := true
	if err := s.start(withWalletOnly); err != nil {
		return err
	}

	if s.opts.WalletUnlockPasswordFile != "" {
		pwdBytes, _ := ioutil.ReadFile(s.opts.WalletUnlockPasswordFile)
		password := string(pwdBytes)

		if err := s.opts.AppConfig.WalletService().Wallet().Unlock(
			context.Background(), password,
		); err != nil {
			return fmt.Errorf("failed to auto unlock wallet: %s", err)
		}

		s.onUnlock(password)
	}

	return nil
}

func (s *service) Stop() {
	if s.password != "" {
		walletSvc := s.opts.AppConfig.WalletService().Wallet()
		walletSvc.Lock(context.Background(), s.password)
	}
	stopMacaroonSvc := true
	s.stop(stopMacaroonSvc)

	s.opts.AppConfig.RepoManager().Close()
	log.Debug("closed connection with database")

	s.opts.AppConfig.PubSubService().Close()
	log.Debug("closed connection with pubsub")

	s.opts.AppConfig.WalletService().Close()
	log.Debug("closed connection with ocean wallet")
}

func (s *service) withMacaroons() bool {
	return !s.opts.NoMacaroons
}

func (s *service) start(withWalletOnly bool) error {
	operatorTlsConfig, err := s.opts.operatorTLSConfig()
	if err != nil {
		return err
	}
	operatorServer, err := s.newOperatorServer(
		operatorTlsConfig, !withWalletOnly,
	)
	if err != nil {
		return err
	}

	var tradeServer *http.Server
	if !withWalletOnly {
		tradeTlsConfig, err := s.opts.tradeTLSConfig()
		if err != nil {
			return err
		}
		tradeServer, err = s.newTradeServer(tradeTlsConfig)
		if err != nil {
			return err
		}
	}

	s.operatorServer = operatorServer
	s.tradeServer = tradeServer

	if s.opts.NoOperatorTls {
		go s.operatorServer.ListenAndServe()
	} else {
		go s.operatorServer.ListenAndServeTLS("", "")
	}
	log.Infof("wallet interface is listening on %s", s.opts.operatorServerAddr())

	if !withWalletOnly {
		if len(s.opts.TradeTLSCert) <= 0 {
			go s.tradeServer.ListenAndServe()
		} else {
			s.tradeServer.ListenAndServeTLS("", "")
		}
		log.Infof("operator interface is listening on %s", s.opts.operatorServerAddr())
		log.Infof("trade interface is listening on %s", s.opts.tradeServerAddr())
	}

	return nil
}

func (s *service) stop(stopMacaroonSvc bool) {
	if s.withMacaroons() && stopMacaroonSvc {
		s.macaroonSvc.Close()
		log.Debug("closed connection with macaroon db")
	}

	s.operatorServer.Shutdown(context.Background())
	log.Debug("stopped operator server")

	if s.tradeServer != nil {
		s.tradeServer.Shutdown(context.Background())
		log.Debug("stopped trade server")
	}
}

func (s *service) newOperatorServer(
	tlsConfig *tls.Config, withOperatorHandler bool,
) (*http.Server, error) {
	serverOpts := []grpc.ServerOption{
		interceptor.UnaryInterceptor(s.macaroonSvc),
		interceptor.StreamInterceptor(s.macaroonSvc),
	}

	creds := insecure.NewCredentials()
	if tlsConfig != nil {
		creds = credentials.NewTLS(tlsConfig)
	}
	serverOpts = append(serverOpts, grpc.Creds(creds))

	var adminMacaroonPath string
	if s.withMacaroons() {
		adminMacaroonPath = filepath.Join(
			s.opts.macaroonsDatadir(), AdminMacaroonFile,
		)
	}

	grpcServer := grpc.NewServer(serverOpts...)

	walletSvc := s.opts.AppConfig.WalletService().Wallet()
	walletHandler := grpchandler.NewWalletHandler(
		walletSvc, s.opts.BuildData, adminMacaroonPath,
		s.onInit, s.onUnlock, s.onLock, s.onChangePwd,
	)
	daemonv2.RegisterWalletServiceServer(
		grpcServer, walletHandler,
	)

	if withOperatorHandler {
		operatorHandler := grpchandler.NewOperatorHandler(
			s.opts.AppConfig.OperatorService(), s.validatePassword,
		)
		daemonv2.RegisterOperatorServiceServer(grpcServer, operatorHandler)
	}

	// grpcweb wrapped server
	grpcWebServer := grpcweb.WrapServer(
		grpcServer,
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	handler := router(grpcServer, grpcWebServer, nil)
	mux := http.NewServeMux()
	mux.Handle("/", handler)

	httpServerHandler := http.Handler(mux)
	if s.opts.NoOperatorTls {
		httpServerHandler = h2c.NewHandler(httpServerHandler, &http2.Server{})
	}

	return &http.Server{
		Addr:      s.opts.operatorServerAddr(),
		Handler:   httpServerHandler,
		TLSConfig: tlsConfig,
	}, nil
}

func (s *service) newTradeServer(tlsConfig *tls.Config) (*http.Server, error) {
	serverOpts := []grpc.ServerOption{
		interceptor.UnaryInterceptor(s.macaroonSvc),
		interceptor.StreamInterceptor(s.macaroonSvc),
	}

	creds := insecure.NewCredentials()
	if tlsConfig != nil {
		creds = credentials.NewTLS(tlsConfig)
	}
	serverOpts = append(serverOpts, grpc.Creds(creds))

	grpcServer := grpc.NewServer(serverOpts...)
	tradeHandler := grpchandler.NewTradeHandler(s.opts.AppConfig.TradeService())
	tdexv1.RegisterTradeServiceServer(grpcServer, tradeHandler)
	transportHandler := grpchandler.NewTransportHandler()
	tdexv1.RegisterTransportServiceServer(grpcServer, transportHandler)

	// grpcweb wrapped server
	grpcWebServer := grpcweb.WrapServer(
		grpcServer,
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	// grpc gateway reverse proxy
	dialOpts := make([]grpc.DialOption, 0)
	if len(s.opts.TradeTLSCert) <= 0 {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(
			// #nosec
			credentials.NewTLS(&tls.Config{
				InsecureSkipVerify: true,
			}),
		))
	}
	conn, err := grpc.DialContext(
		context.Background(), s.opts.tradeClientAddr(), dialOpts...,
	)
	if err != nil {
		return nil, err
	}
	gwmux := runtime.NewServeMux()
	tdexv1.RegisterTransportServiceHandler(context.Background(), gwmux, conn)
	tdexv1.RegisterTradeServiceHandler(context.Background(), gwmux, conn)
	grpcGateway := http.Handler(gwmux)

	handler := router(grpcServer, grpcWebServer, grpcGateway)
	mux := http.NewServeMux()
	mux.Handle("/", handler)

	httpServerHandler := http.Handler(mux)
	if s.opts.NoOperatorTls {
		httpServerHandler = h2c.NewHandler(httpServerHandler, &http2.Server{})
	}

	return &http.Server{
		Addr:      s.opts.tradeServerAddr(),
		Handler:   httpServerHandler,
		TLSConfig: tlsConfig,
	}, nil
}

func (s *service) onInit(password string) {
	s.password = password
}

func (s *service) onUnlock(password string) {
	if s.withMacaroons() {
		pwd := []byte(password)
		if err := s.macaroonSvc.CreateUnlock(&pwd); err != nil {
			log.WithError(err).Warn("failed to unlock macaroon store")
		}
	}

	stopMacaroonSvc := true
	s.stop(!stopMacaroonSvc)

	withWalletOnly := true
	if err := s.start(!withWalletOnly); err != nil {
		log.WithError(err).Warn("failed to load handlers to interface after unlock")
	}

	if s.password == "" {
		s.password = password
	}
}

func (s *service) onLock(_ string) {
	stopMacaroonSvc := true
	s.stop(stopMacaroonSvc)
	withWalletOnly := true
	s.start(withWalletOnly)
}

func (s *service) onChangePwd(oldPassword, newPassword string) {
	if !s.withMacaroons() {
		return
	}
	oldPwd, newPwd := []byte(oldPassword), []byte(newPassword)
	if err := s.macaroonSvc.ChangePassword(oldPwd, newPwd); err != nil {
		log.WithError(err).Warn("failed to change password of macaroon store")
	}
}

func (s *service) validatePassword(pwd string) bool {
	return pwd == s.password
}
