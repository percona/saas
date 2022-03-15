// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: auth/auth_api.proto

package authv1

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

// AuthAPIClient is the client API for AuthAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthAPIClient interface {
	// SignUp creates new user with given email and password.
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	// SignIn checks user authentication, creates session and returns session ID.
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	// SignOut terminates user session.
	SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error)
	// RefreshSession refreshes session timeout.
	RefreshSession(ctx context.Context, in *RefreshSessionRequest, opts ...grpc.CallOption) (*RefreshSessionResponse, error)
	// ResetPassword initiates user's password reset procedure.
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	// GetProfile returns user's email, first name and last name.
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	// UpdateProfile updates user's first name and last name.
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error)
}

type authAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthAPIClient(cc grpc.ClientConnInterface) AuthAPIClient {
	return &authAPIClient{cc}
}

func (c *authAPIClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.v1.AuthAPI/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.v1.AuthAPI/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error) {
	out := new(SignOutResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.v1.AuthAPI/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) RefreshSession(ctx context.Context, in *RefreshSessionRequest, opts ...grpc.CallOption) (*RefreshSessionResponse, error) {
	out := new(RefreshSessionResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.v1.AuthAPI/RefreshSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.v1.AuthAPI/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.v1.AuthAPI/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error) {
	out := new(UpdateProfileResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.auth.v1.AuthAPI/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthAPIServer is the server API for AuthAPI service.
// All implementations must embed UnimplementedAuthAPIServer
// for forward compatibility
type AuthAPIServer interface {
	// SignUp creates new user with given email and password.
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	// SignIn checks user authentication, creates session and returns session ID.
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	// SignOut terminates user session.
	SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error)
	// RefreshSession refreshes session timeout.
	RefreshSession(context.Context, *RefreshSessionRequest) (*RefreshSessionResponse, error)
	// ResetPassword initiates user's password reset procedure.
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	// GetProfile returns user's email, first name and last name.
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	// UpdateProfile updates user's first name and last name.
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error)
	mustEmbedUnimplementedAuthAPIServer()
}

// UnimplementedAuthAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAuthAPIServer struct {
}

func (UnimplementedAuthAPIServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAuthAPIServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthAPIServer) SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (UnimplementedAuthAPIServer) RefreshSession(context.Context, *RefreshSessionRequest) (*RefreshSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshSession not implemented")
}
func (UnimplementedAuthAPIServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthAPIServer) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedAuthAPIServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedAuthAPIServer) mustEmbedUnimplementedAuthAPIServer() {}

// UnsafeAuthAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthAPIServer will
// result in compilation errors.
type UnsafeAuthAPIServer interface {
	mustEmbedUnimplementedAuthAPIServer()
}

func RegisterAuthAPIServer(s grpc.ServiceRegistrar, srv AuthAPIServer) {
	s.RegisterService(&AuthAPI_ServiceDesc, srv)
}

func _AuthAPI_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.v1.AuthAPI/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.v1.AuthAPI/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.v1.AuthAPI/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).SignOut(ctx, req.(*SignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_RefreshSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).RefreshSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.v1.AuthAPI/RefreshSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).RefreshSession(ctx, req.(*RefreshSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.v1.AuthAPI/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.v1.AuthAPI/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.auth.v1.AuthAPI/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthAPI_ServiceDesc is the grpc.ServiceDesc for AuthAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "percona.platform.auth.v1.AuthAPI",
	HandlerType: (*AuthAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _AuthAPI_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _AuthAPI_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _AuthAPI_SignOut_Handler,
		},
		{
			MethodName: "RefreshSession",
			Handler:    _AuthAPI_RefreshSession_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _AuthAPI_ResetPassword_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _AuthAPI_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _AuthAPI_UpdateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth_api.proto",
}
