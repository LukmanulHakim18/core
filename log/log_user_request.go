package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"

	"net/http"
	"strings"
)

const (
	UserInfo = "USER-INFO: [%s]"

	PathURL = "PATH: [%s]"

	Method = "METHOD: [%s]"

	IP = "IP: [%s]"

	Headers = "HEADERS: [%+v]"

	Content = "CONTENT: [%s]"
)

type LogUserRequest struct {
	UserInfo, PathURL, Method, IPAddress, Headers, Content, Info string
}

func getIPClient(w http.ResponseWriter, r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return "127.0.0.1"
}

type BodyContent struct {
	Scope           string `json:"scope"`
	ResponseType    string `json:"response_type"`
	ClientID        string `json:"client_id"`
	RedirectURI     string `json:"redirect_uri"`
	State           string `json:"state"`
	ResponseMode    string `json:"response_mode"`
	Nonce           string `json:"nonce"`
	Display         string `json:"display"`
	Prompt          string `json:"prompt"`
	MaxAge          string `json:"max_age"`
	UILocales       string `json:"ui_locales"`
	IDTokenHint     string `json:"id_token_hint"`
	LoginHint       string `json:"login_hint"`
	AcrValues       string `json:"acr_values"`
	UserID          string `json:"user_id"`
	UserSecret      string `json:"user_secret"`
	OTP             string `json:"otp"`
	Expired         int32  `json:"expired"`
	Verifier        string `json:"verifier"`
	MethodChallenge string `json:"method_challenge"`
	CodeChallenge   string `json:"code_challenge"`
}

func LogRequestClient(w http.ResponseWriter, r *http.Request) LogUserRequest {

	if r.URL.Path != "/" {
		var body []byte
		if r.Body != nil {
			body, _ = ioutil.ReadAll(r.Body)
			r.Body = ioutil.NopCloser(bytes.NewReader(body))

			// set user secret to empty string in log
			var b BodyContent
			if r.URL.Path == "/token/auth" {
				_ = json.Unmarshal(body, &b)
				b.UserSecret = ""
				body, _ = json.Marshal(b)
			}
		}

		token := r.Header.Get("Authorization")
		logResp := LogUserRequest{
			UserInfo:  fmt.Sprintf("Token: %s", token),
			PathURL:   r.URL.Path,
			Method:    r.Method,
			IPAddress: getIPClient(w, r),
			Headers:   fmt.Sprintf("%+v", r.Header),
			Content:   string(body),
		}
		return logResp
	}
	return LogUserRequest{}
}
