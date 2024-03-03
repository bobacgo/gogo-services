package valid

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// ShouldBind 解析请求参数并校验
// T 返回 请求数据
//
// req, err := valid.ShouldBind[model.CreateUserReq](ctx)
//
//	if err != nil {
//		status.Ret(c, status.New(codes.BadRequest).WithError(err))
//	    return
//	}
func ShouldBind[T any](ctx *gin.Context) (T, error) {
	var obj T
	if err := ctx.ShouldBind(&obj); err != nil {
		var errs validator.ValidationErrors
		if !errors.As(err, &errs) { // 非validator.ValidationErrors类型错误直接返回
			return obj, err
		}
		// validator.ValidationErrors类型错误则进行翻译
		if trans == nil {
			InitRequestParamValidate()
		}
		t, _ := trans.GetTranslator(DefaultGetLanguage(ctx))
		msgErr := removeTopStruct(errs.Translate(t))
		return obj, errors.New(msgErr)
	}
	return obj, nil
}

// InitRequestParamValidate 初始化翻译器
// 默认支持 en-英文和 zh-中文、zh_Hant_TW-繁体
// multipleTrans 支持其他国家或地区翻译器
func InitRequestParamValidate(multipleTrans ...TranslationLanguage) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			// skip if tag key says it should be ignored
			if name == "-" {
				return ""
			}
			return name
		})

		registerTrans(validate, multipleTrans...)
	}
}
