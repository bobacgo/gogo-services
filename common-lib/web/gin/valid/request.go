package valid

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func TransErr(ctx *gin.Context, err validator.ValidationErrors) error {
	t, _ := trans.GetTranslator(DefaultGetLanguage(ctx))
	msgErr := removeTopStruct(err.Translate(t))
	return errors.New(msgErr)
}

// InitRequestParamValidate 初始化翻译器
// 默认支持 en-英文和 zh-中文、zh_Hant_TW-繁体
// multipleTrans 支持其他国家或地区翻译器
func InitRequestParamValidate(multipleTrans ...TranslationLanguage) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		validate.SetTagName("validate")
		registerTrans(validate, multipleTrans...)
	}
}
