// Copyright 2020 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: webapp.proto

// 主页HomePage API

package webapp

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

func init() {
	proto.RegisterFile("webapp.proto", fileDescriptor_9abfc33d2d8246e3)
}

var fileDescriptor_9abfc33d2d8246e3 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x95, 0x16, 0xc1, 0xc5, 0xa2, 0x04, 0xa9, 0x12, 0x2f, 0xc1, 0x6b, 0xc9, 0x0e, 0xad,
	0x17, 0x49, 0xf1, 0xd0, 0x82, 0x52, 0x50, 0x4f, 0x15, 0x04, 0x6f, 0xb3, 0x9b, 0xc9, 0x64, 0x4a,
	0x92, 0x19, 0x76, 0xa6, 0x49, 0x73, 0x6b, 0x43, 0xf1, 0x47, 0x6c, 0x4c, 0xa9, 0x51, 0xb1, 0x4d,
	0x2b, 0xd1, 0x2a, 0xa5, 0x68, 0x8b, 0x60, 0x4b, 0x52, 0x6b, 0xf3, 0xc7, 0xb8, 0x33, 0xd9, 0xc5,
	0x43, 0xff, 0x05, 0xc9, 0x4e, 0xa2, 0x6b, 0x8b, 0xa7, 0x7d, 0xfb, 0xe6, 0xf3, 0xbe, 0x6f, 0xde,
	0xee, 0xf7, 0x19, 0x17, 0x0a, 0xc8, 0x82, 0x8c, 0x99, 0xcc, 0xa1, 0x82, 0x46, 0x22, 0x9c, 0x66,
	0x90, 0x09, 0x19, 0x31, 0xf3, 0xa3, 0xa6, 0x3e, 0x89, 0x5e, 0xc3, 0x94, 0xe2, 0x0c, 0x02, 0x01,
	0x61, 0xcd, 0xa6, 0x00, 0xca, 0x32, 0x51, 0xd4, 0x05, 0xd1, 0xd8, 0xc9, 0xc3, 0x24, 0xe2, 0xb6,
	0x43, 0x98, 0xa0, 0x4e, 0x8f, 0x48, 0x60, 0x22, 0xd2, 0xb3, 0x96, 0x69, 0xd3, 0x2c, 0xd0, 0x30,
	0x64, 0x84, 0x9f, 0x0e, 0x01, 0x64, 0x04, 0xa4, 0x85, 0x60, 0x16, 0x4d, 0xf6, 0xd5, 0xa7, 0xc3,
	0xb5, 0x0e, 0xb3, 0xe3, 0xc8, 0xa6, 0xbc, 0xc8, 0x05, 0xea, 0xbd, 0x62, 0x28, 0x50, 0x01, 0x16,
	0x75, 0x77, 0x3b, 0x8e, 0x51, 0x2e, 0xce, 0x0b, 0x10, 0x63, 0xe4, 0x00, 0xca, 0x04, 0xa1, 0x39,
	0x0e, 0x60, 0x2e, 0x47, 0x05, 0x0c, 0x62, 0x2d, 0x3a, 0xf6, 0x75, 0xd0, 0x18, 0x7a, 0x88, 0xac,
	0x09, 0xc6, 0xa6, 0x91, 0x93, 0x27, 0x36, 0x8a, 0xbc, 0x1b, 0x34, 0x06, 0xef, 0x90, 0x0c, 0x8a,
	0x0c, 0x9b, 0xfa, 0x2e, 0x66, 0x7f, 0x1c, 0xf3, 0x76, 0x77, 0xd6, 0xe8, 0xe5, 0x7e, 0xbe, 0xfb,
	0x65, 0xa6, 0x84, 0x60, 0x93, 0x34, 0x59, 0xbc, 0xee, 0x0d, 0xac, 0x4c, 0xb8, 0x03, 0x86, 0x31,
	0xf5, 0xe0, 0xfe, 0x3d, 0x79, 0x50, 0xeb, 0xd4, 0xb6, 0xa2, 0xbf, 0xce, 0x8e, 0xc4, 0x54, 0xfd,
	0xb3, 0x7a, 0xf1, 0x51, 0x56, 0x3f, 0x78, 0xed, 0x9a, 0x5c, 0xdb, 0x50, 0xf5, 0x8a, 0x2c, 0x37,
	0xbc, 0x72, 0xa9, 0x53, 0x6a, 0x79, 0x7b, 0x8b, 0xea, 0x60, 0xd9, 0x6d, 0x37, 0xd4, 0xc2, 0xae,
	0x31, 0x12, 0x93, 0x9f, 0x4a, 0x6a, 0xa3, 0xde, 0xe5, 0x35, 0xd3, 0x5e, 0xf6, 0x1a, 0x4b, 0xb2,
	0xbe, 0x25, 0x37, 0x57, 0xdd, 0xe6, 0xb6, 0x2e, 0x71, 0x5b, 0xcf, 0xdd, 0x66, 0xf9, 0xf8, 0x70,
	0x49, 0x7d, 0xd9, 0x0c, 0xcb, 0x76, 0x56, 0x17, 0xb5, 0x9a, 0xdf, 0xd8, 0xf3, 0xd7, 0xdf, 0xcb,
	0xea, 0xb6, 0xbf, 0xf2, 0x54, 0xee, 0xb6, 0xe4, 0x8f, 0x97, 0xf2, 0x59, 0xc5, 0xdb, 0x69, 0xfb,
	0x6f, 0x76, 0x3a, 0xab, 0x8b, 0xe1, 0xee, 0xc7, 0x87, 0x4b, 0xb2, 0xb5, 0xef, 0x36, 0x2b, 0xb2,
	0x5a, 0x91, 0x07, 0x35, 0xf5, 0xea, 0x9b, 0xdb, 0xdc, 0xf6, 0xd7, 0x1e, 0x77, 0x8e, 0x36, 0xff,
	0xa8, 0x75, 0xca, 0xfb, 0x6a, 0x7e, 0xe1, 0xe7, 0x7c, 0xc9, 0xf8, 0x3b, 0x87, 0xaa, 0x2e, 0xbb,
	0x47, 0x6b, 0xfe, 0xfa, 0x5b, 0x35, 0xbf, 0xa0, 0x5e, 0x3f, 0x71, 0xbf, 0xef, 0x6b, 0xd5, 0xb1,
	0x33, 0x89, 0x2b, 0x02, 0xcd, 0x09, 0x60, 0x73, 0x3e, 0x1e, 0xb3, 0xd3, 0xd0, 0xe1, 0x48, 0xdc,
	0x9a, 0x15, 0xa9, 0xf8, 0xcd, 0xc4, 0x79, 0x92, 0x85, 0x18, 0x01, 0x4c, 0x52, 0x89, 0xab, 0x01,
	0x93, 0x16, 0xd9, 0xcc, 0x49, 0xc8, 0xd0, 0xd0, 0x0c, 0x43, 0xf8, 0x9f, 0x78, 0x18, 0x32, 0x96,
	0x21, 0x76, 0xf0, 0xd7, 0xc0, 0x0c, 0xcc, 0x43, 0x6d, 0xa9, 0xff, 0xe6, 0x2f, 0x86, 0xf3, 0x2c,
	0x99, 0xea, 0x77, 0x67, 0x39, 0x9c, 0x18, 0xd2, 0x21, 0xcf, 0xe3, 0x91, 0xb9, 0x6c, 0x26, 0x71,
	0x29, 0x8c, 0x16, 0x20, 0xcf, 0xf6, 0x1b, 0x17, 0x90, 0xc5, 0x7a, 0xe3, 0xcc, 0x9d, 0xba, 0xe9,
	0x24, 0x78, 0x14, 0x0f, 0x59, 0x91, 0x23, 0xe8, 0xdc, 0x25, 0x39, 0x0c, 0xba, 0xdb, 0x12, 0xf8,
	0x36, 0x3f, 0x0a, 0xf4, 0xb6, 0x8c, 0xeb, 0x87, 0x75, 0x2e, 0x70, 0xd0, 0x8d, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x21, 0xcc, 0x94, 0x91, 0x5f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WebAppServiceClient is the client API for WebAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WebAppServiceClient interface {
	// 静态文件下载
	File(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type webAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebAppServiceClient(cc grpc.ClientConnInterface) WebAppServiceClient {
	return &webAppServiceClient{cc}
}

func (c *webAppServiceClient) File(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/sole.api.v1.webapp.WebAppService/File", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebAppServiceServer is the server API for WebAppService service.
type WebAppServiceServer interface {
	// 静态文件下载
	File(context.Context, *empty.Empty) (*httpbody.HttpBody, error)
}

// UnimplementedWebAppServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWebAppServiceServer struct {
}

func (*UnimplementedWebAppServiceServer) File(ctx context.Context, req *empty.Empty) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method File not implemented")
}

func RegisterWebAppServiceServer(s *grpc.Server, srv WebAppServiceServer) {
	s.RegisterService(&_WebAppService_serviceDesc, srv)
}

func _WebAppService_File_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebAppServiceServer).File(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sole.api.v1.webapp.WebAppService/File",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebAppServiceServer).File(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _WebAppService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sole.api.v1.webapp.WebAppService",
	HandlerType: (*WebAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "File",
			Handler:    _WebAppService_File_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webapp.proto",
}
