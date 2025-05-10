# Redis Client

common redis library that can be used to create client with default bb config

sample code how to use this library

```go
package main

import (
	commonRedis "github.com/LukmanulHakim18/core/redis"
)

func main() {
	bbRedis := commonRedis.NewRedis{
		address:      commonMicro.OsGetString("REDIS_ADDRESS", "localhost"),
		port:        commonMicro.OsGetInt("REDIS_PORT", 6739),
		password: commonMicro.OsGetString("REDIS_PASSWORD", ""),
		db: commonMicro.OsGetInt("REDIS_DB", 0),
	}

	//get client
	_ = bbRedis.Client()

	//ping
	err = bbRedis.Ping()
	if err != nil {
		panic(err)
	}

	//config listener
	err = bbRedis.SetConfigKEA()

	if err != nil {
		panic(err)
    }

}
```
