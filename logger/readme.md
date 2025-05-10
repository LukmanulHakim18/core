# Logger

common logger library that can be used to replace default logger

sample code how to use this library

```go
package main

import (
	"fmt"

	commonLogger "github.com/LukmanulHakim18/core/logger"
	commonMicro "github.com/LukmanulHakim18/core/microservice"
)

func main() {
	conf := commonLogger.LoggerConfig{
		AppName:      commonMicro.OsGetString("APP_NAME", "samplesvc"),
		Level:        commonMicro.OsGetString("LOG_LEVEL", commonLogger.LevelDebug),
		LogDirectory: commonMicro.OsGetString("LOG_DIRECTORY", "/var/log/samplesvc/log/"),
	}

	logger, err := commonLogger.NewLogger(conf)
	if err != nil {
		panic(err)
	}

	logger.Info("sample print log")

	data := map[string]interface{}{
		"hello": "world",
	}
	fields := []commonLogger.Field{}
	for key, val := range data {
		fields = append(fields, commonLogger.Field{
			Key:   key,
			Value: fmt.Sprintf("%v", val),
		})
	}

	logger.Error("sample error with custom field", fields...)
}
```
