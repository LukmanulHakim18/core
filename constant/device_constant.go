package constant

const (
	// for indonesian language
	DEVICE_LANG_ID = "ID"
	// for english language
	DEVICE_LANG_EN = "EN"
)

type DeviceLang string

type DeviceUUID string

func (d DeviceLang) IsEn() bool {
	return d == DEVICE_LANG_EN
}

func (d DeviceLang) IsId() bool {
	return d == DEVICE_LANG_ID
}
