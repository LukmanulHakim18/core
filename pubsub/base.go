package pubsub

import (
	"context"
	"errors"
	"os"

	"cloud.google.com/go/pubsub"
	"go.elastic.co/apm/v2"
)

type Client struct {
	Conn *pubsub.Client
}

func NewPubSub(projectId string) (*Client, error) {
	isEmulator := false
	isCloud := false
	if os.Getenv("PUBSUB_EMULATOR_HOST") != "" {
		isEmulator = true
	}
	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") != "" {
		isCloud = true
	}
	if !isEmulator && !isCloud {
		return nil, errors.New("cannot init pubsub client, need to set PUBSUB_EMULATOR_HOST or GOOGLE_APPLICATION_CREDENTIALS env first")
	}
	client, err := pubsub.NewClient(context.Background(), projectId)
	if err != nil {
		return nil, err
	}
	return &Client{
		Conn: client,
	}, nil
}

func (c *Client) Close() error {
	return c.Conn.Close()
}

func (c *Client) PublishMessage(ctx context.Context, topicName string, message []byte) error {
	topic, err := CheckAndCreateTopic(ctx, topicName, c.Conn)
	if err != nil {
		return err
	}
	span, ctx := apm.StartSpan(ctx, "publish to "+topicName, "pubsub.publisher")
	defer span.End()
	result := topic.Publish(ctx, &pubsub.Message{Data: message})
	if _, err := result.Get(ctx); err != nil {
		return err
	}
	return nil
}
