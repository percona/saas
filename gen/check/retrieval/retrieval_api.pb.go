// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: check/retrieval/retrieval_api.proto

package retrievalv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAllChecksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllChecksRequest) Reset() {
	*x = GetAllChecksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_retrieval_retrieval_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllChecksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllChecksRequest) ProtoMessage() {}

func (x *GetAllChecksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_check_retrieval_retrieval_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllChecksRequest.ProtoReflect.Descriptor instead.
func (*GetAllChecksRequest) Descriptor() ([]byte, []int) {
	return file_check_retrieval_retrieval_api_proto_rawDescGZIP(), []int{0}
}

type GetAllChecksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// YAML formatted checks.
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// Checks signatures.
	Signatures []string `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *GetAllChecksResponse) Reset() {
	*x = GetAllChecksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_retrieval_retrieval_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllChecksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllChecksResponse) ProtoMessage() {}

func (x *GetAllChecksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_retrieval_retrieval_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllChecksResponse.ProtoReflect.Descriptor instead.
func (*GetAllChecksResponse) Descriptor() ([]byte, []int) {
	return file_check_retrieval_retrieval_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllChecksResponse) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *GetAllChecksResponse) GetSignatures() []string {
	if x != nil {
		return x.Signatures
	}
	return nil
}

type GetAllAlertRuleTemplatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllAlertRuleTemplatesRequest) Reset() {
	*x = GetAllAlertRuleTemplatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_retrieval_retrieval_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAlertRuleTemplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAlertRuleTemplatesRequest) ProtoMessage() {}

func (x *GetAllAlertRuleTemplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_check_retrieval_retrieval_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAlertRuleTemplatesRequest.ProtoReflect.Descriptor instead.
func (*GetAllAlertRuleTemplatesRequest) Descriptor() ([]byte, []int) {
	return file_check_retrieval_retrieval_api_proto_rawDescGZIP(), []int{2}
}

type GetAllAlertRuleTemplatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// YAML formatted rules.
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// Rules signatures.
	Signatures []string `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *GetAllAlertRuleTemplatesResponse) Reset() {
	*x = GetAllAlertRuleTemplatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_retrieval_retrieval_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAlertRuleTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAlertRuleTemplatesResponse) ProtoMessage() {}

func (x *GetAllAlertRuleTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_retrieval_retrieval_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAlertRuleTemplatesResponse.ProtoReflect.Descriptor instead.
func (*GetAllAlertRuleTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_check_retrieval_retrieval_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllAlertRuleTemplatesResponse) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *GetAllAlertRuleTemplatesResponse) GetSignatures() []string {
	if x != nil {
		return x.Signatures
	}
	return nil
}

var File_check_retrieval_retrieval_api_proto protoreflect.FileDescriptor

var file_check_retrieval_retrieval_api_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61,
	0x6c, 0x2f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x72, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x4a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x1f, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x56,
	0x0a, 0x20, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c,
	0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x32, 0xaa, 0x05, 0x0a, 0x0c, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x61, 0x6c, 0x41, 0x50, 0x49, 0x12, 0xa8, 0x02, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x38, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f,
	0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x39, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x72, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa2, 0x01,
	0x92, 0x41, 0x7e, 0x12, 0x0e, 0x47, 0x65, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x1a, 0x48, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61,
	0x6c, 0x6c, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73,
	0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x69, 0x65, 0x72, 0x2e, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x2f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0xee, 0x02, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x44, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x75, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x72,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc4, 0x01, 0x92,
	0x41, 0x93, 0x01, 0x12, 0x1c, 0x47, 0x65, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x20, 0x72, 0x75, 0x6c, 0x65, 0x20, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x1a, 0x4f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x61, 0x6c, 0x6c,
	0x20, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x20, 0x72, 0x75, 0x6c, 0x65, 0x20, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x69, 0x65,
	0x72, 0x2e, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x22, 0x22, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x3a, 0x01, 0x2a, 0x42, 0x90, 0x03, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x72, 0x63,
	0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x42,
	0x11, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x2f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x3b, 0x72,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x50, 0x50, 0x43,
	0x52, 0xaa, 0x02, 0x23, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x61, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x23, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e,
	0x61, 0x5c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x5c, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2f,
	0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x5c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x5c, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x27, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x3a, 0x3a, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x61, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x92, 0x41, 0x5a, 0x12, 0x0f, 0x0a, 0x0d,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x20, 0x41, 0x50, 0x49, 0x7a, 0x47, 0x0a,
	0x08, 0x78, 0x2d, 0x72, 0x65, 0x61, 0x64, 0x6d, 0x65, 0x12, 0x3b, 0x2a, 0x39, 0x0a, 0x37, 0x0a,
	0x11, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x22, 0x32, 0x20, 0x0a, 0x06, 0x1a, 0x04, 0x63, 0x75, 0x72, 0x6c, 0x0a, 0x04,
	0x1a, 0x02, 0x67, 0x6f, 0x0a, 0x06, 0x1a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x0a, 0x08, 0x1a, 0x06,
	0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_check_retrieval_retrieval_api_proto_rawDescOnce sync.Once
	file_check_retrieval_retrieval_api_proto_rawDescData = file_check_retrieval_retrieval_api_proto_rawDesc
)

func file_check_retrieval_retrieval_api_proto_rawDescGZIP() []byte {
	file_check_retrieval_retrieval_api_proto_rawDescOnce.Do(func() {
		file_check_retrieval_retrieval_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_check_retrieval_retrieval_api_proto_rawDescData)
	})
	return file_check_retrieval_retrieval_api_proto_rawDescData
}

var file_check_retrieval_retrieval_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_check_retrieval_retrieval_api_proto_goTypes = []interface{}{
	(*GetAllChecksRequest)(nil),              // 0: percona.platform.check.retrieval.v1.GetAllChecksRequest
	(*GetAllChecksResponse)(nil),             // 1: percona.platform.check.retrieval.v1.GetAllChecksResponse
	(*GetAllAlertRuleTemplatesRequest)(nil),  // 2: percona.platform.check.retrieval.v1.GetAllAlertRuleTemplatesRequest
	(*GetAllAlertRuleTemplatesResponse)(nil), // 3: percona.platform.check.retrieval.v1.GetAllAlertRuleTemplatesResponse
}
var file_check_retrieval_retrieval_api_proto_depIdxs = []int32{
	0, // 0: percona.platform.check.retrieval.v1.RetrievalAPI.GetAllChecks:input_type -> percona.platform.check.retrieval.v1.GetAllChecksRequest
	2, // 1: percona.platform.check.retrieval.v1.RetrievalAPI.GetAllAlertRuleTemplates:input_type -> percona.platform.check.retrieval.v1.GetAllAlertRuleTemplatesRequest
	1, // 2: percona.platform.check.retrieval.v1.RetrievalAPI.GetAllChecks:output_type -> percona.platform.check.retrieval.v1.GetAllChecksResponse
	3, // 3: percona.platform.check.retrieval.v1.RetrievalAPI.GetAllAlertRuleTemplates:output_type -> percona.platform.check.retrieval.v1.GetAllAlertRuleTemplatesResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_check_retrieval_retrieval_api_proto_init() }
func file_check_retrieval_retrieval_api_proto_init() {
	if File_check_retrieval_retrieval_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_check_retrieval_retrieval_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllChecksRequest); i {
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
		file_check_retrieval_retrieval_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllChecksResponse); i {
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
		file_check_retrieval_retrieval_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAlertRuleTemplatesRequest); i {
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
		file_check_retrieval_retrieval_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAlertRuleTemplatesResponse); i {
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
			RawDescriptor: file_check_retrieval_retrieval_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_check_retrieval_retrieval_api_proto_goTypes,
		DependencyIndexes: file_check_retrieval_retrieval_api_proto_depIdxs,
		MessageInfos:      file_check_retrieval_retrieval_api_proto_msgTypes,
	}.Build()
	File_check_retrieval_retrieval_api_proto = out.File
	file_check_retrieval_retrieval_api_proto_rawDesc = nil
	file_check_retrieval_retrieval_api_proto_goTypes = nil
	file_check_retrieval_retrieval_api_proto_depIdxs = nil
}
