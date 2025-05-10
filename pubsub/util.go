package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

func CheckAndCreateTopic(ctx context.Context, topicName string, client *pubsub.Client) (*pubsub.Topic, error) {
	topic := client.Topic(topicName)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		created, err := client.CreateTopic(ctx, topicName)
		if err != nil {
			return nil, err
		}
		topic = created
	}
	return topic, nil
}
