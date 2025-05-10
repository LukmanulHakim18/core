package constant

import (
	"fmt"
	"strings"
)

type OperatingSystem string

const (
	OperatingSystemAndroid       OperatingSystem = "android"
	OperatingSystemIOS           OperatingSystem = "ios"
	OperatingSystemAndroidHuawei OperatingSystem = "android_huawei"
)

func (dt OperatingSystem) String() string {
	return string(dt)
}

var DeviceTypeIndex = map[string]OperatingSystem{
	"android":        OperatingSystemAndroid,
	"ios":            OperatingSystemIOS,
	"android huawei": OperatingSystemAndroidHuawei,
}

func OSFromStr(input string) (*OperatingSystem, error) {
	input = strings.ToLower(input)
	dt, ok := DeviceTypeIndex[input]
	if !ok {
		return nil, fmt.Errorf("unknown device type")
	}
	return &dt, nil
}
