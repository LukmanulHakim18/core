package notification

import (
	"context"
	"encoding/json"

	"github.com/LukmanulHakim18/core/constant"
	"github.com/LukmanulHakim18/core/metadata"
)

type PushNotifMessage struct {
	Type             string           `json:"type"`
	ProductType      string           `json:"product_type"`
	Message          string           `json:"message"`
	OrderID          string           `json:"order_id"`
	Title            string           `json:"title"`
	LocalizedMessage LocalizedMessage `json:"localized_message"`
	LocalizedTitle   LocalizedMessage `json:"localized_title"`
}

type LocalizedMessage struct {
	EN string `json:"en"`
	ID string `json:"id"`
}

func (m *PushNotifMessage) BuildMessage(ctx context.Context) []byte {
	deviceLang := metadata.GetDeviceLanguageFromCtx(ctx)

	m.Title = m.LocalizedTitle.CustomLocMessage(deviceLang)
	m.Message = m.LocalizedTitle.CustomLocMessage(deviceLang)

	byt, _ := json.Marshal(m)
	return byt

}

func (lm *LocalizedMessage) CustomLocMessage(lang constant.DeviceLang) string {
	if lang == constant.DEVICE_LANG_ID {
		return lm.ID
	}
	return lm.EN
}
