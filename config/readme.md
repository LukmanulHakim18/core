# Config

common config library that can be used to get config loaded when app running

sample code how to use this library

```go
package main

import (
	"fmt"
	"net/http"
	"sync"

	commonConfig "github.com/LukmanulHakim18/core/config"
	"github.com/joho/godotenv"
)

var defaultConfig = map[string]interface{}{
	"app_name": "sample-app",
	"port":     8080,
}

var (
	onceConfig sync.Once
	appConfig  map[string]commonConfig.Value
)

func main() {
	godotenv.Load()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Hello from %s", GetConfig("app_name").GetString())))
	})

	fmt.Printf("running server at port %d...\n", GetConfig("port").GetInt())
	http.ListenAndServe(fmt.Sprintf(":%d", GetConfig("port").GetInt()), nil)
}

func GetConfig(key string) (val commonConfig.Value) {
	onceConfig.Do(func() {
		appConfig = commonConfig.LoadConfig(defaultConfig)
	})
	if v, ok := appConfig[key]; ok {
		val = v
	}
	return
}
```
