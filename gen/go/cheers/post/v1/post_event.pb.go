// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
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
	Post   *PostEvent_Post     `protobuf:"bytes,2,opt,name=post,proto3" json:"post,omitempty"`
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

func (x *PostEvent) GetPost() *PostEvent_Post {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *PostEvent) GetType() PostEvent_EventType {
	if x != nil {
		return x.Type
	}
	return PostEvent_CREATE
}

type PostEvent_Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorId        string   `protobuf:"bytes,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Caption          string   `protobuf:"bytes,3,opt,name=caption,proto3" json:"caption,omitempty"`
	Address          string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Photos           []string `protobuf:"bytes,6,rep,name=photos,proto3" json:"photos,omitempty"`
	IsCommentEnabled bool     `protobuf:"varint,7,opt,name=is_comment_enabled,json=isCommentEnabled,proto3" json:"is_comment_enabled,omitempty"`
	LocationName     string   `protobuf:"bytes,8,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
	Drink            string   `protobuf:"bytes,9,opt,name=drink,proto3" json:"drink,omitempty"`
	Drunkenness      int64    `protobuf:"varint,10,opt,name=drunkenness,proto3" json:"drunkenness,omitempty"`
	CommentsEnabled  bool     `protobuf:"varint,13,opt,name=comments_enabled,json=commentsEnabled,proto3" json:"comments_enabled,omitempty"`
	ShareEnabled     bool     `protobuf:"varint,14,opt,name=share_enabled,json=shareEnabled,proto3" json:"share_enabled,omitempty"`
}

func (x *PostEvent_Post) Reset() {
	*x = PostEvent_Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_post_v1_post_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostEvent_Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostEvent_Post) ProtoMessage() {}

func (x *PostEvent_Post) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_post_v1_post_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostEvent_Post.ProtoReflect.Descriptor instead.
func (*PostEvent_Post) Descriptor() ([]byte, []int) {
	return file_cheers_post_v1_post_event_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PostEvent_Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostEvent_Post) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *PostEvent_Post) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *PostEvent_Post) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PostEvent_Post) GetPhotos() []string {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *PostEvent_Post) GetIsCommentEnabled() bool {
	if x != nil {
		return x.IsCommentEnabled
	}
	return false
}

func (x *PostEvent_Post) GetLocationName() string {
	if x != nil {
		return x.LocationName
	}
	return ""
}

func (x *PostEvent_Post) GetDrink() string {
	if x != nil {
		return x.Drink
	}
	return ""
}

func (x *PostEvent_Post) GetDrunkenness() int64 {
	if x != nil {
		return x.Drunkenness
	}
	return 0
}

func (x *PostEvent_Post) GetCommentsEnabled() bool {
	if x != nil {
		return x.CommentsEnabled
	}
	return false
}

func (x *PostEvent_Post) GetShareEnabled() bool {
	if x != nil {
		return x.ShareEnabled
	}
	return false
}

var File_cheers_post_v1_post_event_proto protoreflect.FileDescriptor

var file_cheers_post_v1_post_event_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x22, 0xb8, 0x04, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x65,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0xdc, 0x02, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x73, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x64, 0x72, 0x69, 0x6e, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x72, 0x69,
	0x6e, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x72, 0x75, 0x6e, 0x6b, 0x65, 0x6e, 0x6e, 0x65, 0x73,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x72, 0x75, 0x6e, 0x6b, 0x65, 0x6e,
	0x6e, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x22, 0x46, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x03, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x42, 0x3d, 0x50, 0x01,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c,
	0x61, 0x7a, 0x61, 0x72, 0x68, 0x75, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x31,
	0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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
var file_cheers_post_v1_post_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cheers_post_v1_post_event_proto_goTypes = []interface{}{
	(PostEvent_EventType)(0), // 0: cheers.post.v1.PostEvent.EventType
	(*PostEvent)(nil),        // 1: cheers.post.v1.PostEvent
	(*PostEvent_Post)(nil),   // 2: cheers.post.v1.PostEvent.Post
}
var file_cheers_post_v1_post_event_proto_depIdxs = []int32{
	2, // 0: cheers.post.v1.PostEvent.post:type_name -> cheers.post.v1.PostEvent.Post
	0, // 1: cheers.post.v1.PostEvent.type:type_name -> cheers.post.v1.PostEvent.EventType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
		file_cheers_post_v1_post_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostEvent_Post); i {
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
			NumMessages:   2,
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
