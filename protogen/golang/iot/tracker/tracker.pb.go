// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: iot/tracker/tracker.proto

package tracker

import (
	user "github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	datetime "google.golang.org/genproto/googleapis/type/datetime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AlertType int32

const (
	AlertType_Normal    AlertType = 0
	AlertType_Increased AlertType = 1
	AlertType_Decreased AlertType = 2
)

// Enum value maps for AlertType.
var (
	AlertType_name = map[int32]string{
		0: "Normal",
		1: "Increased",
		2: "Decreased",
	}
	AlertType_value = map[string]int32{
		"Normal":    0,
		"Increased": 1,
		"Decreased": 2,
	}
)

func (x AlertType) Enum() *AlertType {
	p := new(AlertType)
	*p = x
	return p
}

func (x AlertType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertType) Descriptor() protoreflect.EnumDescriptor {
	return file_iot_tracker_tracker_proto_enumTypes[0].Descriptor()
}

func (AlertType) Type() protoreflect.EnumType {
	return &file_iot_tracker_tracker_proto_enumTypes[0]
}

func (x AlertType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertType.Descriptor instead.
func (AlertType) EnumDescriptor() ([]byte, []int) {
	return file_iot_tracker_tracker_proto_rawDescGZIP(), []int{0}
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint64             `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Latitude  float64            `protobuf:"fixed64,2,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude float64            `protobuf:"fixed64,3,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	Time      *datetime.DateTime `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	CkId      uint64             `protobuf:"varint,5,opt,name=ck_id,json=ckId,proto3" json:"ck_id,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_tracker_tracker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_iot_tracker_tracker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_iot_tracker_tracker_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Position) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Position) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Position) GetTime() *datetime.DateTime {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Position) GetCkId() uint64 {
	if x != nil {
		return x.CkId
	}
	return 0
}

type CheckpointToAndFrom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To   uint64 `protobuf:"varint,1,opt,name=to,proto3" json:"to,omitempty"`
	From uint64 `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *CheckpointToAndFrom) Reset() {
	*x = CheckpointToAndFrom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_tracker_tracker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckpointToAndFrom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckpointToAndFrom) ProtoMessage() {}

func (x *CheckpointToAndFrom) ProtoReflect() protoreflect.Message {
	mi := &file_iot_tracker_tracker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckpointToAndFrom.ProtoReflect.Descriptor instead.
func (*CheckpointToAndFrom) Descriptor() ([]byte, []int) {
	return file_iot_tracker_tracker_proto_rawDescGZIP(), []int{1}
}

func (x *CheckpointToAndFrom) GetTo() uint64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *CheckpointToAndFrom) GetFrom() uint64 {
	if x != nil {
		return x.From
	}
	return 0
}

type Distance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meter float64 `protobuf:"fixed64,1,opt,name=meter,proto3" json:"meter,omitempty"`
}

func (x *Distance) Reset() {
	*x = Distance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_tracker_tracker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distance) ProtoMessage() {}

func (x *Distance) ProtoReflect() protoreflect.Message {
	mi := &file_iot_tracker_tracker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distance.ProtoReflect.Descriptor instead.
func (*Distance) Descriptor() ([]byte, []int) {
	return file_iot_tracker_tracker_proto_rawDescGZIP(), []int{2}
}

func (x *Distance) GetMeter() float64 {
	if x != nil {
		return x.Meter
	}
	return 0
}

type PulseRateWithUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	PulseRate float32 `protobuf:"fixed32,2,opt,name=pulseRate,proto3" json:"pulseRate,omitempty"`
}

func (x *PulseRateWithUserId) Reset() {
	*x = PulseRateWithUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_tracker_tracker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PulseRateWithUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PulseRateWithUserId) ProtoMessage() {}

func (x *PulseRateWithUserId) ProtoReflect() protoreflect.Message {
	mi := &file_iot_tracker_tracker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PulseRateWithUserId.ProtoReflect.Descriptor instead.
func (*PulseRateWithUserId) Descriptor() ([]byte, []int) {
	return file_iot_tracker_tracker_proto_rawDescGZIP(), []int{3}
}

func (x *PulseRateWithUserId) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PulseRateWithUserId) GetPulseRate() float32 {
	if x != nil {
		return x.PulseRate
	}
	return 0
}

type CheckpointID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CkId uint64 `protobuf:"varint,1,opt,name=ck_id,json=ckId,proto3" json:"ck_id,omitempty"`
}

func (x *CheckpointID) Reset() {
	*x = CheckpointID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_tracker_tracker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckpointID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckpointID) ProtoMessage() {}

func (x *CheckpointID) ProtoReflect() protoreflect.Message {
	mi := &file_iot_tracker_tracker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckpointID.ProtoReflect.Descriptor instead.
func (*CheckpointID) Descriptor() ([]byte, []int) {
	return file_iot_tracker_tracker_proto_rawDescGZIP(), []int{4}
}

func (x *CheckpointID) GetCkId() uint64 {
	if x != nil {
		return x.CkId
	}
	return 0
}

type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alert  AlertType `protobuf:"varint,1,opt,name=alert,proto3,enum=AlertType" json:"alert,omitempty"`
	Advice string    `protobuf:"bytes,2,opt,name=advice,proto3" json:"advice,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_tracker_tracker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_iot_tracker_tracker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_iot_tracker_tracker_proto_rawDescGZIP(), []int{5}
}

func (x *Alert) GetAlert() AlertType {
	if x != nil {
		return x.Alert
	}
	return AlertType_Normal
}

func (x *Alert) GetAdvice() string {
	if x != nil {
		return x.Advice
	}
	return ""
}

var File_iot_tracker_tracker_proto protoreflect.FileDescriptor

var file_iot_tracker_tracker_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6f, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x69, 0x6f, 0x74,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x08, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x13, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x6f, 0x41, 0x6e, 0x64, 0x46, 0x72, 0x6f, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x22, 0x20, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x13, 0x50, 0x75, 0x6c, 0x73, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x6c, 0x73, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x75, 0x6c, 0x73, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x22, 0x23, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x12, 0x20, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0a, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x76, 0x69, 0x63, 0x65, 0x2a, 0x35, 0x0a, 0x09, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65,
	0x64, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x64,
	0x10, 0x02, 0x32, 0xfe, 0x03, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x44,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x46, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x28, 0x01, 0x12, 0x49, 0x0a, 0x0a,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x09, 0x2e, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0d, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x49, 0x44, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x55, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x75, 0x6c, 0x73, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x50, 0x75, 0x6c,
	0x73, 0x65, 0x52, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x1a, 0x06, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x75,
	0x6c, 0x73, 0x65, 0x72, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x28, 0x01, 0x30, 0x01, 0x12, 0x54,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x12, 0x09, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x6d, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x6f, 0x41, 0x6e, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x1a,
	0x09, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x21, 0x12, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f,
	0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x61, 0x76, 0x62, 0x69, 0x73, 0x77, 0x61, 0x73, 0x73, 0x61,
	0x6e, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x2d, 0x70, 0x75, 0x6c, 0x73, 0x65, 0x2d, 0x69,
	0x6f, 0x74, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x69, 0x6f, 0x74, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iot_tracker_tracker_proto_rawDescOnce sync.Once
	file_iot_tracker_tracker_proto_rawDescData = file_iot_tracker_tracker_proto_rawDesc
)

func file_iot_tracker_tracker_proto_rawDescGZIP() []byte {
	file_iot_tracker_tracker_proto_rawDescOnce.Do(func() {
		file_iot_tracker_tracker_proto_rawDescData = protoimpl.X.CompressGZIP(file_iot_tracker_tracker_proto_rawDescData)
	})
	return file_iot_tracker_tracker_proto_rawDescData
}

var file_iot_tracker_tracker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_iot_tracker_tracker_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_iot_tracker_tracker_proto_goTypes = []interface{}{
	(AlertType)(0),              // 0: AlertType
	(*Position)(nil),            // 1: Position
	(*CheckpointToAndFrom)(nil), // 2: CheckpointToAndFrom
	(*Distance)(nil),            // 3: Distance
	(*PulseRateWithUserId)(nil), // 4: PulseRateWithUserId
	(*CheckpointID)(nil),        // 5: CheckpointID
	(*Alert)(nil),               // 6: Alert
	(*datetime.DateTime)(nil),   // 7: google.type.DateTime
	(*user.UserID)(nil),         // 8: UserID
	(*user.Empty)(nil),          // 9: Empty
}
var file_iot_tracker_tracker_proto_depIdxs = []int32{
	7, // 0: Position.time:type_name -> google.type.DateTime
	0, // 1: Alert.alert:type_name -> AlertType
	8, // 2: Tracker.GetLocation:input_type -> UserID
	1, // 3: Tracker.UpdateLocation:input_type -> Position
	1, // 4: Tracker.Checkpoint:input_type -> Position
	4, // 5: Tracker.UpdatePulseRate:input_type -> PulseRateWithUserId
	1, // 6: Tracker.GetRealTimeDistanceCovered:input_type -> Position
	2, // 7: Tracker.GetTotalDistanceBetweenCheckpoint:input_type -> CheckpointToAndFrom
	1, // 8: Tracker.GetLocation:output_type -> Position
	9, // 9: Tracker.UpdateLocation:output_type -> Empty
	5, // 10: Tracker.Checkpoint:output_type -> CheckpointID
	6, // 11: Tracker.UpdatePulseRate:output_type -> Alert
	3, // 12: Tracker.GetRealTimeDistanceCovered:output_type -> Distance
	3, // 13: Tracker.GetTotalDistanceBetweenCheckpoint:output_type -> Distance
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_iot_tracker_tracker_proto_init() }
func file_iot_tracker_tracker_proto_init() {
	if File_iot_tracker_tracker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iot_tracker_tracker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_iot_tracker_tracker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckpointToAndFrom); i {
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
		file_iot_tracker_tracker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distance); i {
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
		file_iot_tracker_tracker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PulseRateWithUserId); i {
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
		file_iot_tracker_tracker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckpointID); i {
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
		file_iot_tracker_tracker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
			RawDescriptor: file_iot_tracker_tracker_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iot_tracker_tracker_proto_goTypes,
		DependencyIndexes: file_iot_tracker_tracker_proto_depIdxs,
		EnumInfos:         file_iot_tracker_tracker_proto_enumTypes,
		MessageInfos:      file_iot_tracker_tracker_proto_msgTypes,
	}.Build()
	File_iot_tracker_tracker_proto = out.File
	file_iot_tracker_tracker_proto_rawDesc = nil
	file_iot_tracker_tracker_proto_goTypes = nil
	file_iot_tracker_tracker_proto_depIdxs = nil
}
