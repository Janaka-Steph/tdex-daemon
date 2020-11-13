package inmemory

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/tdex-network/tdex-daemon/config"
	"github.com/tdex-network/tdex-daemon/internal/core/domain"
	mm "github.com/tdex-network/tdex-daemon/pkg/marketmaking"
	"github.com/tdex-network/tdex-daemon/pkg/marketmaking/formula"
)

func newMockDb() *DbManager {
	config.Set(config.UnspentTtlKey, 2)

	dbManager := NewDbManager()

	if err := insertMarkets(dbManager); err != nil {
		panic(err)
	}
	if err := insertUnspents(dbManager); err != nil {
		panic(err)
	}
	if err := insertTrades(dbManager); err != nil {
		panic(err)
	}

	if err := insertVault(dbManager); err != nil {
		panic(err)
	}

	return dbManager
}

func insertMarkets(db *DbManager) error {
	markets := []domain.Market{
		{
			AccountIndex: 5,
			BaseAsset:    "ah5",
			QuoteAsset:   "qh5",
			Fee:          0,
			FeeAsset:     "",
			Tradable:     true,
			Strategy:     mm.NewStrategyFromFormula(formula.BalancedReserves{}),
			Price:        domain.Prices{},
		},
		{
			AccountIndex: 6,
			BaseAsset:    "ah6",
			QuoteAsset:   "qh6",
			Fee:          0,
			FeeAsset:     "",
			Tradable:     true,
			Strategy:     mm.NewStrategyFromFormula(formula.BalancedReserves{}),
			Price:        domain.Prices{},
		},
		{
			AccountIndex: 7,
			BaseAsset:    "ah7",
			QuoteAsset:   "qh7",
			Fee:          0,
			FeeAsset:     "",
			Tradable:     false,
			Strategy:     mm.NewStrategyFromFormula(formula.BalancedReserves{}),
			Price:        domain.Prices{},
		},
		{
			AccountIndex: 8,
			BaseAsset:    "ah8",
			QuoteAsset:   "qh8",
			Fee:          0,
			FeeAsset:     "",
			Tradable:     false,
			Strategy:     mm.NewStrategyFromFormula(formula.BalancedReserves{}),
			Price:        domain.Prices{},
		},
		{
			AccountIndex: 9,
			BaseAsset:    "ah9",
			QuoteAsset:   "qh9",
			Fee:          0,
			FeeAsset:     "",
			Tradable:     false,
			Strategy:     mm.NewStrategyFromFormula(formula.BalancedReserves{}),
			Price:        domain.Prices{},
		},
	}

	for _, v := range markets {
		db.marketStore.markets[v.AccountIndex] = v
		if len(v.QuoteAsset) > 0 {
			db.marketStore.accountsByAsset[v.QuoteAsset] = v.AccountIndex
		}
	}

	return nil
}

func insertUnspents(db *DbManager) error {
	unspents := []domain.Unspent{
		{
			TxID:         "1",
			VOut:         0,
			Value:        4,
			AssetHash:    "ah",
			Address:      "a",
			Spent:        false,
			Locked:       false,
			ScriptPubKey: nil,
			LockedBy:     nil,
			Confirmed:    true,
		},
		{
			TxID:         "1",
			VOut:         1,
			Value:        2,
			AssetHash:    "ah",
			Address:      "adr",
			Spent:        false,
			Locked:       false,
			ScriptPubKey: nil,
			LockedBy:     nil,
			Confirmed:    true,
		},
		{
			TxID:         "2",
			VOut:         1,
			Value:        4,
			AssetHash:    "ah",
			Address:      "adre",
			Spent:        false,
			Locked:       false,
			ScriptPubKey: nil,
			LockedBy:     nil,
			Confirmed:    false,
		},
		{
			TxID:         "2",
			VOut:         2,
			Value:        9,
			AssetHash:    "ah",
			Address:      "adra",
			Spent:        false,
			Locked:       false,
			ScriptPubKey: nil,
			LockedBy:     nil,
			Confirmed:    false,
		},
		{
			TxID:         "3",
			VOut:         1,
			Value:        4,
			AssetHash:    "ah",
			Address:      "a",
			Spent:        false,
			Locked:       false,
			ScriptPubKey: nil,
			LockedBy:     nil,
			Confirmed:    false,
		},
		{
			TxID:         "3",
			VOut:         0,
			Value:        2,
			AssetHash:    "ah",
			Address:      "a",
			Spent:        false,
			Locked:       false,
			ScriptPubKey: nil,
			LockedBy:     nil,
			Confirmed:    false,
		},
	}
	for _, v := range unspents {
		db.unspentStore.unspents[v.Key()] = v
	}

	return nil
}

func insertVault(db *DbManager) error {
	vault := &domain.Vault{
		EncryptedMnemonic:      "dVoBFte1oeRkPl8Vf8DzBP3PRnzPA3fxtyvDHXFGYAS9MP8V2Sc9nHcQW4PrMkQNnf2uGrDg81dFgBrwqv1n3frXxRBKhp83fSsTm4xqj8+jdwTI3nouFmi1W/O4UqpHdQ62EYoabJQtKpptWO11TFJzw8WF02pfS6git8YjLR4xrnfp2LkOEjSU9CI82ZasF46WZFKcpeUJTAsxU/03ONpAdwwEsC96f1KAvh8tqaO0yLDOcmPf8a5B82jefgncCRrt32kCpbpIE4YiCFrqqdUHXKH+",
		PassphraseHash:         []byte("pass"),
		Accounts:               map[int]*domain.Account{},
		AccountAndKeyByAddress: map[string]domain.AccountAndKey{},
	}

	db.vaultStore.vault = vault
	return nil
}

func insertTrades(db *DbManager) error {
	tradeID1, _ := uuid.Parse("cc913d4e-174e-449c-82b4-e848d57cbf2e")
	tradeID2, _ := uuid.Parse("5440a53e-58d2-4e3d-8380-20410e687589")
	tradeID3, _ := uuid.Parse("2a12e2a0-d99c-4bd3-ad99-03dd926ae080")

	trades := []domain.Trade{
		{
			ID:               tradeID1,
			MarketQuoteAsset: "mqa1",
			SwapRequest: domain.Swap{
				ID: "1",
			},
			SwapAccept: domain.Swap{
				ID: "2",
			},
			SwapComplete: domain.Swap{
				ID: "3",
			},
			SwapFail: domain.Swap{
				ID: "4",
			},
		},
		{
			ID:               tradeID2,
			MarketQuoteAsset: "mqa2",
			SwapRequest: domain.Swap{
				ID: "11",
			},
			SwapAccept: domain.Swap{
				ID: "21",
			},
			SwapComplete: domain.Swap{
				ID: "31",
			},
			SwapFail: domain.Swap{
				ID: "41",
			},
		},
		{
			ID:               tradeID3,
			MarketQuoteAsset: "mqa2",
			TxID:             "424",
			SwapRequest: domain.Swap{
				ID: "12",
			},
			SwapAccept: domain.Swap{
				ID: "22",
			},
			SwapComplete: domain.Swap{
				ID: "32",
			},
			SwapFail: domain.Swap{
				ID: "42",
			},
		},
	}

	for _, v := range trades {
		db.tradeStore.trades[v.ID] = v

		if v.MarketQuoteAsset != "" {
			if db.tradeStore.tradesByMarket[v.MarketQuoteAsset] == nil {
				array := []uuid.UUID{v.ID}
				db.tradeStore.tradesByMarket[v.MarketQuoteAsset] = array
			}

			db.tradeStore.tradesByMarket[v.MarketQuoteAsset] = append(db.tradeStore.tradesByMarket[v.MarketQuoteAsset], v.ID)
		}

		db.tradeStore.tradesBySwapAcceptID[v.SwapAccept.ID] = v.ID
	}
	return nil
}

var (
	hexCharset  = "0123456789abcdef"
	addrCharset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"
	seededRand  = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func randUnspents() []domain.Unspent {
	numUnspents := randInt(1, 4)
	unspents := make([]domain.Unspent, numUnspents)
	for i := range unspents {
		unspents[i] = domain.Unspent{
			TxID:            randStr(32),
			VOut:            uint32(randInt(0, 15)),
			Value:           uint64(randInt(1, 100000000000)),
			AssetHash:       randStr(32),
			ValueCommitment: "08" + randStr(32),
			AssetCommitment: "0b" + randStr(32),
			ScriptPubKey:    append([]byte{0, 20}, randBytes(20)...),
			Nonce:           append([]byte{2}, randBytes(32)...),
			RangeProof:      make([]byte, 4174),
			SurjectionProof: make([]byte, 64),
			Address:         randAddr(),
			Confirmed:       true,
		}
	}
	return unspents
}

func randInt(min, max int) int {
	return seededRand.Intn(max-min+1) + min
}

func randAddr() string {
	return "el1qq" + string(_randBytes(48, addrCharset))
}

func randStr(length int) string {
	return string(randBytes(length))
}

func randBytes(length int) []byte {
	return _randBytes(length, hexCharset)
}

func _randBytes(length int, charset string) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = hexCharset[randInt(0, len(hexCharset)-1)]
	}
	return b
}
