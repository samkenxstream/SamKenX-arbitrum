// Code generated by protoc-gen-go. DO NOT EDIT.
// source: valmessage.proto

package valmessage

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	value "github.com/offchainlabs/arb-avm/value"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Address struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{0}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}

func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}

func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}

func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}

func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type TokenTypeBuf struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenTypeBuf) Reset()         { *m = TokenTypeBuf{} }
func (m *TokenTypeBuf) String() string { return proto.CompactTextString(m) }
func (*TokenTypeBuf) ProtoMessage()    {}
func (*TokenTypeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{1}
}

func (m *TokenTypeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenTypeBuf.Unmarshal(m, b)
}

func (m *TokenTypeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenTypeBuf.Marshal(b, m, deterministic)
}

func (m *TokenTypeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenTypeBuf.Merge(m, src)
}

func (m *TokenTypeBuf) XXX_Size() int {
	return xxx_messageInfo_TokenTypeBuf.Size(m)
}

func (m *TokenTypeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenTypeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_TokenTypeBuf proto.InternalMessageInfo

func (m *TokenTypeBuf) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type VMConfiguration struct {
	GracePeriod           uint64               `protobuf:"varint,1,opt,name=grace_period,json=gracePeriod,proto3" json:"grace_period,omitempty"`
	EscrowRequired        *value.BigIntegerBuf `protobuf:"bytes,2,opt,name=escrow_required,json=escrowRequired,proto3" json:"escrow_required,omitempty"`
	EscrowCurrency        *Address             `protobuf:"bytes,3,opt,name=escrow_currency,json=escrowCurrency,proto3" json:"escrow_currency,omitempty"`
	AssertKeys            []*Address           `protobuf:"bytes,4,rep,name=assert_keys,json=assertKeys,proto3" json:"assert_keys,omitempty"`
	MaxExecutionStepCount uint32               `protobuf:"varint,5,opt,name=max_execution_step_count,json=maxExecutionStepCount,proto3" json:"max_execution_step_count,omitempty"`
	Owner                 *Address             `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}             `json:"-"`
	XXX_unrecognized      []byte               `json:"-"`
	XXX_sizecache         int32                `json:"-"`
}

func (m *VMConfiguration) Reset()         { *m = VMConfiguration{} }
func (m *VMConfiguration) String() string { return proto.CompactTextString(m) }
func (*VMConfiguration) ProtoMessage()    {}
func (*VMConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{2}
}

func (m *VMConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMConfiguration.Unmarshal(m, b)
}

func (m *VMConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMConfiguration.Marshal(b, m, deterministic)
}

func (m *VMConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMConfiguration.Merge(m, src)
}

func (m *VMConfiguration) XXX_Size() int {
	return xxx_messageInfo_VMConfiguration.Size(m)
}

func (m *VMConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_VMConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_VMConfiguration proto.InternalMessageInfo

func (m *VMConfiguration) GetGracePeriod() uint64 {
	if m != nil {
		return m.GracePeriod
	}
	return 0
}

func (m *VMConfiguration) GetEscrowRequired() *value.BigIntegerBuf {
	if m != nil {
		return m.EscrowRequired
	}
	return nil
}

func (m *VMConfiguration) GetEscrowCurrency() *Address {
	if m != nil {
		return m.EscrowCurrency
	}
	return nil
}

func (m *VMConfiguration) GetAssertKeys() []*Address {
	if m != nil {
		return m.AssertKeys
	}
	return nil
}

func (m *VMConfiguration) GetMaxExecutionStepCount() uint32 {
	if m != nil {
		return m.MaxExecutionStepCount
	}
	return 0
}

func (m *VMConfiguration) GetOwner() *Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func init() {
	proto.RegisterType((*Address)(nil), "valmessage.Address")
	proto.RegisterType((*TokenTypeBuf)(nil), "valmessage.TokenTypeBuf")
	proto.RegisterType((*VMConfiguration)(nil), "valmessage.VMConfiguration")
}

func init() { proto.RegisterFile("valmessage.proto", fileDescriptor_b34ccd35396e2606) }

var fileDescriptor_b34ccd35396e2606 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0x49, 0xff, 0x3e, 0x98, 0xf6, 0xb3, 0x3a, 0x56, 0x08, 0x6e, 0x8c, 0xc5, 0x45, 0x5d,
	0x98, 0x48, 0x15, 0xdc, 0xe8, 0xc2, 0x14, 0x17, 0x22, 0x82, 0xc4, 0xe2, 0xc2, 0x4d, 0x98, 0x4c,
	0x4e, 0xd2, 0xa1, 0xc9, 0x4c, 0x9c, 0x9f, 0xb6, 0xb9, 0x46, 0x6f, 0x4a, 0x9a, 0xb4, 0xd6, 0x4d,
	0x37, 0x81, 0xbc, 0xef, 0xf3, 0x0c, 0xe7, 0xcc, 0xa0, 0xc3, 0x05, 0xc9, 0x72, 0x50, 0x8a, 0xa4,
	0xe0, 0x16, 0x52, 0x68, 0x81, 0xd1, 0x2e, 0x39, 0x3d, 0x5a, 0x90, 0xcc, 0x80, 0x57, 0x7d, 0xeb,
	0x7a, 0x78, 0x86, 0xfe, 0x3d, 0xc6, 0xb1, 0x04, 0xa5, 0xf0, 0x00, 0xb5, 0xab, 0xc6, 0xb6, 0x1c,
	0x6b, 0xd4, 0x0b, 0xea, 0x9f, 0xe1, 0x05, 0xea, 0x4d, 0xc5, 0x1c, 0xf8, 0xb4, 0x2c, 0xc0, 0x37,
	0xc9, 0x1e, 0xea, 0xbb, 0x81, 0xfa, 0x1f, 0xaf, 0x13, 0xc1, 0x13, 0x96, 0x1a, 0x49, 0x34, 0x13,
	0x1c, 0x9f, 0xa3, 0x5e, 0x2a, 0x09, 0x85, 0xb0, 0x00, 0xc9, 0x44, 0x5c, 0x09, 0xad, 0xa0, 0x5b,
	0x65, 0x6f, 0x55, 0x84, 0x1f, 0x50, 0x1f, 0x14, 0x95, 0x62, 0x19, 0x4a, 0xf8, 0x32, 0x4c, 0x42,
	0x6c, 0x37, 0x1c, 0x6b, 0xd4, 0x1d, 0x0f, 0xdc, 0x7a, 0x48, 0x9f, 0xa5, 0xcf, 0x5c, 0x43, 0x0a,
	0xd2, 0x37, 0x49, 0x70, 0x50, 0xc3, 0xc1, 0x86, 0xc5, 0xf7, 0xbf, 0x3a, 0x35, 0x52, 0x02, 0xa7,
	0xa5, 0xdd, 0xac, 0xf4, 0x63, 0xf7, 0xcf, 0x3d, 0x6c, 0xf6, 0xdb, 0xda, 0x93, 0x0d, 0x8a, 0x6f,
	0x51, 0x97, 0x28, 0x05, 0x52, 0x87, 0x73, 0x28, 0x95, 0xdd, 0x72, 0x9a, 0xfb, 0x4c, 0x54, 0x73,
	0x2f, 0x50, 0x2a, 0x7c, 0x87, 0xec, 0x9c, 0xac, 0x42, 0x58, 0x01, 0x35, 0xeb, 0x35, 0x43, 0xa5,
	0xa1, 0x08, 0xa9, 0x30, 0x5c, 0xdb, 0x6d, 0xc7, 0x1a, 0xfd, 0x0f, 0x4e, 0x72, 0xb2, 0x7a, 0xda,
	0xd6, 0xef, 0x1a, 0x8a, 0xc9, 0xba, 0xc4, 0x97, 0xa8, 0x2d, 0x96, 0x1c, 0xa4, 0xdd, 0xd9, 0x3f,
	0x62, 0x4d, 0xf8, 0xe3, 0xcf, 0xeb, 0x94, 0xe9, 0x99, 0x89, 0x5c, 0x2a, 0x72, 0x4f, 0x24, 0x09,
	0x9d, 0x11, 0xc6, 0x33, 0x12, 0x29, 0x8f, 0xc8, 0xe8, 0x6a, 0x41, 0x32, 0x16, 0x13, 0x2d, 0xa4,
	0xb7, 0x3b, 0x22, 0xea, 0x54, 0xef, 0x79, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0x99, 0xe0, 0x87,
	0x04, 0x02, 0x02, 0x00, 0x00,
}
