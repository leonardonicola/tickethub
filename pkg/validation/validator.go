package validation

import (
	"fmt"

	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func Validate(dto interface{}) validator.ValidationErrorsTranslations {
	errs := validate.Struct(dto)
	if errs != nil {
		return errs.(validator.ValidationErrors).Translate(trans)
	}
	return nil
}

func InitValidator() error {
	portuguese := pt_BR.New()
	validate = validator.New()
	uni := ut.New(portuguese, portuguese)
	trans, _ = uni.GetTranslator("pt_BR")
	if err := pt_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		return fmt.Errorf("registering default translation: %v", err)
	}
	return nil
}
