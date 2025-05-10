package log

import (
	"io"
	"log"
	"os"
	"time"
)

type defaultLogWriter struct{}

func (lw *defaultLogWriter) Write(data []byte) (int, error) {
	log.Printf("%s", data)
	return len(data), nil
}

//NewDefaultLogWriter return default log writer
func NewDefaultLogWriter() io.Writer {
	return &defaultLogWriter{}
}

type logFileWriter struct {
	file      string
	logChan   chan []byte
	maxBuffer int
	interval  time.Duration
	ticker    *time.Ticker
	stopChan  chan interface{}
}

func (lw *logFileWriter) Write(log []byte) (int, error) {
	n := len(log)
	lw.logChan <- log
	return n, nil
}

func (lw *logFileWriter) run() error {
	go func() {
		var logs []byte
		last := time.Now()
		for {
			select {
			case log := <-lw.logChan:
				logs = append(logs, log...)
				if len(logs) >= lw.maxBuffer {
					lw.flush(logs)
					logs = nil
					last = time.Now()
				}
			case <-lw.ticker.C:
				now := time.Now()
				if last.Add(lw.interval).Before(now) && len(logs) > 0 {
					lw.flush(logs)
					logs = nil
					last = time.Now()
				}
			case <-lw.stopChan:
				if len(logs) > 0 {
					lw.flush(logs)
					logs = nil
				}
				return
			}
		}
	}()

	return nil
}

func (lw *logFileWriter) flush(logs []byte) (n int, err error) {
	f, err := os.OpenFile(lw.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	n, err = f.Write(logs)
	return
}

//NewFileWriter returns new file writer
func NewFileWriter(file string, max, interval int) io.Writer {
	duration := time.Duration(interval) * time.Second
	writer := &logFileWriter{file: file,
		logChan:   make(chan []byte),
		maxBuffer: max,
		interval:  duration,
		ticker:    time.NewTicker(duration),
		stopChan:  make(chan interface{}, 1)}
	writer.run()
	return writer
}
