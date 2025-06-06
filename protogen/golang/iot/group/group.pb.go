// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: iot/group/group.proto

package group

import (
	user "github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GroupId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GId uint64 `protobuf:"varint,1,opt,name=g_id,json=gId,proto3" json:"g_id,omitempty"`
}

func (x *GroupId) Reset() {
	*x = GroupId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_group_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupId) ProtoMessage() {}

func (x *GroupId) ProtoReflect() protoreflect.Message {
	mi := &file_iot_group_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupId.ProtoReflect.Descriptor instead.
func (*GroupId) Descriptor() ([]byte, []int) {
	return file_iot_group_group_proto_rawDescGZIP(), []int{0}
}

func (x *GroupId) GetGId() uint64 {
	if x != nil {
		return x.GId
	}
	return 0
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GId     uint64         `protobuf:"varint,1,opt,name=g_id,json=gId,proto3" json:"g_id,omitempty"`
	Name    string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Members []*user.UserID `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_group_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_iot_group_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_iot_group_group_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetGId() uint64 {
	if x != nil {
		return x.GId
	}
	return 0
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetMembers() []*user.UserID {
	if x != nil {
		return x.Members
	}
	return nil
}

type UserAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	GroupId uint64 `protobuf:"varint,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *UserAdd) Reset() {
	*x = UserAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_group_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdd) ProtoMessage() {}

func (x *UserAdd) ProtoReflect() protoreflect.Message {
	mi := &file_iot_group_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdd.ProtoReflect.Descriptor instead.
func (*UserAdd) Descriptor() ([]byte, []int) {
	return file_iot_group_group_proto_rawDescGZIP(), []int{2}
}

func (x *UserAdd) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserAdd) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

var File_iot_group_group_proto protoreflect.FileDescriptor

var file_iot_group_group_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6f, 0x74, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x69, 0x6f, 0x74, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x07, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x11, 0x0a, 0x04, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x67, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x11, 0x0a, 0x04, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x3b, 0x0a, 0x07, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x32, 0xc1, 0x02, 0x0a, 0x0c, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x06, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e,
	0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x01, 0x2a, 0x12, 0x3a,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x06, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x1a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2f, 0x7b, 0x67, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x39, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x08, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x7b,
	0x67, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x36, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x08, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x1a, 0x06, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x7b, 0x67, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x4d, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x08, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x7d, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x7d, 0x42, 0x4f, 0x5a, 0x4d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x61,
	0x76, 0x62, 0x69, 0x73, 0x77, 0x61, 0x73, 0x73, 0x61, 0x6e, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x74,
	0x68, 0x2d, 0x70, 0x75, 0x6c, 0x73, 0x65, 0x2d, 0x69, 0x6f, 0x74, 0x2d, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x69, 0x6f, 0x74, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iot_group_group_proto_rawDescOnce sync.Once
	file_iot_group_group_proto_rawDescData = file_iot_group_group_proto_rawDesc
)

func file_iot_group_group_proto_rawDescGZIP() []byte {
	file_iot_group_group_proto_rawDescOnce.Do(func() {
		file_iot_group_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_iot_group_group_proto_rawDescData)
	})
	return file_iot_group_group_proto_rawDescData
}

var file_iot_group_group_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_iot_group_group_proto_goTypes = []interface{}{
	(*GroupId)(nil),     // 0: GroupId
	(*Group)(nil),       // 1: Group
	(*UserAdd)(nil),     // 2: UserAdd
	(*user.UserID)(nil), // 3: UserID
	(*user.Empty)(nil),  // 4: Empty
}
var file_iot_group_group_proto_depIdxs = []int32{
	3, // 0: Group.members:type_name -> UserID
	1, // 1: GroupManager.CreateGroup:input_type -> Group
	1, // 2: GroupManager.UpdateGroup:input_type -> Group
	0, // 3: GroupManager.DeleteGroup:input_type -> GroupId
	0, // 4: GroupManager.GetGroup:input_type -> GroupId
	2, // 5: GroupManager.AddUserToGroup:input_type -> UserAdd
	4, // 6: GroupManager.CreateGroup:output_type -> Empty
	4, // 7: GroupManager.UpdateGroup:output_type -> Empty
	4, // 8: GroupManager.DeleteGroup:output_type -> Empty
	1, // 9: GroupManager.GetGroup:output_type -> Group
	4, // 10: GroupManager.AddUserToGroup:output_type -> Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_iot_group_group_proto_init() }
func file_iot_group_group_proto_init() {
	if File_iot_group_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iot_group_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupId); i {
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
		file_iot_group_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_iot_group_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAdd); i {
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
			RawDescriptor: file_iot_group_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iot_group_group_proto_goTypes,
		DependencyIndexes: file_iot_group_group_proto_depIdxs,
		MessageInfos:      file_iot_group_group_proto_msgTypes,
	}.Build()
	File_iot_group_group_proto = out.File
	file_iot_group_group_proto_rawDesc = nil
	file_iot_group_group_proto_goTypes = nil
	file_iot_group_group_proto_depIdxs = nil
}
