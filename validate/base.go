package validate

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
	"goblog/utils/maps"
)

var trans ut.Translator

// TransInit loca 通常取决于 http 请求头的 'Accept-Language'
func TransInit(local string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(enT, zhT, enT)

		var o bool
		trans, o = uni.GetTranslator(local)
		if !o {
			return fmt.Errorf("uni.GetTranslator(%s) failed", local)
		}
		//register translate
		// 注册翻译器
		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = chTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

type Base struct{}

// Check 校验
func (e *Base) Check(obj interface{}) (error, bool) {
	validate := validator.New()
	if err := validate.Struct(&obj); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err, false
		}
		var transfer maps.Transfer
		return fmt.Errorf(transfer.MapToJson(errs.Translate(trans))), false
	}

	return nil, true
}
