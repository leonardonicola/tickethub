package handler

import (
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_translations "github.com/go-playground/validator/v10/translations/pt_BR"
	"github.com/leonardonicola/tickethub/config"
)

var (
	log      *config.Logger
	validate *validator.Validate
	trans    ut.Translator
)

func Validate(dto interface{}) validator.ValidationErrorsTranslations {
	errs := validate.Struct(dto)
	if errs != nil {
		return errs.(*validator.ValidationErrors).Translate(trans)
	}
	return nil
}

func InitHandler() {
	log = config.NewLogger()
	portuguese := pt_BR.New()
	validate = validator.New()
	uni := ut.New(portuguese, portuguese)
	trans, _ = uni.GetTranslator("pt_BR")
	pt_translations.RegisterDefaultTranslations(validate, trans)
}
