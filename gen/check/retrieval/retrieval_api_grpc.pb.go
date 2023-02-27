// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: check/retrieval/retrieval_api.proto

package retrievalv1

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

// RetrievalAPIClient is the client API for RetrievalAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RetrievalAPIClient interface {
	// GetAllChecks returns all available checks.
	GetAllChecks(ctx context.Context, in *GetAllChecksRequest, opts ...grpc.CallOption) (*GetAllChecksResponse, error)
	// GetAllAdvisors returns all available advisors.
	GetAllAdvisors(ctx context.Context, in *GetAllAdvisorsRequest, opts ...grpc.CallOption) (*GetAllAdvisorsResponse, error)
	// GetAllAlertRuleTemplates returns all available alert rule templates.
	GetAllAlertRuleTemplates(ctx context.Context, in *GetAllAlertRuleTemplatesRequest, opts ...grpc.CallOption) (*GetAllAlertRuleTemplatesResponse, error)
}

type retrievalAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewRetrievalAPIClient(cc grpc.ClientConnInterface) RetrievalAPIClient {
	return &retrievalAPIClient{cc}
}

func (c *retrievalAPIClient) GetAllChecks(ctx context.Context, in *GetAllChecksRequest, opts ...grpc.CallOption) (*GetAllChecksResponse, error) {
	out := new(GetAllChecksResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.check.retrieval.v1.RetrievalAPI/GetAllChecks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrievalAPIClient) GetAllAdvisors(ctx context.Context, in *GetAllAdvisorsRequest, opts ...grpc.CallOption) (*GetAllAdvisorsResponse, error) {
	out := new(GetAllAdvisorsResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.check.retrieval.v1.RetrievalAPI/GetAllAdvisors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrievalAPIClient) GetAllAlertRuleTemplates(ctx context.Context, in *GetAllAlertRuleTemplatesRequest, opts ...grpc.CallOption) (*GetAllAlertRuleTemplatesResponse, error) {
	out := new(GetAllAlertRuleTemplatesResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.check.retrieval.v1.RetrievalAPI/GetAllAlertRuleTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RetrievalAPIServer is the server API for RetrievalAPI service.
// All implementations must embed UnimplementedRetrievalAPIServer
// for forward compatibility
type RetrievalAPIServer interface {
	// GetAllChecks returns all available checks.
	GetAllChecks(context.Context, *GetAllChecksRequest) (*GetAllChecksResponse, error)
	// GetAllAdvisors returns all available advisors.
	GetAllAdvisors(context.Context, *GetAllAdvisorsRequest) (*GetAllAdvisorsResponse, error)
	// GetAllAlertRuleTemplates returns all available alert rule templates.
	GetAllAlertRuleTemplates(context.Context, *GetAllAlertRuleTemplatesRequest) (*GetAllAlertRuleTemplatesResponse, error)
	mustEmbedUnimplementedRetrievalAPIServer()
}

// UnimplementedRetrievalAPIServer must be embedded to have forward compatible implementations.
type UnimplementedRetrievalAPIServer struct {
}

func (UnimplementedRetrievalAPIServer) GetAllChecks(context.Context, *GetAllChecksRequest) (*GetAllChecksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllChecks not implemented")
}
func (UnimplementedRetrievalAPIServer) GetAllAdvisors(context.Context, *GetAllAdvisorsRequest) (*GetAllAdvisorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAdvisors not implemented")
}
func (UnimplementedRetrievalAPIServer) GetAllAlertRuleTemplates(context.Context, *GetAllAlertRuleTemplatesRequest) (*GetAllAlertRuleTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAlertRuleTemplates not implemented")
}
func (UnimplementedRetrievalAPIServer) mustEmbedUnimplementedRetrievalAPIServer() {}

// UnsafeRetrievalAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RetrievalAPIServer will
// result in compilation errors.
type UnsafeRetrievalAPIServer interface {
	mustEmbedUnimplementedRetrievalAPIServer()
}

func RegisterRetrievalAPIServer(s grpc.ServiceRegistrar, srv RetrievalAPIServer) {
	s.RegisterService(&RetrievalAPI_ServiceDesc, srv)
}

func _RetrievalAPI_GetAllChecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllChecksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalAPIServer).GetAllChecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.check.retrieval.v1.RetrievalAPI/GetAllChecks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalAPIServer).GetAllChecks(ctx, req.(*GetAllChecksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RetrievalAPI_GetAllAdvisors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAdvisorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalAPIServer).GetAllAdvisors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.check.retrieval.v1.RetrievalAPI/GetAllAdvisors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalAPIServer).GetAllAdvisors(ctx, req.(*GetAllAdvisorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RetrievalAPI_GetAllAlertRuleTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAlertRuleTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalAPIServer).GetAllAlertRuleTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.check.retrieval.v1.RetrievalAPI/GetAllAlertRuleTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalAPIServer).GetAllAlertRuleTemplates(ctx, req.(*GetAllAlertRuleTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RetrievalAPI_ServiceDesc is the grpc.ServiceDesc for RetrievalAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RetrievalAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "percona.platform.check.retrieval.v1.RetrievalAPI",
	HandlerType: (*RetrievalAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllChecks",
			Handler:    _RetrievalAPI_GetAllChecks_Handler,
		},
		{
			MethodName: "GetAllAdvisors",
			Handler:    _RetrievalAPI_GetAllAdvisors_Handler,
		},
		{
			MethodName: "GetAllAlertRuleTemplates",
			Handler:    _RetrievalAPI_GetAllAlertRuleTemplates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "check/retrieval/retrieval_api.proto",
}
