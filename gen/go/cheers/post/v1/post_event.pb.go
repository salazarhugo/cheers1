// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: cheers/post/v1/post_event.proto

package post

import (
	post "github.com/salazarhugo/cheers1/gen/go/cheers/type/post"
	user "github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
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

type PostEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*PostEvent_Create
	//	*PostEvent_Like
	//	*PostEvent_Delete
	Event isPostEvent_Event `protobuf_oneof:"event"`
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

func (m *PostEvent) GetEvent() isPostEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *PostEvent) GetCreate() *CreatePost {
	if x, ok := x.GetEvent().(*PostEvent_Create); ok {
		return x.Create
	}
	return nil
}

func (x *PostEvent) GetLike() *LikePost {
	if x, ok := x.GetEvent().(*PostEvent_Like); ok {
		return x.Like
	}
	return nil
}

func (x *PostEvent) GetDelete() *DeletePost {
	if x, ok := x.GetEvent().(*PostEvent_Delete); ok {
		return x.Delete
	}
	return nil
}

type isPostEvent_Event interface {
	isPostEvent_Event()
}

type PostEvent_Create struct {
	Create *CreatePost `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type PostEvent_Like struct {
	Like *LikePost `protobuf:"bytes,2,opt,name=like,proto3,oneof"`
}

type PostEvent_Delete struct {
	Delete *DeletePost `protobuf:"bytes,3,opt,name=delete,proto3,oneof"`
}

func (*PostEvent_Create) isPostEvent_Event() {}

func (*PostEvent_Like) isPostEvent_Event() {}

func (*PostEvent_Delete) isPostEvent_Event() {}

type CreatePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *post.Post     `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	User *user.UserItem `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreatePost) Reset() {
	*x = CreatePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_post_v1_post_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePost) ProtoMessage() {}

func (x *CreatePost) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreatePost.ProtoReflect.Descriptor instead.
func (*CreatePost) Descriptor() ([]byte, []int) {
	return file_cheers_post_v1_post_event_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePost) GetPost() *post.Post {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *CreatePost) GetUser() *user.UserItem {
	if x != nil {
		return x.User
	}
	return nil
}

type LikePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *LikePost) Reset() {
	*x = LikePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_post_v1_post_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikePost) ProtoMessage() {}

func (x *LikePost) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_post_v1_post_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikePost.ProtoReflect.Descriptor instead.
func (*LikePost) Descriptor() ([]byte, []int) {
	return file_cheers_post_v1_post_event_proto_rawDescGZIP(), []int{2}
}

func (x *LikePost) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *LikePost) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeletePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender *user.UserItem `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *DeletePost) Reset() {
	*x = DeletePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_post_v1_post_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePost) ProtoMessage() {}

func (x *DeletePost) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_post_v1_post_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePost.ProtoReflect.Descriptor instead.
func (*DeletePost) Descriptor() ([]byte, []int) {
	return file_cheers_post_v1_post_event_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePost) GetSender() *user.UserItem {
	if x != nil {
		return x.Sender
	}
	return nil
}

var File_cheers_post_v1_post_event_proto protoreflect.FileDescriptor

var file_cheers_post_v1_post_event_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x1a, 0x21, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0,
	0x01, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x06,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63,
	0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x69,
	0x6b, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x5e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x3c, 0x0a, 0x08, 0x4c, 0x69, 0x6b, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x3b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x3d, 0x50, 0x01,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c,
	0x61, 0x7a, 0x61, 0x72, 0x68, 0x75, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x31,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70,
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

var file_cheers_post_v1_post_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cheers_post_v1_post_event_proto_goTypes = []interface{}{
	(*PostEvent)(nil),     // 0: cheers.post.v1.PostEvent
	(*CreatePost)(nil),    // 1: cheers.post.v1.CreatePost
	(*LikePost)(nil),      // 2: cheers.post.v1.LikePost
	(*DeletePost)(nil),    // 3: cheers.post.v1.DeletePost
	(*post.Post)(nil),     // 4: cheers.type.Post
	(*user.UserItem)(nil), // 5: cheers.type.UserItem
}
var file_cheers_post_v1_post_event_proto_depIdxs = []int32{
	1, // 0: cheers.post.v1.PostEvent.create:type_name -> cheers.post.v1.CreatePost
	2, // 1: cheers.post.v1.PostEvent.like:type_name -> cheers.post.v1.LikePost
	3, // 2: cheers.post.v1.PostEvent.delete:type_name -> cheers.post.v1.DeletePost
	4, // 3: cheers.post.v1.CreatePost.post:type_name -> cheers.type.Post
	5, // 4: cheers.post.v1.CreatePost.user:type_name -> cheers.type.UserItem
	5, // 5: cheers.post.v1.DeletePost.sender:type_name -> cheers.type.UserItem
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cheers_post_v1_post_event_proto_init() }
func file_cheers_post_v1_post_event_proto_init() {
	if File_cheers_post_v1_post_event_proto != nil {
		return
	}
	file_cheers_post_v1_post_service_proto_init()
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
			switch v := v.(*CreatePost); i {
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
		file_cheers_post_v1_post_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikePost); i {
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
		file_cheers_post_v1_post_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePost); i {
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
	file_cheers_post_v1_post_event_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PostEvent_Create)(nil),
		(*PostEvent_Like)(nil),
		(*PostEvent_Delete)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cheers_post_v1_post_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cheers_post_v1_post_event_proto_goTypes,
		DependencyIndexes: file_cheers_post_v1_post_event_proto_depIdxs,
		MessageInfos:      file_cheers_post_v1_post_event_proto_msgTypes,
	}.Build()
	File_cheers_post_v1_post_event_proto = out.File
	file_cheers_post_v1_post_event_proto_rawDesc = nil
	file_cheers_post_v1_post_event_proto_goTypes = nil
	file_cheers_post_v1_post_event_proto_depIdxs = nil
}
