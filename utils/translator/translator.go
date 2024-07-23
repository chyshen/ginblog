// @Author scy
// @Time 2024/7/24 0:40
// @File translator.go

package translator

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"

	"reflect"
)

var Trans ut.Translator

// TransInit 翻译验证信息  local 通常取决于 http 请求头的 'Accept-Language'
func TransInit(local string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(enT, zhT, enT)

		var o bool
		Trans, o = uni.GetTranslator(local)
		fmt.Printf("local: %v\n", local)
		if !o {
			return fmt.Errorf("uni.GetTranslator(%s) failed", local)
		}
		//register translate
		// 注册翻译器
		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, Trans)
			v.RegisterTagNameFunc(func(field reflect.StructField) string {
				label := field.Tag.Get("label")
				if label == "" {
					return field.Name
				}
				return label
			})
		default:
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		return
	}
	return
}
