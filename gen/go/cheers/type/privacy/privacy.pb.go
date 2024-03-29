// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: cheers/type/privacy/privacy.proto

package privacy

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

type Privacy int32

const (
	Privacy_FRIENDS Privacy = 0
	Privacy_PRIVATE Privacy = 1
	Privacy_PUBLIC  Privacy = 2
	Privacy_GROUP   Privacy = 3
)

// Enum value maps for Privacy.
var (
	Privacy_name = map[int32]string{
		0: "FRIENDS",
		1: "PRIVATE",
		2: "PUBLIC",
		3: "GROUP",
	}
	Privacy_value = map[string]int32{
		"FRIENDS": 0,
		"PRIVATE": 1,
		"PUBLIC":  2,
		"GROUP":   3,
	}
)

func (x Privacy) Enum() *Privacy {
	p := new(Privacy)
	*p = x
	return p
}

func (x Privacy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Privacy) Descriptor() protoreflect.EnumDescriptor {
	return file_cheers_type_privacy_privacy_proto_enumTypes[0].Descriptor()
}

func (Privacy) Type() protoreflect.EnumType {
	return &file_cheers_type_privacy_privacy_proto_enumTypes[0]
}

func (x Privacy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Privacy.Descriptor instead.
func (Privacy) EnumDescriptor() ([]byte, []int) {
	return file_cheers_type_privacy_privacy_proto_rawDescGZIP(), []int{0}
}

var File_cheers_type_privacy_privacy_proto protoreflect.FileDescriptor

var file_cheers_type_privacy_privacy_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2a, 0x3a, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x46,
	0x52, 0x49, 0x45, 0x4e, 0x44, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56,
	0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x03, 0x42, 0x43, 0x5a, 0x41,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x61, 0x7a,
	0x61, 0x72, 0x68, 0x75, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x31, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x3b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cheers_type_privacy_privacy_proto_rawDescOnce sync.Once
	file_cheers_type_privacy_privacy_proto_rawDescData = file_cheers_type_privacy_privacy_proto_rawDesc
)

func file_cheers_type_privacy_privacy_proto_rawDescGZIP() []byte {
	file_cheers_type_privacy_privacy_proto_rawDescOnce.Do(func() {
		file_cheers_type_privacy_privacy_proto_rawDescData = protoimpl.X.CompressGZIP(file_cheers_type_privacy_privacy_proto_rawDescData)
	})
	return file_cheers_type_privacy_privacy_proto_rawDescData
}

var file_cheers_type_privacy_privacy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cheers_type_privacy_privacy_proto_goTypes = []interface{}{
	(Privacy)(0), // 0: cheers.type.Privacy
}
var file_cheers_type_privacy_privacy_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cheers_type_privacy_privacy_proto_init() }
func file_cheers_type_privacy_privacy_proto_init() {
	if File_cheers_type_privacy_privacy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cheers_type_privacy_privacy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cheers_type_privacy_privacy_proto_goTypes,
		DependencyIndexes: file_cheers_type_privacy_privacy_proto_depIdxs,
		EnumInfos:         file_cheers_type_privacy_privacy_proto_enumTypes,
	}.Build()
	File_cheers_type_privacy_privacy_proto = out.File
	file_cheers_type_privacy_privacy_proto_rawDesc = nil
	file_cheers_type_privacy_privacy_proto_goTypes = nil
	file_cheers_type_privacy_privacy_proto_depIdxs = nil
}
