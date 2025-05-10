package http

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"google.golang.org/grpc/metadata"
)

func TestMetadataToHeader(t *testing.T) {
	ctx := context.Background()
	data := map[string]string{
		"Key1": "value1",
		"Key2": "value2",
	}
	md := metadata.New(data)
	ctx = metadata.NewIncomingContext(ctx, md)
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want http.Header
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			want: map[string][]string{
				"Key1": []string{"value1"},
				"Key2": []string{"value2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MetadataToHttpHeader(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MetadataToHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
