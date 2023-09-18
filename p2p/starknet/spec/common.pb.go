// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: p2p/proto/common.proto

package spec

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

type Iteration_Direction int32

const (
	Iteration_Forward  Iteration_Direction = 0
	Iteration_Backward Iteration_Direction = 1
)

// Enum value maps for Iteration_Direction.
var (
	Iteration_Direction_name = map[int32]string{
		0: "Forward",
		1: "Backward",
	}
	Iteration_Direction_value = map[string]int32{
		"Forward":  0,
		"Backward": 1,
	}
)

func (x Iteration_Direction) Enum() *Iteration_Direction {
	p := new(Iteration_Direction)
	*p = x
	return p
}

func (x Iteration_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Iteration_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_p2p_proto_common_proto_enumTypes[0].Descriptor()
}

func (Iteration_Direction) Type() protoreflect.EnumType {
	return &file_p2p_proto_common_proto_enumTypes[0]
}

func (x Iteration_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Iteration_Direction.Descriptor instead.
func (Iteration_Direction) EnumDescriptor() ([]byte, []int) {
	return file_p2p_proto_common_proto_rawDescGZIP(), []int{8, 0}
}

type Fin_Error int32

const (
	Fin_busy     Fin_Error = 0
	Fin_too_much Fin_Error = 1
	Fin_unknown  Fin_Error = 2
	Fin_pruned   Fin_Error = 3
)

// Enum value maps for Fin_Error.
var (
	Fin_Error_name = map[int32]string{
		0: "busy",
		1: "too_much",
		2: "unknown",
		3: "pruned",
	}
	Fin_Error_value = map[string]int32{
		"busy":     0,
		"too_much": 1,
		"unknown":  2,
		"pruned":   3,
	}
)

func (x Fin_Error) Enum() *Fin_Error {
	p := new(Fin_Error)
	*p = x
	return p
}

func (x Fin_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Fin_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_p2p_proto_common_proto_enumTypes[1].Descriptor()
}

func (Fin_Error) Type() protoreflect.EnumType {
	return &file_p2p_proto_common_proto_enumTypes[1]
}

func (x Fin_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Fin_Error.Descriptor instead.
func (Fin_Error) EnumDescriptor() ([]byte, []int) {
	return file_p2p_proto_common_proto_rawDescGZIP(), []int{9, 0}
}

type Felt252 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elements []byte `protobuf:"bytes,1,opt,name=elements,proto3" json:"elements,omitempty"`
}

func (x *Felt252) Reset() {
	*x = Felt252{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Felt252) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Felt252) ProtoMessage() {}

func (x *Felt252) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Felt252.ProtoReflect.Descriptor instead.
func (*Felt252) Descriptor() ([]byte, []int) {
	return file_p2p_proto_common_proto_rawDescGZIP(), []int{0}
}

func (x *Felt252) GetElements() []byte {
	if x != nil {
		return x.Elements
	}
	return nil
}

type Hash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elements []byte `protobuf:"bytes,1,opt,name=elements,proto3" json:"elements,omitempty"`
}

func (x *Hash) Reset() {
	*x = Hash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hash) ProtoMessage() {}

func (x *Hash) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hash.ProtoReflect.Descriptor instead.
func (*Hash) Descriptor() ([]byte, []int) {
	return file_p2p_proto_common_proto_rawDescGZIP(), []int{1}
}

func (x *Hash) GetElements() []byte {
	if x != nil {
		return x.Elements
	}
	return nil
}

type Hashes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Hash `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Hashes) Reset() {
	*x = Hashes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hashes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hashes) ProtoMessage() {}

func (x *Hashes) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hashes.ProtoReflect.Descriptor instead.
func (*Hashes) Descriptor() ([]byte, []int) {
	return file_p2p_proto_common_proto_rawDescGZIP(), []int{2}
}

func (x *Hashes) GetItems() []*Hash {
	if x != nil {
		return x.Items
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elements []byte `protobuf:"bytes,1,opt,name=elements,proto3" json:"elements,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_p2p_proto_common_proto_rawDescGZIP(), []int{3}
}

func (x *Address) GetElements() []byte {
	if x != nil {
		return x.Elements
	}
	return nil
}

type PeerID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PeerID) Reset() {
	*x = PeerID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerID) ProtoMessage() {}

func (x *PeerID) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerID.ProtoReflect.Descriptor instead.
func (*PeerID) Descriptor() ([]byte, []int) {
	return file_p2p_proto_common_proto_rawDescGZIP(), []int{4}
}

func (x *PeerID) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type ConsensusSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	R *Felt252 `protobuf:"bytes,1,opt,name=r,proto3" json:"r,omitempty"`
	S *Felt252 `protobuf:"bytes,2,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *ConsensusSignature) Reset() {
	*x = ConsensusSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusSignature) ProtoMessage() {}

func (x *ConsensusSignature) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusSignature.ProtoReflect.Descriptor instead.
func (*ConsensusSignature) Descriptor() ([]byte, []int) {
	return file_p2p_proto_common_proto_rawDescGZIP(), []int{5}
}

func (x *ConsensusSignature) GetR() *Felt252 {
	if x != nil {
		return x.R
	}
	return nil
}

func (x *ConsensusSignature) GetS() *Felt252 {
	if x != nil {
		return x.S
	}
	return nil
}

type Merkle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NLeaves uint32 `protobuf:"varint,1,opt,name=n_leaves,json=nLeaves,proto3" json:"n_leaves,omitempty"` // needed to know the height, so as to how many nodes to expect in a proof.
	// and also when receiving all leaves, how many to expect
	Root *Hash `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
}

func (x *Merkle) Reset() {
	*x = Merkle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Merkle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Merkle) ProtoMessage() {}

func (x *Merkle) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Merkle.ProtoReflect.Descriptor instead.
func (*Merkle) Descriptor() ([]byte, []int) {
	return file_p2p_proto_common_proto_rawDescGZIP(), []int{6}
}

func (x *Merkle) GetNLeaves() uint32 {
	if x != nil {
		return x.NLeaves
	}
	return 0
}

func (x *Merkle) GetRoot() *Hash {
	if x != nil {
		return x.Root
	}
	return nil
}

type Patricia struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint32 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Root   *Hash  `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
}

func (x *Patricia) Reset() {
	*x = Patricia{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Patricia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Patricia) ProtoMessage() {}

func (x *Patricia) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Patricia.ProtoReflect.Descriptor instead.
func (*Patricia) Descriptor() ([]byte, []int) {
	return file_p2p_proto_common_proto_rawDescGZIP(), []int{7}
}

func (x *Patricia) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Patricia) GetRoot() *Hash {
	if x != nil {
		return x.Root
	}
	return nil
}

type Iteration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartBlock uint64              `protobuf:"varint,1,opt,name=start_block,json=startBlock,proto3" json:"start_block,omitempty"` // exclude start from the result
	Direction  Iteration_Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=Iteration_Direction" json:"direction,omitempty"`
	Limit      uint64              `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Step       uint64              `protobuf:"varint,4,opt,name=step,proto3" json:"step,omitempty"` // to allow interleaving from several nodes
}

func (x *Iteration) Reset() {
	*x = Iteration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Iteration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Iteration) ProtoMessage() {}

func (x *Iteration) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Iteration.ProtoReflect.Descriptor instead.
func (*Iteration) Descriptor() ([]byte, []int) {
	return file_p2p_proto_common_proto_rawDescGZIP(), []int{8}
}

func (x *Iteration) GetStartBlock() uint64 {
	if x != nil {
		return x.StartBlock
	}
	return 0
}

func (x *Iteration) GetDirection() Iteration_Direction {
	if x != nil {
		return x.Direction
	}
	return Iteration_Forward
}

func (x *Iteration) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Iteration) GetStep() uint64 {
	if x != nil {
		return x.Step
	}
	return 0
}

// mark the end of a stream of messages
// TBD: may not be required if we open a stream per request.
type Fin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Fin_Error `protobuf:"varint,1,opt,name=error,proto3,enum=Fin_Error,oneof" json:"error,omitempty"`
}

func (x *Fin) Reset() {
	*x = Fin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fin) ProtoMessage() {}

func (x *Fin) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fin.ProtoReflect.Descriptor instead.
func (*Fin) Descriptor() ([]byte, []int) {
	return file_p2p_proto_common_proto_rawDescGZIP(), []int{9}
}

func (x *Fin) GetError() Fin_Error {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return Fin_busy
}

var File_p2p_proto_common_proto protoreflect.FileDescriptor

var file_p2p_proto_common_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x07, 0x46, 0x65, 0x6c, 0x74,
	0x32, 0x35, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x22, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x25, 0x0a, 0x06, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x1b, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x25, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x18, 0x0a, 0x06, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x12, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x16, 0x0a, 0x01, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x46,
	0x65, 0x6c, 0x74, 0x32, 0x35, 0x32, 0x52, 0x01, 0x72, 0x12, 0x16, 0x0a, 0x01, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x46, 0x65, 0x6c, 0x74, 0x32, 0x35, 0x32, 0x52, 0x01,
	0x73, 0x22, 0x3e, 0x0a, 0x06, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e,
	0x5f, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x04, 0x72, 0x6f, 0x6f,
	0x74, 0x22, 0x3d, 0x0a, 0x08, 0x50, 0x61, 0x74, 0x72, 0x69, 0x63, 0x69, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74,
	0x22, 0xb2, 0x01, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x32, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0x26, 0x0a,
	0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x61, 0x63, 0x6b, 0x77,
	0x61, 0x72, 0x64, 0x10, 0x01, 0x22, 0x70, 0x0a, 0x03, 0x46, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x46, 0x69,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x88, 0x01, 0x01, 0x22, 0x38, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x08, 0x0a, 0x04,
	0x62, 0x75, 0x73, 0x79, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x74, 0x6f, 0x6f, 0x5f, 0x6d, 0x75,
	0x63, 0x68, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x70, 0x72, 0x75, 0x6e, 0x65, 0x64, 0x10, 0x03, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_p2p_proto_common_proto_rawDescOnce sync.Once
	file_p2p_proto_common_proto_rawDescData = file_p2p_proto_common_proto_rawDesc
)

func file_p2p_proto_common_proto_rawDescGZIP() []byte {
	file_p2p_proto_common_proto_rawDescOnce.Do(func() {
		file_p2p_proto_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_proto_common_proto_rawDescData)
	})
	return file_p2p_proto_common_proto_rawDescData
}

var file_p2p_proto_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_p2p_proto_common_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_p2p_proto_common_proto_goTypes = []interface{}{
	(Iteration_Direction)(0),   // 0: Iteration.Direction
	(Fin_Error)(0),             // 1: Fin.Error
	(*Felt252)(nil),            // 2: Felt252
	(*Hash)(nil),               // 3: Hash
	(*Hashes)(nil),             // 4: Hashes
	(*Address)(nil),            // 5: Address
	(*PeerID)(nil),             // 6: PeerID
	(*ConsensusSignature)(nil), // 7: ConsensusSignature
	(*Merkle)(nil),             // 8: Merkle
	(*Patricia)(nil),           // 9: Patricia
	(*Iteration)(nil),          // 10: Iteration
	(*Fin)(nil),                // 11: Fin
}
var file_p2p_proto_common_proto_depIdxs = []int32{
	3, // 0: Hashes.items:type_name -> Hash
	2, // 1: ConsensusSignature.r:type_name -> Felt252
	2, // 2: ConsensusSignature.s:type_name -> Felt252
	3, // 3: Merkle.root:type_name -> Hash
	3, // 4: Patricia.root:type_name -> Hash
	0, // 5: Iteration.direction:type_name -> Iteration.Direction
	1, // 6: Fin.error:type_name -> Fin.Error
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_p2p_proto_common_proto_init() }
func file_p2p_proto_common_proto_init() {
	if File_p2p_proto_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_p2p_proto_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Felt252); i {
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
		file_p2p_proto_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hash); i {
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
		file_p2p_proto_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hashes); i {
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
		file_p2p_proto_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_p2p_proto_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerID); i {
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
		file_p2p_proto_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusSignature); i {
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
		file_p2p_proto_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Merkle); i {
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
		file_p2p_proto_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Patricia); i {
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
		file_p2p_proto_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Iteration); i {
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
		file_p2p_proto_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fin); i {
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
	file_p2p_proto_common_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_p2p_proto_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_p2p_proto_common_proto_goTypes,
		DependencyIndexes: file_p2p_proto_common_proto_depIdxs,
		EnumInfos:         file_p2p_proto_common_proto_enumTypes,
		MessageInfos:      file_p2p_proto_common_proto_msgTypes,
	}.Build()
	File_p2p_proto_common_proto = out.File
	file_p2p_proto_common_proto_rawDesc = nil
	file_p2p_proto_common_proto_goTypes = nil
	file_p2p_proto_common_proto_depIdxs = nil
}
