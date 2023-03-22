package mapper

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func MapToProto(out proto.Message, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(bytes, out)
	if err != nil {
		return err
	}
	return nil
}

func ProtoToMap(m proto.Message) (map[string]interface{}, error) {
	bytes, err := protojson.Marshal(m)
	if err != nil {
		return nil, err
	}
	var res = make(map[string]interface{}, 0)
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
