// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: proto/user-event.proto

package usereventpb

import (
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

type UserEventType int32

const (
	UserEventType_FOLLOW       UserEventType = 0
	UserEventType_STORY_LIKE   UserEventType = 1
	UserEventType_POST_LIKE    UserEventType = 2
	UserEventType_COMMENT      UserEventType = 3
	UserEventType_MENTION      UserEventType = 4
	UserEventType_CREATE_POST  UserEventType = 5
	UserEventType_CREATE_EVENT UserEventType = 6
	UserEventType_CREATE_STORY UserEventType = 7
)

// Enum value maps for UserEventType.
var (
	UserEventType_name = map[int32]string{
		0: "FOLLOW",
		1: "STORY_LIKE",
		2: "POST_LIKE",
		3: "COMMENT",
		4: "MENTION",
		5: "CREATE_POST",
		6: "CREATE_EVENT",
		7: "CREATE_STORY",
	}
	UserEventType_value = map[string]int32{
		"FOLLOW":       0,
		"STORY_LIKE":   1,
		"POST_LIKE":    2,
		"COMMENT":      3,
		"MENTION":      4,
		"CREATE_POST":  5,
		"CREATE_EVENT": 6,
		"CREATE_STORY": 7,
	}
)

func (x UserEventType) Enum() *UserEventType {
	p := new(UserEventType)
	*p = x
	return p
}

func (x UserEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_user_event_proto_enumTypes[0].Descriptor()
}

func (UserEventType) Type() protoreflect.EnumType {
	return &file_proto_user_event_proto_enumTypes[0]
}

func (x UserEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserEventType.Descriptor instead.
func (UserEventType) EnumDescriptor() ([]byte, []int) {
	return file_proto_user_event_proto_rawDescGZIP(), []int{0}
}

type UserEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   UserEventType `protobuf:"varint,1,opt,name=type,proto3,enum=usereventpb.UserEventType" json:"type,omitempty"`
	UserId string        `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Time        int64         `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	OtherUserId string        `protobuf:"bytes,4,opt,name=otherUserId,proto3" json:"otherUserId,omitempty"`
	PostId      string        `protobuf:"bytes,5,opt,name=postId,proto3" json:"postId,omitempty"`
	EventId     string        `protobuf:"bytes,6,opt,name=eventId,proto3" json:"eventId,omitempty"`
	StoryId     string        `protobuf:"bytes,7,opt,name=storyId,proto3" json:"storyId,omitempty"`
}

func (x *UserEvent) Reset() {
	*x = UserEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEvent) ProtoMessage() {}

func (x *UserEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEvent.ProtoReflect.Descriptor instead.
func (*UserEvent) Descriptor() ([]byte, []int) {
	return file_proto_user_event_proto_rawDescGZIP(), []int{0}
}

func (x *UserEvent) GetType() UserEventType {
	if x != nil {
		return x.Type
	}
	return UserEventType_FOLLOW
}

func (x *UserEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserEvent) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *UserEvent) GetOtherUserId() string {
	if x != nil {
		return x.OtherUserId
	}
	return ""
}

func (x *UserEvent) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *UserEvent) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *UserEvent) GetStoryId() string {
	if x != nil {
		return x.StoryId
	}
	return ""
}

var File_proto_user_event_proto protoreflect.FileDescriptor

var file_proto_user_event_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x22, 0xd5, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x2a, 0x89, 0x01,
	0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x4f, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53,
	0x54, 0x4f, 0x52, 0x59, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x50,
	0x4f, 0x53, 0x54, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f,
	0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x4e, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x50,
	0x4f, 0x53, 0x54, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x07, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2f, 0x72,
	0x65, 0x73, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_event_proto_rawDescOnce sync.Once
	file_proto_user_event_proto_rawDescData = file_proto_user_event_proto_rawDesc
)

func file_proto_user_event_proto_rawDescGZIP() []byte {
	file_proto_user_event_proto_rawDescOnce.Do(func() {
		file_proto_user_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_event_proto_rawDescData)
	})
	return file_proto_user_event_proto_rawDescData
}

var file_proto_user_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_user_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_user_event_proto_goTypes = []interface{}{
	(UserEventType)(0), // 0: usereventpb.UserEventType
	(*UserEvent)(nil),  // 1: usereventpb.UserEvent
}
var file_proto_user_event_proto_depIdxs = []int32{
	0, // 0: usereventpb.UserEvent.type:type_name -> usereventpb.UserEventType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_user_event_proto_init() }
func file_proto_user_event_proto_init() {
	if File_proto_user_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEvent); i {
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
			RawDescriptor: file_proto_user_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_user_event_proto_goTypes,
		DependencyIndexes: file_proto_user_event_proto_depIdxs,
		EnumInfos:         file_proto_user_event_proto_enumTypes,
		MessageInfos:      file_proto_user_event_proto_msgTypes,
	}.Build()
	File_proto_user_event_proto = out.File
	file_proto_user_event_proto_rawDesc = nil
	file_proto_user_event_proto_goTypes = nil
	file_proto_user_event_proto_depIdxs = nil
}
