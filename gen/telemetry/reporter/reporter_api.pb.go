// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: telemetry/reporter/reporter_api.proto

package reporterv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "github.com/percona-platform/saas/gen/validator"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One or more events to report.
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_telemetry_reporter_reporter_api_proto_rawDescGZIP(), []int{0}
}

func (x *ReportRequest) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type ReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportResponse) Reset() {
	*x = ReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResponse) ProtoMessage() {}

func (x *ReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResponse.ProtoReflect.Descriptor instead.
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return file_telemetry_reporter_reporter_api_proto_rawDescGZIP(), []int{1}
}

var File_telemetry_reporter_reporter_api_proto protoreflect.FileDescriptor

var file_telemetry_reporter_reporter_api_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x19, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x65, 0x72, 0x63,
	0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x60, 0x01,
	0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xbc, 0x02, 0x0a, 0x0b, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0xac, 0x02, 0x0a, 0x06, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x35, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x70,
	0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb2, 0x01, 0x92, 0x41, 0x8f, 0x01, 0x12, 0x06, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x1a, 0x40, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x69, 0x73, 0x20, 0x75,
	0x73, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x20, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x74,
	0x6f, 0x20, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x64, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x4a, 0x1f, 0x0a, 0x03, 0x32, 0x30, 0x30,
	0x12, 0x18, 0x22, 0x16, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x02, 0x7b, 0x7d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19,
	0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x9f, 0x03, 0x0a, 0x2a, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61,
	0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x50, 0x50, 0x54, 0x52, 0xaa, 0x02, 0x26, 0x50, 0x65,
	0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x26, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x5c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x5c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x32,
	0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x5c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5c, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x2a, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x3a, 0x3a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x92,
	0x41, 0x59, 0x12, 0x0e, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x20, 0x41,
	0x50, 0x49, 0x7a, 0x47, 0x0a, 0x08, 0x78, 0x2d, 0x72, 0x65, 0x61, 0x64, 0x6d, 0x65, 0x12, 0x3b,
	0x2a, 0x39, 0x0a, 0x37, 0x0a, 0x11, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2d, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x22, 0x32, 0x20, 0x0a, 0x06, 0x1a, 0x04, 0x63,
	0x75, 0x72, 0x6c, 0x0a, 0x04, 0x1a, 0x02, 0x67, 0x6f, 0x0a, 0x06, 0x1a, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x0a, 0x08, 0x1a, 0x06, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_telemetry_reporter_reporter_api_proto_rawDescOnce sync.Once
	file_telemetry_reporter_reporter_api_proto_rawDescData = file_telemetry_reporter_reporter_api_proto_rawDesc
)

func file_telemetry_reporter_reporter_api_proto_rawDescGZIP() []byte {
	file_telemetry_reporter_reporter_api_proto_rawDescOnce.Do(func() {
		file_telemetry_reporter_reporter_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_telemetry_reporter_reporter_api_proto_rawDescData)
	})
	return file_telemetry_reporter_reporter_api_proto_rawDescData
}

var file_telemetry_reporter_reporter_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_telemetry_reporter_reporter_api_proto_goTypes = []interface{}{
	(*ReportRequest)(nil),  // 0: percona.platform.telemetry.reporter.v1.ReportRequest
	(*ReportResponse)(nil), // 1: percona.platform.telemetry.reporter.v1.ReportResponse
	(*Event)(nil),          // 2: percona.platform.telemetry.reporter.v1.Event
}
var file_telemetry_reporter_reporter_api_proto_depIdxs = []int32{
	2, // 0: percona.platform.telemetry.reporter.v1.ReportRequest.events:type_name -> percona.platform.telemetry.reporter.v1.Event
	0, // 1: percona.platform.telemetry.reporter.v1.ReporterAPI.Report:input_type -> percona.platform.telemetry.reporter.v1.ReportRequest
	1, // 2: percona.platform.telemetry.reporter.v1.ReporterAPI.Report:output_type -> percona.platform.telemetry.reporter.v1.ReportResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_telemetry_reporter_reporter_api_proto_init() }
func file_telemetry_reporter_reporter_api_proto_init() {
	if File_telemetry_reporter_reporter_api_proto != nil {
		return
	}
	file_telemetry_reporter_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_telemetry_reporter_reporter_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRequest); i {
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
		file_telemetry_reporter_reporter_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_telemetry_reporter_reporter_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_telemetry_reporter_reporter_api_proto_goTypes,
		DependencyIndexes: file_telemetry_reporter_reporter_api_proto_depIdxs,
		MessageInfos:      file_telemetry_reporter_reporter_api_proto_msgTypes,
	}.Build()
	File_telemetry_reporter_reporter_api_proto = out.File
	file_telemetry_reporter_reporter_api_proto_rawDesc = nil
	file_telemetry_reporter_reporter_api_proto_goTypes = nil
	file_telemetry_reporter_reporter_api_proto_depIdxs = nil
}
