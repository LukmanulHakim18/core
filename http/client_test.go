package http_test

import (
	"context"
	"net/http"
	"testing"

	commHttp "github.com/LukmanulHakim18/core/http"
)

func TestClient(t *testing.T) {
	ep := commHttp.NewEndpoint("/token/7364612d313638343931353232392d4f47524956e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", nil, "GET")
	cli := commHttp.NewClient("http://localhost:1407")

	h := http.Header{}
	h.Add("App-Version", "6.2.0")
	h.Add("host", "http://localhost:1407")

	httpResp, err := cli.Exec(context.Background(), ep, h, nil)
	if err != nil {
		t.Fail()
		return
	}

	var result map[string]interface{}

	if err := commHttp.MappingResponse(httpResp, &result, &result); err != nil {
		t.Log(err)
	}
	t.Log(result)

}
