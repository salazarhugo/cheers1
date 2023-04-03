// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: cheers/activity/v1/activity_service.proto

package activity

import (
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

type Activity_ActivityType int32

const (
	Activity_POST_LIKED           Activity_ActivityType = 0
	Activity_STORY_LIKED          Activity_ActivityType = 1
	Activity_POST_COMMENTED       Activity_ActivityType = 2
	Activity_MENTION_POST_CAPTION Activity_ActivityType = 3
	Activity_MENTION_POST_COMMENT Activity_ActivityType = 4
	Activity_FRIEND_ADDED         Activity_ActivityType = 5
	Activity_COMMENT_LIKED        Activity_ActivityType = 6
	Activity_FOLLOW               Activity_ActivityType = 7
)

// Enum value maps for Activity_ActivityType.
var (
	Activity_ActivityType_name = map[int32]string{
		0: "POST_LIKED",
		1: "STORY_LIKED",
		2: "POST_COMMENTED",
		3: "MENTION_POST_CAPTION",
		4: "MENTION_POST_COMMENT",
		5: "FRIEND_ADDED",
		6: "COMMENT_LIKED",
		7: "FOLLOW",
	}
	Activity_ActivityType_value = map[string]int32{
		"POST_LIKED":           0,
		"STORY_LIKED":          1,
		"POST_COMMENTED":       2,
		"MENTION_POST_CAPTION": 3,
		"MENTION_POST_COMMENT": 4,
		"FRIEND_ADDED":         5,
		"COMMENT_LIKED":        6,
		"FOLLOW":               7,
	}
)

func (x Activity_ActivityType) Enum() *Activity_ActivityType {
	p := new(Activity_ActivityType)
	*p = x
	return p
}

func (x Activity_ActivityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Activity_ActivityType) Descriptor() protoreflect.EnumDescriptor {
	return file_cheers_activity_v1_activity_service_proto_enumTypes[0].Descriptor()
}

func (Activity_ActivityType) Type() protoreflect.EnumType {
	return &file_cheers_activity_v1_activity_service_proto_enumTypes[0]
}

func (x Activity_ActivityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Activity_ActivityType.Descriptor instead.
func (Activity_ActivityType) EnumDescriptor() ([]byte, []int) {
	return file_cheers_activity_v1_activity_service_proto_rawDescGZIP(), []int{2, 0}
}

type ListActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListActivityRequest) Reset() {
	*x = ListActivityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_activity_v1_activity_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActivityRequest) ProtoMessage() {}

func (x *ListActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_activity_v1_activity_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActivityRequest.ProtoReflect.Descriptor instead.
func (*ListActivityRequest) Descriptor() ([]byte, []int) {
	return file_cheers_activity_v1_activity_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListActivityRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activities []*Activity `protobuf:"bytes,1,rep,name=activities,proto3" json:"activities,omitempty"`
}

func (x *ListActivityResponse) Reset() {
	*x = ListActivityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_activity_v1_activity_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActivityResponse) ProtoMessage() {}

func (x *ListActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_activity_v1_activity_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActivityResponse.ProtoReflect.Descriptor instead.
func (*ListActivityResponse) Descriptor() ([]byte, []int) {
	return file_cheers_activity_v1_activity_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListActivityResponse) GetActivities() []*Activity {
	if x != nil {
		return x.Activities
	}
	return nil
}

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	Type         Activity_ActivityType `protobuf:"varint,1,opt,name=type,proto3,enum=cheers.activity.v1.Activity_ActivityType" json:"type,omitempty"`
	Text         string                `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Picture      string                `protobuf:"bytes,3,opt,name=picture,proto3" json:"picture,omitempty"`
	UserId       string                `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Timestamp    int64                 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	MediaPicture string                `protobuf:"bytes,6,opt,name=media_picture,json=mediaPicture,proto3" json:"media_picture,omitempty"`
	MediaId      string                `protobuf:"bytes,7,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	Username     string                `protobuf:"bytes,9,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_activity_v1_activity_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_activity_v1_activity_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_cheers_activity_v1_activity_service_proto_rawDescGZIP(), []int{2}
}

func (x *Activity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Activity) GetType() Activity_ActivityType {
	if x != nil {
		return x.Type
	}
	return Activity_POST_LIKED
}

func (x *Activity) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Activity) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Activity) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Activity) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Activity) GetMediaPicture() string {
	if x != nil {
		return x.MediaPicture
	}
	return ""
}

func (x *Activity) GetMediaId() string {
	if x != nil {
		return x.MediaId
	}
	return ""
}

func (x *Activity) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_cheers_activity_v1_activity_service_proto protoreflect.FileDescriptor

var file_cheers_activity_v1_activity_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x68, 0x65,
	0x65, 0x72, 0x73, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x54, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x68, 0x65, 0x65,
	0x72, 0x73, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x22, 0xc5, 0x03, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29,
	0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x70, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xa8, 0x01, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45,
	0x4e, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x4e, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03,
	0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x4e, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4f, 0x53, 0x54,
	0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x52,
	0x49, 0x45, 0x4e, 0x44, 0x5f, 0x41, 0x44, 0x44, 0x45, 0x44, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x44, 0x10, 0x06, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x4f, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x07, 0x32, 0x8c, 0x01, 0x0a, 0x0f,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x79, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12,
	0x27, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x45, 0x50, 0x01, 0x5a, 0x41,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x61, 0x7a,
	0x61, 0x72, 0x68, 0x75, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x31, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cheers_activity_v1_activity_service_proto_rawDescOnce sync.Once
	file_cheers_activity_v1_activity_service_proto_rawDescData = file_cheers_activity_v1_activity_service_proto_rawDesc
)

func file_cheers_activity_v1_activity_service_proto_rawDescGZIP() []byte {
	file_cheers_activity_v1_activity_service_proto_rawDescOnce.Do(func() {
		file_cheers_activity_v1_activity_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cheers_activity_v1_activity_service_proto_rawDescData)
	})
	return file_cheers_activity_v1_activity_service_proto_rawDescData
}

var file_cheers_activity_v1_activity_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cheers_activity_v1_activity_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cheers_activity_v1_activity_service_proto_goTypes = []interface{}{
	(Activity_ActivityType)(0),   // 0: cheers.activity.v1.Activity.ActivityType
	(*ListActivityRequest)(nil),  // 1: cheers.activity.v1.ListActivityRequest
	(*ListActivityResponse)(nil), // 2: cheers.activity.v1.ListActivityResponse
	(*Activity)(nil),             // 3: cheers.activity.v1.Activity
}
var file_cheers_activity_v1_activity_service_proto_depIdxs = []int32{
	3, // 0: cheers.activity.v1.ListActivityResponse.activities:type_name -> cheers.activity.v1.Activity
	0, // 1: cheers.activity.v1.Activity.type:type_name -> cheers.activity.v1.Activity.ActivityType
	1, // 2: cheers.activity.v1.ActivityService.ListActivity:input_type -> cheers.activity.v1.ListActivityRequest
	2, // 3: cheers.activity.v1.ActivityService.ListActivity:output_type -> cheers.activity.v1.ListActivityResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cheers_activity_v1_activity_service_proto_init() }
func file_cheers_activity_v1_activity_service_proto_init() {
	if File_cheers_activity_v1_activity_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cheers_activity_v1_activity_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActivityRequest); i {
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
		file_cheers_activity_v1_activity_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActivityResponse); i {
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
		file_cheers_activity_v1_activity_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
			RawDescriptor: file_cheers_activity_v1_activity_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cheers_activity_v1_activity_service_proto_goTypes,
		DependencyIndexes: file_cheers_activity_v1_activity_service_proto_depIdxs,
		EnumInfos:         file_cheers_activity_v1_activity_service_proto_enumTypes,
		MessageInfos:      file_cheers_activity_v1_activity_service_proto_msgTypes,
	}.Build()
	File_cheers_activity_v1_activity_service_proto = out.File
	file_cheers_activity_v1_activity_service_proto_rawDesc = nil
	file_cheers_activity_v1_activity_service_proto_goTypes = nil
	file_cheers_activity_v1_activity_service_proto_depIdxs = nil
}
