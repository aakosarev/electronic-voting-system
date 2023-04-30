// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: votingVerifier.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VotingVerifierClient is the client API for VotingVerifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VotingVerifierClient interface {
	GetPublicKeyForVotingID(ctx context.Context, in *GetPublicKeyForVotingIDRequest, opts ...grpc.CallOption) (*GetPublicKeyForVotingIDResponse, error)
	SignBlindedToken(ctx context.Context, in *SignBlindedTokenRequest, opts ...grpc.CallOption) (*SignBlindedTokenResponse, error)
}

type votingVerifierClient struct {
	cc grpc.ClientConnInterface
}

func NewVotingVerifierClient(cc grpc.ClientConnInterface) VotingVerifierClient {
	return &votingVerifierClient{cc}
}

func (c *votingVerifierClient) GetPublicKeyForVotingID(ctx context.Context, in *GetPublicKeyForVotingIDRequest, opts ...grpc.CallOption) (*GetPublicKeyForVotingIDResponse, error) {
	out := new(GetPublicKeyForVotingIDResponse)
	err := c.cc.Invoke(ctx, "/VotingVerifier/GetPublicKeyForVotingID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *votingVerifierClient) SignBlindedToken(ctx context.Context, in *SignBlindedTokenRequest, opts ...grpc.CallOption) (*SignBlindedTokenResponse, error) {
	out := new(SignBlindedTokenResponse)
	err := c.cc.Invoke(ctx, "/VotingVerifier/SignBlindedToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VotingVerifierServer is the server API for VotingVerifier service.
// All implementations must embed UnimplementedVotingVerifierServer
// for forward compatibility
type VotingVerifierServer interface {
	GetPublicKeyForVotingID(context.Context, *GetPublicKeyForVotingIDRequest) (*GetPublicKeyForVotingIDResponse, error)
	SignBlindedToken(context.Context, *SignBlindedTokenRequest) (*SignBlindedTokenResponse, error)
	mustEmbedUnimplementedVotingVerifierServer()
}

// UnimplementedVotingVerifierServer must be embedded to have forward compatible implementations.
type UnimplementedVotingVerifierServer struct {
}

func (UnimplementedVotingVerifierServer) GetPublicKeyForVotingID(context.Context, *GetPublicKeyForVotingIDRequest) (*GetPublicKeyForVotingIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKeyForVotingID not implemented")
}
func (UnimplementedVotingVerifierServer) SignBlindedToken(context.Context, *SignBlindedTokenRequest) (*SignBlindedTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignBlindedToken not implemented")
}
func (UnimplementedVotingVerifierServer) mustEmbedUnimplementedVotingVerifierServer() {}

// UnsafeVotingVerifierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VotingVerifierServer will
// result in compilation errors.
type UnsafeVotingVerifierServer interface {
	mustEmbedUnimplementedVotingVerifierServer()
}

func RegisterVotingVerifierServer(s grpc.ServiceRegistrar, srv VotingVerifierServer) {
	s.RegisterService(&VotingVerifier_ServiceDesc, srv)
}

func _VotingVerifier_GetPublicKeyForVotingID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicKeyForVotingIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VotingVerifierServer).GetPublicKeyForVotingID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VotingVerifier/GetPublicKeyForVotingID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VotingVerifierServer).GetPublicKeyForVotingID(ctx, req.(*GetPublicKeyForVotingIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VotingVerifier_SignBlindedToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignBlindedTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VotingVerifierServer).SignBlindedToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VotingVerifier/SignBlindedToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VotingVerifierServer).SignBlindedToken(ctx, req.(*SignBlindedTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VotingVerifier_ServiceDesc is the grpc.ServiceDesc for VotingVerifier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VotingVerifier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VotingVerifier",
	HandlerType: (*VotingVerifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublicKeyForVotingID",
			Handler:    _VotingVerifier_GetPublicKeyForVotingID_Handler,
		},
		{
			MethodName: "SignBlindedToken",
			Handler:    _VotingVerifier_SignBlindedToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "votingVerifier.proto",
}