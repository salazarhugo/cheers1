package utils

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
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
