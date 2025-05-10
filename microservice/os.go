package microservice

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// get system env variable
func GetOsEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		fmt.Println("INFO", "The env "+name+" not set")
	}
	return value
}

// GetString returns config value for `key` as string. If no value is
// found, returns `_default“
func OsGetString(key, _default string) string {
	val := os.Getenv(key)
	if val == "" {
		return _default
	}
	return val
}

// GetBool returns config value for `key` as bool. If no value is found,
// returns `_default“
func OsGetBool(key string, _default bool) bool {
	val, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return _default
	}
	return val
}

// GetFloat returns config value for `key` as float64. If no value is
// found, returns `_default`
func OsGetFloat(key string, _default float64) float64 {
	val, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		return _default
	}
	return val
}

// GetInt returns config value for `key` as int64. If no value is found,
// returns _default
func OsGetInt(key string, _default int64) int64 {
	val, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil {
		return _default
	}
	return val
}

// GetUInt returns config value for `key` as uint64. If no value is found,
// returns _default
func OsGetUInt(key string, _default uint64) uint64 {
	val, err := strconv.ParseUint(os.Getenv(key), 10, 64)
	if err != nil {
		return _default
	}
	return val
}

// GetDuration returns config value for `key` as duration. If no value is found,
// returns _default
func OsGetDuration(key string, _default time.Duration) time.Duration {
	val, err := time.ParseDuration(os.Getenv(key))
	if err != nil {
		return _default
	}
	return val
}

// IsProductionMode returns whether or not application is configured to run
// on production mode.
func IsProductionMode() bool {
	return os.Getenv("RUNMODE") == "production"
}

func IsTestMode() bool {
	return os.Getenv("RUNMODE") == "test"
}

func IsDebugMode() bool {
	return OsGetBool("DEBUG", false)
}
