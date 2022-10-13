// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: cheers/post/v1/post_service.proto

package post

import (
	post "github.com/salazarhugo/cheers1/genproto/cheers/type/post"
	user "github.com/salazarhugo/cheers1/genproto/cheers/type/user"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *post.Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_post_v1_post_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_post_v1_post_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_cheers_post_v1_post_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePostRequest) GetPost() *post.Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type UpdatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *post.Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *UpdatePostRequest) Reset() {
	*x = UpdatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_post_v1_post_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostRequest) ProtoMessage() {}

func (x *UpdatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_post_v1_post_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return file_cheers_post_v1_post_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePostRequest) GetPost() *post.Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type DeletePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_post_v1_post_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_post_v1_post_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_cheers_post_v1_post_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeletePostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent    string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListPostRequest) Reset() {
	*x = ListPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_post_v1_post_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostRequest) ProtoMessage() {}

func (x *ListPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_post_v1_post_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostRequest.ProtoReflect.Descriptor instead.
func (*ListPostRequest) Descriptor() ([]byte, []int) {
	return file_cheers_post_v1_post_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListPostRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListPostRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPostRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts         []*PostResponse `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	NextPageToken string          `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPostResponse) Reset() {
	*x = ListPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_post_v1_post_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostResponse) ProtoMessage() {}

func (x *ListPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_post_v1_post_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostResponse.ProtoReflect.Descriptor instead.
func (*ListPostResponse) Descriptor() ([]byte, []int) {
	return file_cheers_post_v1_post_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListPostResponse) GetPosts() []*PostResponse {
	if x != nil {
		return x.Posts
	}
	return nil
}

func (x *ListPostResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type PostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post            *post.Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	User            *user.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	LikeCount       int64      `protobuf:"varint,3,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	CommentCount    int64      `protobuf:"varint,4,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	HasLiked        bool       `protobuf:"varint,5,opt,name=has_liked,json=hasLiked,proto3" json:"has_liked,omitempty"`
	CommentsEnabled bool       `protobuf:"varint,6,opt,name=comments_enabled,json=commentsEnabled,proto3" json:"comments_enabled,omitempty"`
	ReshareEnabled  bool       `protobuf:"varint,8,opt,name=reshare_enabled,json=reshareEnabled,proto3" json:"reshare_enabled,omitempty"`
	MediaCount      int64      `protobuf:"varint,7,opt,name=media_count,json=mediaCount,proto3" json:"media_count,omitempty"`
}

func (x *PostResponse) Reset() {
	*x = PostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_post_v1_post_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostResponse) ProtoMessage() {}

func (x *PostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_post_v1_post_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostResponse.ProtoReflect.Descriptor instead.
func (*PostResponse) Descriptor() ([]byte, []int) {
	return file_cheers_post_v1_post_service_proto_rawDescGZIP(), []int{5}
}

func (x *PostResponse) GetPost() *post.Post {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *PostResponse) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PostResponse) GetLikeCount() int64 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *PostResponse) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *PostResponse) GetHasLiked() bool {
	if x != nil {
		return x.HasLiked
	}
	return false
}

func (x *PostResponse) GetCommentsEnabled() bool {
	if x != nil {
		return x.CommentsEnabled
	}
	return false
}

func (x *PostResponse) GetReshareEnabled() bool {
	if x != nil {
		return x.ReshareEnabled
	}
	return false
}

func (x *PostResponse) GetMediaCount() int64 {
	if x != nil {
		return x.MediaCount
	}
	return 0
}

var File_cheers_post_v1_post_service_proto protoreflect.FileDescriptor

var file_cheers_post_v1_post_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x70,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x68, 0x65, 0x65,
	0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x65, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6e,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb2,
	0x02, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x12, 0x29,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x32, 0xad, 0x02, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x3f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x61, 0x7a, 0x61, 0x72, 0x68, 0x75, 0x67, 0x6f, 0x2f,
	0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cheers_post_v1_post_service_proto_rawDescOnce sync.Once
	file_cheers_post_v1_post_service_proto_rawDescData = file_cheers_post_v1_post_service_proto_rawDesc
)

func file_cheers_post_v1_post_service_proto_rawDescGZIP() []byte {
	file_cheers_post_v1_post_service_proto_rawDescOnce.Do(func() {
		file_cheers_post_v1_post_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cheers_post_v1_post_service_proto_rawDescData)
	})
	return file_cheers_post_v1_post_service_proto_rawDescData
}

var file_cheers_post_v1_post_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cheers_post_v1_post_service_proto_goTypes = []interface{}{
	(*CreatePostRequest)(nil), // 0: cheers.post.v1.CreatePostRequest
	(*UpdatePostRequest)(nil), // 1: cheers.post.v1.UpdatePostRequest
	(*DeletePostRequest)(nil), // 2: cheers.post.v1.DeletePostRequest
	(*ListPostRequest)(nil),   // 3: cheers.post.v1.ListPostRequest
	(*ListPostResponse)(nil),  // 4: cheers.post.v1.ListPostResponse
	(*PostResponse)(nil),      // 5: cheers.post.v1.PostResponse
	(*post.Post)(nil),         // 6: cheers.type.Post
	(*user.User)(nil),         // 7: cheers.type.User
	(*emptypb.Empty)(nil),     // 8: google.protobuf.Empty
}
var file_cheers_post_v1_post_service_proto_depIdxs = []int32{
	6, // 0: cheers.post.v1.CreatePostRequest.post:type_name -> cheers.type.Post
	6, // 1: cheers.post.v1.UpdatePostRequest.post:type_name -> cheers.type.Post
	5, // 2: cheers.post.v1.ListPostResponse.posts:type_name -> cheers.post.v1.PostResponse
	6, // 3: cheers.post.v1.PostResponse.post:type_name -> cheers.type.Post
	7, // 4: cheers.post.v1.PostResponse.user:type_name -> cheers.type.User
	0, // 5: cheers.post.v1.PostService.CreatePost:input_type -> cheers.post.v1.CreatePostRequest
	3, // 6: cheers.post.v1.PostService.ListPost:input_type -> cheers.post.v1.ListPostRequest
	1, // 7: cheers.post.v1.PostService.UpdatePost:input_type -> cheers.post.v1.UpdatePostRequest
	2, // 8: cheers.post.v1.PostService.DeletePost:input_type -> cheers.post.v1.DeletePostRequest
	6, // 9: cheers.post.v1.PostService.CreatePost:output_type -> cheers.type.Post
	4, // 10: cheers.post.v1.PostService.ListPost:output_type -> cheers.post.v1.ListPostResponse
	6, // 11: cheers.post.v1.PostService.UpdatePost:output_type -> cheers.type.Post
	8, // 12: cheers.post.v1.PostService.DeletePost:output_type -> google.protobuf.Empty
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cheers_post_v1_post_service_proto_init() }
func file_cheers_post_v1_post_service_proto_init() {
	if File_cheers_post_v1_post_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cheers_post_v1_post_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
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
		file_cheers_post_v1_post_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostRequest); i {
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
		file_cheers_post_v1_post_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostRequest); i {
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
		file_cheers_post_v1_post_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostRequest); i {
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
		file_cheers_post_v1_post_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostResponse); i {
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
		file_cheers_post_v1_post_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostResponse); i {
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
			RawDescriptor: file_cheers_post_v1_post_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cheers_post_v1_post_service_proto_goTypes,
		DependencyIndexes: file_cheers_post_v1_post_service_proto_depIdxs,
		MessageInfos:      file_cheers_post_v1_post_service_proto_msgTypes,
	}.Build()
	File_cheers_post_v1_post_service_proto = out.File
	file_cheers_post_v1_post_service_proto_rawDesc = nil
	file_cheers_post_v1_post_service_proto_goTypes = nil
	file_cheers_post_v1_post_service_proto_depIdxs = nil
}
