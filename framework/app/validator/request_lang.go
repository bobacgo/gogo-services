package validator

import (
	"context"
)

var LanguageCtxKey = "language"

type GetRequestLanguageFunc func(ctx context.Context) string

var DefaultGetLanguage GetRequestLanguageFunc = func(ctx context.Context) string {
	lang, _ := ctx.Value(LanguageCtxKey).(string)
	if lang == "" {
		lang = "en"
	}
	return lang
}
