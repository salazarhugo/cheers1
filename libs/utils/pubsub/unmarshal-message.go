package pubsub

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
)

type Message struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

// UnmarshalPubSubMessage Unmarshal pub/sub message into proto.Message `event`
func UnmarshalPubSubMessage(
	r *http.Request,
	event proto.Message,
) error {
	var m Message
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	// byte slice unmarshalling handles base64 decoding.
	if err := json.Unmarshal(body, &m); err != nil {
		return err
	}

	err = proto.Unmarshal(m.Message.Data, event)
	if err != nil {
		return err
	}

	return nil
}
