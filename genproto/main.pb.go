// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: cheers/api/v1/main.proto

package genproto

import (
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

type DeletePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePartyRequest) Reset() {
	*x = DeletePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_api_v1_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePartyRequest) ProtoMessage() {}

func (x *DeletePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_api_v1_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePartyRequest.ProtoReflect.Descriptor instead.
func (*DeletePartyRequest) Descriptor() ([]byte, []int) {
	return file_cheers_api_v1_main_proto_rawDescGZIP(), []int{0}
}

func (x *DeletePartyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdatePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePartyRequest) Reset() {
	*x = UpdatePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_api_v1_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePartyRequest) ProtoMessage() {}

func (x *UpdatePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_api_v1_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePartyRequest.ProtoReflect.Descriptor instead.
func (*UpdatePartyRequest) Descriptor() ([]byte, []int) {
	return file_cheers_api_v1_main_proto_rawDescGZIP(), []int{1}
}

type UpdatePartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePartyResponse) Reset() {
	*x = UpdatePartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_api_v1_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePartyResponse) ProtoMessage() {}

func (x *UpdatePartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_api_v1_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePartyResponse.ProtoReflect.Descriptor instead.
func (*UpdatePartyResponse) Descriptor() ([]byte, []int) {
	return file_cheers_api_v1_main_proto_rawDescGZIP(), []int{2}
}

type CreatePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Party *Party `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
}

func (x *CreatePartyRequest) Reset() {
	*x = CreatePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_api_v1_main_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePartyRequest) ProtoMessage() {}

func (x *CreatePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_api_v1_main_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePartyRequest.ProtoReflect.Descriptor instead.
func (*CreatePartyRequest) Descriptor() ([]byte, []int) {
	return file_cheers_api_v1_main_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePartyRequest) GetParty() *Party {
	if x != nil {
		return x.Party
	}
	return nil
}

type CreatePartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePartyResponse) Reset() {
	*x = CreatePartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_api_v1_main_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePartyResponse) ProtoMessage() {}

func (x *CreatePartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_api_v1_main_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePartyResponse.ProtoReflect.Descriptor instead.
func (*CreatePartyResponse) Descriptor() ([]byte, []int) {
	return file_cheers_api_v1_main_proto_rawDescGZIP(), []int{4}
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Identification:
	//	*GetUserRequest_UserId
	//	*GetUserRequest_Username
	Identification isGetUserRequest_Identification `protobuf_oneof:"identification"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_api_v1_main_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_api_v1_main_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_cheers_api_v1_main_proto_rawDescGZIP(), []int{5}
}

func (m *GetUserRequest) GetIdentification() isGetUserRequest_Identification {
	if m != nil {
		return m.Identification
	}
	return nil
}

func (x *GetUserRequest) GetUserId() string {
	if x, ok := x.GetIdentification().(*GetUserRequest_UserId); ok {
		return x.UserId
	}
	return ""
}

func (x *GetUserRequest) GetUsername() string {
	if x, ok := x.GetIdentification().(*GetUserRequest_Username); ok {
		return x.Username
	}
	return ""
}

type isGetUserRequest_Identification interface {
	isGetUserRequest_Identification()
}

type GetUserRequest_UserId struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3,oneof"`
}

type GetUserRequest_Username struct {
	Username string `protobuf:"bytes,2,opt,name=username,proto3,oneof"`
}

func (*GetUserRequest_UserId) isGetUserRequest_Identification() {}

func (*GetUserRequest_Username) isGetUserRequest_Identification() {}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  cheers.User user = 1;
	PostCount      int64  `protobuf:"varint,2,opt,name=post_count,json=postCount,proto3" json:"post_count,omitempty"`
	FollowBack     bool   `protobuf:"varint,3,opt,name=follow_back,json=followBack,proto3" json:"follow_back,omitempty"`
	FollowingCount int64  `protobuf:"varint,4,opt,name=following_count,json=followingCount,proto3" json:"following_count,omitempty"`
	FollowersCount int64  `protobuf:"varint,5,opt,name=followers_count,json=followersCount,proto3" json:"followers_count,omitempty"`
	StoryState     string `protobuf:"bytes,6,opt,name=story_state,json=storyState,proto3" json:"story_state,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_api_v1_main_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_api_v1_main_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_cheers_api_v1_main_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserResponse) GetPostCount() int64 {
	if x != nil {
		return x.PostCount
	}
	return 0
}

func (x *GetUserResponse) GetFollowBack() bool {
	if x != nil {
		return x.FollowBack
	}
	return false
}

func (x *GetUserResponse) GetFollowingCount() int64 {
	if x != nil {
		return x.FollowingCount
	}
	return 0
}

func (x *GetUserResponse) GetFollowersCount() int64 {
	if x != nil {
		return x.FollowersCount
	}
	return 0
}

func (x *GetUserResponse) GetStoryState() string {
	if x != nil {
		return x.StoryState
	}
	return ""
}

var File_cheers_api_v1_main_proto protoreflect.FileDescriptor

var file_cheers_api_v1_main_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x68, 0x65, 0x65,
	0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x12, 0x63, 0x68, 0x65, 0x65, 0x72,
	0x73, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc4, 0x01, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x42, 0x61, 0x63, 0x6b,
	0x12, 0x27, 0x0a, 0x0f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x32, 0xdb, 0x01, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6e, 0x12, 0x48, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x28, 0x50, 0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x61, 0x7a, 0x61, 0x72, 0x68, 0x75, 0x67, 0x6f, 0x2f, 0x63, 0x68,
	0x65, 0x65, 0x72, 0x73, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cheers_api_v1_main_proto_rawDescOnce sync.Once
	file_cheers_api_v1_main_proto_rawDescData = file_cheers_api_v1_main_proto_rawDesc
)

func file_cheers_api_v1_main_proto_rawDescGZIP() []byte {
	file_cheers_api_v1_main_proto_rawDescOnce.Do(func() {
		file_cheers_api_v1_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_cheers_api_v1_main_proto_rawDescData)
	})
	return file_cheers_api_v1_main_proto_rawDescData
}

var file_cheers_api_v1_main_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cheers_api_v1_main_proto_goTypes = []interface{}{
	(*DeletePartyRequest)(nil),  // 0: cheers.api.v1.DeletePartyRequest
	(*UpdatePartyRequest)(nil),  // 1: cheers.api.v1.UpdatePartyRequest
	(*UpdatePartyResponse)(nil), // 2: cheers.api.v1.UpdatePartyResponse
	(*CreatePartyRequest)(nil),  // 3: cheers.api.v1.CreatePartyRequest
	(*CreatePartyResponse)(nil), // 4: cheers.api.v1.CreatePartyResponse
	(*GetUserRequest)(nil),      // 5: cheers.api.v1.GetUserRequest
	(*GetUserResponse)(nil),     // 6: cheers.api.v1.GetUserResponse
	(*Party)(nil),               // 7: cheers.Party
	(*emptypb.Empty)(nil),       // 8: google.protobuf.Empty
}
var file_cheers_api_v1_main_proto_depIdxs = []int32{
	7, // 0: cheers.api.v1.CreatePartyRequest.party:type_name -> cheers.Party
	5, // 1: cheers.api.v1.Main.GetUser:input_type -> cheers.api.v1.GetUserRequest
	3, // 2: cheers.api.v1.Main.CreateParty:input_type -> cheers.api.v1.CreatePartyRequest
	0, // 3: cheers.api.v1.Main.DeleteParty:input_type -> cheers.api.v1.DeletePartyRequest
	6, // 4: cheers.api.v1.Main.GetUser:output_type -> cheers.api.v1.GetUserResponse
	7, // 5: cheers.api.v1.Main.CreateParty:output_type -> cheers.Party
	8, // 6: cheers.api.v1.Main.DeleteParty:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cheers_api_v1_main_proto_init() }
func file_cheers_api_v1_main_proto_init() {
	if File_cheers_api_v1_main_proto != nil {
		return
	}
	File_cheers_party_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cheers_api_v1_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePartyRequest); i {
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
		file_cheers_api_v1_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePartyRequest); i {
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
		file_cheers_api_v1_main_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePartyResponse); i {
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
		file_cheers_api_v1_main_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePartyRequest); i {
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
		file_cheers_api_v1_main_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePartyResponse); i {
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
		file_cheers_api_v1_main_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_cheers_api_v1_main_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
	file_cheers_api_v1_main_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*GetUserRequest_UserId)(nil),
		(*GetUserRequest_Username)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cheers_api_v1_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cheers_api_v1_main_proto_goTypes,
		DependencyIndexes: file_cheers_api_v1_main_proto_depIdxs,
		MessageInfos:      file_cheers_api_v1_main_proto_msgTypes,
	}.Build()
	File_cheers_api_v1_main_proto = out.File
	file_cheers_api_v1_main_proto_rawDesc = nil
	file_cheers_api_v1_main_proto_goTypes = nil
	file_cheers_api_v1_main_proto_depIdxs = nil
}
