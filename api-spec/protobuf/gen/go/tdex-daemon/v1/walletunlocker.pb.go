// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: tdex-daemon/v1/walletunlocker.proto

package daemonv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InitWalletReply_Status int32

const (
	InitWalletReply_PROCESSING InitWalletReply_Status = 0
	InitWalletReply_DONE       InitWalletReply_Status = 1
)

// Enum value maps for InitWalletReply_Status.
var (
	InitWalletReply_Status_name = map[int32]string{
		0: "PROCESSING",
		1: "DONE",
	}
	InitWalletReply_Status_value = map[string]int32{
		"PROCESSING": 0,
		"DONE":       1,
	}
)

func (x InitWalletReply_Status) Enum() *InitWalletReply_Status {
	p := new(InitWalletReply_Status)
	*p = x
	return p
}

func (x InitWalletReply_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InitWalletReply_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_tdex_daemon_v1_walletunlocker_proto_enumTypes[0].Descriptor()
}

func (InitWalletReply_Status) Type() protoreflect.EnumType {
	return &file_tdex_daemon_v1_walletunlocker_proto_enumTypes[0]
}

func (x InitWalletReply_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InitWalletReply_Status.Descriptor instead.
func (InitWalletReply_Status) EnumDescriptor() ([]byte, []int) {
	return file_tdex_daemon_v1_walletunlocker_proto_rawDescGZIP(), []int{3, 0}
}

type GenSeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GenSeedRequest) Reset() {
	*x = GenSeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenSeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenSeedRequest) ProtoMessage() {}

func (x *GenSeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenSeedRequest.ProtoReflect.Descriptor instead.
func (*GenSeedRequest) Descriptor() ([]byte, []int) {
	return file_tdex_daemon_v1_walletunlocker_proto_rawDescGZIP(), []int{0}
}

type GenSeedReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeedMnemonic []string `protobuf:"bytes,1,rep,name=seed_mnemonic,json=seedMnemonic,proto3" json:"seed_mnemonic,omitempty"`
}

func (x *GenSeedReply) Reset() {
	*x = GenSeedReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenSeedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenSeedReply) ProtoMessage() {}

func (x *GenSeedReply) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenSeedReply.ProtoReflect.Descriptor instead.
func (*GenSeedReply) Descriptor() ([]byte, []int) {
	return file_tdex_daemon_v1_walletunlocker_proto_rawDescGZIP(), []int{1}
}

func (x *GenSeedReply) GetSeedMnemonic() []string {
	if x != nil {
		return x.SeedMnemonic
	}
	return nil
}

type InitWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	// wallet_password is the passphrase that should be used to encrypt the
	// wallet. This MUST be at least 8 chars in length. After creation, this
	// password is required to unlock the daemon.
	WalletPassword []byte `protobuf:"bytes,1,opt,name=wallet_password,json=walletPassword,proto3" json:"wallet_password,omitempty"`
	//
	// seed_mnemonic is a 24-word mnemonic that encodes a prior seed obtained by the
	// user. This MUST be a generated by the GenSeed method
	SeedMnemonic []string `protobuf:"bytes,2,rep,name=seed_mnemonic,json=seedMnemonic,proto3" json:"seed_mnemonic,omitempty"`
	//
	// the flag to let the daemon restore existing funds for the wallet.
	Restore bool `protobuf:"varint,3,opt,name=restore,proto3" json:"restore,omitempty"`
}

func (x *InitWalletRequest) Reset() {
	*x = InitWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitWalletRequest) ProtoMessage() {}

func (x *InitWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitWalletRequest.ProtoReflect.Descriptor instead.
func (*InitWalletRequest) Descriptor() ([]byte, []int) {
	return file_tdex_daemon_v1_walletunlocker_proto_rawDescGZIP(), []int{2}
}

func (x *InitWalletRequest) GetWalletPassword() []byte {
	if x != nil {
		return x.WalletPassword
	}
	return nil
}

func (x *InitWalletRequest) GetSeedMnemonic() []string {
	if x != nil {
		return x.SeedMnemonic
	}
	return nil
}

func (x *InitWalletRequest) GetRestore() bool {
	if x != nil {
		return x.Restore
	}
	return false
}

type InitWalletReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account int32                  `protobuf:"varint,1,opt,name=account,proto3" json:"account,omitempty"`
	Status  InitWalletReply_Status `protobuf:"varint,2,opt,name=status,proto3,enum=tdex.daemon.v1.InitWalletReply_Status" json:"status,omitempty"`
	Data    string                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *InitWalletReply) Reset() {
	*x = InitWalletReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitWalletReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitWalletReply) ProtoMessage() {}

func (x *InitWalletReply) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitWalletReply.ProtoReflect.Descriptor instead.
func (*InitWalletReply) Descriptor() ([]byte, []int) {
	return file_tdex_daemon_v1_walletunlocker_proto_rawDescGZIP(), []int{3}
}

func (x *InitWalletReply) GetAccount() int32 {
	if x != nil {
		return x.Account
	}
	return 0
}

func (x *InitWalletReply) GetStatus() InitWalletReply_Status {
	if x != nil {
		return x.Status
	}
	return InitWalletReply_PROCESSING
}

func (x *InitWalletReply) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type UnlockWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	// wallet_password should be the current valid passphrase for the daemon. This
	// will be required to decrypt on-disk material that the daemon requires to
	// function properly.
	WalletPassword []byte `protobuf:"bytes,1,opt,name=wallet_password,json=walletPassword,proto3" json:"wallet_password,omitempty"`
}

func (x *UnlockWalletRequest) Reset() {
	*x = UnlockWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockWalletRequest) ProtoMessage() {}

func (x *UnlockWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockWalletRequest.ProtoReflect.Descriptor instead.
func (*UnlockWalletRequest) Descriptor() ([]byte, []int) {
	return file_tdex_daemon_v1_walletunlocker_proto_rawDescGZIP(), []int{4}
}

func (x *UnlockWalletRequest) GetWalletPassword() []byte {
	if x != nil {
		return x.WalletPassword
	}
	return nil
}

type UnlockWalletReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnlockWalletReply) Reset() {
	*x = UnlockWalletReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockWalletReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockWalletReply) ProtoMessage() {}

func (x *UnlockWalletReply) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockWalletReply.ProtoReflect.Descriptor instead.
func (*UnlockWalletReply) Descriptor() ([]byte, []int) {
	return file_tdex_daemon_v1_walletunlocker_proto_rawDescGZIP(), []int{5}
}

type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	// current_password should be the current valid passphrase used to unlock the
	// daemon.
	CurrentPassword []byte `protobuf:"bytes,1,opt,name=current_password,json=currentPassword,proto3" json:"current_password,omitempty"`
	//
	// new_password should be the new passphrase that will be needed to unlock the
	// daemon.
	NewPassword []byte `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_tdex_daemon_v1_walletunlocker_proto_rawDescGZIP(), []int{6}
}

func (x *ChangePasswordRequest) GetCurrentPassword() []byte {
	if x != nil {
		return x.CurrentPassword
	}
	return nil
}

func (x *ChangePasswordRequest) GetNewPassword() []byte {
	if x != nil {
		return x.NewPassword
	}
	return nil
}

type ChangePasswordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangePasswordReply) Reset() {
	*x = ChangePasswordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordReply) ProtoMessage() {}

func (x *ChangePasswordReply) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordReply.ProtoReflect.Descriptor instead.
func (*ChangePasswordReply) Descriptor() ([]byte, []int) {
	return file_tdex_daemon_v1_walletunlocker_proto_rawDescGZIP(), []int{7}
}

type IsReadyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IsReadyRequest) Reset() {
	*x = IsReadyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsReadyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsReadyRequest) ProtoMessage() {}

func (x *IsReadyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsReadyRequest.ProtoReflect.Descriptor instead.
func (*IsReadyRequest) Descriptor() ([]byte, []int) {
	return file_tdex_daemon_v1_walletunlocker_proto_rawDescGZIP(), []int{8}
}

type IsReadyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Initialized bool `protobuf:"varint,1,opt,name=initialized,proto3" json:"initialized,omitempty"`
	Unlocked    bool `protobuf:"varint,2,opt,name=unlocked,proto3" json:"unlocked,omitempty"`
	Synced      bool `protobuf:"varint,3,opt,name=synced,proto3" json:"synced,omitempty"`
}

func (x *IsReadyReply) Reset() {
	*x = IsReadyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsReadyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsReadyReply) ProtoMessage() {}

func (x *IsReadyReply) ProtoReflect() protoreflect.Message {
	mi := &file_tdex_daemon_v1_walletunlocker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsReadyReply.ProtoReflect.Descriptor instead.
func (*IsReadyReply) Descriptor() ([]byte, []int) {
	return file_tdex_daemon_v1_walletunlocker_proto_rawDescGZIP(), []int{9}
}

func (x *IsReadyReply) GetInitialized() bool {
	if x != nil {
		return x.Initialized
	}
	return false
}

func (x *IsReadyReply) GetUnlocked() bool {
	if x != nil {
		return x.Unlocked
	}
	return false
}

func (x *IsReadyReply) GetSynced() bool {
	if x != nil {
		return x.Synced
	}
	return false
}

var File_tdex_daemon_v1_walletunlocker_proto protoreflect.FileDescriptor

var file_tdex_daemon_v1_walletunlocker_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x64, 0x65, 0x78, 0x2d, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x64, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x53, 0x65, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x53, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x65, 0x64, 0x5f,
	0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x65, 0x65, 0x64, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x22, 0x7b, 0x0a, 0x11,
	0x49, 0x6e, 0x69, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65,
	0x65, 0x64, 0x5f, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x65, 0x65, 0x64, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x0f, 0x49, 0x6e,
	0x69, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x64,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53,
	0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x22,
	0x3e, 0x0a, 0x13, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x13, 0x0a, 0x11, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x65, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x64, 0x0a, 0x0c, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x32, 0xac, 0x03, 0x0a, 0x0e, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x47, 0x0a,
	0x07, 0x47, 0x65, 0x6e, 0x53, 0x65, 0x65, 0x64, 0x12, 0x1e, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x53, 0x65, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x53, 0x65, 0x65,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x52, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x64, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x64,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x12, 0x56, 0x0a, 0x0c, 0x55, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x74, 0x64, 0x65,
	0x78, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x5c, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x64, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74, 0x64,
	0x65, 0x78, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x47, 0x0a, 0x07, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x1e, 0x2e, 0x74, 0x64,
	0x65, 0x78, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x64,
	0x65, 0x78, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0xd9, 0x01, 0x0a, 0x12, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x64, 0x65, 0x78, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x42, 0x13, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x64, 0x65, 0x78, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x74, 0x64, 0x65, 0x78, 0x2d, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x64, 0x65, 0x78, 0x2d, 0x64, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x54, 0x44, 0x58, 0xaa, 0x02, 0x0e, 0x54, 0x64, 0x65, 0x78, 0x2e, 0x44, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x54, 0x64, 0x65, 0x78, 0x5c, 0x44, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x54, 0x64, 0x65, 0x78, 0x5c, 0x44, 0x61, 0x65,
	0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x10, 0x54, 0x64, 0x65, 0x78, 0x3a, 0x3a, 0x44, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tdex_daemon_v1_walletunlocker_proto_rawDescOnce sync.Once
	file_tdex_daemon_v1_walletunlocker_proto_rawDescData = file_tdex_daemon_v1_walletunlocker_proto_rawDesc
)

func file_tdex_daemon_v1_walletunlocker_proto_rawDescGZIP() []byte {
	file_tdex_daemon_v1_walletunlocker_proto_rawDescOnce.Do(func() {
		file_tdex_daemon_v1_walletunlocker_proto_rawDescData = protoimpl.X.CompressGZIP(file_tdex_daemon_v1_walletunlocker_proto_rawDescData)
	})
	return file_tdex_daemon_v1_walletunlocker_proto_rawDescData
}

var file_tdex_daemon_v1_walletunlocker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tdex_daemon_v1_walletunlocker_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_tdex_daemon_v1_walletunlocker_proto_goTypes = []interface{}{
	(InitWalletReply_Status)(0),   // 0: tdex.daemon.v1.InitWalletReply.Status
	(*GenSeedRequest)(nil),        // 1: tdex.daemon.v1.GenSeedRequest
	(*GenSeedReply)(nil),          // 2: tdex.daemon.v1.GenSeedReply
	(*InitWalletRequest)(nil),     // 3: tdex.daemon.v1.InitWalletRequest
	(*InitWalletReply)(nil),       // 4: tdex.daemon.v1.InitWalletReply
	(*UnlockWalletRequest)(nil),   // 5: tdex.daemon.v1.UnlockWalletRequest
	(*UnlockWalletReply)(nil),     // 6: tdex.daemon.v1.UnlockWalletReply
	(*ChangePasswordRequest)(nil), // 7: tdex.daemon.v1.ChangePasswordRequest
	(*ChangePasswordReply)(nil),   // 8: tdex.daemon.v1.ChangePasswordReply
	(*IsReadyRequest)(nil),        // 9: tdex.daemon.v1.IsReadyRequest
	(*IsReadyReply)(nil),          // 10: tdex.daemon.v1.IsReadyReply
}
var file_tdex_daemon_v1_walletunlocker_proto_depIdxs = []int32{
	0,  // 0: tdex.daemon.v1.InitWalletReply.status:type_name -> tdex.daemon.v1.InitWalletReply.Status
	1,  // 1: tdex.daemon.v1.WalletUnlocker.GenSeed:input_type -> tdex.daemon.v1.GenSeedRequest
	3,  // 2: tdex.daemon.v1.WalletUnlocker.InitWallet:input_type -> tdex.daemon.v1.InitWalletRequest
	5,  // 3: tdex.daemon.v1.WalletUnlocker.UnlockWallet:input_type -> tdex.daemon.v1.UnlockWalletRequest
	7,  // 4: tdex.daemon.v1.WalletUnlocker.ChangePassword:input_type -> tdex.daemon.v1.ChangePasswordRequest
	9,  // 5: tdex.daemon.v1.WalletUnlocker.IsReady:input_type -> tdex.daemon.v1.IsReadyRequest
	2,  // 6: tdex.daemon.v1.WalletUnlocker.GenSeed:output_type -> tdex.daemon.v1.GenSeedReply
	4,  // 7: tdex.daemon.v1.WalletUnlocker.InitWallet:output_type -> tdex.daemon.v1.InitWalletReply
	6,  // 8: tdex.daemon.v1.WalletUnlocker.UnlockWallet:output_type -> tdex.daemon.v1.UnlockWalletReply
	8,  // 9: tdex.daemon.v1.WalletUnlocker.ChangePassword:output_type -> tdex.daemon.v1.ChangePasswordReply
	10, // 10: tdex.daemon.v1.WalletUnlocker.IsReady:output_type -> tdex.daemon.v1.IsReadyReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_tdex_daemon_v1_walletunlocker_proto_init() }
func file_tdex_daemon_v1_walletunlocker_proto_init() {
	if File_tdex_daemon_v1_walletunlocker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tdex_daemon_v1_walletunlocker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenSeedRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tdex_daemon_v1_walletunlocker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenSeedReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tdex_daemon_v1_walletunlocker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitWalletRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tdex_daemon_v1_walletunlocker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitWalletReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tdex_daemon_v1_walletunlocker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockWalletRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tdex_daemon_v1_walletunlocker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockWalletReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tdex_daemon_v1_walletunlocker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tdex_daemon_v1_walletunlocker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tdex_daemon_v1_walletunlocker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsReadyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tdex_daemon_v1_walletunlocker_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsReadyReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tdex_daemon_v1_walletunlocker_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tdex_daemon_v1_walletunlocker_proto_goTypes,
		DependencyIndexes: file_tdex_daemon_v1_walletunlocker_proto_depIdxs,
		EnumInfos:         file_tdex_daemon_v1_walletunlocker_proto_enumTypes,
		MessageInfos:      file_tdex_daemon_v1_walletunlocker_proto_msgTypes,
	}.Build()
	File_tdex_daemon_v1_walletunlocker_proto = out.File
	file_tdex_daemon_v1_walletunlocker_proto_rawDesc = nil
	file_tdex_daemon_v1_walletunlocker_proto_goTypes = nil
	file_tdex_daemon_v1_walletunlocker_proto_depIdxs = nil
}
