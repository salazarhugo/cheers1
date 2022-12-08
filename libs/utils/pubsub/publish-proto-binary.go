package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"google.golang.org/protobuf/proto"
)

// PublishProtoWithBinaryEncoding Publish Proto Message with BINARY encoding
func PublishProtoWithBinaryEncoding(topicID string, message proto.Message) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	t := client.Topic(topicID)

	var msg []byte
	msg, err = proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("proto.Marshal err: %v", err)
	}

	result := t.Publish(ctx, &pubsub.Message{
		Data: msg,
	})
	_, err = result.Get(ctx)
	if err != nil {
		return fmt.Errorf("result.Get: %v", err)
	}
	fmt.Printf("Published proto message %s\n", string(msg))

	return nil
}
