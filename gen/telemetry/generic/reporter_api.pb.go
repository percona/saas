// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        (unknown)
// source: telemetry/generic/reporter_api.proto

package genericv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ProductFamily represents Percona product family.
type ProductFamily int32

const (
	ProductFamily_PRODUCT_FAMILY_INVALID         ProductFamily = 0
	ProductFamily_PRODUCT_FAMILY_PS              ProductFamily = 1
	ProductFamily_PRODUCT_FAMILY_PXC             ProductFamily = 2
	ProductFamily_PRODUCT_FAMILY_PSMDB           ProductFamily = 3
	ProductFamily_PRODUCT_FAMILY_PBM             ProductFamily = 4
	ProductFamily_PRODUCT_FAMILY_POSTGRESQL      ProductFamily = 5
	ProductFamily_PRODUCT_FAMILY_PMM             ProductFamily = 6
	ProductFamily_PRODUCT_FAMILY_EVEREST         ProductFamily = 7
	ProductFamily_PRODUCT_FAMILY_PERCONA_TOOLKIT ProductFamily = 8
	ProductFamily_PRODUCT_FAMILY_PXB             ProductFamily = 9
)

// Enum value maps for ProductFamily.
var (
	ProductFamily_name = map[int32]string{
		0: "PRODUCT_FAMILY_INVALID",
		1: "PRODUCT_FAMILY_PS",
		2: "PRODUCT_FAMILY_PXC",
		3: "PRODUCT_FAMILY_PSMDB",
		4: "PRODUCT_FAMILY_PBM",
		5: "PRODUCT_FAMILY_POSTGRESQL",
		6: "PRODUCT_FAMILY_PMM",
		7: "PRODUCT_FAMILY_EVEREST",
		8: "PRODUCT_FAMILY_PERCONA_TOOLKIT",
		9: "PRODUCT_FAMILY_PXB",
	}
	ProductFamily_value = map[string]int32{
		"PRODUCT_FAMILY_INVALID":         0,
		"PRODUCT_FAMILY_PS":              1,
		"PRODUCT_FAMILY_PXC":             2,
		"PRODUCT_FAMILY_PSMDB":           3,
		"PRODUCT_FAMILY_PBM":             4,
		"PRODUCT_FAMILY_POSTGRESQL":      5,
		"PRODUCT_FAMILY_PMM":             6,
		"PRODUCT_FAMILY_EVEREST":         7,
		"PRODUCT_FAMILY_PERCONA_TOOLKIT": 8,
		"PRODUCT_FAMILY_PXB":             9,
	}
)

func (x ProductFamily) Enum() *ProductFamily {
	p := new(ProductFamily)
	*p = x
	return p
}

func (x ProductFamily) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductFamily) Descriptor() protoreflect.EnumDescriptor {
	return file_telemetry_generic_reporter_api_proto_enumTypes[0].Descriptor()
}

func (ProductFamily) Type() protoreflect.EnumType {
	return &file_telemetry_generic_reporter_api_proto_enumTypes[0]
}

func (x ProductFamily) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductFamily.Descriptor instead.
func (ProductFamily) EnumDescriptor() ([]byte, []int) {
	return file_telemetry_generic_reporter_api_proto_rawDescGZIP(), []int{0}
}

// GenericReport contains set of metrics and service information.
// The exact metrics pairs (metric name, metric value) is defined by each product.
type GenericReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Report UUID.
	// Required.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Time when this event was generated by the sender.
	// Required.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Product instance unique identifier. UUID.
	// Required.
	InstanceId string `protobuf:"bytes,3,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// Product family.
	// Mandatory.
	ProductFamily ProductFamily `protobuf:"varint,4,opt,name=product_family,json=productFamily,proto3,enum=percona.platform.telemetry.generic.v1.ProductFamily" json:"product_family,omitempty"`
	// Metrics collection.
	// Optional.
	Metrics []*GenericReport_Metric `protobuf:"bytes,7,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *GenericReport) Reset() {
	*x = GenericReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_generic_reporter_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericReport) ProtoMessage() {}

func (x *GenericReport) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_generic_reporter_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericReport.ProtoReflect.Descriptor instead.
func (*GenericReport) Descriptor() ([]byte, []int) {
	return file_telemetry_generic_reporter_api_proto_rawDescGZIP(), []int{0}
}

func (x *GenericReport) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GenericReport) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *GenericReport) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *GenericReport) GetProductFamily() ProductFamily {
	if x != nil {
		return x.ProductFamily
	}
	return ProductFamily_PRODUCT_FAMILY_INVALID
}

func (x *GenericReport) GetMetrics() []*GenericReport_Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One or more metric reports.
	Reports []*GenericReport `protobuf:"bytes,1,rep,name=reports,proto3" json:"reports,omitempty"`
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_generic_reporter_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_generic_reporter_api_proto_msgTypes[1]
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
	return file_telemetry_generic_reporter_api_proto_rawDescGZIP(), []int{1}
}

func (x *ReportRequest) GetReports() []*GenericReport {
	if x != nil {
		return x.Reports
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
		mi := &file_telemetry_generic_reporter_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResponse) ProtoMessage() {}

func (x *ReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_generic_reporter_api_proto_msgTypes[2]
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
	return file_telemetry_generic_reporter_api_proto_rawDescGZIP(), []int{2}
}

type GenericReport_Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GenericReport_Metric) Reset() {
	*x = GenericReport_Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_generic_reporter_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericReport_Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericReport_Metric) ProtoMessage() {}

func (x *GenericReport_Metric) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_generic_reporter_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericReport_Metric.ProtoReflect.Descriptor instead.
func (*GenericReport_Metric) Descriptor() ([]byte, []int) {
	return file_telemetry_generic_reporter_api_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GenericReport_Metric) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GenericReport_Metric) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_telemetry_generic_reporter_api_proto protoreflect.FileDescriptor

var file_telemetry_generic_reporter_api_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f,
	0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x03, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x58, 0x01, 0x90, 0x01, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f,
	0x05, 0x58, 0x01, 0x90, 0x01, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x64, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x70, 0x65, 0x72,
	0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x88, 0x01, 0x01, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x55, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x70, 0x65, 0x72, 0x63,
	0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x1a,
	0x30, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x5f, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x4e, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x9b, 0x02, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43,
	0x54, 0x5f, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x46, 0x41,
	0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x50, 0x53, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x4f,
	0x44, 0x55, 0x43, 0x54, 0x5f, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x50, 0x58, 0x43, 0x10,
	0x02, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x46, 0x41, 0x4d,
	0x49, 0x4c, 0x59, 0x5f, 0x50, 0x53, 0x4d, 0x44, 0x42, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x50,
	0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x50, 0x42,
	0x4d, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x46,
	0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45, 0x53, 0x51, 0x4c,
	0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x46, 0x41,
	0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x50, 0x4d, 0x4d, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x45, 0x56, 0x45,
	0x52, 0x45, 0x53, 0x54, 0x10, 0x07, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43,
	0x54, 0x5f, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x50, 0x45, 0x52, 0x43, 0x4f, 0x4e, 0x41,
	0x5f, 0x54, 0x4f, 0x4f, 0x4c, 0x4b, 0x49, 0x54, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x50, 0x58, 0x42,
	0x10, 0x09, 0x32, 0xdf, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x41,
	0x50, 0x49, 0x12, 0xcf, 0x02, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x34, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x70, 0x65, 0x72,
	0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xd0, 0x01, 0x92, 0x41, 0xa6, 0x01, 0x12, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x50, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x74,
	0x6f, 0x20, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x20, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x64,
	0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x4a, 0x1f, 0x0a,
	0x03, 0x32, 0x30, 0x30, 0x12, 0x18, 0x22, 0x16, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x02, 0x7b, 0x7d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x42, 0xd0, 0x03, 0x92, 0x41, 0x90, 0x01, 0x12, 0x45, 0x0a, 0x1e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x20, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x20, 0x41, 0x50, 0x49, 0x12, 0x23, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x20, 0x41, 0x50, 0x49, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x64, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x7a, 0x47, 0x0a, 0x08, 0x78, 0x2d, 0x72, 0x65, 0x61, 0x64, 0x6d, 0x65, 0x12, 0x3b,
	0x2a, 0x39, 0x0a, 0x37, 0x0a, 0x11, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2d, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x22, 0x32, 0x20, 0x0a, 0x06, 0x1a, 0x04, 0x63,
	0x75, 0x72, 0x6c, 0x0a, 0x04, 0x1a, 0x02, 0x67, 0x6f, 0x0a, 0x06, 0x1a, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x0a, 0x08, 0x1a, 0x06, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x0a, 0x29, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2d, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x3b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x76, 0x31,
	0xa2, 0x02, 0x04, 0x50, 0x50, 0x54, 0x47, 0xaa, 0x02, 0x25, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e,
	0x61, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x25, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x5c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x5c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5c, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x31, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e,
	0x61, 0x5c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x54, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x5c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x29, 0x50, 0x65,
	0x72, 0x63, 0x6f, 0x6e, 0x61, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a,
	0x3a, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_telemetry_generic_reporter_api_proto_rawDescOnce sync.Once
	file_telemetry_generic_reporter_api_proto_rawDescData = file_telemetry_generic_reporter_api_proto_rawDesc
)

func file_telemetry_generic_reporter_api_proto_rawDescGZIP() []byte {
	file_telemetry_generic_reporter_api_proto_rawDescOnce.Do(func() {
		file_telemetry_generic_reporter_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_telemetry_generic_reporter_api_proto_rawDescData)
	})
	return file_telemetry_generic_reporter_api_proto_rawDescData
}

var file_telemetry_generic_reporter_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_telemetry_generic_reporter_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_telemetry_generic_reporter_api_proto_goTypes = []interface{}{
	(ProductFamily)(0),            // 0: percona.platform.telemetry.generic.v1.ProductFamily
	(*GenericReport)(nil),         // 1: percona.platform.telemetry.generic.v1.GenericReport
	(*ReportRequest)(nil),         // 2: percona.platform.telemetry.generic.v1.ReportRequest
	(*ReportResponse)(nil),        // 3: percona.platform.telemetry.generic.v1.ReportResponse
	(*GenericReport_Metric)(nil),  // 4: percona.platform.telemetry.generic.v1.GenericReport.Metric
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_telemetry_generic_reporter_api_proto_depIdxs = []int32{
	5, // 0: percona.platform.telemetry.generic.v1.GenericReport.create_time:type_name -> google.protobuf.Timestamp
	0, // 1: percona.platform.telemetry.generic.v1.GenericReport.product_family:type_name -> percona.platform.telemetry.generic.v1.ProductFamily
	4, // 2: percona.platform.telemetry.generic.v1.GenericReport.metrics:type_name -> percona.platform.telemetry.generic.v1.GenericReport.Metric
	1, // 3: percona.platform.telemetry.generic.v1.ReportRequest.reports:type_name -> percona.platform.telemetry.generic.v1.GenericReport
	2, // 4: percona.platform.telemetry.generic.v1.ReporterAPI.GenericReport:input_type -> percona.platform.telemetry.generic.v1.ReportRequest
	3, // 5: percona.platform.telemetry.generic.v1.ReporterAPI.GenericReport:output_type -> percona.platform.telemetry.generic.v1.ReportResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_telemetry_generic_reporter_api_proto_init() }
func file_telemetry_generic_reporter_api_proto_init() {
	if File_telemetry_generic_reporter_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_telemetry_generic_reporter_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericReport); i {
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
		file_telemetry_generic_reporter_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_telemetry_generic_reporter_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_telemetry_generic_reporter_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericReport_Metric); i {
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
			RawDescriptor: file_telemetry_generic_reporter_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_telemetry_generic_reporter_api_proto_goTypes,
		DependencyIndexes: file_telemetry_generic_reporter_api_proto_depIdxs,
		EnumInfos:         file_telemetry_generic_reporter_api_proto_enumTypes,
		MessageInfos:      file_telemetry_generic_reporter_api_proto_msgTypes,
	}.Build()
	File_telemetry_generic_reporter_api_proto = out.File
	file_telemetry_generic_reporter_api_proto_rawDesc = nil
	file_telemetry_generic_reporter_api_proto_goTypes = nil
	file_telemetry_generic_reporter_api_proto_depIdxs = nil
}
