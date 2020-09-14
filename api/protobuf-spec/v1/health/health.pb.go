// Code generated by protoc-gen-go. DO NOT EDIT.
// source: health.proto

// 健康监测服务 API

package health

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Status int32

const (
	Status_ok Status = 0
)

var Status_name = map[int32]string{
	0: "ok",
}

var Status_value = map[string]int32{
	"ok": 0,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fdbebe66dda7cb29, []int{0}
}

type AliveResponse struct {
	// Status always contains "ok".
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=sole.api.v1.health.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AliveResponse) Reset()         { *m = AliveResponse{} }
func (m *AliveResponse) String() string { return proto.CompactTextString(m) }
func (*AliveResponse) ProtoMessage()    {}
func (*AliveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbebe66dda7cb29, []int{0}
}

func (m *AliveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AliveResponse.Unmarshal(m, b)
}
func (m *AliveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AliveResponse.Marshal(b, m, deterministic)
}
func (m *AliveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AliveResponse.Merge(m, src)
}
func (m *AliveResponse) XXX_Size() int {
	return xxx_messageInfo_AliveResponse.Size(m)
}
func (m *AliveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AliveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AliveResponse proto.InternalMessageInfo

func (m *AliveResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_ok
}

type ReadyResponse struct {
	// Status always contains "ok".
	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=sole.api.v1.health.Status" json:"status,omitempty"`
	// Errors contains a list of errors that caused the not ready status.
	Errors               map[string]string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReadyResponse) Reset()         { *m = ReadyResponse{} }
func (m *ReadyResponse) String() string { return proto.CompactTextString(m) }
func (*ReadyResponse) ProtoMessage()    {}
func (*ReadyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbebe66dda7cb29, []int{1}
}

func (m *ReadyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadyResponse.Unmarshal(m, b)
}
func (m *ReadyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadyResponse.Marshal(b, m, deterministic)
}
func (m *ReadyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadyResponse.Merge(m, src)
}
func (m *ReadyResponse) XXX_Size() int {
	return xxx_messageInfo_ReadyResponse.Size(m)
}
func (m *ReadyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadyResponse proto.InternalMessageInfo

func (m *ReadyResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_ok
}

func (m *ReadyResponse) GetErrors() map[string]string {
	if m != nil {
		return m.Errors
	}
	return nil
}

type VersionResponse struct {
	// 服务软件版本
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbebe66dda7cb29, []int{2}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterEnum("sole.api.v1.health.Status", Status_name, Status_value)
	proto.RegisterType((*AliveResponse)(nil), "sole.api.v1.health.AliveResponse")
	proto.RegisterType((*ReadyResponse)(nil), "sole.api.v1.health.ReadyResponse")
	proto.RegisterMapType((map[string]string)(nil), "sole.api.v1.health.ReadyResponse.ErrorsEntry")
	proto.RegisterType((*VersionResponse)(nil), "sole.api.v1.health.VersionResponse")
}

func init() {
	proto.RegisterFile("health.proto", fileDescriptor_fdbebe66dda7cb29)
}

var fileDescriptor_fdbebe66dda7cb29 = []byte{
	// 1198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0x41, 0x6c, 0x13, 0x47,
	0x14, 0xc5, 0x36, 0x0e, 0x62, 0x80, 0x12, 0xb6, 0x08, 0x45, 0x6e, 0xd5, 0x9a, 0xed, 0x05, 0x15,
	0x79, 0x17, 0xdc, 0x4b, 0xeb, 0x9e, 0x8c, 0x14, 0x29, 0x55, 0x5b, 0x09, 0x05, 0xd4, 0x43, 0x6f,
	0x1b, 0x7b, 0x70, 0xb6, 0x38, 0xde, 0xed, 0xee, 0xda, 0xd4, 0x42, 0x48, 0x8e, 0x9b, 0x04, 0xe3,
	0xd8, 0xd8, 0x75, 0x93, 0x10, 0xd9, 0x06, 0x83, 0x48, 0xd2, 0x28, 0x0a, 0x76, 0x10, 0x10, 0x1c,
	0xc7, 0x09, 0x97, 0x1e, 0x7a, 0x68, 0x2f, 0xbd, 0x57, 0x3d, 0xb0, 0x33, 0x6b, 0x4b, 0x95, 0xac,
	0x1e, 0x7a, 0xaf, 0x76, 0x67, 0xed, 0x38, 0xc1, 0x39, 0x54, 0x55, 0xa5, 0xaa, 0xea, 0x69, 0x67,
	0xde, 0xfc, 0xff, 0xe7, 0xcd, 0xfb, 0x6f, 0x46, 0x0b, 0x8e, 0x8e, 0x42, 0x2e, 0xa8, 0x8c, 0x32,
	0xa2, 0x24, 0x28, 0x02, 0x45, 0xc9, 0x42, 0x10, 0x32, 0x9c, 0xc8, 0x33, 0x91, 0xf3, 0x0c, 0x59,
	0x71, 0xbc, 0x11, 0x10, 0x84, 0x40, 0x10, 0xb2, 0x46, 0xc4, 0x48, 0xf8, 0x0a, 0x0b, 0xc7, 0x44,
	0x25, 0x4a, 0x12, 0x1c, 0xce, 0xbd, 0x8b, 0x7e, 0x28, 0xfb, 0x24, 0x5e, 0x54, 0x04, 0xc9, 0x8c,
	0xf0, 0x04, 0x78, 0x65, 0x34, 0x3c, 0xc2, 0xf8, 0x84, 0x31, 0x96, 0x04, 0x73, 0x22, 0x2f, 0xbf,
	0x3a, 0x64, 0x39, 0x91, 0x67, 0x47, 0x15, 0x45, 0x1c, 0x11, 0xfc, 0xed, 0xea, 0x97, 0xba, 0x73,
	0x25, 0xd1, 0xe7, 0x82, 0x3e, 0x41, 0x8e, 0xca, 0x0a, 0x34, 0xa7, 0x01, 0x4e, 0x81, 0xd7, 0xb8,
	0x28, 0xd9, 0xdd, 0xe7, 0x0a, 0xc0, 0x90, 0x4b, 0xbe, 0xc6, 0x05, 0x02, 0x50, 0x62, 0x05, 0x51,
	0xe1, 0x85, 0x90, 0xcc, 0x72, 0xa1, 0x90, 0xa0, 0x70, 0xc6, 0x98, 0x14, 0xa5, 0xaf, 0x80, 0x63,
	0xde, 0x20, 0x1f, 0x81, 0xc3, 0x50, 0x16, 0x85, 0x90, 0x0c, 0x29, 0x37, 0xe8, 0x93, 0x15, 0x4e,
	0x09, 0xcb, 0x03, 0x16, 0xa7, 0xe5, 0xcc, 0x6b, 0x6e, 0x07, 0xf3, 0xaa, 0x0a, 0xcc, 0x25, 0x23,
	0x62, 0xd8, 0x8c, 0xf4, 0x9c, 0xce, 0x79, 0xdf, 0x02, 0x87, 0x7f, 0xb6, 0x98, 0x73, 0xf7, 0x09,
	0xea, 0xf8, 0x75, 0x9a, 0x8c, 0x69, 0x0f, 0x2d, 0x5c, 0xa5, 0x6f, 0xd0, 0xbf, 0x58, 0xc0, 0xb1,
	0x61, 0xc8, 0xf9, 0xa3, 0x7f, 0x67, 0x23, 0x6a, 0x10, 0xf4, 0x41, 0x49, 0x12, 0x24, 0x79, 0xc0,
	0xea, 0xb4, 0x9d, 0x39, 0xe2, 0x76, 0xf5, 0xca, 0xd9, 0xb5, 0x0d, 0x33, 0x68, 0xc4, 0x0f, 0x86,
	0x14, 0x29, 0x3a, 0x6c, 0x26, 0x3b, 0x3e, 0x00, 0x47, 0xba, 0x60, 0xaa, 0x1f, 0xd8, 0xae, 0xc2,
	0xa8, 0x41, 0xe3, 0xf0, 0xb0, 0x3e, 0xa4, 0x4e, 0x02, 0x7b, 0x84, 0x0b, 0x86, 0xe1, 0x80, 0xd5,
	0xc0, 0xc8, 0xc4, 0x63, 0x7d, 0xdf, 0xe2, 0x39, 0x95, 0xf3, 0xbe, 0xde, 0xeb, 0x7c, 0x9f, 0x80,
	0xe3, 0x9f, 0x41, 0x49, 0xe6, 0x85, 0x50, 0xe7, 0x80, 0x03, 0xe0, 0x50, 0x84, 0x40, 0x66, 0xe9,
	0xf6, 0xd4, 0xf3, 0x76, 0xce, 0xfb, 0xa6, 0xdb, 0x41, 0x0d, 0x5c, 0xa7, 0x4d, 0x84, 0xf6, 0xd0,
	0x7e, 0x18, 0x71, 0x8d, 0x71, 0xb2, 0x02, 0x25, 0xfa, 0xc6, 0xbb, 0xfd, 0xa0, 0x8f, 0x9c, 0x9c,
	0xea, 0x03, 0x56, 0xe1, 0x6a, 0xff, 0x01, 0x77, 0xa9, 0x1f, 0x1c, 0x1b, 0x32, 0xce, 0x77, 0x09,
	0x4a, 0x11, 0xde, 0x07, 0xa9, 0xf4, 0x41, 0x60, 0x37, 0x5a, 0x47, 0x9d, 0x62, 0x88, 0x69, 0x98,
	0xb6, 0xef, 0x98, 0x41, 0xdd, 0x94, 0x8e, 0xd3, 0xbd, 0xd4, 0xd9, 0xd5, 0x6d, 0xfa, 0x07, 0x5b,
	0xce, 0xfb, 0xbd, 0x0d, 0x1c, 0x45, 0xe3, 0x8b, 0xa8, 0xf6, 0x5c, 0x2b, 0x7c, 0x83, 0x9f, 0x4d,
	0x3b, 0x4a, 0xb6, 0xb3, 0x4e, 0x5c, 0x5c, 0xc5, 0xb7, 0x17, 0x51, 0xe6, 0x21, 0x7e, 0x18, 0xc3,
	0xcf, 0xa6, 0x71, 0x31, 0x8d, 0x92, 0x0f, 0x1a, 0xc9, 0xb8, 0x16, 0xdf, 0xd0, 0x92, 0xeb, 0x38,
	0x36, 0xde, 0xaa, 0xa7, 0x50, 0xfa, 0x49, 0xa3, 0xb2, 0x48, 0x40, 0x3c, 0x5f, 0x41, 0xd9, 0x25,
	0xfc, 0xe8, 0x21, 0xaa, 0x56, 0x51, 0xb6, 0x82, 0x92, 0xcb, 0xad, 0x7a, 0x41, 0xdd, 0x9a, 0x52,
	0xab, 0x69, 0x94, 0x9a, 0xc4, 0xd3, 0xab, 0xea, 0xf6, 0xdd, 0xc6, 0xb3, 0x39, 0x2d, 0x3f, 0x81,
	0x26, 0xd7, 0xd5, 0xcd, 0x39, 0x52, 0x90, 0x94, 0x02, 0x67, 0x9d, 0x68, 0xeb, 0x5b, 0x7d, 0x47,
	0x03, 0x1c, 0xba, 0x7c, 0xf9, 0x22, 0x19, 0x91, 0x42, 0x38, 0x91, 0x45, 0xc9, 0x12, 0xca, 0xde,
	0x76, 0xb6, 0xea, 0xa9, 0xc6, 0x8b, 0x19, 0x54, 0xb8, 0x47, 0xf2, 0xb4, 0xfb, 0xe3, 0xee, 0x73,
	0xe7, 0x5e, 0xc6, 0xe2, 0xa4, 0x00, 0xba, 0x95, 0xee, 0xe0, 0x68, 0x63, 0x5d, 0xad, 0xa6, 0xdb,
	0xdc, 0xcd, 0xc2, 0xf8, 0x56, 0xac, 0x37, 0x8b, 0xfc, 0x44, 0xe3, 0x45, 0xb6, 0xf1, 0x20, 0xb5,
	0x43, 0x67, 0x29, 0x8e, 0xef, 0x15, 0x3b, 0x89, 0x64, 0x15, 0x15, 0x97, 0xd1, 0x42, 0x5e, 0xad,
	0xae, 0x90, 0x13, 0xab, 0x1b, 0xd3, 0x6a, 0x35, 0xd9, 0xaa, 0xa7, 0xf0, 0xa3, 0x05, 0x22, 0x95,
	0xc9, 0x2d, 0x3f, 0x61, 0xaa, 0xd9, 0x21, 0x93, 0x59, 0x69, 0xe6, 0x12, 0xa8, 0xb2, 0x41, 0x48,
	0x36, 0xca, 0x2f, 0x9a, 0x77, 0xca, 0x5a, 0x7e, 0xa2, 0x5b, 0x54, 0x5d, 0x4e, 0x83, 0x33, 0xca,
	0xa4, 0x51, 0x6d, 0x06, 0xcf, 0x3e, 0x55, 0xab, 0x2b, 0xcd, 0xc2, 0x94, 0xb6, 0xbd, 0xb0, 0xa7,
	0xe0, 0xcb, 0x58, 0xdc, 0x7d, 0xc0, 0xd3, 0xcf, 0x89, 0x62, 0x90, 0xf7, 0x19, 0x77, 0x99, 0xfd,
	0x42, 0x16, 0x42, 0xd4, 0x4f, 0x36, 0x60, 0x37, 0xac, 0xff, 0xd7, 0xfc, 0xb0, 0xeb, 0xb6, 0xd0,
	0x09, 0x5b, 0xce, 0x3b, 0xb1, 0xd7, 0x0f, 0xbf, 0x5a, 0x7b, 0xf8, 0x61, 0x75, 0x3f, 0x4b, 0x74,
	0xe3, 0xc4, 0x15, 0x28, 0x53, 0xd1, 0x66, 0x96, 0xcd, 0xd5, 0x36, 0x62, 0x5a, 0xa5, 0xb2, 0x81,
	0x16, 0xe6, 0x70, 0x26, 0xab, 0x6e, 0x17, 0x48, 0x62, 0xab, 0x5e, 0x20, 0x86, 0xc1, 0xc5, 0x55,
	0x52, 0x64, 0xc7, 0x0f, 0xdd, 0x5d, 0x42, 0x99, 0xe4, 0x7f, 0xbb, 0x9f, 0xdb, 0x56, 0x70, 0xc8,
	0x7c, 0x52, 0xf6, 0xed, 0xe8, 0x3b, 0xbd, 0x3a, 0xba, 0xe7, 0x1d, 0xa2, 0xff, 0xb0, 0xe4, 0xbc,
	0xbf, 0x5b, 0xf6, 0xf4, 0xf4, 0x47, 0x4b, 0x57, 0x4f, 0xc9, 0x41, 0xbb, 0xe9, 0x37, 0xb6, 0x2a,
	0xea, 0xe6, 0xba, 0x76, 0x2b, 0x81, 0x8b, 0xab, 0xff, 0x7e, 0xfd, 0x40, 0x4f, 0x01, 0x7f, 0xb3,
	0x83, 0x13, 0x9f, 0x42, 0x45, 0xe2, 0x7d, 0xf2, 0x45, 0x49, 0x18, 0x83, 0xca, 0x28, 0x0c, 0xcb,
	0xfb, 0x4a, 0x79, 0xb2, 0x8d, 0xeb, 0x62, 0x0e, 0x29, 0x8a, 0x78, 0x41, 0xf0, 0x47, 0xe9, 0x65,
	0x7b, 0xce, 0xbb, 0x60, 0x07, 0xc7, 0xd1, 0x76, 0xd9, 0xf4, 0x56, 0xf9, 0x81, 0x96, 0x9d, 0x72,
	0x54, 0x0f, 0xfe, 0x7f, 0x25, 0xfe, 0xd9, 0x96, 0x9e, 0x75, 0xe2, 0x99, 0x0a, 0x4e, 0x8d, 0x8b,
	0xa4, 0x75, 0x61, 0xd9, 0xe9, 0xbd, 0xf8, 0x11, 0x21, 0xd1, 0xaa, 0xa7, 0x3a, 0x28, 0x9e, 0xaf,
	0xa8, 0xd5, 0x98, 0xf6, 0x5d, 0xba, 0x59, 0x98, 0x52, 0xab, 0x8f, 0x50, 0x3d, 0x83, 0x6a, 0x4b,
	0xcd, 0x9b, 0x19, 0xed, 0xc9, 0xa6, 0xb6, 0x59, 0x6a, 0xd5, 0x53, 0xda, 0xcc, 0xb2, 0x5a, 0xbb,
	0xdd, 0x28, 0xaf, 0xa1, 0xad, 0x59, 0x94, 0xa9, 0x34, 0xc6, 0xf3, 0x28, 0x9b, 0xd0, 0x99, 0xcd,
	0xae, 0xe1, 0x74, 0x59, 0x97, 0x76, 0x29, 0xde, 0x9c, 0x2f, 0xa1, 0xc4, 0x1d, 0x9d, 0x84, 0xc1,
	0x1e, 0x3f, 0x7f, 0x8c, 0x6a, 0x4b, 0xad, 0x7a, 0xaa, 0x99, 0xcf, 0xa0, 0xc9, 0xaf, 0xb5, 0xa7,
	0xf7, 0xd1, 0xe4, 0xa2, 0x1e, 0x73, 0xb3, 0xd6, 0x9c, 0x2f, 0x35, 0x62, 0x29, 0x5c, 0x78, 0x8a,
	0xe7, 0xd6, 0x5a, 0xf5, 0x82, 0xde, 0x07, 0x5d, 0xda, 0xca, 0x73, 0xfc, 0x38, 0xae, 0x56, 0x57,
	0xf0, 0xec, 0x1a, 0x49, 0xd3, 0x17, 0x08, 0x8a, 0x12, 0x6b, 0xf8, 0xee, 0x62, 0x27, 0x09, 0xd0,
	0x43, 0x00, 0x88, 0x3b, 0x76, 0xf4, 0xe8, 0x3f, 0x79, 0xb2, 0x87, 0x65, 0x77, 0x30, 0x86, 0x17,
	0x58, 0xbf, 0xe0, 0x93, 0xbb, 0x20, 0x36, 0xc8, 0x29, 0x50, 0x56, 0xd8, 0x2f, 0xc3, 0x50, 0x8a,
	0xf2, 0xa1, 0x80, 0xf1, 0x7b, 0xe8, 0x3e, 0xe0, 0x01, 0x0a, 0xfc, 0x4a, 0x61, 0xc5, 0x20, 0xc7,
	0x87, 0x2e, 0xb0, 0x9f, 0xbb, 0xba, 0x7e, 0x10, 0x65, 0xc8, 0x49, 0x1f, 0xeb, 0x91, 0xfa, 0x03,
	0x61, 0x84, 0x47, 0xce, 0xb3, 0xe4, 0x81, 0xf8, 0x90, 0x7c, 0x46, 0xfa, 0x8c, 0x1b, 0xf0, 0xde,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xec, 0xee, 0x78, 0xb0, 0xf5, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthServiceClient interface {
	// 节点启动状态检测
	Alive(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AliveResponse, error)
	// 节点就绪状态监测
	Ready(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadyResponse, error)
	// 服务版本查询
	Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	// Prometheus监控
	MetricsPrometheus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) Alive(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AliveResponse, error) {
	out := new(AliveResponse)
	err := c.cc.Invoke(ctx, "/sole.api.v1.health.HealthService/Alive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) Ready(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadyResponse, error) {
	out := new(ReadyResponse)
	err := c.cc.Invoke(ctx, "/sole.api.v1.health.HealthService/Ready", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/sole.api.v1.health.HealthService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) MetricsPrometheus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/sole.api.v1.health.HealthService/MetricsPrometheus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
type HealthServiceServer interface {
	// 节点启动状态检测
	Alive(context.Context, *empty.Empty) (*AliveResponse, error)
	// 节点就绪状态监测
	Ready(context.Context, *empty.Empty) (*ReadyResponse, error)
	// 服务版本查询
	Version(context.Context, *empty.Empty) (*VersionResponse, error)
	// Prometheus监控
	MetricsPrometheus(context.Context, *empty.Empty) (*httpbody.HttpBody, error)
}

// UnimplementedHealthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServiceServer struct {
}

func (*UnimplementedHealthServiceServer) Alive(ctx context.Context, req *empty.Empty) (*AliveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alive not implemented")
}
func (*UnimplementedHealthServiceServer) Ready(ctx context.Context, req *empty.Empty) (*ReadyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ready not implemented")
}
func (*UnimplementedHealthServiceServer) Version(ctx context.Context, req *empty.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (*UnimplementedHealthServiceServer) MetricsPrometheus(ctx context.Context, req *empty.Empty) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MetricsPrometheus not implemented")
}

func RegisterHealthServiceServer(s *grpc.Server, srv HealthServiceServer) {
	s.RegisterService(&_HealthService_serviceDesc, srv)
}

func _HealthService_Alive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Alive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sole.api.v1.health.HealthService/Alive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Alive(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sole.api.v1.health.HealthService/Ready",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Ready(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sole.api.v1.health.HealthService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Version(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_MetricsPrometheus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).MetricsPrometheus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sole.api.v1.health.HealthService/MetricsPrometheus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).MetricsPrometheus(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sole.api.v1.health.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Alive",
			Handler:    _HealthService_Alive_Handler,
		},
		{
			MethodName: "Ready",
			Handler:    _HealthService_Ready_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _HealthService_Version_Handler,
		},
		{
			MethodName: "MetricsPrometheus",
			Handler:    _HealthService_MetricsPrometheus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health.proto",
}
