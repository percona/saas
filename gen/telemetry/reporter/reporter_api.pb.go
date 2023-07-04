// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: telemetry/reporter/reporter_api.proto

package reporterv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	pmm "github.com/percona-platform/saas/gen/telemetry/events/pmm"
	fsp "github.com/percona-platform/saas/gen/utils/fsp"
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
	// One or more metrics events to report.
	// Only 'events' or 'metrics' field must be set.
	Metrics []*pmm.ServerMetric `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
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

func (x *ReportRequest) GetMetrics() []*pmm.ServerMetric {
	if x != nil {
		return x.Metrics
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

type PMMServerMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the metric requested by the client.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	// Value of the metric requested by the client.
	MetricValue string `protobuf:"bytes,2,opt,name=metric_value,json=metricValue,proto3" json:"metric_value,omitempty"`
}

func (x *PMMServerMetric) Reset() {
	*x = PMMServerMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PMMServerMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PMMServerMetric) ProtoMessage() {}

func (x *PMMServerMetric) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PMMServerMetric.ProtoReflect.Descriptor instead.
func (*PMMServerMetric) Descriptor() ([]byte, []int) {
	return file_telemetry_reporter_reporter_api_proto_rawDescGZIP(), []int{2}
}

func (x *PMMServerMetric) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (x *PMMServerMetric) GetMetricValue() string {
	if x != nil {
		return x.MetricValue
	}
	return ""
}

// PMMMetricEvent represents telemetry event containing metrics sent by PMM.
type PMMMetricEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PMM Server UUID.
	PmmServerTelemetryId string `protobuf:"bytes,1,opt,name=pmm_server_telemetry_id,json=pmmServerTelemetryId,proto3" json:"pmm_server_telemetry_id,omitempty"`
	// Portal Organization UUID (Can be empty).
	PortalOrgId string `protobuf:"bytes,2,opt,name=portal_org_id,json=portalOrgId,proto3" json:"portal_org_id,omitempty"`
	// Time when this event was created and received.
	EventTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// Array of server metrics requested by the client.
	Metrics []*PMMServerMetric `protobuf:"bytes,4,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *PMMMetricEvent) Reset() {
	*x = PMMMetricEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PMMMetricEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PMMMetricEvent) ProtoMessage() {}

func (x *PMMMetricEvent) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PMMMetricEvent.ProtoReflect.Descriptor instead.
func (*PMMMetricEvent) Descriptor() ([]byte, []int) {
	return file_telemetry_reporter_reporter_api_proto_rawDescGZIP(), []int{3}
}

func (x *PMMMetricEvent) GetPmmServerTelemetryId() string {
	if x != nil {
		return x.PmmServerTelemetryId
	}
	return ""
}

func (x *PMMMetricEvent) GetPortalOrgId() string {
	if x != nil {
		return x.PortalOrgId
	}
	return ""
}

func (x *PMMMetricEvent) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *PMMMetricEvent) GetMetrics() []*PMMServerMetric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type SearchEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of metrics requested by the client.
	Metrics []string `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// Time range of events to search between.
	StartsAt   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=starts_at,json=startsAt,proto3" json:"starts_at,omitempty"`
	EndsAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=ends_at,json=endsAt,proto3" json:"ends_at,omitempty"`
	PageParams *fsp.PageParams        `protobuf:"bytes,4,opt,name=page_params,json=pageParams,proto3" json:"page_params,omitempty"`
}

func (x *SearchEventRequest) Reset() {
	*x = SearchEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchEventRequest) ProtoMessage() {}

func (x *SearchEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchEventRequest.ProtoReflect.Descriptor instead.
func (*SearchEventRequest) Descriptor() ([]byte, []int) {
	return file_telemetry_reporter_reporter_api_proto_rawDescGZIP(), []int{4}
}

func (x *SearchEventRequest) GetMetrics() []string {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *SearchEventRequest) GetStartsAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartsAt
	}
	return nil
}

func (x *SearchEventRequest) GetEndsAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndsAt
	}
	return nil
}

func (x *SearchEventRequest) GetPageParams() *fsp.PageParams {
	if x != nil {
		return x.PageParams
	}
	return nil
}

type SearchEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PMM metric events.
	Events []*PMMMetricEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	// pagination details.
	PageTotals *fsp.PageTotals `protobuf:"bytes,2,opt,name=page_totals,json=pageTotals,proto3" json:"page_totals,omitempty"`
}

func (x *SearchEventResponse) Reset() {
	*x = SearchEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchEventResponse) ProtoMessage() {}

func (x *SearchEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchEventResponse.ProtoReflect.Descriptor instead.
func (*SearchEventResponse) Descriptor() ([]byte, []int) {
	return file_telemetry_reporter_reporter_api_proto_rawDescGZIP(), []int{5}
}

func (x *SearchEventResponse) GetEvents() []*PMMMetricEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *SearchEventResponse) GetPageTotals() *fsp.PageTotals {
	if x != nil {
		return x.PageTotals
	}
	return nil
}

var File_telemetry_reporter_reporter_api_proto protoreflect.FileDescriptor

var file_telemetry_reporter_reporter_api_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74,
	0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x66, 0x73,
	0x70, 0x2f, 0x66, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x0d,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x50, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x6d, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x0a, 0x0f, 0x50, 0x4d, 0x4d, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0xf9, 0x01, 0x0a, 0x0e, 0x50, 0x4d, 0x4d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x35, 0x0a, 0x17, 0x70, 0x6d, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x6d, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x65, 0x72, 0x63,
	0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x4d, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x12,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x37, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x73, 0x41, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x73, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x73, 0x41, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x66, 0x73, 0x70, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x70, 0x65, 0x72,
	0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x4d, 0x4d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x0b, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x66, 0x73, 0x70, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x73, 0x32, 0x83, 0x05, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x41,
	0x50, 0x49, 0x12, 0xac, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x35, 0x2e,
	0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb2, 0x01, 0x92,
	0x41, 0x8f, 0x01, 0x12, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x40, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x20, 0x69, 0x73, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x20, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x20,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x64, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x4a, 0x1f, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x18, 0x22, 0x16, 0x0a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x02,
	0x7b, 0x7d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0xc4, 0x02, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x3a, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e,
	0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xbb, 0x01, 0x92, 0x41, 0x91,
	0x01, 0x12, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x5e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x65, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x20, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x61,
	0x6e, 0x64, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x74, 0x68, 0x65, 0x6d, 0x20,
	0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x32, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x3a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x9f, 0x03, 0x92, 0x41, 0x59, 0x12, 0x0e,
	0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x20, 0x41, 0x50, 0x49, 0x7a, 0x47,
	0x0a, 0x08, 0x78, 0x2d, 0x72, 0x65, 0x61, 0x64, 0x6d, 0x65, 0x12, 0x3b, 0x2a, 0x39, 0x0a, 0x37,
	0x0a, 0x11, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x22, 0x32, 0x20, 0x0a, 0x06, 0x1a, 0x04, 0x63, 0x75, 0x72, 0x6c, 0x0a,
	0x04, 0x1a, 0x02, 0x67, 0x6f, 0x0a, 0x06, 0x1a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x0a, 0x08, 0x1a,
	0x06, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x72,
	0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x42, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x41, 0x70, 0x69,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2d, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x3b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x76, 0x31, 0xa2,
	0x02, 0x04, 0x50, 0x50, 0x54, 0x52, 0xaa, 0x02, 0x26, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61,
	0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x26, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x5c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x5c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5c, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x32, 0x50, 0x65, 0x72, 0x63, 0x6f,
	0x6e, 0x61, 0x5c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x54, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x5c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2a,
	0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x3a, 0x3a, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_telemetry_reporter_reporter_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_telemetry_reporter_reporter_api_proto_goTypes = []interface{}{
	(*ReportRequest)(nil),         // 0: percona.platform.telemetry.reporter.v1.ReportRequest
	(*ReportResponse)(nil),        // 1: percona.platform.telemetry.reporter.v1.ReportResponse
	(*PMMServerMetric)(nil),       // 2: percona.platform.telemetry.reporter.v1.PMMServerMetric
	(*PMMMetricEvent)(nil),        // 3: percona.platform.telemetry.reporter.v1.PMMMetricEvent
	(*SearchEventRequest)(nil),    // 4: percona.platform.telemetry.reporter.v1.SearchEventRequest
	(*SearchEventResponse)(nil),   // 5: percona.platform.telemetry.reporter.v1.SearchEventResponse
	(*Event)(nil),                 // 6: percona.platform.telemetry.reporter.v1.Event
	(*pmm.ServerMetric)(nil),      // 7: percona.platform.telemetry.events.pmm.v1.ServerMetric
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*fsp.PageParams)(nil),        // 9: utils.fsp.PageParams
	(*fsp.PageTotals)(nil),        // 10: utils.fsp.PageTotals
}
var file_telemetry_reporter_reporter_api_proto_depIdxs = []int32{
	6,  // 0: percona.platform.telemetry.reporter.v1.ReportRequest.events:type_name -> percona.platform.telemetry.reporter.v1.Event
	7,  // 1: percona.platform.telemetry.reporter.v1.ReportRequest.metrics:type_name -> percona.platform.telemetry.events.pmm.v1.ServerMetric
	8,  // 2: percona.platform.telemetry.reporter.v1.PMMMetricEvent.event_time:type_name -> google.protobuf.Timestamp
	2,  // 3: percona.platform.telemetry.reporter.v1.PMMMetricEvent.metrics:type_name -> percona.platform.telemetry.reporter.v1.PMMServerMetric
	8,  // 4: percona.platform.telemetry.reporter.v1.SearchEventRequest.starts_at:type_name -> google.protobuf.Timestamp
	8,  // 5: percona.platform.telemetry.reporter.v1.SearchEventRequest.ends_at:type_name -> google.protobuf.Timestamp
	9,  // 6: percona.platform.telemetry.reporter.v1.SearchEventRequest.page_params:type_name -> utils.fsp.PageParams
	3,  // 7: percona.platform.telemetry.reporter.v1.SearchEventResponse.events:type_name -> percona.platform.telemetry.reporter.v1.PMMMetricEvent
	10, // 8: percona.platform.telemetry.reporter.v1.SearchEventResponse.page_totals:type_name -> utils.fsp.PageTotals
	0,  // 9: percona.platform.telemetry.reporter.v1.ReporterAPI.Report:input_type -> percona.platform.telemetry.reporter.v1.ReportRequest
	4,  // 10: percona.platform.telemetry.reporter.v1.ReporterAPI.SearchEvent:input_type -> percona.platform.telemetry.reporter.v1.SearchEventRequest
	1,  // 11: percona.platform.telemetry.reporter.v1.ReporterAPI.Report:output_type -> percona.platform.telemetry.reporter.v1.ReportResponse
	5,  // 12: percona.platform.telemetry.reporter.v1.ReporterAPI.SearchEvent:output_type -> percona.platform.telemetry.reporter.v1.SearchEventResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
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
		file_telemetry_reporter_reporter_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PMMServerMetric); i {
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
		file_telemetry_reporter_reporter_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PMMMetricEvent); i {
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
		file_telemetry_reporter_reporter_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchEventRequest); i {
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
		file_telemetry_reporter_reporter_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchEventResponse); i {
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
			NumMessages:   6,
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
