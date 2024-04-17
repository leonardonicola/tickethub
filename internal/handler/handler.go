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
	trans    ut.Translator
	validate *validator.Validate
)

func InitHandler() {
	log = config.NewLogger()
	portuguese := pt_BR.New()
	validate = validator.New()
	uni := ut.New(portuguese, portuguese)
	trans, _ = uni.GetTranslator("pt_BR")
	pt_translations.RegisterDefaultTranslations(validate, trans)
}
