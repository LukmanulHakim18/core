package notification

import (
	"context"
	"encoding/json"
)

type Email struct {
	From       string   `json:"from"`
	Recipients []string `json:"recipients"`
	Cc         []string `json:"cc"`
	Subject    string   `json:"subject"`
	Content    string   `json:"content"`
}

func (m *Email) BuildMessage(ctx context.Context) []byte {
	byt, _ := json.Marshal(m)
	return byt

}
