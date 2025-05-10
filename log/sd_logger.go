package log

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"sync"

	logkit "github.com/go-kit/kit/log"
)

const (
	FORMAT_SD_JSON = `{"message": "%s", "severity": "%s"}`
)

type logfmtEncoder struct {
	*Encoder
	buf bytes.Buffer
}

func (l *logfmtEncoder) Reset() {
	l.Encoder.Reset()
	l.buf.Reset()
}

var logfmtEncoderPool = sync.Pool{
	New: func() interface{} {
		var enc logfmtEncoder
		enc.Encoder = NewEncoder(&enc.buf)
		return &enc
	},
}

type logfmtLogger struct {
	w io.Writer
}

// NewLogfmtLogger returns a logger that encodes keyvals to the Writer in
// logfmt format. Each log event produces no more than one call to w.Write.
// The passed Writer must be safe for concurrent use by multiple goroutines if
// the returned Logger will be used concurrently.
func NewSDLogger(w io.Writer) logkit.Logger {
	return &logfmtLogger{w}
}

func (l logfmtLogger) Log(keyvals ...interface{}) error {
	enc := logfmtEncoderPool.Get().(*logfmtEncoder)
	enc.Reset()
	defer logfmtEncoderPool.Put(enc)

	if err := enc.EncodeKeyvals(keyvals...); err != nil {
		return err
	}

	// Add newline to the end of the buffer
	if err := enc.EndRecord(); err != nil {
		return err
	}

	payload := strings.TrimSuffix(string(enc.buf.Bytes()), "\n")
	kvp := firstKVP(payload)

	x := strings.Split(kvp, "=")
	if x[0] == "level" {
		payload = fmt.Sprintf(FORMAT_SD_JSON, payload, x[1])
	}

	payload = payload + "\n"

	// The Logger interface requires implementations to be safe for concurrent
	// use by multiple goroutines. For this implementation that means making
	// only one call to l.w.Write() for each call to Log.
	if _, err := l.w.Write([]byte(payload)); err != nil {
		return err
	}
	return nil
}

func firstKVP(value string) string {
	for i := range value {
		if value[i] == ' ' {
			return value[0:i]
		}
	}
	return value
}
