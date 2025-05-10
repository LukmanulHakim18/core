package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// * Transform proto message to query params.
// nested message not supported
func MessageToQueryParams(msg protoreflect.ProtoMessage) (url.Values, error) {
	queryParams := url.Values{}
	msgType := msg.ProtoReflect().Descriptor()
	for i := 0; i < msgType.Fields().Len(); i++ {
		fd := msgType.Fields().Get(i)
		value := msg.ProtoReflect().Get(fd)
		if fd.IsList() {
			for i := 0; i < value.List().Len(); i++ {
				queryParams.Add(string(fd.Name()), value.List().Get(i).String())
			}
		} else {
			queryParams.Add(string(fd.Name()), value.String())
		}
	}
	return queryParams, nil
}

// * Mapping response and add wrapper for wrap response message.
func MappingResponse(response *http.Response, targetSuccess, targetError any) error {
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.Status)
	if response.StatusCode < 200 || response.StatusCode > 299 {
		if err := json.Unmarshal(body, targetError); err != nil {
			return fmt.Errorf("%d:%s", response.StatusCode, string(body))
		}
	}
	if response.StatusCode == http.StatusNoContent {
		return nil
	}
	err := json.Unmarshal(body, &targetSuccess)
	return err
}

func MetadataToHttpHeader(ctx context.Context) http.Header {
	header := http.Header{}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return header
	}
	for k, v := range md {
		for _, v2 := range v {
			header.Add(k, v2)
		}
	}
	return header
}
