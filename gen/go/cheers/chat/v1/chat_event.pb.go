// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: cheers/chat/v1/chat_event.proto

package chat

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

type ChatEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*ChatEvent_Create
	Event isChatEvent_Event `protobuf_oneof:"event"`
}

func (x *ChatEvent) Reset() {
	*x = ChatEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_chat_v1_chat_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatEvent) ProtoMessage() {}

func (x *ChatEvent) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_chat_v1_chat_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatEvent.ProtoReflect.Descriptor instead.
func (*ChatEvent) Descriptor() ([]byte, []int) {
	return file_cheers_chat_v1_chat_event_proto_rawDescGZIP(), []int{0}
}

func (m *ChatEvent) GetEvent() isChatEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *ChatEvent) GetCreate() *CreateMessage {
	if x, ok := x.GetEvent().(*ChatEvent_Create); ok {
		return x.Create
	}
	return nil
}

type isChatEvent_Event interface {
	isChatEvent_Event()
}

type ChatEvent_Create struct {
	Create *CreateMessage `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

func (*ChatEvent_Create) isChatEvent_Event() {}

type CreateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message         `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Sender  *user.UserItem   `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Members []*user.UserItem `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	Room    *Room            `protobuf:"bytes,4,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *CreateMessage) Reset() {
	*x = CreateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_chat_v1_chat_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessage) ProtoMessage() {}

func (x *CreateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_chat_v1_chat_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessage.ProtoReflect.Descriptor instead.
func (*CreateMessage) Descriptor() ([]byte, []int) {
	return file_cheers_chat_v1_chat_event_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMessage) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *CreateMessage) GetSender() *user.UserItem {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *CreateMessage) GetMembers() []*user.UserItem {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *CreateMessage) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

var File_cheers_chat_v1_chat_event_proto protoreflect.FileDescriptor

var file_cheers_chat_v1_chat_event_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76,
	0x31, 0x1a, 0x21, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4d, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x37,
	0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0xcc, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x42,
	0x3d, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x61, 0x6c, 0x61, 0x7a, 0x61, 0x72, 0x68, 0x75, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65,
	0x72, 0x73, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72,
	0x73, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cheers_chat_v1_chat_event_proto_rawDescOnce sync.Once
	file_cheers_chat_v1_chat_event_proto_rawDescData = file_cheers_chat_v1_chat_event_proto_rawDesc
)

func file_cheers_chat_v1_chat_event_proto_rawDescGZIP() []byte {
	file_cheers_chat_v1_chat_event_proto_rawDescOnce.Do(func() {
		file_cheers_chat_v1_chat_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_cheers_chat_v1_chat_event_proto_rawDescData)
	})
	return file_cheers_chat_v1_chat_event_proto_rawDescData
}

var file_cheers_chat_v1_chat_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cheers_chat_v1_chat_event_proto_goTypes = []interface{}{
	(*ChatEvent)(nil),     // 0: cheers.chat.v1.ChatEvent
	(*CreateMessage)(nil), // 1: cheers.chat.v1.CreateMessage
	(*Message)(nil),       // 2: cheers.chat.v1.Message
	(*user.UserItem)(nil), // 3: cheers.type.UserItem
	(*Room)(nil),          // 4: cheers.chat.v1.Room
}
var file_cheers_chat_v1_chat_event_proto_depIdxs = []int32{
	1, // 0: cheers.chat.v1.ChatEvent.create:type_name -> cheers.chat.v1.CreateMessage
	2, // 1: cheers.chat.v1.CreateMessage.message:type_name -> cheers.chat.v1.Message
	3, // 2: cheers.chat.v1.CreateMessage.sender:type_name -> cheers.type.UserItem
	3, // 3: cheers.chat.v1.CreateMessage.members:type_name -> cheers.type.UserItem
	4, // 4: cheers.chat.v1.CreateMessage.room:type_name -> cheers.chat.v1.Room
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cheers_chat_v1_chat_event_proto_init() }
func file_cheers_chat_v1_chat_event_proto_init() {
	if File_cheers_chat_v1_chat_event_proto != nil {
		return
	}
	file_cheers_chat_v1_chat_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cheers_chat_v1_chat_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatEvent); i {
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
		file_cheers_chat_v1_chat_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessage); i {
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
	file_cheers_chat_v1_chat_event_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ChatEvent_Create)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cheers_chat_v1_chat_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cheers_chat_v1_chat_event_proto_goTypes,
		DependencyIndexes: file_cheers_chat_v1_chat_event_proto_depIdxs,
		MessageInfos:      file_cheers_chat_v1_chat_event_proto_msgTypes,
	}.Build()
	File_cheers_chat_v1_chat_event_proto = out.File
	file_cheers_chat_v1_chat_event_proto_rawDesc = nil
	file_cheers_chat_v1_chat_event_proto_goTypes = nil
	file_cheers_chat_v1_chat_event_proto_depIdxs = nil
}
