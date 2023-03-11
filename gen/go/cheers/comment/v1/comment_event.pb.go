// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: cheers/comment/v1/comment_event.proto

package comment

import (
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

type CommentEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*CommentEvent_Created
	//	*CommentEvent_Deleted
	//	*CommentEvent_Liked
	Event isCommentEvent_Event `protobuf_oneof:"event"`
}

func (x *CommentEvent) Reset() {
	*x = CommentEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_comment_v1_comment_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentEvent) ProtoMessage() {}

func (x *CommentEvent) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_comment_v1_comment_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentEvent.ProtoReflect.Descriptor instead.
func (*CommentEvent) Descriptor() ([]byte, []int) {
	return file_cheers_comment_v1_comment_event_proto_rawDescGZIP(), []int{0}
}

func (m *CommentEvent) GetEvent() isCommentEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *CommentEvent) GetCreated() *CreatedComment {
	if x, ok := x.GetEvent().(*CommentEvent_Created); ok {
		return x.Created
	}
	return nil
}

func (x *CommentEvent) GetDeleted() *DeletedComment {
	if x, ok := x.GetEvent().(*CommentEvent_Deleted); ok {
		return x.Deleted
	}
	return nil
}

func (x *CommentEvent) GetLiked() *LikedComment {
	if x, ok := x.GetEvent().(*CommentEvent_Liked); ok {
		return x.Liked
	}
	return nil
}

type isCommentEvent_Event interface {
	isCommentEvent_Event()
}

type CommentEvent_Created struct {
	Created *CreatedComment `protobuf:"bytes,1,opt,name=created,proto3,oneof"`
}

type CommentEvent_Deleted struct {
	Deleted *DeletedComment `protobuf:"bytes,2,opt,name=deleted,proto3,oneof"`
}

type CommentEvent_Liked struct {
	Liked *LikedComment `protobuf:"bytes,3,opt,name=liked,proto3,oneof"`
}

func (*CommentEvent_Created) isCommentEvent_Event() {}

func (*CommentEvent_Deleted) isCommentEvent_Event() {}

func (*CommentEvent_Liked) isCommentEvent_Event() {}

type LikedComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId string `protobuf:"bytes,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *LikedComment) Reset() {
	*x = LikedComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_comment_v1_comment_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikedComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikedComment) ProtoMessage() {}

func (x *LikedComment) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_comment_v1_comment_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikedComment.ProtoReflect.Descriptor instead.
func (*LikedComment) Descriptor() ([]byte, []int) {
	return file_cheers_comment_v1_comment_event_proto_rawDescGZIP(), []int{1}
}

func (x *LikedComment) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

func (x *LikedComment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreatedComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment  *Comment       `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	User     *user.UserItem `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Mentions []string       `protobuf:"bytes,3,rep,name=mentions,proto3" json:"mentions,omitempty"`
}

func (x *CreatedComment) Reset() {
	*x = CreatedComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_comment_v1_comment_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedComment) ProtoMessage() {}

func (x *CreatedComment) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_comment_v1_comment_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedComment.ProtoReflect.Descriptor instead.
func (*CreatedComment) Descriptor() ([]byte, []int) {
	return file_cheers_comment_v1_comment_event_proto_rawDescGZIP(), []int{2}
}

func (x *CreatedComment) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *CreatedComment) GetUser() *user.UserItem {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CreatedComment) GetMentions() []string {
	if x != nil {
		return x.Mentions
	}
	return nil
}

type DeletedComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment       `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	User    *user.UserItem `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *DeletedComment) Reset() {
	*x = DeletedComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_comment_v1_comment_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletedComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletedComment) ProtoMessage() {}

func (x *DeletedComment) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_comment_v1_comment_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletedComment.ProtoReflect.Descriptor instead.
func (*DeletedComment) Descriptor() ([]byte, []int) {
	return file_cheers_comment_v1_comment_event_proto_rawDescGZIP(), []int{3}
}

func (x *DeletedComment) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *DeletedComment) GetUser() *user.UserItem {
	if x != nil {
		return x.User
	}
	return nil
}

var File_cheers_comment_v1_comment_event_proto protoreflect.FileDescriptor

var file_cheers_comment_v1_comment_event_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x63, 0x68, 0x65, 0x65,
	0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xce, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x3d, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12,
	0x37, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x48,
	0x00, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x71, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63,
	0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x29, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x42, 0x43, 0x50, 0x01,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c,
	0x61, 0x7a, 0x61, 0x72, 0x68, 0x75, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x31,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cheers_comment_v1_comment_event_proto_rawDescOnce sync.Once
	file_cheers_comment_v1_comment_event_proto_rawDescData = file_cheers_comment_v1_comment_event_proto_rawDesc
)

func file_cheers_comment_v1_comment_event_proto_rawDescGZIP() []byte {
	file_cheers_comment_v1_comment_event_proto_rawDescOnce.Do(func() {
		file_cheers_comment_v1_comment_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_cheers_comment_v1_comment_event_proto_rawDescData)
	})
	return file_cheers_comment_v1_comment_event_proto_rawDescData
}

var file_cheers_comment_v1_comment_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cheers_comment_v1_comment_event_proto_goTypes = []interface{}{
	(*CommentEvent)(nil),   // 0: cheers.comment.v1.CommentEvent
	(*LikedComment)(nil),   // 1: cheers.comment.v1.LikedComment
	(*CreatedComment)(nil), // 2: cheers.comment.v1.CreatedComment
	(*DeletedComment)(nil), // 3: cheers.comment.v1.DeletedComment
	(*Comment)(nil),        // 4: cheers.comment.v1.Comment
	(*user.UserItem)(nil),  // 5: cheers.type.UserItem
}
var file_cheers_comment_v1_comment_event_proto_depIdxs = []int32{
	2, // 0: cheers.comment.v1.CommentEvent.created:type_name -> cheers.comment.v1.CreatedComment
	3, // 1: cheers.comment.v1.CommentEvent.deleted:type_name -> cheers.comment.v1.DeletedComment
	1, // 2: cheers.comment.v1.CommentEvent.liked:type_name -> cheers.comment.v1.LikedComment
	4, // 3: cheers.comment.v1.CreatedComment.comment:type_name -> cheers.comment.v1.Comment
	5, // 4: cheers.comment.v1.CreatedComment.user:type_name -> cheers.type.UserItem
	4, // 5: cheers.comment.v1.DeletedComment.comment:type_name -> cheers.comment.v1.Comment
	5, // 6: cheers.comment.v1.DeletedComment.user:type_name -> cheers.type.UserItem
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cheers_comment_v1_comment_event_proto_init() }
func file_cheers_comment_v1_comment_event_proto_init() {
	if File_cheers_comment_v1_comment_event_proto != nil {
		return
	}
	file_cheers_comment_v1_comment_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cheers_comment_v1_comment_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentEvent); i {
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
		file_cheers_comment_v1_comment_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikedComment); i {
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
		file_cheers_comment_v1_comment_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedComment); i {
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
		file_cheers_comment_v1_comment_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletedComment); i {
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
	file_cheers_comment_v1_comment_event_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CommentEvent_Created)(nil),
		(*CommentEvent_Deleted)(nil),
		(*CommentEvent_Liked)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cheers_comment_v1_comment_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cheers_comment_v1_comment_event_proto_goTypes,
		DependencyIndexes: file_cheers_comment_v1_comment_event_proto_depIdxs,
		MessageInfos:      file_cheers_comment_v1_comment_event_proto_msgTypes,
	}.Build()
	File_cheers_comment_v1_comment_event_proto = out.File
	file_cheers_comment_v1_comment_event_proto_rawDesc = nil
	file_cheers_comment_v1_comment_event_proto_goTypes = nil
	file_cheers_comment_v1_comment_event_proto_depIdxs = nil
}
