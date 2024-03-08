package validator

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

var Default *validator.Validate

func init() {
	if Default == nil {
		Default = validator.New()
		registerTrans(Default)
	}
}

// Init 初始化翻译器
// 默认支持 en-英文和 zh-中文、zh_Hant_TW-繁体
// multipleTrans 支持其他国家或地区翻译器
func Init(multipleTrans ...TranslationLanguage) {
	Default = validator.New()
	// 修改gin框架中的Validator引擎属性，实现自定制
	// 注册一个获取json tag的自定义方法
	registerTrans(Default, multipleTrans...)
}

func Struct(obj any) error {
	return TransErr(Default.Struct(obj))
}

func StructCtx(ctx context.Context, obj any) error {
	return TransErrCtx(ctx, Default.StructCtx(ctx, obj))
}

func TransErrZh(err error) error {
	return TransErrLocale("zh", err)
}

func TransErr(err error) error {
	return TransErrLocale("en", err)
}

func TransErrCtx(ctx context.Context, err error) error {
	return TransErrLocale(DefaultGetLanguage(ctx), err)
}

func TransErrLocale(locale string, err error) error {
	var verr validator.ValidationErrors
	if !errors.As(err, &verr) {
		return err
	}
	t, _ := trans.GetTranslator(locale)
	msgErr := removeTopStruct(verr.Translate(t))
	return errors.New(msgErr)
}

// 去掉结构体名称前缀
func removeTopStruct(fields map[string]string) string {
	msgErrs := strings.Builder{}
	for _, err := range fields {
		msgErrs.WriteString(err)
		msgErrs.WriteString(", ")
	}
	return strings.TrimRight(msgErrs.String(), ", ")
}
