// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: telemetry/events/pmm/server_uptime_event.proto

package pmmv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/mwitkow/go-proto-validators"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DistributionMethod represents PMM Server distribution method.
type DistributionMethod int32

const (
	// DISTRIBUTION_METHOD_INVALID represents unknown distribution method.
	DistributionMethod_DISTRIBUTION_METHOD_INVALID DistributionMethod = 0
	// DOCKER represents Docker distribution method.
	DistributionMethod_DOCKER DistributionMethod = 1
	// OVF represents Open Virtualization Format / Open Virtual Appliance distribution method.
	DistributionMethod_OVF DistributionMethod = 2
	// AMI represents Amazon Machine Image distribution method.
	DistributionMethod_AMI DistributionMethod = 3
	// AZURE represents Azure distribution method.
	DistributionMethod_AZURE DistributionMethod = 4
	// DO represents Digital Ocean distribution method.
	DistributionMethod_DO DistributionMethod = 5
)

// Enum value maps for DistributionMethod.
var (
	DistributionMethod_name = map[int32]string{
		0: "DISTRIBUTION_METHOD_INVALID",
		1: "DOCKER",
		2: "OVF",
		3: "AMI",
		4: "AZURE",
		5: "DO",
	}
	DistributionMethod_value = map[string]int32{
		"DISTRIBUTION_METHOD_INVALID": 0,
		"DOCKER":                      1,
		"OVF":                         2,
		"AMI":                         3,
		"AZURE":                       4,
		"DO":                          5,
	}
)

func (x DistributionMethod) Enum() *DistributionMethod {
	p := new(DistributionMethod)
	*p = x
	return p
}

func (x DistributionMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DistributionMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_telemetry_events_pmm_server_uptime_event_proto_enumTypes[0].Descriptor()
}

func (DistributionMethod) Type() protoreflect.EnumType {
	return &file_telemetry_events_pmm_server_uptime_event_proto_enumTypes[0]
}

func (x DistributionMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DistributionMethod.Descriptor instead.
func (DistributionMethod) EnumDescriptor() ([]byte, []int) {
	return file_telemetry_events_pmm_server_uptime_event_proto_rawDescGZIP(), []int{0}
}

// ServerUptimeEvent contains basic information about PMM Server version,
// uptime, and distribution method.
type ServerUptimeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PMM Server unique identifier.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// PMM Server version.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// PMM Server uptime.
	UpDuration *durationpb.Duration `protobuf:"bytes,3,opt,name=up_duration,json=upDuration,proto3" json:"up_duration,omitempty"`
	// PMM Server distribution method.
	DistributionMethod DistributionMethod `protobuf:"varint,4,opt,name=distribution_method,json=distributionMethod,proto3,enum=percona.platform.telemetry.events.pmm.v1.DistributionMethod" json:"distribution_method,omitempty"`
	// PMM Server STT enabled.
	SttEnabled *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=stt_enabled,json=sttEnabled,proto3" json:"stt_enabled,omitempty"`
	// PMM Server IA enabled.
	IaEnabled *wrapperspb.BoolValue `protobuf:"bytes,6,opt,name=ia_enabled,json=iaEnabled,proto3" json:"ia_enabled,omitempty"`
}

func (x *ServerUptimeEvent) Reset() {
	*x = ServerUptimeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_events_pmm_server_uptime_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerUptimeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerUptimeEvent) ProtoMessage() {}

func (x *ServerUptimeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_events_pmm_server_uptime_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerUptimeEvent.ProtoReflect.Descriptor instead.
func (*ServerUptimeEvent) Descriptor() ([]byte, []int) {
	return file_telemetry_events_pmm_server_uptime_event_proto_rawDescGZIP(), []int{0}
}

func (x *ServerUptimeEvent) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ServerUptimeEvent) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServerUptimeEvent) GetUpDuration() *durationpb.Duration {
	if x != nil {
		return x.UpDuration
	}
	return nil
}

func (x *ServerUptimeEvent) GetDistributionMethod() DistributionMethod {
	if x != nil {
		return x.DistributionMethod
	}
	return DistributionMethod_DISTRIBUTION_METHOD_INVALID
}

func (x *ServerUptimeEvent) GetSttEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.SttEnabled
	}
	return nil
}

func (x *ServerUptimeEvent) GetIaEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.IaEnabled
	}
	return nil
}

// ServerMetrics contains PMM Server metrics.
// There are mandatory metrics and optional metric pairs.
// The exact metrics pairs (metric name, metric value) is defined by PMM Server.
type ServerMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event UUID.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Time when this event was received by the sender.
	Time *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// PMM Server unique Telemetry identifier.
	// Mandatory.
	PmmServerTelemetryId []byte `protobuf:"bytes,3,opt,name=pmm_server_telemetry_id,json=pmmServerTelemetryId,proto3" json:"pmm_server_telemetry_id,omitempty"`
	// PMM Server version.
	// Mandatory.
	PmmServerVersion string `protobuf:"bytes,4,opt,name=pmm_server_version,json=pmmServerVersion,proto3" json:"pmm_server_version,omitempty"`
	// PMM Server uptime.
	// Mandatory.
	UpDuration *durationpb.Duration `protobuf:"bytes,5,opt,name=up_duration,json=upDuration,proto3" json:"up_duration,omitempty"`
	// PMM Server distribution method.
	// Mandatory.
	DistributionMethod DistributionMethod `protobuf:"varint,6,opt,name=distribution_method,json=distributionMethod,proto3,enum=percona.platform.telemetry.events.pmm.v1.DistributionMethod" json:"distribution_method,omitempty"`
	// PMM server metrics collection.
	// Optional.
	Metrics []*ServerMetric_Metric `protobuf:"bytes,7,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *ServerMetric) Reset() {
	*x = ServerMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_events_pmm_server_uptime_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMetric) ProtoMessage() {}

func (x *ServerMetric) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_events_pmm_server_uptime_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMetric.ProtoReflect.Descriptor instead.
func (*ServerMetric) Descriptor() ([]byte, []int) {
	return file_telemetry_events_pmm_server_uptime_event_proto_rawDescGZIP(), []int{1}
}

func (x *ServerMetric) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ServerMetric) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *ServerMetric) GetPmmServerTelemetryId() []byte {
	if x != nil {
		return x.PmmServerTelemetryId
	}
	return nil
}

func (x *ServerMetric) GetPmmServerVersion() string {
	if x != nil {
		return x.PmmServerVersion
	}
	return ""
}

func (x *ServerMetric) GetUpDuration() *durationpb.Duration {
	if x != nil {
		return x.UpDuration
	}
	return nil
}

func (x *ServerMetric) GetDistributionMethod() DistributionMethod {
	if x != nil {
		return x.DistributionMethod
	}
	return DistributionMethod_DISTRIBUTION_METHOD_INVALID
}

func (x *ServerMetric) GetMetrics() []*ServerMetric_Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type ServerMetric_Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ServerMetric_Metric) Reset() {
	*x = ServerMetric_Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_events_pmm_server_uptime_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMetric_Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMetric_Metric) ProtoMessage() {}

func (x *ServerMetric_Metric) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_events_pmm_server_uptime_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMetric_Metric.ProtoReflect.Descriptor instead.
func (*ServerMetric_Metric) Descriptor() ([]byte, []int) {
	return file_telemetry_events_pmm_server_uptime_event_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ServerMetric_Metric) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ServerMetric_Metric) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_telemetry_events_pmm_server_uptime_event_proto protoreflect.FileDescriptor

var file_telemetry_events_pmm_server_uptime_event_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x28, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x6d, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67,
	0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x10, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x70, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0b, 0x75, 0x70, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0a, 0x75, 0x70,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x76, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x6d, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x88, 0x01, 0x01, 0x52, 0x12, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x74, 0x74, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0a, 0x73, 0x74, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x69, 0x61, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x69,
	0x61, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x9c, 0x04, 0x0a, 0x0c, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x10, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x20, 0x01, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x17, 0x70, 0x6d,
	0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xe2, 0xdf, 0x1f,
	0x03, 0x80, 0x01, 0x10, 0x52, 0x14, 0x70, 0x6d, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x12, 0x70, 0x6d,
	0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x70, 0x04, 0x52, 0x10,
	0x70, 0x6d, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x42, 0x0a, 0x0b, 0x75, 0x70, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x76, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x3c, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x6d, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42,
	0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x88, 0x01, 0x01, 0x52, 0x12, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x57, 0x0a, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e,
	0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x6d, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x1a, 0x30, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x66, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1f, 0x0a,
	0x1b, 0x44, 0x49, 0x53, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x4f, 0x43, 0x4b, 0x45, 0x52, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x56,
	0x46, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4d, 0x49, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x5a, 0x55, 0x52, 0x45, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x44, 0x4f, 0x10, 0x05, 0x42,
	0xd2, 0x02, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x6d, 0x6d, 0x2e, 0x76, 0x31,
	0x42, 0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2d, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6d, 0x6d, 0x3b, 0x70, 0x6d, 0x6d, 0x76, 0x31, 0xa2,
	0x02, 0x05, 0x50, 0x50, 0x54, 0x45, 0x50, 0xaa, 0x02, 0x28, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e,
	0x61, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x6d, 0x6d, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x28, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x5c, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5c,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x50, 0x6d, 0x6d, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x34,
	0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x5c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x5c, 0x50, 0x6d, 0x6d, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2d, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x3a, 0x3a,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x50, 0x6d, 0x6d,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_telemetry_events_pmm_server_uptime_event_proto_rawDescOnce sync.Once
	file_telemetry_events_pmm_server_uptime_event_proto_rawDescData = file_telemetry_events_pmm_server_uptime_event_proto_rawDesc
)

func file_telemetry_events_pmm_server_uptime_event_proto_rawDescGZIP() []byte {
	file_telemetry_events_pmm_server_uptime_event_proto_rawDescOnce.Do(func() {
		file_telemetry_events_pmm_server_uptime_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_telemetry_events_pmm_server_uptime_event_proto_rawDescData)
	})
	return file_telemetry_events_pmm_server_uptime_event_proto_rawDescData
}

var file_telemetry_events_pmm_server_uptime_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_telemetry_events_pmm_server_uptime_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_telemetry_events_pmm_server_uptime_event_proto_goTypes = []interface{}{
	(DistributionMethod)(0),       // 0: percona.platform.telemetry.events.pmm.v1.DistributionMethod
	(*ServerUptimeEvent)(nil),     // 1: percona.platform.telemetry.events.pmm.v1.ServerUptimeEvent
	(*ServerMetric)(nil),          // 2: percona.platform.telemetry.events.pmm.v1.ServerMetric
	(*ServerMetric_Metric)(nil),   // 3: percona.platform.telemetry.events.pmm.v1.ServerMetric.Metric
	(*durationpb.Duration)(nil),   // 4: google.protobuf.Duration
	(*wrapperspb.BoolValue)(nil),  // 5: google.protobuf.BoolValue
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_telemetry_events_pmm_server_uptime_event_proto_depIdxs = []int32{
	4, // 0: percona.platform.telemetry.events.pmm.v1.ServerUptimeEvent.up_duration:type_name -> google.protobuf.Duration
	0, // 1: percona.platform.telemetry.events.pmm.v1.ServerUptimeEvent.distribution_method:type_name -> percona.platform.telemetry.events.pmm.v1.DistributionMethod
	5, // 2: percona.platform.telemetry.events.pmm.v1.ServerUptimeEvent.stt_enabled:type_name -> google.protobuf.BoolValue
	5, // 3: percona.platform.telemetry.events.pmm.v1.ServerUptimeEvent.ia_enabled:type_name -> google.protobuf.BoolValue
	6, // 4: percona.platform.telemetry.events.pmm.v1.ServerMetric.time:type_name -> google.protobuf.Timestamp
	4, // 5: percona.platform.telemetry.events.pmm.v1.ServerMetric.up_duration:type_name -> google.protobuf.Duration
	0, // 6: percona.platform.telemetry.events.pmm.v1.ServerMetric.distribution_method:type_name -> percona.platform.telemetry.events.pmm.v1.DistributionMethod
	3, // 7: percona.platform.telemetry.events.pmm.v1.ServerMetric.metrics:type_name -> percona.platform.telemetry.events.pmm.v1.ServerMetric.Metric
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_telemetry_events_pmm_server_uptime_event_proto_init() }
func file_telemetry_events_pmm_server_uptime_event_proto_init() {
	if File_telemetry_events_pmm_server_uptime_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_telemetry_events_pmm_server_uptime_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerUptimeEvent); i {
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
		file_telemetry_events_pmm_server_uptime_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMetric); i {
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
		file_telemetry_events_pmm_server_uptime_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMetric_Metric); i {
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
			RawDescriptor: file_telemetry_events_pmm_server_uptime_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_telemetry_events_pmm_server_uptime_event_proto_goTypes,
		DependencyIndexes: file_telemetry_events_pmm_server_uptime_event_proto_depIdxs,
		EnumInfos:         file_telemetry_events_pmm_server_uptime_event_proto_enumTypes,
		MessageInfos:      file_telemetry_events_pmm_server_uptime_event_proto_msgTypes,
	}.Build()
	File_telemetry_events_pmm_server_uptime_event_proto = out.File
	file_telemetry_events_pmm_server_uptime_event_proto_rawDesc = nil
	file_telemetry_events_pmm_server_uptime_event_proto_goTypes = nil
	file_telemetry_events_pmm_server_uptime_event_proto_depIdxs = nil
}
