package error

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.elastic.co/apm/v2"
)

type APMReporter struct {
	ErrorFormat *Error
	Culprit     string
	ctx         context.Context
}

func GetAPMReporter(err *Error) *APMReporter {
	return &APMReporter{
		ErrorFormat: err,
		Culprit:     "APM Reporter",
	}
}

type ErrorNews struct {
	HttpStatus string `json:"http_status"`
	GrpcStatus string `json:"grpc_status"`
	Message    string `json:"error_message"`
}

func (ar *APMReporter) Report(ctx context.Context) {
	en := MakeErrorNews(ar.ErrorFormat)
	e := apm.CaptureError(ctx, en)
	if e != nil {
		return
	}
	e.Culprit = ar.Culprit
	e.Send()
}

func MakeErrorNews(e *Error) *ErrorNews {
	if e == nil {
		return &ErrorNews{
			HttpStatus: "",
			GrpcStatus: "",
			Message:    "unknown error",
		}
	}
	return &ErrorNews{
		HttpStatus: fmt.Sprintf("[%d]-%s", e.StatusCode, http.StatusText(e.StatusCode)),
		GrpcStatus: fmt.Sprintf("[%d]-%s", e.GrpcCode(), e.GrpcCode().String()),
		Message:    e.ErrorMessage,
	}

}
func (e *ErrorNews) Error() string {
	byt, _ := json.Marshal(e)
	return string(byt)
}
