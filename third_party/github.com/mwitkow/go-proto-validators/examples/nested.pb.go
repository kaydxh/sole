// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/nested.proto

package validator_examples

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
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

type InnerMessage struct {
	// some_integer can only be in range (1, 100).
	SomeInteger          int32    `protobuf:"varint,1,opt,name=some_integer,json=someInteger,proto3" json:"some_integer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InnerMessage) Reset()         { *m = InnerMessage{} }
func (m *InnerMessage) String() string { return proto.CompactTextString(m) }
func (*InnerMessage) ProtoMessage()    {}
func (*InnerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_48ff2d59662eadb1, []int{0}
}

func (m *InnerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerMessage.Unmarshal(m, b)
}
func (m *InnerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerMessage.Marshal(b, m, deterministic)
}
func (m *InnerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerMessage.Merge(m, src)
}
func (m *InnerMessage) XXX_Size() int {
	return xxx_messageInfo_InnerMessage.Size(m)
}
func (m *InnerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InnerMessage proto.InternalMessageInfo

func (m *InnerMessage) GetSomeInteger() int32 {
	if m != nil {
		return m.SomeInteger
	}
	return 0
}

type OuterMessage struct {
	// important_string must be a lowercase alpha-numeric of 5 to 30 characters (RE2 syntax).
	ImportantString string `protobuf:"bytes,1,opt,name=important_string,json=importantString,proto3" json:"important_string,omitempty"`
	// proto3 doesn't have `required`, the `msg_exist` enforces presence of InnerMessage.
	Inner                *InnerMessage `protobuf:"bytes,2,opt,name=inner,proto3" json:"inner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *OuterMessage) Reset()         { *m = OuterMessage{} }
func (m *OuterMessage) String() string { return proto.CompactTextString(m) }
func (*OuterMessage) ProtoMessage()    {}
func (*OuterMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_48ff2d59662eadb1, []int{1}
}

func (m *OuterMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OuterMessage.Unmarshal(m, b)
}
func (m *OuterMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OuterMessage.Marshal(b, m, deterministic)
}
func (m *OuterMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OuterMessage.Merge(m, src)
}
func (m *OuterMessage) XXX_Size() int {
	return xxx_messageInfo_OuterMessage.Size(m)
}
func (m *OuterMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OuterMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OuterMessage proto.InternalMessageInfo

func (m *OuterMessage) GetImportantString() string {
	if m != nil {
		return m.ImportantString
	}
	return ""
}

func (m *OuterMessage) GetInner() *InnerMessage {
	if m != nil {
		return m.Inner
	}
	return nil
}

func init() {
	proto.RegisterType((*InnerMessage)(nil), "validator.examples.InnerMessage")
	proto.RegisterType((*OuterMessage)(nil), "validator.examples.OuterMessage")
}

func init() { proto.RegisterFile("examples/nested.proto", fileDescriptor_48ff2d59662eadb1) }

var fileDescriptor_48ff2d59662eadb1 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0xcf, 0x4b, 0x2d, 0x2e, 0x49, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x2a, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x83, 0x29, 0x90,
	0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0xcf, 0x2c,
	0xc9, 0xce, 0x2f, 0xd7, 0x4f, 0xcf, 0xd7, 0x05, 0x6b, 0xd0, 0x85, 0xab, 0x2f, 0xd6, 0x47, 0x68,
	0x05, 0x4b, 0x29, 0x59, 0x73, 0xf1, 0x78, 0xe6, 0xe5, 0xa5, 0x16, 0xf9, 0xa6, 0x16, 0x17, 0x27,
	0xa6, 0xa7, 0x0a, 0x69, 0x73, 0xf1, 0x14, 0xe7, 0xe7, 0xa6, 0xc6, 0x67, 0xe6, 0x95, 0xa4, 0xa6,
	0xa7, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x3a, 0x71, 0x3c, 0xba, 0x2f, 0xcf, 0x22, 0xc0,
	0x20, 0x91, 0x12, 0xc4, 0x0d, 0x92, 0xf5, 0x84, 0x48, 0x2a, 0xf5, 0x32, 0x72, 0xf1, 0xf8, 0x97,
	0x96, 0x20, 0x74, 0xdb, 0x72, 0x09, 0x64, 0xe6, 0x16, 0xe4, 0x17, 0x95, 0x24, 0xe6, 0x95, 0xc4,
	0x17, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x83, 0x4d, 0xe0, 0x74, 0x12, 0x7a, 0x74, 0x5f, 0x9e, 0x8f,
	0x8b, 0x27, 0x2e, 0x3a, 0x51, 0xb7, 0x2a, 0xb6, 0xda, 0x48, 0xc7, 0xb4, 0x56, 0x25, 0x88, 0x1f,
	0xae, 0x36, 0x18, 0xac, 0x54, 0xc8, 0x8e, 0x8b, 0x35, 0x13, 0xe4, 0x18, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0x6e, 0x23, 0x05, 0x3d, 0x4c, 0x8f, 0xea, 0x21, 0xbb, 0xd6, 0x89, 0xed, 0xd1, 0x7d, 0x79,
	0x26, 0x05, 0xc6, 0x20, 0x88, 0xb6, 0x24, 0x36, 0xb0, 0x9f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x08, 0x68, 0xd9, 0xd5, 0x38, 0x01, 0x00, 0x00,
}