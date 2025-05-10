package pubsub

import (
	"context"
	"encoding/json"
	"strings"

	"cloud.google.com/go/pubsub"
	"google.golang.org/grpc/metadata"
)

type (
	BaseMsg struct {
		MetaData map[string]string `json:"meta_data"`
	}
	SubsTopicType struct {
		SubscriptionId string
		Topic          string
		Consumer       ConsumerInterface
	}

	ConsumerInterface interface {
		Serve(ctx context.Context, msg *pubsub.Message)
	}
)

func NewBaseMsg() BaseMsg {
	return BaseMsg{}
}

// deprecated// Deprecated: Use package marshaler.
func (b *BaseMsg) Parse(ctx context.Context, msgData []byte, res interface{}) (context.Context, error) {
	var err = json.Unmarshal(msgData, res)
	if err != nil {
		return ctx, err
	}

	err = json.Unmarshal(msgData, &b)
	if err != nil {
		return ctx, err
	}

	if len(b.MetaData) > 0 {
		ctx = b.metaDataParse(ctx)
	}

	return ctx, err
}

func (b *BaseMsg) metaDataParse(ctx context.Context) context.Context {
	var newMetaData = map[string]string{}
	for i, v := range b.MetaData {
		i := strings.ToLower(i)
		switch i {
		case "device-type", "device_type":
			newMetaData["Device-Type"] = v
		case "app-version", "app_version":
			newMetaData["App-Version"] = v
		}
	}
	//add meta data / header to context
	md := metadata.New(newMetaData)
	return metadata.NewIncomingContext(ctx, md)
}
