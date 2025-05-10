package lang

import (
	"context"

	"github.com/LukmanulHakim18/core/constant"
	"github.com/LukmanulHakim18/core/metadata"
)

func CustomTextByLanguage(ctx context.Context, textEn, textId string) string {
	lang := metadata.GetDeviceLanguageFromCtx(ctx)
	if lang == constant.DEVICE_LANG_ID {
		return textId
	}
	return textEn
}
