package utils

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"salazar/cheers/payment/proto/ticketpb"
)

func MapToTicket(m map[string]interface{}) (*ticketpb.Ticket, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	s := &ticketpb.Ticket{}
	err = protojson.Unmarshal(b, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func MapToProto(proto proto.Message, m map[string]interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = protojson.Unmarshal(b, proto)
	if err != nil {
		return err
	}

	return nil
}
