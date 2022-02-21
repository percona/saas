// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: telemetry/reporter/reporter_api.proto

package reporterv1

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

// ReporterAPIClient is the client API for ReporterAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReporterAPIClient interface {
	// Report submits several telemetry events to the server.
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
}

type reporterAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewReporterAPIClient(cc grpc.ClientConnInterface) ReporterAPIClient {
	return &reporterAPIClient{cc}
}

func (c *reporterAPIClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.telemetry.reporter.v1.ReporterAPI/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReporterAPIServer is the server API for ReporterAPI service.
// All implementations must embed UnimplementedReporterAPIServer
// for forward compatibility
type ReporterAPIServer interface {
	// Report submits several telemetry events to the server.
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
	mustEmbedUnimplementedReporterAPIServer()
}

// UnimplementedReporterAPIServer must be embedded to have forward compatible implementations.
type UnimplementedReporterAPIServer struct {
}

func (UnimplementedReporterAPIServer) Report(context.Context, *ReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}
func (UnimplementedReporterAPIServer) mustEmbedUnimplementedReporterAPIServer() {}

// UnsafeReporterAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReporterAPIServer will
// result in compilation errors.
type UnsafeReporterAPIServer interface {
	mustEmbedUnimplementedReporterAPIServer()
}

func RegisterReporterAPIServer(s grpc.ServiceRegistrar, srv ReporterAPIServer) {
	s.RegisterService(&ReporterAPI_ServiceDesc, srv)
}

func _ReporterAPI_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReporterAPIServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.telemetry.reporter.v1.ReporterAPI/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReporterAPIServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReporterAPI_ServiceDesc is the grpc.ServiceDesc for ReporterAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReporterAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "percona.platform.telemetry.reporter.v1.ReporterAPI",
	HandlerType: (*ReporterAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _ReporterAPI_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "telemetry/reporter/reporter_api.proto",
}
