// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: telemetry/reporter/event.proto

package reporterv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/mwitkow/go-proto-validators"
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

// AnyEvent represents a serialized telemetry event from the events sub-packages.
// It is similar to any.Any, but:
// * is not a well-known type, so no special support from the library (it is a good thing in our case);
// * may contain a JSON variant of serialization.
type AnyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message.
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Must be a valid serialized protobuf binary of the above specified type.
	// Only binary or json field must be set.
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3" json:"binary,omitempty"`
	// Must be a valid serialized protobuf JSON of the above specified type.
	// Only binary or json field must be set.
	Json string `protobuf:"bytes,3,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *AnyEvent) Reset() {
	*x = AnyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_reporter_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyEvent) ProtoMessage() {}

func (x *AnyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_reporter_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyEvent.ProtoReflect.Descriptor instead.
func (*AnyEvent) Descriptor() ([]byte, []int) {
	return file_telemetry_reporter_event_proto_rawDescGZIP(), []int{0}
}

func (x *AnyEvent) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *AnyEvent) GetBinary() []byte {
	if x != nil {
		return x.Binary
	}
	return nil
}

func (x *AnyEvent) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

// Event contains original event and additional information added by the sender.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event UUID.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Time when this event was received by the sender.
	Time *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// Original event with a well-known type URL.
	Event *AnyEvent `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_reporter_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_reporter_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_telemetry_reporter_event_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Event) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Event) GetEvent() *AnyEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_telemetry_reporter_event_proto protoreflect.FileDescriptor

var file_telemetry_reporter_event_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x26, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x41, 0x6e, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x5a,
	0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x3f, 0xe2, 0xdf, 0x1f, 0x3b, 0x0a, 0x39, 0x5e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61,
	0x5c, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x5c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x2e, 0x5c,
	0x77, 0x2b, 0x5c, 0x2e, 0x76, 0x31, 0x5c, 0x2e, 0x5c, 0x77, 0x2b, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x24, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0xa8, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xe2, 0xdf,
	0x1f, 0x03, 0x80, 0x01, 0x10, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x4e, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x79, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x42, 0xbd, 0x02, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e,
	0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f,
	0x6e, 0x61, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3b, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x50, 0x50, 0x54, 0x52, 0xaa, 0x02, 0x26,
	0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x26, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61,
	0x5c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x5c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x32, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x5c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x5c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5c, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2a, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x3a, 0x3a,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x3a, 0x3a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_telemetry_reporter_event_proto_rawDescOnce sync.Once
	file_telemetry_reporter_event_proto_rawDescData = file_telemetry_reporter_event_proto_rawDesc
)

func file_telemetry_reporter_event_proto_rawDescGZIP() []byte {
	file_telemetry_reporter_event_proto_rawDescOnce.Do(func() {
		file_telemetry_reporter_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_telemetry_reporter_event_proto_rawDescData)
	})
	return file_telemetry_reporter_event_proto_rawDescData
}

var file_telemetry_reporter_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_telemetry_reporter_event_proto_goTypes = []interface{}{
	(*AnyEvent)(nil),              // 0: percona.platform.telemetry.reporter.v1.AnyEvent
	(*Event)(nil),                 // 1: percona.platform.telemetry.reporter.v1.Event
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_telemetry_reporter_event_proto_depIdxs = []int32{
	2, // 0: percona.platform.telemetry.reporter.v1.Event.time:type_name -> google.protobuf.Timestamp
	0, // 1: percona.platform.telemetry.reporter.v1.Event.event:type_name -> percona.platform.telemetry.reporter.v1.AnyEvent
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_telemetry_reporter_event_proto_init() }
func file_telemetry_reporter_event_proto_init() {
	if File_telemetry_reporter_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_telemetry_reporter_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyEvent); i {
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
		file_telemetry_reporter_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_telemetry_reporter_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_telemetry_reporter_event_proto_goTypes,
		DependencyIndexes: file_telemetry_reporter_event_proto_depIdxs,
		MessageInfos:      file_telemetry_reporter_event_proto_msgTypes,
	}.Build()
	File_telemetry_reporter_event_proto = out.File
	file_telemetry_reporter_event_proto_rawDesc = nil
	file_telemetry_reporter_event_proto_goTypes = nil
	file_telemetry_reporter_event_proto_depIdxs = nil
}
