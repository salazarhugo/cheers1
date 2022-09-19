package utils

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
)

func PublishPayment(email string) error {
	projectID := "cheers-a275e"
	topicID := "payment-topic"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("subscriptions.NewClient: %v", err)
	}
	defer client.Close()

	topic := client.Topic(topicID)

	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(email),
	})
	_, err = res.Get(ctx)
	if err != nil {
		return fmt.Errorf("publish: %v", err)
	}

	//if totalErrors > 0 {
	//	return fmt.Errorf("%d of %d messages did not publish successfully", totalErrors, n)
	//}
	return nil
}
