package error

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LukmanulHakim18/core/constant"
	meta "github.com/LukmanulHakim18/core/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errDetails "google.golang.org/genproto/googleapis/rpc/errdetails"
)

type Message struct {
	English   string `json:"en"`
	Indonesia string `json:"id"`
}

type Error struct {
	DeviceLang       constant.DeviceLang `json:"-"`
	StatusCode       int                 `json:"-"`
	ErrorCode        string              `json:"error_code"`
	ErrorMessage     string              `json:"error_message"`
	ErrorField       string              `json:"error_field,omitempty"`
	LocalizedMessage Message             `json:"localized_message"`
	Data             []Data              `json:"data,omitempty"`
	ErrorData        []Data              `json:"error_data,omitempty"`
}
type Data struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type LocalizedErrorData struct {
	LocalizedTitle       LocalizedText `json:"localized_title"`
	LocalizedDescription LocalizedText `json:"localized_description"`
}

type LocalizedText struct {
	En string `json:"en"`
	Id string `json:"id"`
}

func (e *Error) Error() string {
	if e.DeviceLang == constant.DEVICE_LANG_ID {
		return e.LocalizedMessage.Indonesia
	}
	return e.LocalizedMessage.English
}

func (e *Error) WithData(data []Data) *Error {
	return &Error{
		StatusCode:   e.StatusCode,
		ErrorCode:    e.ErrorCode,
		ErrorMessage: e.ErrorMessage,
		LocalizedMessage: Message{
			English:   e.LocalizedMessage.English,
			Indonesia: e.LocalizedMessage.Indonesia,
		},
		Data: data,
	}
}

func (e *Error) WithErrorData(errCode string, data map[string]string) *Error {
	err := &Error{
		StatusCode:   e.StatusCode,
		ErrorCode:    errCode,
		ErrorMessage: e.ErrorMessage,
		LocalizedMessage: Message{
			English:   e.LocalizedMessage.English,
			Indonesia: e.LocalizedMessage.Indonesia,
		},
	}
	for k, v := range data {
		err.ErrorData = append(err.ErrorData, Data{
			Key:   k,
			Value: v,
		})
	}

	return err

}

func (e *Error) WithParameter(parameter ...interface{}) *Error {
	return &Error{
		StatusCode:   e.StatusCode,
		ErrorCode:    e.ErrorCode,
		ErrorMessage: fmt.Sprintf(e.ErrorMessage, parameter...),
		LocalizedMessage: Message{
			English:   fmt.Sprintf(e.LocalizedMessage.English, parameter...),
			Indonesia: fmt.Sprintf(e.LocalizedMessage.Indonesia, parameter...),
		},
	}
}

func (e Error) WithField(field string) *Error {
	e.ErrorField = field
	return &e
}

func NewError(code, message, english, indonesia string) *Error {
	return &Error{
		ErrorCode:    code,
		ErrorMessage: message,
		LocalizedMessage: Message{
			English:   english,
			Indonesia: indonesia,
		},
	}
}

func NewErrorWithStatus(status int, code, message, english, indonesia string) *Error {
	return &Error{
		StatusCode:   status,
		ErrorCode:    code,
		ErrorMessage: message,
		LocalizedMessage: Message{
			English:   english,
			Indonesia: indonesia,
		},
	}
}

func GetInvalidParameterMessage(err *Error, parameter ...interface{}) *Error {
	return &Error{
		StatusCode:   err.StatusCode,
		ErrorCode:    err.ErrorCode,
		ErrorMessage: fmt.Sprintf(err.ErrorMessage, parameter),
		LocalizedMessage: Message{
			English:   fmt.Sprintf(err.LocalizedMessage.English, parameter),
			Indonesia: fmt.Sprintf(err.LocalizedMessage.Indonesia, parameter),
		},
	}
}

func GetCustomMessageWithParameters(err *Error, paramID, paramEN string) *Error {
	return &Error{
		StatusCode:   err.StatusCode,
		ErrorCode:    err.ErrorCode,
		ErrorMessage: fmt.Sprintf(err.ErrorMessage, paramEN),
		LocalizedMessage: Message{
			English:   fmt.Sprintf(err.LocalizedMessage.English, paramEN),
			Indonesia: fmt.Sprintf(err.LocalizedMessage.Indonesia, paramID),
		},
	}
}

func GetCustomMessageWithArgs(err *Error, args ...interface{}) *Error {
	return &Error{
		StatusCode:   err.StatusCode,
		ErrorCode:    err.ErrorCode,
		ErrorMessage: fmt.Sprintf(err.ErrorMessage, args...),
		LocalizedMessage: Message{
			English:   fmt.Sprintf(err.LocalizedMessage.English, args...),
			Indonesia: fmt.Sprintf(err.LocalizedMessage.Indonesia, args...),
		},
	}
}

func GetUnauthorizedAccess(errCode string) *Error {
	return &Error{
		StatusCode: http.StatusUnauthorized,
		ErrorCode:  errCode,
		LocalizedMessage: Message{
			English:   "user-info not exist",
			Indonesia: "user-info tidak ditemukan",
		},
	}
}

func (e Error) GrpcCode() codes.Code {
	switch e.StatusCode {
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusGatewayTimeout:
		return codes.DeadlineExceeded
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusNotImplemented:
		return codes.Unimplemented
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	case http.StatusInternalServerError:
		return codes.Internal
	default:
		return codes.Unknown
	}
}

func (e *Error) BuildError(ctx context.Context) error {
	e.DeviceLang = meta.GetDeviceLanguageFromCtx(ctx)

	// Message error will dynamic base on DeviceLang
	st := status.New(e.GrpcCode(), e.Error())

	// Set Reason as ErrorCode and Domain is StatusCode
	errorCode := &errDetails.ErrorInfo{Reason: e.ErrorCode, Domain: strconv.Itoa(e.StatusCode)}
	if e.ErrorData != nil {
		byteData, _ := json.Marshal(e.ErrorData)
		errorCode.Metadata = map[string]string{
			"error_data": string(byteData),
		}
	}

	// set localization message for error
	en := &errDetails.LocalizedMessage{Locale: constant.DEVICE_LANG_EN, Message: e.LocalizedMessage.English}
	id := &errDetails.LocalizedMessage{Locale: constant.DEVICE_LANG_ID, Message: e.LocalizedMessage.Indonesia}

	// set data error
	data := []*errDetails.BadRequest_FieldViolation{}
	for _, v := range e.ErrorData {
		data = append(data, &errDetails.BadRequest_FieldViolation{
			Field:       v.Key,
			Description: v.Value,
		})
	}
	badRequest := &errDetails.BadRequest{
		FieldViolations: data,
	}

	st, _ = st.WithDetails(en, id, errorCode, badRequest)

	// report
	reporter := GetAPMReporter(e)
	reporter.Report(ctx)

	return st.Err()
}
