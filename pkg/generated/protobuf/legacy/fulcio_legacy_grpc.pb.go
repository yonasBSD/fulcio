//
// Copyright 2022 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.1
// source: fulcio_legacy.proto

package legacy

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CA_CreateSigningCertificate_FullMethodName = "/dev.sigstore.fulcio.v1beta.CA/CreateSigningCertificate"
	CA_GetRootCertificate_FullMethodName       = "/dev.sigstore.fulcio.v1beta.CA/GetRootCertificate"
)

// CAClient is the client API for CA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// This implements the pre-GA HTTP-based Fulcio API.
// This interface is deprecated and will only receive backports of security-related features - clients should prefer the GA GRPC interface!
type CAClient interface {
	// Deprecated: Do not use.
	//
	// Returns an X509 certificate created by the Fulcio certificate authority for the given request parameters
	CreateSigningCertificate(ctx context.Context, in *CreateSigningCertificateRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// Deprecated: Do not use.
	//
	// Returns the public key that can be used to validate the signed tree head
	GetRootCertificate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type cAClient struct {
	cc grpc.ClientConnInterface
}

func NewCAClient(cc grpc.ClientConnInterface) CAClient {
	return &cAClient{cc}
}

// Deprecated: Do not use.
func (c *cAClient) CreateSigningCertificate(ctx context.Context, in *CreateSigningCertificateRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, CA_CreateSigningCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *cAClient) GetRootCertificate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, CA_GetRootCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CAServer is the server API for CA service.
// All implementations must embed UnimplementedCAServer
// for forward compatibility.
//
// This implements the pre-GA HTTP-based Fulcio API.
// This interface is deprecated and will only receive backports of security-related features - clients should prefer the GA GRPC interface!
type CAServer interface {
	// Deprecated: Do not use.
	//
	// Returns an X509 certificate created by the Fulcio certificate authority for the given request parameters
	CreateSigningCertificate(context.Context, *CreateSigningCertificateRequest) (*httpbody.HttpBody, error)
	// Deprecated: Do not use.
	//
	// Returns the public key that can be used to validate the signed tree head
	GetRootCertificate(context.Context, *emptypb.Empty) (*httpbody.HttpBody, error)
	mustEmbedUnimplementedCAServer()
}

// UnimplementedCAServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCAServer struct{}

func (UnimplementedCAServer) CreateSigningCertificate(context.Context, *CreateSigningCertificateRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSigningCertificate not implemented")
}
func (UnimplementedCAServer) GetRootCertificate(context.Context, *emptypb.Empty) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRootCertificate not implemented")
}
func (UnimplementedCAServer) mustEmbedUnimplementedCAServer() {}
func (UnimplementedCAServer) testEmbeddedByValue()            {}

// UnsafeCAServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CAServer will
// result in compilation errors.
type UnsafeCAServer interface {
	mustEmbedUnimplementedCAServer()
}

func RegisterCAServer(s grpc.ServiceRegistrar, srv CAServer) {
	// If the following call pancis, it indicates UnimplementedCAServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CA_ServiceDesc, srv)
}

func _CA_CreateSigningCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSigningCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServer).CreateSigningCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CA_CreateSigningCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServer).CreateSigningCertificate(ctx, req.(*CreateSigningCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CA_GetRootCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServer).GetRootCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CA_GetRootCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServer).GetRootCertificate(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CA_ServiceDesc is the grpc.ServiceDesc for CA service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CA_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.sigstore.fulcio.v1beta.CA",
	HandlerType: (*CAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSigningCertificate",
			Handler:    _CA_CreateSigningCertificate_Handler,
		},
		{
			MethodName: "GetRootCertificate",
			Handler:    _CA_GetRootCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fulcio_legacy.proto",
}
