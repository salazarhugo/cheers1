// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: cheers/post/v1/post_event.proto

package post

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

type PostEvent_EventType int32

const (
	PostEvent_CREATE  PostEvent_EventType = 0
	PostEvent_UPDATE  PostEvent_EventType = 1
	PostEvent_DELETE  PostEvent_EventType = 2
	PostEvent_LIKE    PostEvent_EventType = 3
	PostEvent_COMMENT PostEvent_EventType = 4
)

// Enum value maps for PostEvent_EventType.
var (
	PostEvent_EventType_name = map[int32]string{
		0: "CREATE",
		1: "UPDATE",
		2: "DELETE",
		3: "LIKE",
		4: "COMMENT",
	}
	PostEvent_EventType_value = map[string]int32{
		"CREATE":  0,
		"UPDATE":  1,
		"DELETE":  2,
		"LIKE":    3,
		"COMMENT": 4,
	}
)

func (x PostEvent_EventType) Enum() *PostEvent_EventType {
	p := new(PostEvent_EventType)
	*p = x
	return p
}

func (x PostEvent_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PostEvent_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_cheers_post_v1_post_event_proto_enumTypes[0].Descriptor()
}

func (PostEvent_EventType) Type() protoreflect.EnumType {
	return &file_cheers_post_v1_post_event_proto_enumTypes[0]
}

func (x PostEvent_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PostEvent_EventType.Descriptor instead.
func (PostEvent_EventType) EnumDescriptor() ([]byte, []int) {
	return file_cheers_post_v1_post_event_proto_rawDescGZIP(), []int{0, 0}
}

type PostEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string              `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostId string              `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Type   PostEvent_EventType `protobuf:"varint,3,opt,name=type,proto3,enum=cheers.post.v1.PostEvent_EventType" json:"type,omitempty"`
}

func (x *PostEvent) Reset() {
	*x = PostEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_post_v1_post_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostEvent) ProtoMessage() {}

func (x *PostEvent) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_post_v1_post_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostEvent.ProtoReflect.Descriptor instead.
func (*PostEvent) Descriptor() ([]byte, []int) {
	return file_cheers_post_v1_post_event_proto_rawDescGZIP(), []int{0}
}

func (x *PostEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PostEvent) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *PostEvent) GetType() PostEvent_EventType {
	if x != nil {
		return x.Type
	}
	return PostEvent_CREATE
}

var File_cheers_post_v1_post_event_proto protoreflect.FileDescriptor

var file_cheers_post_v1_post_event_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x22, 0xbe, 0x01, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x23, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x46, 0x0a, 0x09, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c,
	0x49, 0x4b, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x04, 0x42, 0x3f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x61, 0x7a, 0x61, 0x72, 0x68, 0x75, 0x67, 0x6f, 0x2f, 0x63,
	0x68, 0x65, 0x65, 0x72, 0x73, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cheers_post_v1_post_event_proto_rawDescOnce sync.Once
	file_cheers_post_v1_post_event_proto_rawDescData = file_cheers_post_v1_post_event_proto_rawDesc
)

func file_cheers_post_v1_post_event_proto_rawDescGZIP() []byte {
	file_cheers_post_v1_post_event_proto_rawDescOnce.Do(func() {
		file_cheers_post_v1_post_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_cheers_post_v1_post_event_proto_rawDescData)
	})
	return file_cheers_post_v1_post_event_proto_rawDescData
}

var file_cheers_post_v1_post_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cheers_post_v1_post_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cheers_post_v1_post_event_proto_goTypes = []interface{}{
	(PostEvent_EventType)(0), // 0: cheers.post.v1.PostEvent.EventType
	(*PostEvent)(nil),        // 1: cheers.post.v1.PostEvent
}
var file_cheers_post_v1_post_event_proto_depIdxs = []int32{
	0, // 0: cheers.post.v1.PostEvent.type:type_name -> cheers.post.v1.PostEvent.EventType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cheers_post_v1_post_event_proto_init() }
func file_cheers_post_v1_post_event_proto_init() {
	if File_cheers_post_v1_post_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cheers_post_v1_post_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostEvent); i {
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
			RawDescriptor: file_cheers_post_v1_post_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cheers_post_v1_post_event_proto_goTypes,
		DependencyIndexes: file_cheers_post_v1_post_event_proto_depIdxs,
		EnumInfos:         file_cheers_post_v1_post_event_proto_enumTypes,
		MessageInfos:      file_cheers_post_v1_post_event_proto_msgTypes,
	}.Build()
	File_cheers_post_v1_post_event_proto = out.File
	file_cheers_post_v1_post_event_proto_rawDesc = nil
	file_cheers_post_v1_post_event_proto_goTypes = nil
	file_cheers_post_v1_post_event_proto_depIdxs = nil
}
