package config

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/LukmanulHakim18/core/microservice"
)

type Value struct {
	value interface{}
}

type Getter interface {
	Get() interface{}
	GetString() string
	GetBool() bool
	GetFloat() float64
	GetInt() int64
	GetUint() uint64
	GetDuration() time.Duration
	GetList() []string
}

func (v Value) Get() interface{} {
	return v.value
}

func (v Value) GetString() string {
	if !reflect.ValueOf(v.value).IsValid() {
		return ""
	}
	if reflect.ValueOf(v.value).IsZero() {
		return ""
	}
	return fmt.Sprintf("%s", v.value)
}

func (v Value) GetBool() bool {
	if val, err := strconv.ParseBool(fmt.Sprintf("%s", v.value)); err == nil {
		return val
	}
	return false
}

func (v Value) GetFloat() float64 {
	if val, err := strconv.ParseFloat(fmt.Sprintf("%s", v.value), 64); err == nil {
		return val
	}
	return 0.0
}

func (v Value) GetInt() int64 {
	if val, err := strconv.ParseInt(fmt.Sprintf("%s", v.value), 10, 64); err == nil {
		return val
	}
	return 0
}

func (v Value) GetUint() uint64 {
	if val, err := strconv.ParseUint(fmt.Sprintf("%s", v.value), 10, 64); err == nil {
		return val
	}
	return 0
}

func (v Value) GetList() []string {
	slice := strings.Split(v.value.(string), ",")
	for i, val := range slice {
		slice[i] = strings.TrimSpace(val)
	}
	return slice
}

func (v Value) GetDuration() time.Duration {
	if parsed, err := time.ParseDuration(fmt.Sprintf("%s", v.value)); err == nil {
		return parsed
	}
	return 0
}

// GetMapOfBoolean returns config value for `key` as map of boolean. If no value is
// found, returns a map with `defaultKey` as the key, `defaultVal` as the value
func (v Value) GetMapOfBoolean(value bool) map[string]bool {
	_map := map[string]bool{}

	arr := v.GetList()
	if len(arr) == 0 {
		return _map
	}

	for _, key := range arr {
		_map[key] = value
	}

	return _map
}

func NewValue(val interface{}) Value {
	return Value{value: val}
}

func (v *Value) SetValue(val interface{}) {
	v.value = val
}

func LoadConfig(conf map[string]interface{}) map[string]Value {
	retConfig := map[string]Value{}
	for key, value := range conf {
		retConfig[key] = Value{value: microservice.OsGetString(strings.ToUpper(key), fmt.Sprintf("%v", value))}
	}
	return retConfig
}
