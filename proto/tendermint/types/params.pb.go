// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: tendermint/types/params.proto

package types

import (
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ConsensusParams contains consensus critical parameters that determine the
// validity of blocks.
type ConsensusParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block     *BlockParams     `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Evidence  *EvidenceParams  `protobuf:"bytes,2,opt,name=evidence,proto3" json:"evidence,omitempty"`
	Validator *ValidatorParams `protobuf:"bytes,3,opt,name=validator,proto3" json:"validator,omitempty"`
	Version   *VersionParams   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ConsensusParams) Reset() {
	*x = ConsensusParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusParams) ProtoMessage() {}

func (x *ConsensusParams) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusParams.ProtoReflect.Descriptor instead.
func (*ConsensusParams) Descriptor() ([]byte, []int) {
	return file_tendermint_types_params_proto_rawDescGZIP(), []int{0}
}

func (x *ConsensusParams) GetBlock() *BlockParams {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *ConsensusParams) GetEvidence() *EvidenceParams {
	if x != nil {
		return x.Evidence
	}
	return nil
}

func (x *ConsensusParams) GetValidator() *ValidatorParams {
	if x != nil {
		return x.Validator
	}
	return nil
}

func (x *ConsensusParams) GetVersion() *VersionParams {
	if x != nil {
		return x.Version
	}
	return nil
}

// BlockParams contains limits on the block size.
type BlockParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Max block size, in bytes.
	// Note: must be greater than 0
	MaxBytes int64 `protobuf:"varint,1,opt,name=max_bytes,json=maxBytes,proto3" json:"max_bytes,omitempty"`
	// Max gas per block.
	// Note: must be greater or equal to -1
	MaxGas int64 `protobuf:"varint,2,opt,name=max_gas,json=maxGas,proto3" json:"max_gas,omitempty"`
	// Minimum time increment between consecutive blocks (in milliseconds) If the
	// block header timestamp is ahead of the system clock, decrease this value.
	//
	// Not exposed to the application.
	TimeIotaMs int64 `protobuf:"varint,3,opt,name=time_iota_ms,json=timeIotaMs,proto3" json:"time_iota_ms,omitempty"`
}

func (x *BlockParams) Reset() {
	*x = BlockParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_params_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockParams) ProtoMessage() {}

func (x *BlockParams) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_params_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockParams.ProtoReflect.Descriptor instead.
func (*BlockParams) Descriptor() ([]byte, []int) {
	return file_tendermint_types_params_proto_rawDescGZIP(), []int{1}
}

func (x *BlockParams) GetMaxBytes() int64 {
	if x != nil {
		return x.MaxBytes
	}
	return 0
}

func (x *BlockParams) GetMaxGas() int64 {
	if x != nil {
		return x.MaxGas
	}
	return 0
}

func (x *BlockParams) GetTimeIotaMs() int64 {
	if x != nil {
		return x.TimeIotaMs
	}
	return 0
}

// EvidenceParams determine how we handle evidence of malfeasance.
type EvidenceParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Max age of evidence, in blocks.
	//
	// The basic formula for calculating this is: MaxAgeDuration / {average block
	// time}.
	MaxAgeNumBlocks int64 `protobuf:"varint,1,opt,name=max_age_num_blocks,json=maxAgeNumBlocks,proto3" json:"max_age_num_blocks,omitempty"`
	// Max age of evidence, in time.
	//
	// It should correspond with an app's "unbonding period" or other similar
	// mechanism for handling [Nothing-At-Stake
	// attacks](https://github.com/ethereum/wiki/wiki/Proof-of-Stake-FAQ#what-is-the-nothing-at-stake-problem-and-how-can-it-be-fixed).
	MaxAgeDuration *durationpb.Duration `protobuf:"bytes,2,opt,name=max_age_duration,json=maxAgeDuration,proto3" json:"max_age_duration,omitempty"`
	// This sets the maximum size of total evidence in bytes that can be committed in a single block.
	// and should fall comfortably under the max block bytes.
	// Default is 1048576 or 1MB
	MaxBytes int64 `protobuf:"varint,3,opt,name=max_bytes,json=maxBytes,proto3" json:"max_bytes,omitempty"`
}

func (x *EvidenceParams) Reset() {
	*x = EvidenceParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_params_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvidenceParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvidenceParams) ProtoMessage() {}

func (x *EvidenceParams) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_params_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvidenceParams.ProtoReflect.Descriptor instead.
func (*EvidenceParams) Descriptor() ([]byte, []int) {
	return file_tendermint_types_params_proto_rawDescGZIP(), []int{2}
}

func (x *EvidenceParams) GetMaxAgeNumBlocks() int64 {
	if x != nil {
		return x.MaxAgeNumBlocks
	}
	return 0
}

func (x *EvidenceParams) GetMaxAgeDuration() *durationpb.Duration {
	if x != nil {
		return x.MaxAgeDuration
	}
	return nil
}

func (x *EvidenceParams) GetMaxBytes() int64 {
	if x != nil {
		return x.MaxBytes
	}
	return 0
}

// ValidatorParams restrict the public key types validators can use.
// NOTE: uses ABCI pubkey naming, not Amino names.
type ValidatorParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKeyTypes []string `protobuf:"bytes,1,rep,name=pub_key_types,json=pubKeyTypes,proto3" json:"pub_key_types,omitempty"`
}

func (x *ValidatorParams) Reset() {
	*x = ValidatorParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_params_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorParams) ProtoMessage() {}

func (x *ValidatorParams) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_params_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorParams.ProtoReflect.Descriptor instead.
func (*ValidatorParams) Descriptor() ([]byte, []int) {
	return file_tendermint_types_params_proto_rawDescGZIP(), []int{3}
}

func (x *ValidatorParams) GetPubKeyTypes() []string {
	if x != nil {
		return x.PubKeyTypes
	}
	return nil
}

// VersionParams contains the ABCI application version.
type VersionParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppVersion uint64 `protobuf:"varint,1,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
}

func (x *VersionParams) Reset() {
	*x = VersionParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_params_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionParams) ProtoMessage() {}

func (x *VersionParams) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_params_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionParams.ProtoReflect.Descriptor instead.
func (*VersionParams) Descriptor() ([]byte, []int) {
	return file_tendermint_types_params_proto_rawDescGZIP(), []int{4}
}

func (x *VersionParams) GetAppVersion() uint64 {
	if x != nil {
		return x.AppVersion
	}
	return 0
}

// HashedParams is a subset of ConsensusParams.
//
// It is hashed into the Header.ConsensusHash.
type HashedParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockMaxBytes int64 `protobuf:"varint,1,opt,name=block_max_bytes,json=blockMaxBytes,proto3" json:"block_max_bytes,omitempty"`
	BlockMaxGas   int64 `protobuf:"varint,2,opt,name=block_max_gas,json=blockMaxGas,proto3" json:"block_max_gas,omitempty"`
}

func (x *HashedParams) Reset() {
	*x = HashedParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_params_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashedParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashedParams) ProtoMessage() {}

func (x *HashedParams) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_params_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashedParams.ProtoReflect.Descriptor instead.
func (*HashedParams) Descriptor() ([]byte, []int) {
	return file_tendermint_types_params_proto_rawDescGZIP(), []int{5}
}

func (x *HashedParams) GetBlockMaxBytes() int64 {
	if x != nil {
		return x.BlockMaxBytes
	}
	return 0
}

func (x *HashedParams) GetBlockMaxGas() int64 {
	if x != nil {
		return x.BlockMaxGas
	}
	return 0
}

var File_tendermint_types_params_proto protoreflect.FileDescriptor

var file_tendermint_types_params_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x39, 0x0a, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x42, 0x0a, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x76, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00,
	0x52, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x3f, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x65, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6d, 0x61, 0x78, 0x47, 0x61, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x69, 0x6f, 0x74, 0x61, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74,
	0x69, 0x6d, 0x65, 0x49, 0x6f, 0x74, 0x61, 0x4d, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x0e, 0x45, 0x76,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x12,
	0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x4d, 0x0a, 0x10, 0x6d, 0x61, 0x78,
	0x5f, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08,
	0xc8, 0xde, 0x1f, 0x00, 0x98, 0xdf, 0x1f, 0x01, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x08, 0xb8, 0xa0,
	0x1f, 0x01, 0xe8, 0xa0, 0x1f, 0x01, 0x22, 0x3a, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x70,
	0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x08, 0xb8, 0xa0, 0x1f, 0x01, 0xe8, 0xa0,
	0x1f, 0x01, 0x22, 0x5a, 0x0a, 0x0c, 0x48, 0x61, 0x73, 0x68, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x61, 0x78, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4d, 0x61, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x61, 0x78, 0x47, 0x61, 0x73, 0x42, 0x3c,
	0xa8, 0xe2, 0x1e, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x59, 0x75, 0x6e, 0x77, 0x65, 0x69, 0x4d, 0x61, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tendermint_types_params_proto_rawDescOnce sync.Once
	file_tendermint_types_params_proto_rawDescData = file_tendermint_types_params_proto_rawDesc
)

func file_tendermint_types_params_proto_rawDescGZIP() []byte {
	file_tendermint_types_params_proto_rawDescOnce.Do(func() {
		file_tendermint_types_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_tendermint_types_params_proto_rawDescData)
	})
	return file_tendermint_types_params_proto_rawDescData
}

var file_tendermint_types_params_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tendermint_types_params_proto_goTypes = []interface{}{
	(*ConsensusParams)(nil),     // 0: tendermint.types.ConsensusParams
	(*BlockParams)(nil),         // 1: tendermint.types.BlockParams
	(*EvidenceParams)(nil),      // 2: tendermint.types.EvidenceParams
	(*ValidatorParams)(nil),     // 3: tendermint.types.ValidatorParams
	(*VersionParams)(nil),       // 4: tendermint.types.VersionParams
	(*HashedParams)(nil),        // 5: tendermint.types.HashedParams
	(*durationpb.Duration)(nil), // 6: google.protobuf.Duration
}
var file_tendermint_types_params_proto_depIdxs = []int32{
	1, // 0: tendermint.types.ConsensusParams.block:type_name -> tendermint.types.BlockParams
	2, // 1: tendermint.types.ConsensusParams.evidence:type_name -> tendermint.types.EvidenceParams
	3, // 2: tendermint.types.ConsensusParams.validator:type_name -> tendermint.types.ValidatorParams
	4, // 3: tendermint.types.ConsensusParams.version:type_name -> tendermint.types.VersionParams
	6, // 4: tendermint.types.EvidenceParams.max_age_duration:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tendermint_types_params_proto_init() }
func file_tendermint_types_params_proto_init() {
	if File_tendermint_types_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tendermint_types_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusParams); i {
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
		file_tendermint_types_params_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockParams); i {
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
		file_tendermint_types_params_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvidenceParams); i {
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
		file_tendermint_types_params_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorParams); i {
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
		file_tendermint_types_params_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionParams); i {
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
		file_tendermint_types_params_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashedParams); i {
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
			RawDescriptor: file_tendermint_types_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tendermint_types_params_proto_goTypes,
		DependencyIndexes: file_tendermint_types_params_proto_depIdxs,
		MessageInfos:      file_tendermint_types_params_proto_msgTypes,
	}.Build()
	File_tendermint_types_params_proto = out.File
	file_tendermint_types_params_proto_rawDesc = nil
	file_tendermint_types_params_proto_goTypes = nil
	file_tendermint_types_params_proto_depIdxs = nil
}
