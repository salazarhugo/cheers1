package utils

import (
	"cloud.google.com/go/pubsub"
	"context"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
	"salazar/cheers/user/usereventpb"
)

func PublishProtoMessages(state *usereventpb.UserEvent) error {
	projectID := "cheers-a275e"
	topicID := "user-event-topic"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("subscriptions.NewClient: %v", err)
		return err
	}

	// Get the topic encoding type.
	t := client.Topic(topicID)
	cfg, err := t.Config(ctx)
	if err != nil {
		log.Printf("topic.Config err: %v", err)
		return err
	}
	encoding := cfg.SchemaSettings.Encoding

	var msg []byte
	switch encoding {
	case pubsub.EncodingBinary:
		msg, err = proto.Marshal(state)
		if err != nil {
			log.Printf("proto.Marshal err: %v", err)
			return err
		}
	case pubsub.EncodingJSON:
		msg, err = protojson.Marshal(state)
		if err != nil {
			log.Printf("protojson.Marshal err: %v", err)
			return err
		}
	default:
		log.Printf("invalid encoding: %v", encoding)
		return err
	}

	result := t.Publish(ctx, &pubsub.Message{
		Data: msg,
	})
	_, err = result.Get(ctx)
	if err != nil {
		log.Printf("result.Get: %v", err)
		return err
	}
	log.Printf("Published proto message with %#v encoding: %s\n", encoding, string(msg))
	return nil
}
