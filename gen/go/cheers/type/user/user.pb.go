// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: cheers/type/user/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StoryState int32

const (
	StoryState_EMPTY    StoryState = 0
	StoryState_SEEN     StoryState = 1
	StoryState_NOT_SEEN StoryState = 2
	StoryState_LOADING  StoryState = 3
)

// Enum value maps for StoryState.
var (
	StoryState_name = map[int32]string{
		0: "EMPTY",
		1: "SEEN",
		2: "NOT_SEEN",
		3: "LOADING",
	}
	StoryState_value = map[string]int32{
		"EMPTY":    0,
		"SEEN":     1,
		"NOT_SEEN": 2,
		"LOADING":  3,
	}
)

func (x StoryState) Enum() *StoryState {
	p := new(StoryState)
	*p = x
	return p
}

func (x StoryState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StoryState) Descriptor() protoreflect.EnumDescriptor {
	return file_cheers_type_user_user_proto_enumTypes[0].Descriptor()
}

func (StoryState) Type() protoreflect.EnumType {
	return &file_cheers_type_user_user_proto_enumTypes[0]
}

func (x StoryState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StoryState.Descriptor instead.
func (StoryState) EnumDescriptor() ([]byte, []int) {
	return file_cheers_type_user_user_proto_rawDescGZIP(), []int{0}
}

type Gender int32

const (
	Gender_MALE   Gender = 0
	Gender_FEMALE Gender = 1
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "MALE",
		1: "FEMALE",
	}
	Gender_value = map[string]int32{
		"MALE":   0,
		"FEMALE": 1,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_cheers_type_user_user_proto_enumTypes[1].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_cheers_type_user_user_proto_enumTypes[1]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_cheers_type_user_user_proto_rawDescGZIP(), []int{1}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email              string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Verified           bool                   `protobuf:"varint,4,opt,name=verified,proto3" json:"verified,omitempty"`
	Username           string                 `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Picture            string                 `protobuf:"bytes,6,opt,name=picture,proto3" json:"picture,omitempty"`
	Bio                string                 `protobuf:"bytes,7,opt,name=bio,proto3" json:"bio,omitempty"`
	Website            string                 `protobuf:"bytes,8,opt,name=website,proto3" json:"website,omitempty"`
	Birthday           string                 `protobuf:"bytes,12,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender             Gender                 `protobuf:"varint,13,opt,name=gender,proto3,enum=cheers.type.Gender" json:"gender,omitempty"`
	PhoneNumber        string                 `protobuf:"bytes,9,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	CreateTime         *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	RegistrationTokens []string               `protobuf:"bytes,11,rep,name=registration_tokens,json=registrationTokens,proto3" json:"registration_tokens,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_type_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_type_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_cheers_type_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *User) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *User) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *User) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *User) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_MALE
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *User) GetRegistrationTokens() []string {
	if x != nil {
		return x.RegistrationTokens
	}
	return nil
}

type UserItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Username    string     `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Verified    bool       `protobuf:"varint,4,opt,name=verified,proto3" json:"verified,omitempty"`
	Picture     string     `protobuf:"bytes,5,opt,name=picture,proto3" json:"picture,omitempty"`
	HasFollowed bool       `protobuf:"varint,6,opt,name=has_followed,json=hasFollowed,proto3" json:"has_followed,omitempty"`
	StoryState  StoryState `protobuf:"varint,7,opt,name=story_state,json=storyState,proto3,enum=cheers.type.StoryState" json:"story_state,omitempty"`
}

func (x *UserItem) Reset() {
	*x = UserItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_type_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserItem) ProtoMessage() {}

func (x *UserItem) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_type_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserItem.ProtoReflect.Descriptor instead.
func (*UserItem) Descriptor() ([]byte, []int) {
	return file_cheers_type_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserItem) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserItem) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *UserItem) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *UserItem) GetHasFollowed() bool {
	if x != nil {
		return x.HasFollowed
	}
	return false
}

func (x *UserItem) GetStoryState() StoryState {
	if x != nil {
		return x.StoryState
	}
	return StoryState_EMPTY
}

var File_cheers_type_user_user_proto protoreflect.FileDescriptor

var file_cheers_type_user_user_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63,
	0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x03, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62,
	0x69, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x61, 0x73,
	0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x68, 0x61, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x0b,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x3c, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x53, 0x45, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x54,
	0x5f, 0x53, 0x45, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x4f, 0x41, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x2a, 0x1e, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x08,
	0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41,
	0x4c, 0x45, 0x10, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x61, 0x7a, 0x61, 0x72, 0x68, 0x75, 0x67, 0x6f, 0x2f, 0x63,
	0x68, 0x65, 0x65, 0x72, 0x73, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68,
	0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75,
	0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cheers_type_user_user_proto_rawDescOnce sync.Once
	file_cheers_type_user_user_proto_rawDescData = file_cheers_type_user_user_proto_rawDesc
)

func file_cheers_type_user_user_proto_rawDescGZIP() []byte {
	file_cheers_type_user_user_proto_rawDescOnce.Do(func() {
		file_cheers_type_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_cheers_type_user_user_proto_rawDescData)
	})
	return file_cheers_type_user_user_proto_rawDescData
}

var file_cheers_type_user_user_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cheers_type_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cheers_type_user_user_proto_goTypes = []interface{}{
	(StoryState)(0),               // 0: cheers.type.StoryState
	(Gender)(0),                   // 1: cheers.type.Gender
	(*User)(nil),                  // 2: cheers.type.User
	(*UserItem)(nil),              // 3: cheers.type.UserItem
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_cheers_type_user_user_proto_depIdxs = []int32{
	1, // 0: cheers.type.User.gender:type_name -> cheers.type.Gender
	4, // 1: cheers.type.User.create_time:type_name -> google.protobuf.Timestamp
	0, // 2: cheers.type.UserItem.story_state:type_name -> cheers.type.StoryState
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cheers_type_user_user_proto_init() }
func file_cheers_type_user_user_proto_init() {
	if File_cheers_type_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cheers_type_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_cheers_type_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserItem); i {
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
			RawDescriptor: file_cheers_type_user_user_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cheers_type_user_user_proto_goTypes,
		DependencyIndexes: file_cheers_type_user_user_proto_depIdxs,
		EnumInfos:         file_cheers_type_user_user_proto_enumTypes,
		MessageInfos:      file_cheers_type_user_user_proto_msgTypes,
	}.Build()
	File_cheers_type_user_user_proto = out.File
	file_cheers_type_user_user_proto_rawDesc = nil
	file_cheers_type_user_user_proto_goTypes = nil
	file_cheers_type_user_user_proto_depIdxs = nil
}
