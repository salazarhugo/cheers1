package utils

import (
	"encoding/json"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func MapToTicket(m map[string]interface{}) (*empty.Empty, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	s := &empty.Empty{}
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
