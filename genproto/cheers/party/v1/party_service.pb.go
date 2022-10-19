// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: cheers/party/v1/party_service.proto

package party

import (
	party "github.com/salazarhugo/cheers1/genproto/cheers/type/party"
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

type PartyUserResponse int32

const (
	PartyUserResponse_GOING          PartyUserResponse = 0
	PartyUserResponse_INTERESTED     PartyUserResponse = 1
	PartyUserResponse_NOT_INTERESTED PartyUserResponse = 2
)

// Enum value maps for PartyUserResponse.
var (
	PartyUserResponse_name = map[int32]string{
		0: "GOING",
		1: "INTERESTED",
		2: "NOT_INTERESTED",
	}
	PartyUserResponse_value = map[string]int32{
		"GOING":          0,
		"INTERESTED":     1,
		"NOT_INTERESTED": 2,
	}
)

func (x PartyUserResponse) Enum() *PartyUserResponse {
	p := new(PartyUserResponse)
	*p = x
	return p
}

func (x PartyUserResponse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PartyUserResponse) Descriptor() protoreflect.EnumDescriptor {
	return file_cheers_party_v1_party_service_proto_enumTypes[0].Descriptor()
}

func (PartyUserResponse) Type() protoreflect.EnumType {
	return &file_cheers_party_v1_party_service_proto_enumTypes[0]
}

func (x PartyUserResponse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PartyUserResponse.Descriptor instead.
func (PartyUserResponse) EnumDescriptor() ([]byte, []int) {
	return file_cheers_party_v1_party_service_proto_rawDescGZIP(), []int{0}
}

type CreatePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Party *party.Party `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
}

func (x *CreatePartyRequest) Reset() {
	*x = CreatePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_party_v1_party_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePartyRequest) ProtoMessage() {}

func (x *CreatePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_party_v1_party_service_proto_msgTypes[0]
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
	return file_cheers_party_v1_party_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePartyRequest) GetParty() *party.Party {
	if x != nil {
		return x.Party
	}
	return nil
}

type GetPartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPartyRequest) Reset() {
	*x = GetPartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_party_v1_party_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartyRequest) ProtoMessage() {}

func (x *GetPartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_party_v1_party_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartyRequest.ProtoReflect.Descriptor instead.
func (*GetPartyRequest) Descriptor() ([]byte, []int) {
	return file_cheers_party_v1_party_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetPartyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdatePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Party *party.Party `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
}

func (x *UpdatePartyRequest) Reset() {
	*x = UpdatePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_party_v1_party_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePartyRequest) ProtoMessage() {}

func (x *UpdatePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_party_v1_party_service_proto_msgTypes[2]
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
	return file_cheers_party_v1_party_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePartyRequest) GetParty() *party.Party {
	if x != nil {
		return x.Party
	}
	return nil
}

type DeletePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePartyRequest) Reset() {
	*x = DeletePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_party_v1_party_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePartyRequest) ProtoMessage() {}

func (x *DeletePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_party_v1_party_service_proto_msgTypes[3]
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
	return file_cheers_party_v1_party_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePartyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FeedPartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent    string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *FeedPartyRequest) Reset() {
	*x = FeedPartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_party_v1_party_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedPartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedPartyRequest) ProtoMessage() {}

func (x *FeedPartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_party_v1_party_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedPartyRequest.ProtoReflect.Descriptor instead.
func (*FeedPartyRequest) Descriptor() ([]byte, []int) {
	return file_cheers_party_v1_party_service_proto_rawDescGZIP(), []int{4}
}

func (x *FeedPartyRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *FeedPartyRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FeedPartyRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type FeedPartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parties       []*PartyResponse `protobuf:"bytes,1,rep,name=parties,proto3" json:"parties,omitempty"`
	NextPageToken string           `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *FeedPartyResponse) Reset() {
	*x = FeedPartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_party_v1_party_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedPartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedPartyResponse) ProtoMessage() {}

func (x *FeedPartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_party_v1_party_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedPartyResponse.ProtoReflect.Descriptor instead.
func (*FeedPartyResponse) Descriptor() ([]byte, []int) {
	return file_cheers_party_v1_party_service_proto_rawDescGZIP(), []int{5}
}

func (x *FeedPartyResponse) GetParties() []*PartyResponse {
	if x != nil {
		return x.Parties
	}
	return nil
}

func (x *FeedPartyResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type PartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Party           *party.Party      `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
	Creator         *user.User        `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	GoingCount      int64             `protobuf:"varint,3,opt,name=going_count,json=goingCount,proto3" json:"going_count,omitempty"`
	InterestedCount int64             `protobuf:"varint,4,opt,name=interested_count,json=interestedCount,proto3" json:"interested_count,omitempty"`
	InvitedCount    int64             `protobuf:"varint,5,opt,name=invited_count,json=invitedCount,proto3" json:"invited_count,omitempty"`
	IsGoing         bool              `protobuf:"varint,6,opt,name=is_going,json=isGoing,proto3" json:"is_going,omitempty"`
	IsInterested    bool              `protobuf:"varint,7,opt,name=is_interested,json=isInterested,proto3" json:"is_interested,omitempty"`
	IsCreator       bool              `protobuf:"varint,8,opt,name=is_creator,json=isCreator,proto3" json:"is_creator,omitempty"`
	UserResponse    PartyUserResponse `protobuf:"varint,9,opt,name=user_response,json=userResponse,proto3,enum=cheers.party.v1.PartyUserResponse" json:"user_response,omitempty"`
}

func (x *PartyResponse) Reset() {
	*x = PartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_party_v1_party_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyResponse) ProtoMessage() {}

func (x *PartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_party_v1_party_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyResponse.ProtoReflect.Descriptor instead.
func (*PartyResponse) Descriptor() ([]byte, []int) {
	return file_cheers_party_v1_party_service_proto_rawDescGZIP(), []int{6}
}

func (x *PartyResponse) GetParty() *party.Party {
	if x != nil {
		return x.Party
	}
	return nil
}

func (x *PartyResponse) GetCreator() *user.User {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *PartyResponse) GetGoingCount() int64 {
	if x != nil {
		return x.GoingCount
	}
	return 0
}

func (x *PartyResponse) GetInterestedCount() int64 {
	if x != nil {
		return x.InterestedCount
	}
	return 0
}

func (x *PartyResponse) GetInvitedCount() int64 {
	if x != nil {
		return x.InvitedCount
	}
	return 0
}

func (x *PartyResponse) GetIsGoing() bool {
	if x != nil {
		return x.IsGoing
	}
	return false
}

func (x *PartyResponse) GetIsInterested() bool {
	if x != nil {
		return x.IsInterested
	}
	return false
}

func (x *PartyResponse) GetIsCreator() bool {
	if x != nil {
		return x.IsCreator
	}
	return false
}

func (x *PartyResponse) GetUserResponse() PartyUserResponse {
	if x != nil {
		return x.UserResponse
	}
	return PartyUserResponse_GOING
}

var File_cheers_party_v1_party_service_proto protoreflect.FileDescriptor

var file_cheers_party_v1_party_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x22,
	0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x3e, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x05, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x66, 0x0a, 0x10, 0x46, 0x65, 0x65, 0x64,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x75, 0x0a, 0x11, 0x46, 0x65, 0x65, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xff, 0x02, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x05, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x73, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x47, 0x6f, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x47, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x42, 0x0a, 0x11, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x47, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x54,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0x8c, 0x03,
	0x0a, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x23, 0x2e,
	0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x4c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x12, 0x20, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x12, 0x23, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x23, 0x2e, 0x63, 0x68,
	0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x52, 0x0a, 0x09, 0x46, 0x65, 0x65, 0x64,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x41, 0x50, 0x01,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c,
	0x61, 0x7a, 0x61, 0x72, 0x68, 0x75, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x31,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73,
	0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x61, 0x72, 0x74, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cheers_party_v1_party_service_proto_rawDescOnce sync.Once
	file_cheers_party_v1_party_service_proto_rawDescData = file_cheers_party_v1_party_service_proto_rawDesc
)

func file_cheers_party_v1_party_service_proto_rawDescGZIP() []byte {
	file_cheers_party_v1_party_service_proto_rawDescOnce.Do(func() {
		file_cheers_party_v1_party_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cheers_party_v1_party_service_proto_rawDescData)
	})
	return file_cheers_party_v1_party_service_proto_rawDescData
}

var file_cheers_party_v1_party_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cheers_party_v1_party_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cheers_party_v1_party_service_proto_goTypes = []interface{}{
	(PartyUserResponse)(0),     // 0: cheers.party.v1.PartyUserResponse
	(*CreatePartyRequest)(nil), // 1: cheers.party.v1.CreatePartyRequest
	(*GetPartyRequest)(nil),    // 2: cheers.party.v1.GetPartyRequest
	(*UpdatePartyRequest)(nil), // 3: cheers.party.v1.UpdatePartyRequest
	(*DeletePartyRequest)(nil), // 4: cheers.party.v1.DeletePartyRequest
	(*FeedPartyRequest)(nil),   // 5: cheers.party.v1.FeedPartyRequest
	(*FeedPartyResponse)(nil),  // 6: cheers.party.v1.FeedPartyResponse
	(*PartyResponse)(nil),      // 7: cheers.party.v1.PartyResponse
	(*party.Party)(nil),        // 8: cheers.type.Party
	(*user.User)(nil),          // 9: cheers.type.User
	(*emptypb.Empty)(nil),      // 10: google.protobuf.Empty
}
var file_cheers_party_v1_party_service_proto_depIdxs = []int32{
	8,  // 0: cheers.party.v1.CreatePartyRequest.party:type_name -> cheers.type.Party
	8,  // 1: cheers.party.v1.UpdatePartyRequest.party:type_name -> cheers.type.Party
	7,  // 2: cheers.party.v1.FeedPartyResponse.parties:type_name -> cheers.party.v1.PartyResponse
	8,  // 3: cheers.party.v1.PartyResponse.party:type_name -> cheers.type.Party
	9,  // 4: cheers.party.v1.PartyResponse.creator:type_name -> cheers.type.User
	0,  // 5: cheers.party.v1.PartyResponse.user_response:type_name -> cheers.party.v1.PartyUserResponse
	1,  // 6: cheers.party.v1.PartyService.CreateParty:input_type -> cheers.party.v1.CreatePartyRequest
	2,  // 7: cheers.party.v1.PartyService.GetParty:input_type -> cheers.party.v1.GetPartyRequest
	3,  // 8: cheers.party.v1.PartyService.UpdateParty:input_type -> cheers.party.v1.UpdatePartyRequest
	4,  // 9: cheers.party.v1.PartyService.DeleteParty:input_type -> cheers.party.v1.DeletePartyRequest
	5,  // 10: cheers.party.v1.PartyService.FeedParty:input_type -> cheers.party.v1.FeedPartyRequest
	8,  // 11: cheers.party.v1.PartyService.CreateParty:output_type -> cheers.type.Party
	7,  // 12: cheers.party.v1.PartyService.GetParty:output_type -> cheers.party.v1.PartyResponse
	8,  // 13: cheers.party.v1.PartyService.UpdateParty:output_type -> cheers.type.Party
	10, // 14: cheers.party.v1.PartyService.DeleteParty:output_type -> google.protobuf.Empty
	6,  // 15: cheers.party.v1.PartyService.FeedParty:output_type -> cheers.party.v1.FeedPartyResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_cheers_party_v1_party_service_proto_init() }
func file_cheers_party_v1_party_service_proto_init() {
	if File_cheers_party_v1_party_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cheers_party_v1_party_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_cheers_party_v1_party_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartyRequest); i {
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
		file_cheers_party_v1_party_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_cheers_party_v1_party_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_cheers_party_v1_party_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedPartyRequest); i {
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
		file_cheers_party_v1_party_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedPartyResponse); i {
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
		file_cheers_party_v1_party_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyResponse); i {
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
			RawDescriptor: file_cheers_party_v1_party_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cheers_party_v1_party_service_proto_goTypes,
		DependencyIndexes: file_cheers_party_v1_party_service_proto_depIdxs,
		EnumInfos:         file_cheers_party_v1_party_service_proto_enumTypes,
		MessageInfos:      file_cheers_party_v1_party_service_proto_msgTypes,
	}.Build()
	File_cheers_party_v1_party_service_proto = out.File
	file_cheers_party_v1_party_service_proto_rawDesc = nil
	file_cheers_party_v1_party_service_proto_goTypes = nil
	file_cheers_party_v1_party_service_proto_depIdxs = nil
}
