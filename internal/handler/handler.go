package handler

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_translations "github.com/go-playground/validator/v10/translations/pt_BR"
	"github.com/leonardonicola/tickethub/config"
)

var (
	log   *config.Logger
	trans ut.Translator
)

func InitHandler() {
	log = config.NewLogger()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		portuguese := pt_BR.New()
		uni := ut.New(portuguese, portuguese)
		trans, _ = uni.GetTranslator("pt_BR")
		pt_translations.RegisterDefaultTranslations(v, trans)
	}
}
