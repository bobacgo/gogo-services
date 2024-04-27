package validator

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	zhTwTranslations "github.com/go-playground/validator/v10/translations/zh_tw"
	"golang.org/x/exp/maps"
)

// Trans 定义一个全局翻译器T
var trans *ut.UniversalTranslator

type TranslationLanguage struct {
	Lt           locales.Translator
	RegisterFunc func(*validator.Validate, ut.Translator) error
}

func registerTrans(validate *validator.Validate, multipleTrans ...TranslationLanguage) {
	tMap := map[locales.Translator]func(*validator.Validate, ut.Translator) error{
		en.New():         enTranslations.RegisterDefaultTranslations,
		zh.New():         zhTranslations.RegisterDefaultTranslations,
		zh_Hant_TW.New(): zhTwTranslations.RegisterDefaultTranslations,
	}
	for _, tran := range multipleTrans {
		tMap[tran.Lt] = tran.RegisterFunc
	}
	ts := maps.Keys(tMap)
	trans = ut.New(ts[0], ts...) // 默认 en
	for t, register := range tMap {
		lt, _ := trans.GetTranslator(t.Locale())
		if err := register(validate, lt); err != nil { // 注册多语言翻译器
			panic(err)
		}
	}
}
