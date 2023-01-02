// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: cheers/ticket/v1/ticket_service.proto

package ticket

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

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Price of the ticket
	Price int64 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	// Number of available tickets to be sold
	Quantity        uint32 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Name            string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description     string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	PartyId         string `protobuf:"bytes,6,opt,name=partyId,proto3" json:"partyId,omitempty"`
	PartyName       string `protobuf:"bytes,10,opt,name=party_name,json=partyName,proto3" json:"party_name,omitempty"`
	PartyStartTime  int64  `protobuf:"varint,11,opt,name=party_start_time,json=partyStartTime,proto3" json:"party_start_time,omitempty"`
	PartyEndTime    int64  `protobuf:"varint,12,opt,name=party_end_time,json=partyEndTime,proto3" json:"party_end_time,omitempty"`
	LocationName    string `protobuf:"bytes,13,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
	Organizer       string `protobuf:"bytes,14,opt,name=organizer,proto3" json:"organizer,omitempty"`
	PaymentIntentId string `protobuf:"bytes,7,opt,name=paymentIntentId,proto3" json:"paymentIntentId,omitempty"`
	UserId          string `protobuf:"bytes,8,opt,name=userId,proto3" json:"userId,omitempty"`
	Validated       bool   `protobuf:"varint,9,opt,name=validated,proto3" json:"validated,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_cheers_ticket_v1_ticket_service_proto_rawDescGZIP(), []int{0}
}

func (x *Ticket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ticket) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Ticket) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Ticket) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Ticket) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Ticket) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *Ticket) GetPartyName() string {
	if x != nil {
		return x.PartyName
	}
	return ""
}

func (x *Ticket) GetPartyStartTime() int64 {
	if x != nil {
		return x.PartyStartTime
	}
	return 0
}

func (x *Ticket) GetPartyEndTime() int64 {
	if x != nil {
		return x.PartyEndTime
	}
	return 0
}

func (x *Ticket) GetLocationName() string {
	if x != nil {
		return x.LocationName
	}
	return ""
}

func (x *Ticket) GetOrganizer() string {
	if x != nil {
		return x.Organizer
	}
	return ""
}

func (x *Ticket) GetPaymentIntentId() string {
	if x != nil {
		return x.PaymentIntentId
	}
	return ""
}

func (x *Ticket) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Ticket) GetValidated() bool {
	if x != nil {
		return x.Validated
	}
	return false
}

type CreateTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *CreateTicketRequest) Reset() {
	*x = CreateTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTicketRequest) ProtoMessage() {}

func (x *CreateTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTicketRequest.ProtoReflect.Descriptor instead.
func (*CreateTicketRequest) Descriptor() ([]byte, []int) {
	return file_cheers_ticket_v1_ticket_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTicketRequest) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type CreateTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *CreateTicketResponse) Reset() {
	*x = CreateTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTicketResponse) ProtoMessage() {}

func (x *CreateTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTicketResponse.ProtoReflect.Descriptor instead.
func (*CreateTicketResponse) Descriptor() ([]byte, []int) {
	return file_cheers_ticket_v1_ticket_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTicketResponse) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type GetTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId string `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
}

func (x *GetTicketRequest) Reset() {
	*x = GetTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicketRequest) ProtoMessage() {}

func (x *GetTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicketRequest.ProtoReflect.Descriptor instead.
func (*GetTicketRequest) Descriptor() ([]byte, []int) {
	return file_cheers_ticket_v1_ticket_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetTicketRequest) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

type GetTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *GetTicketResponse) Reset() {
	*x = GetTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicketResponse) ProtoMessage() {}

func (x *GetTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicketResponse.ProtoReflect.Descriptor instead.
func (*GetTicketResponse) Descriptor() ([]byte, []int) {
	return file_cheers_ticket_v1_ticket_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetTicketResponse) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type UpdateTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *UpdateTicketRequest) Reset() {
	*x = UpdateTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTicketRequest) ProtoMessage() {}

func (x *UpdateTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTicketRequest.ProtoReflect.Descriptor instead.
func (*UpdateTicketRequest) Descriptor() ([]byte, []int) {
	return file_cheers_ticket_v1_ticket_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTicketRequest) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type UpdateTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *UpdateTicketResponse) Reset() {
	*x = UpdateTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTicketResponse) ProtoMessage() {}

func (x *UpdateTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTicketResponse.ProtoReflect.Descriptor instead.
func (*UpdateTicketResponse) Descriptor() ([]byte, []int) {
	return file_cheers_ticket_v1_ticket_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTicketResponse) GetTicket() *Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

type DeleteTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId string `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
}

func (x *DeleteTicketRequest) Reset() {
	*x = DeleteTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTicketRequest) ProtoMessage() {}

func (x *DeleteTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTicketRequest.ProtoReflect.Descriptor instead.
func (*DeleteTicketRequest) Descriptor() ([]byte, []int) {
	return file_cheers_ticket_v1_ticket_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTicketRequest) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

type DeleteTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTicketResponse) Reset() {
	*x = DeleteTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTicketResponse) ProtoMessage() {}

func (x *DeleteTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTicketResponse.ProtoReflect.Descriptor instead.
func (*DeleteTicketResponse) Descriptor() ([]byte, []int) {
	return file_cheers_ticket_v1_ticket_service_proto_rawDescGZIP(), []int{8}
}

type ListTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Filter:
	//
	//	*ListTicketRequest_PartyId
	//	*ListTicketRequest_UserId
	Filter isListTicketRequest_Filter `protobuf_oneof:"filter"`
}

func (x *ListTicketRequest) Reset() {
	*x = ListTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTicketRequest) ProtoMessage() {}

func (x *ListTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTicketRequest.ProtoReflect.Descriptor instead.
func (*ListTicketRequest) Descriptor() ([]byte, []int) {
	return file_cheers_ticket_v1_ticket_service_proto_rawDescGZIP(), []int{9}
}

func (m *ListTicketRequest) GetFilter() isListTicketRequest_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (x *ListTicketRequest) GetPartyId() string {
	if x, ok := x.GetFilter().(*ListTicketRequest_PartyId); ok {
		return x.PartyId
	}
	return ""
}

func (x *ListTicketRequest) GetUserId() string {
	if x, ok := x.GetFilter().(*ListTicketRequest_UserId); ok {
		return x.UserId
	}
	return ""
}

type isListTicketRequest_Filter interface {
	isListTicketRequest_Filter()
}

type ListTicketRequest_PartyId struct {
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3,oneof"`
}

type ListTicketRequest_UserId struct {
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3,oneof"`
}

func (*ListTicketRequest_PartyId) isListTicketRequest_Filter() {}

func (*ListTicketRequest_UserId) isListTicketRequest_Filter() {}

type ListTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tickets []*Ticket `protobuf:"bytes,1,rep,name=tickets,proto3" json:"tickets,omitempty"`
}

func (x *ListTicketResponse) Reset() {
	*x = ListTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTicketResponse) ProtoMessage() {}

func (x *ListTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cheers_ticket_v1_ticket_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTicketResponse.ProtoReflect.Descriptor instead.
func (*ListTicketResponse) Descriptor() ([]byte, []int) {
	return file_cheers_ticket_v1_ticket_service_proto_rawDescGZIP(), []int{10}
}

func (x *ListTicketResponse) GetTickets() []*Ticket {
	if x != nil {
		return x.Tickets
	}
	return nil
}

var File_cheers_ticket_v1_ticket_service_proto protoreflect.FileDescriptor

var file_cheers_ticket_v1_ticket_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x03, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x0e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x79, 0x45, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x47, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a,
	0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22,
	0x48, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x22, 0x47, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x48, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x55, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x08, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x32, 0xe2, 0x04, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x75, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x68, 0x65,
	0x65, 0x72, 0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x75, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x22, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x68,
	0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0x75, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x25, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x32, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a,
	0x17, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6c, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x68,
	0x65, 0x65, 0x72, 0x73, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x41, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x61, 0x7a, 0x61, 0x72, 0x68, 0x75,
	0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x63, 0x68, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f,
	0x76, 0x31, 0x3b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cheers_ticket_v1_ticket_service_proto_rawDescOnce sync.Once
	file_cheers_ticket_v1_ticket_service_proto_rawDescData = file_cheers_ticket_v1_ticket_service_proto_rawDesc
)

func file_cheers_ticket_v1_ticket_service_proto_rawDescGZIP() []byte {
	file_cheers_ticket_v1_ticket_service_proto_rawDescOnce.Do(func() {
		file_cheers_ticket_v1_ticket_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cheers_ticket_v1_ticket_service_proto_rawDescData)
	})
	return file_cheers_ticket_v1_ticket_service_proto_rawDescData
}

var file_cheers_ticket_v1_ticket_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_cheers_ticket_v1_ticket_service_proto_goTypes = []interface{}{
	(*Ticket)(nil),               // 0: cheers.ticket.v1.Ticket
	(*CreateTicketRequest)(nil),  // 1: cheers.ticket.v1.CreateTicketRequest
	(*CreateTicketResponse)(nil), // 2: cheers.ticket.v1.CreateTicketResponse
	(*GetTicketRequest)(nil),     // 3: cheers.ticket.v1.GetTicketRequest
	(*GetTicketResponse)(nil),    // 4: cheers.ticket.v1.GetTicketResponse
	(*UpdateTicketRequest)(nil),  // 5: cheers.ticket.v1.UpdateTicketRequest
	(*UpdateTicketResponse)(nil), // 6: cheers.ticket.v1.UpdateTicketResponse
	(*DeleteTicketRequest)(nil),  // 7: cheers.ticket.v1.DeleteTicketRequest
	(*DeleteTicketResponse)(nil), // 8: cheers.ticket.v1.DeleteTicketResponse
	(*ListTicketRequest)(nil),    // 9: cheers.ticket.v1.ListTicketRequest
	(*ListTicketResponse)(nil),   // 10: cheers.ticket.v1.ListTicketResponse
}
var file_cheers_ticket_v1_ticket_service_proto_depIdxs = []int32{
	0,  // 0: cheers.ticket.v1.CreateTicketRequest.ticket:type_name -> cheers.ticket.v1.Ticket
	0,  // 1: cheers.ticket.v1.CreateTicketResponse.ticket:type_name -> cheers.ticket.v1.Ticket
	0,  // 2: cheers.ticket.v1.GetTicketResponse.ticket:type_name -> cheers.ticket.v1.Ticket
	0,  // 3: cheers.ticket.v1.UpdateTicketRequest.ticket:type_name -> cheers.ticket.v1.Ticket
	0,  // 4: cheers.ticket.v1.UpdateTicketResponse.ticket:type_name -> cheers.ticket.v1.Ticket
	0,  // 5: cheers.ticket.v1.ListTicketResponse.tickets:type_name -> cheers.ticket.v1.Ticket
	1,  // 6: cheers.ticket.v1.TicketService.CreateTicket:input_type -> cheers.ticket.v1.CreateTicketRequest
	3,  // 7: cheers.ticket.v1.TicketService.GetTicket:input_type -> cheers.ticket.v1.GetTicketRequest
	5,  // 8: cheers.ticket.v1.TicketService.UpdateTicket:input_type -> cheers.ticket.v1.UpdateTicketRequest
	7,  // 9: cheers.ticket.v1.TicketService.DeleteTicket:input_type -> cheers.ticket.v1.DeleteTicketRequest
	9,  // 10: cheers.ticket.v1.TicketService.ListTicket:input_type -> cheers.ticket.v1.ListTicketRequest
	2,  // 11: cheers.ticket.v1.TicketService.CreateTicket:output_type -> cheers.ticket.v1.CreateTicketResponse
	4,  // 12: cheers.ticket.v1.TicketService.GetTicket:output_type -> cheers.ticket.v1.GetTicketResponse
	6,  // 13: cheers.ticket.v1.TicketService.UpdateTicket:output_type -> cheers.ticket.v1.UpdateTicketResponse
	8,  // 14: cheers.ticket.v1.TicketService.DeleteTicket:output_type -> cheers.ticket.v1.DeleteTicketResponse
	10, // 15: cheers.ticket.v1.TicketService.ListTicket:output_type -> cheers.ticket.v1.ListTicketResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_cheers_ticket_v1_ticket_service_proto_init() }
func file_cheers_ticket_v1_ticket_service_proto_init() {
	if File_cheers_ticket_v1_ticket_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cheers_ticket_v1_ticket_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
		file_cheers_ticket_v1_ticket_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTicketRequest); i {
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
		file_cheers_ticket_v1_ticket_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTicketResponse); i {
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
		file_cheers_ticket_v1_ticket_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTicketRequest); i {
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
		file_cheers_ticket_v1_ticket_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTicketResponse); i {
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
		file_cheers_ticket_v1_ticket_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTicketRequest); i {
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
		file_cheers_ticket_v1_ticket_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTicketResponse); i {
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
		file_cheers_ticket_v1_ticket_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTicketRequest); i {
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
		file_cheers_ticket_v1_ticket_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTicketResponse); i {
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
		file_cheers_ticket_v1_ticket_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTicketRequest); i {
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
		file_cheers_ticket_v1_ticket_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTicketResponse); i {
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
	file_cheers_ticket_v1_ticket_service_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*ListTicketRequest_PartyId)(nil),
		(*ListTicketRequest_UserId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cheers_ticket_v1_ticket_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cheers_ticket_v1_ticket_service_proto_goTypes,
		DependencyIndexes: file_cheers_ticket_v1_ticket_service_proto_depIdxs,
		MessageInfos:      file_cheers_ticket_v1_ticket_service_proto_msgTypes,
	}.Build()
	File_cheers_ticket_v1_ticket_service_proto = out.File
	file_cheers_ticket_v1_ticket_service_proto_rawDesc = nil
	file_cheers_ticket_v1_ticket_service_proto_goTypes = nil
	file_cheers_ticket_v1_ticket_service_proto_depIdxs = nil
}
