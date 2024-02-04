// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: cheers/note/v1/note_service.proto

package note

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

type ListFriendNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListFriendNoteRequest) Reset() {
	*x = ListFriendNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_note_v1_note_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendNoteRequest) ProtoMessage() {}

func (x *ListFriendNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_note_v1_note_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendNoteRequest.ProtoReflect.Descriptor instead.
func (*ListFriendNoteRequest) Descriptor() ([]byte, []int) {
	return file_cheers_note_v1_note_service_proto_rawDescGZIP(), []int{0}
}

type ListFriendNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Note `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListFriendNoteResponse) Reset() {
	*x = ListFriendNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_note_v1_note_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendNoteResponse) ProtoMessage() {}

func (x *ListFriendNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_note_v1_note_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendNoteResponse.ProtoReflect.Descriptor instead.
func (*ListFriendNoteResponse) Descriptor() ([]byte, []int) {
	return file_cheers_note_v1_note_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListFriendNoteResponse) GetItems() []*Note {
	if x != nil {
		return x.Items
	}
	return nil
}

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Picture  string `protobuf:"bytes,5,opt,name=picture,proto3" json:"picture,omitempty"`
	Created  int64  `protobuf:"varint,6,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_note_v1_note_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_note_v1_note_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_cheers_note_v1_note_service_proto_rawDescGZIP(), []int{2}
}

func (x *Note) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Note) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Note) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Note) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Note) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Note) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

type CreateNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateNoteRequest) Reset() {
	*x = CreateNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_note_v1_note_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteRequest) ProtoMessage() {}

func (x *CreateNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_note_v1_note_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteRequest.ProtoReflect.Descriptor instead.
func (*CreateNoteRequest) Descriptor() ([]byte, []int) {
	return file_cheers_note_v1_note_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNoteRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CreateNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note *Note `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *CreateNoteResponse) Reset() {
	*x = CreateNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_note_v1_note_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteResponse) ProtoMessage() {}

func (x *CreateNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_note_v1_note_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteResponse.ProtoReflect.Descriptor instead.
func (*CreateNoteResponse) Descriptor() ([]byte, []int) {
	return file_cheers_note_v1_note_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateNoteResponse) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type DeleteNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *DeleteNoteRequest) Reset() {
	*x = DeleteNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_note_v1_note_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNoteRequest) ProtoMessage() {}

func (x *DeleteNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_note_v1_note_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNoteRequest.ProtoReflect.Descriptor instead.
func (*DeleteNoteRequest) Descriptor() ([]byte, []int) {
	return file_cheers_note_v1_note_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteNoteRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type DeleteNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteNoteResponse) Reset() {
	*x = DeleteNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_note_v1_note_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNoteResponse) ProtoMessage() {}

func (x *DeleteNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_note_v1_note_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNoteResponse.ProtoReflect.Descriptor instead.
func (*DeleteNoteResponse) Descriptor() ([]byte, []int) {
	return file_cheers_note_v1_note_service_proto_rawDescGZIP(), []int{6}
}

var File_cheers_note_v1_note_service_proto protoreflect.FileDescriptor

var file_cheers_note_v1_note_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4e,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x16, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x97, 0x01, 0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x3e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73,
	0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x14, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xdc, 0x02, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x69, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65,
	0x12, 0x21, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a,
	0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x66, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x68,
	0x65, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x2a, 0x09, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x7a, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73,
	0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x73, 0x42, 0x3d, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x61, 0x7a, 0x61, 0x72, 0x68, 0x75, 0x67, 0x6f, 0x2f, 0x63, 0x68,
	0x65, 0x65, 0x72, 0x73, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65,
	0x65, 0x72, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x6f, 0x74, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cheers_note_v1_note_service_proto_rawDescOnce sync.Once
	file_cheers_note_v1_note_service_proto_rawDescData = file_cheers_note_v1_note_service_proto_rawDesc
)

func file_cheers_note_v1_note_service_proto_rawDescGZIP() []byte {
	file_cheers_note_v1_note_service_proto_rawDescOnce.Do(func() {
		file_cheers_note_v1_note_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cheers_note_v1_note_service_proto_rawDescData)
	})
	return file_cheers_note_v1_note_service_proto_rawDescData
}

var file_cheers_note_v1_note_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cheers_note_v1_note_service_proto_goTypes = []interface{}{
	(*ListFriendNoteRequest)(nil),  // 0: cheers.note.v1.ListFriendNoteRequest
	(*ListFriendNoteResponse)(nil), // 1: cheers.note.v1.ListFriendNoteResponse
	(*Note)(nil),                   // 2: cheers.note.v1.Note
	(*CreateNoteRequest)(nil),      // 3: cheers.note.v1.CreateNoteRequest
	(*CreateNoteResponse)(nil),     // 4: cheers.note.v1.CreateNoteResponse
	(*DeleteNoteRequest)(nil),      // 5: cheers.note.v1.DeleteNoteRequest
	(*DeleteNoteResponse)(nil),     // 6: cheers.note.v1.DeleteNoteResponse
}
var file_cheers_note_v1_note_service_proto_depIdxs = []int32{
	2, // 0: cheers.note.v1.ListFriendNoteResponse.items:type_name -> cheers.note.v1.Note
	2, // 1: cheers.note.v1.CreateNoteResponse.note:type_name -> cheers.note.v1.Note
	3, // 2: cheers.note.v1.NoteService.CreateNote:input_type -> cheers.note.v1.CreateNoteRequest
	5, // 3: cheers.note.v1.NoteService.DeleteNote:input_type -> cheers.note.v1.DeleteNoteRequest
	0, // 4: cheers.note.v1.NoteService.ListFriendNote:input_type -> cheers.note.v1.ListFriendNoteRequest
	4, // 5: cheers.note.v1.NoteService.CreateNote:output_type -> cheers.note.v1.CreateNoteResponse
	6, // 6: cheers.note.v1.NoteService.DeleteNote:output_type -> cheers.note.v1.DeleteNoteResponse
	1, // 7: cheers.note.v1.NoteService.ListFriendNote:output_type -> cheers.note.v1.ListFriendNoteResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cheers_note_v1_note_service_proto_init() }
func file_cheers_note_v1_note_service_proto_init() {
	if File_cheers_note_v1_note_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cheers_note_v1_note_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFriendNoteRequest); i {
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
		file_cheers_note_v1_note_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFriendNoteResponse); i {
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
		file_cheers_note_v1_note_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
		file_cheers_note_v1_note_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNoteRequest); i {
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
		file_cheers_note_v1_note_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNoteResponse); i {
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
		file_cheers_note_v1_note_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNoteRequest); i {
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
		file_cheers_note_v1_note_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNoteResponse); i {
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
			RawDescriptor: file_cheers_note_v1_note_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cheers_note_v1_note_service_proto_goTypes,
		DependencyIndexes: file_cheers_note_v1_note_service_proto_depIdxs,
		MessageInfos:      file_cheers_note_v1_note_service_proto_msgTypes,
	}.Build()
	File_cheers_note_v1_note_service_proto = out.File
	file_cheers_note_v1_note_service_proto_rawDesc = nil
	file_cheers_note_v1_note_service_proto_goTypes = nil
	file_cheers_note_v1_note_service_proto_depIdxs = nil
}
