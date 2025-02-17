// event_service.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: grpc/event_service.proto

package eventservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PushEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required Attributes
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Source      string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"` // URI-reference
	SpecVersion string `protobuf:"bytes,3,opt,name=spec_version,json=specVersion,proto3" json:"spec_version,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// Optional & Extension Attributes
	Attributes map[string]*PushEventRequest_CloudEventAttributeValue `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// -- CloudEvent Data (Bytes, Text, or Proto)
	//
	// Types that are assignable to Data:
	//
	//	*PushEventRequest_BinaryData
	//	*PushEventRequest_TextData
	//	*PushEventRequest_ProtoData
	Data isPushEventRequest_Data `protobuf_oneof:"data"`
}

func (x *PushEventRequest) Reset() {
	*x = PushEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_event_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushEventRequest) ProtoMessage() {}

func (x *PushEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_event_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushEventRequest.ProtoReflect.Descriptor instead.
func (*PushEventRequest) Descriptor() ([]byte, []int) {
	return file_grpc_event_service_proto_rawDescGZIP(), []int{0}
}

func (x *PushEventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PushEventRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *PushEventRequest) GetSpecVersion() string {
	if x != nil {
		return x.SpecVersion
	}
	return ""
}

func (x *PushEventRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PushEventRequest) GetAttributes() map[string]*PushEventRequest_CloudEventAttributeValue {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (m *PushEventRequest) GetData() isPushEventRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *PushEventRequest) GetBinaryData() []byte {
	if x, ok := x.GetData().(*PushEventRequest_BinaryData); ok {
		return x.BinaryData
	}
	return nil
}

func (x *PushEventRequest) GetTextData() string {
	if x, ok := x.GetData().(*PushEventRequest_TextData); ok {
		return x.TextData
	}
	return ""
}

func (x *PushEventRequest) GetProtoData() *anypb.Any {
	if x, ok := x.GetData().(*PushEventRequest_ProtoData); ok {
		return x.ProtoData
	}
	return nil
}

type isPushEventRequest_Data interface {
	isPushEventRequest_Data()
}

type PushEventRequest_BinaryData struct {
	BinaryData []byte `protobuf:"bytes,6,opt,name=binary_data,json=binaryData,proto3,oneof"`
}

type PushEventRequest_TextData struct {
	TextData string `protobuf:"bytes,7,opt,name=text_data,json=textData,proto3,oneof"`
}

type PushEventRequest_ProtoData struct {
	ProtoData *anypb.Any `protobuf:"bytes,8,opt,name=proto_data,json=protoData,proto3,oneof"`
}

func (*PushEventRequest_BinaryData) isPushEventRequest_Data() {}

func (*PushEventRequest_TextData) isPushEventRequest_Data() {}

func (*PushEventRequest_ProtoData) isPushEventRequest_Data() {}

type PushEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PushEventResponse) Reset() {
	*x = PushEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_event_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushEventResponse) ProtoMessage() {}

func (x *PushEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_event_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushEventResponse.ProtoReflect.Descriptor instead.
func (*PushEventResponse) Descriptor() ([]byte, []int) {
	return file_grpc_event_service_proto_rawDescGZIP(), []int{1}
}

func (x *PushEventResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PushEventRequest_CloudEventAttributeValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Attr:
	//
	//	*PushEventRequest_CloudEventAttributeValue_CeBoolean
	//	*PushEventRequest_CloudEventAttributeValue_CeInteger
	//	*PushEventRequest_CloudEventAttributeValue_CeString
	//	*PushEventRequest_CloudEventAttributeValue_CeBytes
	//	*PushEventRequest_CloudEventAttributeValue_CeUri
	//	*PushEventRequest_CloudEventAttributeValue_CeUriRef
	//	*PushEventRequest_CloudEventAttributeValue_CeTimestamp
	Attr isPushEventRequest_CloudEventAttributeValue_Attr `protobuf_oneof:"attr"`
}

func (x *PushEventRequest_CloudEventAttributeValue) Reset() {
	*x = PushEventRequest_CloudEventAttributeValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_event_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushEventRequest_CloudEventAttributeValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushEventRequest_CloudEventAttributeValue) ProtoMessage() {}

func (x *PushEventRequest_CloudEventAttributeValue) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_event_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushEventRequest_CloudEventAttributeValue.ProtoReflect.Descriptor instead.
func (*PushEventRequest_CloudEventAttributeValue) Descriptor() ([]byte, []int) {
	return file_grpc_event_service_proto_rawDescGZIP(), []int{0, 1}
}

func (m *PushEventRequest_CloudEventAttributeValue) GetAttr() isPushEventRequest_CloudEventAttributeValue_Attr {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (x *PushEventRequest_CloudEventAttributeValue) GetCeBoolean() bool {
	if x, ok := x.GetAttr().(*PushEventRequest_CloudEventAttributeValue_CeBoolean); ok {
		return x.CeBoolean
	}
	return false
}

func (x *PushEventRequest_CloudEventAttributeValue) GetCeInteger() int32 {
	if x, ok := x.GetAttr().(*PushEventRequest_CloudEventAttributeValue_CeInteger); ok {
		return x.CeInteger
	}
	return 0
}

func (x *PushEventRequest_CloudEventAttributeValue) GetCeString() string {
	if x, ok := x.GetAttr().(*PushEventRequest_CloudEventAttributeValue_CeString); ok {
		return x.CeString
	}
	return ""
}

func (x *PushEventRequest_CloudEventAttributeValue) GetCeBytes() []byte {
	if x, ok := x.GetAttr().(*PushEventRequest_CloudEventAttributeValue_CeBytes); ok {
		return x.CeBytes
	}
	return nil
}

func (x *PushEventRequest_CloudEventAttributeValue) GetCeUri() string {
	if x, ok := x.GetAttr().(*PushEventRequest_CloudEventAttributeValue_CeUri); ok {
		return x.CeUri
	}
	return ""
}

func (x *PushEventRequest_CloudEventAttributeValue) GetCeUriRef() string {
	if x, ok := x.GetAttr().(*PushEventRequest_CloudEventAttributeValue_CeUriRef); ok {
		return x.CeUriRef
	}
	return ""
}

func (x *PushEventRequest_CloudEventAttributeValue) GetCeTimestamp() *timestamppb.Timestamp {
	if x, ok := x.GetAttr().(*PushEventRequest_CloudEventAttributeValue_CeTimestamp); ok {
		return x.CeTimestamp
	}
	return nil
}

type isPushEventRequest_CloudEventAttributeValue_Attr interface {
	isPushEventRequest_CloudEventAttributeValue_Attr()
}

type PushEventRequest_CloudEventAttributeValue_CeBoolean struct {
	CeBoolean bool `protobuf:"varint,1,opt,name=ce_boolean,json=ceBoolean,proto3,oneof"`
}

type PushEventRequest_CloudEventAttributeValue_CeInteger struct {
	CeInteger int32 `protobuf:"varint,2,opt,name=ce_integer,json=ceInteger,proto3,oneof"`
}

type PushEventRequest_CloudEventAttributeValue_CeString struct {
	CeString string `protobuf:"bytes,3,opt,name=ce_string,json=ceString,proto3,oneof"`
}

type PushEventRequest_CloudEventAttributeValue_CeBytes struct {
	CeBytes []byte `protobuf:"bytes,4,opt,name=ce_bytes,json=ceBytes,proto3,oneof"`
}

type PushEventRequest_CloudEventAttributeValue_CeUri struct {
	CeUri string `protobuf:"bytes,5,opt,name=ce_uri,json=ceUri,proto3,oneof"`
}

type PushEventRequest_CloudEventAttributeValue_CeUriRef struct {
	CeUriRef string `protobuf:"bytes,6,opt,name=ce_uri_ref,json=ceUriRef,proto3,oneof"`
}

type PushEventRequest_CloudEventAttributeValue_CeTimestamp struct {
	CeTimestamp *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=ce_timestamp,json=ceTimestamp,proto3,oneof"`
}

func (*PushEventRequest_CloudEventAttributeValue_CeBoolean) isPushEventRequest_CloudEventAttributeValue_Attr() {
}

func (*PushEventRequest_CloudEventAttributeValue_CeInteger) isPushEventRequest_CloudEventAttributeValue_Attr() {
}

func (*PushEventRequest_CloudEventAttributeValue_CeString) isPushEventRequest_CloudEventAttributeValue_Attr() {
}

func (*PushEventRequest_CloudEventAttributeValue_CeBytes) isPushEventRequest_CloudEventAttributeValue_Attr() {
}

func (*PushEventRequest_CloudEventAttributeValue_CeUri) isPushEventRequest_CloudEventAttributeValue_Attr() {
}

func (*PushEventRequest_CloudEventAttributeValue_CeUriRef) isPushEventRequest_CloudEventAttributeValue_Attr() {
}

func (*PushEventRequest_CloudEventAttributeValue_CeTimestamp) isPushEventRequest_CloudEventAttributeValue_Attr() {
}

var File_grpc_event_service_proto protoreflect.FileDescriptor

var file_grpc_event_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x05, 0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x0a, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x09, 0x74,
	0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x74, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x76, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x9a, 0x02, 0x0a, 0x18, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x65, 0x5f, 0x62, 0x6f, 0x6f,
	0x6c, 0x65, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x63, 0x65,
	0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x65, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x63,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x09, 0x63, 0x65, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x63,
	0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x08, 0x63, 0x65, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x63, 0x65, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x06, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x63, 0x65, 0x55, 0x72, 0x69, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x63, 0x65, 0x55, 0x72, 0x69, 0x52, 0x65, 0x66, 0x12, 0x3f, 0x0a,
	0x0c, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48,
	0x00, 0x52, 0x0b, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06,
	0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2d,
	0x0a, 0x11, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x5e, 0x0a,
	0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a,
	0x09, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a,
	0x0e, 0x2e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_event_service_proto_rawDescOnce sync.Once
	file_grpc_event_service_proto_rawDescData = file_grpc_event_service_proto_rawDesc
)

func file_grpc_event_service_proto_rawDescGZIP() []byte {
	file_grpc_event_service_proto_rawDescOnce.Do(func() {
		file_grpc_event_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_event_service_proto_rawDescData)
	})
	return file_grpc_event_service_proto_rawDescData
}

var file_grpc_event_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_event_service_proto_goTypes = []any{
	(*PushEventRequest)(nil),  // 0: eventservice.PushEventRequest
	(*PushEventResponse)(nil), // 1: eventservice.PushEventResponse
	nil,                       // 2: eventservice.PushEventRequest.AttributesEntry
	(*PushEventRequest_CloudEventAttributeValue)(nil), // 3: eventservice.PushEventRequest.CloudEventAttributeValue
	(*anypb.Any)(nil),             // 4: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_grpc_event_service_proto_depIdxs = []int32{
	2, // 0: eventservice.PushEventRequest.attributes:type_name -> eventservice.PushEventRequest.AttributesEntry
	4, // 1: eventservice.PushEventRequest.proto_data:type_name -> google.protobuf.Any
	3, // 2: eventservice.PushEventRequest.AttributesEntry.value:type_name -> eventservice.PushEventRequest.CloudEventAttributeValue
	5, // 3: eventservice.PushEventRequest.CloudEventAttributeValue.ce_timestamp:type_name -> google.protobuf.Timestamp
	0, // 4: eventservice.EventService.PushEvent:input_type -> eventservice.PushEventRequest
	1, // 5: eventservice.EventService.PushEvent:output_type -> eventservice.PushEventResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_grpc_event_service_proto_init() }
func file_grpc_event_service_proto_init() {
	if File_grpc_event_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_event_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PushEventRequest); i {
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
		file_grpc_event_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PushEventResponse); i {
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
		file_grpc_event_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PushEventRequest_CloudEventAttributeValue); i {
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
	file_grpc_event_service_proto_msgTypes[0].OneofWrappers = []any{
		(*PushEventRequest_BinaryData)(nil),
		(*PushEventRequest_TextData)(nil),
		(*PushEventRequest_ProtoData)(nil),
	}
	file_grpc_event_service_proto_msgTypes[3].OneofWrappers = []any{
		(*PushEventRequest_CloudEventAttributeValue_CeBoolean)(nil),
		(*PushEventRequest_CloudEventAttributeValue_CeInteger)(nil),
		(*PushEventRequest_CloudEventAttributeValue_CeString)(nil),
		(*PushEventRequest_CloudEventAttributeValue_CeBytes)(nil),
		(*PushEventRequest_CloudEventAttributeValue_CeUri)(nil),
		(*PushEventRequest_CloudEventAttributeValue_CeUriRef)(nil),
		(*PushEventRequest_CloudEventAttributeValue_CeTimestamp)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_event_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_event_service_proto_goTypes,
		DependencyIndexes: file_grpc_event_service_proto_depIdxs,
		MessageInfos:      file_grpc_event_service_proto_msgTypes,
	}.Build()
	File_grpc_event_service_proto = out.File
	file_grpc_event_service_proto_rawDesc = nil
	file_grpc_event_service_proto_goTypes = nil
	file_grpc_event_service_proto_depIdxs = nil
}
