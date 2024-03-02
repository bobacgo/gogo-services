package valid

import "github.com/gin-gonic/gin"

var LanguageCtxKey = "language"

type GetRequestLanguageFunc func(ctx *gin.Context) string

var DefaultGetLanguage GetRequestLanguageFunc = func(ctx *gin.Context) string {
	lang := ctx.GetString(LanguageCtxKey)
	if lang == "" {
		lang = "en"
	}
	return lang
}
