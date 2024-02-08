// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: cheers/type/post/post.proto

package postpb

import (
	audio "github.com/salazarhugo/cheers1/gen/go/cheers/type/audio"
	privacy "github.com/salazarhugo/cheers1/gen/go/cheers/type/privacy"
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

type MediaType int32

const (
	MediaType_IMAGE MediaType = 0
	MediaType_VIDEO MediaType = 1
)

// Enum value maps for MediaType.
var (
	MediaType_name = map[int32]string{
		0: "IMAGE",
		1: "VIDEO",
	}
	MediaType_value = map[string]int32{
		"IMAGE": 0,
		"VIDEO": 1,
	}
)

func (x MediaType) Enum() *MediaType {
	p := new(MediaType)
	*p = x
	return p
}

func (x MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_cheers_type_post_post_proto_enumTypes[0].Descriptor()
}

func (MediaType) Type() protoreflect.EnumType {
	return &file_cheers_type_post_post_proto_enumTypes[0]
}

func (x MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaType.Descriptor instead.
func (MediaType) EnumDescriptor() ([]byte, []int) {
	return file_cheers_type_post_post_proto_rawDescGZIP(), []int{0}
}

type PostRatio int32

const (
	PostRatio_RATIO_16_9 PostRatio = 0
	PostRatio_RATIO_1_1  PostRatio = 1
	PostRatio_RATIO_4_5  PostRatio = 2
)

// Enum value maps for PostRatio.
var (
	PostRatio_name = map[int32]string{
		0: "RATIO_16_9",
		1: "RATIO_1_1",
		2: "RATIO_4_5",
	}
	PostRatio_value = map[string]int32{
		"RATIO_16_9": 0,
		"RATIO_1_1":  1,
		"RATIO_4_5":  2,
	}
)

func (x PostRatio) Enum() *PostRatio {
	p := new(PostRatio)
	*p = x
	return p
}

func (x PostRatio) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PostRatio) Descriptor() protoreflect.EnumDescriptor {
	return file_cheers_type_post_post_proto_enumTypes[1].Descriptor()
}

func (PostRatio) Type() protoreflect.EnumType {
	return &file_cheers_type_post_post_proto_enumTypes[1]
}

func (x PostRatio) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PostRatio.Descriptor instead.
func (PostRatio) EnumDescriptor() ([]byte, []int) {
	return file_cheers_type_post_post_proto_rawDescGZIP(), []int{1}
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorId    string          `protobuf:"bytes,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Caption      string          `protobuf:"bytes,3,opt,name=caption,proto3" json:"caption,omitempty"`
	Address      string          `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Privacy      privacy.Privacy `protobuf:"varint,5,opt,name=privacy,proto3,enum=cheers.type.Privacy" json:"privacy,omitempty"`
	PostMedia    []*PostMedia    `protobuf:"bytes,6,rep,name=post_media,json=postMedia,proto3" json:"post_media,omitempty"`
	LocationName string          `protobuf:"bytes,8,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
	Drink        *Drink          `protobuf:"bytes,9,opt,name=drink,proto3" json:"drink,omitempty"`
	Drunkenness  int64           `protobuf:"varint,10,opt,name=drunkenness,proto3" json:"drunkenness,omitempty"`
	CreateTime   int64           `protobuf:"varint,12,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CanComment   bool            `protobuf:"varint,13,opt,name=can_comment,json=canComment,proto3" json:"can_comment,omitempty"`
	CanShare     bool            `protobuf:"varint,14,opt,name=can_share,json=canShare,proto3" json:"can_share,omitempty"`
	Ratio        PostRatio       `protobuf:"varint,15,opt,name=ratio,proto3,enum=cheers.type.PostRatio" json:"ratio,omitempty"`
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude float64 `protobuf:"fixed64,16,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude             float64      `protobuf:"fixed64,17,opt,name=longitude,proto3" json:"longitude,omitempty"`
	LastCommentText       string       `protobuf:"bytes,18,opt,name=last_comment_text,json=lastCommentText,proto3" json:"last_comment_text,omitempty"`
	LastCommentUsername   string       `protobuf:"bytes,19,opt,name=last_comment_username,json=lastCommentUsername,proto3" json:"last_comment_username,omitempty"`
	LastCommentCreateTime int64        `protobuf:"varint,20,opt,name=last_comment_create_time,json=lastCommentCreateTime,proto3" json:"last_comment_create_time,omitempty"`
	Audio                 *audio.Audio `protobuf:"bytes,21,opt,name=audio,proto3" json:"audio,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_type_post_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_type_post_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_cheers_type_post_post_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Post) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *Post) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Post) GetPrivacy() privacy.Privacy {
	if x != nil {
		return x.Privacy
	}
	return privacy.Privacy(0)
}

func (x *Post) GetPostMedia() []*PostMedia {
	if x != nil {
		return x.PostMedia
	}
	return nil
}

func (x *Post) GetLocationName() string {
	if x != nil {
		return x.LocationName
	}
	return ""
}

func (x *Post) GetDrink() *Drink {
	if x != nil {
		return x.Drink
	}
	return nil
}

func (x *Post) GetDrunkenness() int64 {
	if x != nil {
		return x.Drunkenness
	}
	return 0
}

func (x *Post) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Post) GetCanComment() bool {
	if x != nil {
		return x.CanComment
	}
	return false
}

func (x *Post) GetCanShare() bool {
	if x != nil {
		return x.CanShare
	}
	return false
}

func (x *Post) GetRatio() PostRatio {
	if x != nil {
		return x.Ratio
	}
	return PostRatio_RATIO_16_9
}

func (x *Post) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Post) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Post) GetLastCommentText() string {
	if x != nil {
		return x.LastCommentText
	}
	return ""
}

func (x *Post) GetLastCommentUsername() string {
	if x != nil {
		return x.LastCommentUsername
	}
	return ""
}

func (x *Post) GetLastCommentCreateTime() int64 {
	if x != nil {
		return x.LastCommentCreateTime
	}
	return 0
}

func (x *Post) GetAudio() *audio.Audio {
	if x != nil {
		return x.Audio
	}
	return nil
}

type PostMedia struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessibilityCaption string          `protobuf:"bytes,1,opt,name=accessibility_caption,json=accessibilityCaption,proto3" json:"accessibility_caption,omitempty"`
	MediaType            MediaType       `protobuf:"varint,2,opt,name=media_type,json=mediaType,proto3,enum=cheers.type.MediaType" json:"media_type,omitempty"`
	ImageVersions        []*ImageVersion `protobuf:"bytes,3,rep,name=image_versions,json=imageVersions,proto3" json:"image_versions,omitempty"`
}

func (x *PostMedia) Reset() {
	*x = PostMedia{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_type_post_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostMedia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostMedia) ProtoMessage() {}

func (x *PostMedia) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_type_post_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostMedia.ProtoReflect.Descriptor instead.
func (*PostMedia) Descriptor() ([]byte, []int) {
	return file_cheers_type_post_post_proto_rawDescGZIP(), []int{1}
}

func (x *PostMedia) GetAccessibilityCaption() string {
	if x != nil {
		return x.AccessibilityCaption
	}
	return ""
}

func (x *PostMedia) GetMediaType() MediaType {
	if x != nil {
		return x.MediaType
	}
	return MediaType_IMAGE
}

func (x *PostMedia) GetImageVersions() []*ImageVersion {
	if x != nil {
		return x.ImageVersions
	}
	return nil
}

type ImageVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Ref    string `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	Width  int64  `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height int64  `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *ImageVersion) Reset() {
	*x = ImageVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_type_post_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageVersion) ProtoMessage() {}

func (x *ImageVersion) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_type_post_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageVersion.ProtoReflect.Descriptor instead.
func (*ImageVersion) Descriptor() ([]byte, []int) {
	return file_cheers_type_post_post_proto_rawDescGZIP(), []int{2}
}

func (x *ImageVersion) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ImageVersion) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *ImageVersion) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ImageVersion) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type Drink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *Drink) Reset() {
	*x = Drink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_type_post_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Drink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Drink) ProtoMessage() {}

func (x *Drink) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_type_post_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Drink.ProtoReflect.Descriptor instead.
func (*Drink) Descriptor() ([]byte, []int) {
	return file_cheers_type_post_post_proto_rawDescGZIP(), []int{3}
}

func (x *Drink) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Drink) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Drink) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

var File_cheers_type_post_post_proto protoreflect.FileDescriptor

var file_cheers_type_post_post_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63,
	0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x21, 0x63, 0x68, 0x65, 0x65,
	0x72, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63,
	0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x05, 0x0a,
	0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x65, 0x65,
	0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x52,
	0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x35, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63,
	0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12,
	0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x64, 0x72, 0x69, 0x6e, 0x6b, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x44, 0x72, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x64, 0x72, 0x69, 0x6e, 0x6b, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x72, 0x75, 0x6e, 0x6b, 0x65, 0x6e, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x72, 0x75, 0x6e, 0x6b, 0x65, 0x6e, 0x6e, 0x65, 0x73, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x61, 0x6e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12,
	0x2c, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x52, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x6f, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x09, 0x50,
	0x6f, 0x73, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x33, 0x0a, 0x15, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a,
	0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x60, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3f, 0x0a, 0x05, 0x44, 0x72, 0x69, 0x6e,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x2a, 0x21, 0x0a, 0x09, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x01, 0x2a, 0x39, 0x0a, 0x09,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x5f, 0x31, 0x36, 0x5f, 0x39, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x5f, 0x31, 0x5f, 0x31, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x5f, 0x34, 0x5f, 0x35, 0x10, 0x02, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x61, 0x7a, 0x61, 0x72, 0x68, 0x75, 0x67,
	0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cheers_type_post_post_proto_rawDescOnce sync.Once
	file_cheers_type_post_post_proto_rawDescData = file_cheers_type_post_post_proto_rawDesc
)

func file_cheers_type_post_post_proto_rawDescGZIP() []byte {
	file_cheers_type_post_post_proto_rawDescOnce.Do(func() {
		file_cheers_type_post_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_cheers_type_post_post_proto_rawDescData)
	})
	return file_cheers_type_post_post_proto_rawDescData
}

var file_cheers_type_post_post_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cheers_type_post_post_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cheers_type_post_post_proto_goTypes = []interface{}{
	(MediaType)(0),       // 0: cheers.type.MediaType
	(PostRatio)(0),       // 1: cheers.type.PostRatio
	(*Post)(nil),         // 2: cheers.type.Post
	(*PostMedia)(nil),    // 3: cheers.type.PostMedia
	(*ImageVersion)(nil), // 4: cheers.type.ImageVersion
	(*Drink)(nil),        // 5: cheers.type.Drink
	(privacy.Privacy)(0), // 6: cheers.type.Privacy
	(*audio.Audio)(nil),  // 7: cheers.type.Audio
}
var file_cheers_type_post_post_proto_depIdxs = []int32{
	6, // 0: cheers.type.Post.privacy:type_name -> cheers.type.Privacy
	3, // 1: cheers.type.Post.post_media:type_name -> cheers.type.PostMedia
	5, // 2: cheers.type.Post.drink:type_name -> cheers.type.Drink
	1, // 3: cheers.type.Post.ratio:type_name -> cheers.type.PostRatio
	7, // 4: cheers.type.Post.audio:type_name -> cheers.type.Audio
	0, // 5: cheers.type.PostMedia.media_type:type_name -> cheers.type.MediaType
	4, // 6: cheers.type.PostMedia.image_versions:type_name -> cheers.type.ImageVersion
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cheers_type_post_post_proto_init() }
func file_cheers_type_post_post_proto_init() {
	if File_cheers_type_post_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cheers_type_post_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_cheers_type_post_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostMedia); i {
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
		file_cheers_type_post_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageVersion); i {
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
		file_cheers_type_post_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Drink); i {
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
			RawDescriptor: file_cheers_type_post_post_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cheers_type_post_post_proto_goTypes,
		DependencyIndexes: file_cheers_type_post_post_proto_depIdxs,
		EnumInfos:         file_cheers_type_post_post_proto_enumTypes,
		MessageInfos:      file_cheers_type_post_post_proto_msgTypes,
	}.Build()
	File_cheers_type_post_post_proto = out.File
	file_cheers_type_post_post_proto_rawDesc = nil
	file_cheers_type_post_post_proto_goTypes = nil
	file_cheers_type_post_post_proto_depIdxs = nil
}
