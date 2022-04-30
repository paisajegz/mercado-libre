package translation

import (
	"log"

	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translation "github.com/go-playground/validator/v10/translations/es"
)

func TranslationValidation(v *validator.Validate) *ut.Translator {

	es := es.New()
	uni := ut.New(es, es)

	trans, found := uni.GetTranslator("es")

	if !found {
		log.Fatal("Traductor no encontrado")
	}

	if err := es_translation.RegisterDefaultTranslations(v, trans); err != nil {
		log.Fatal(err)
	}

	v.RegisterTranslation("isBool", trans, func(ut ut.Translator) error {
		return ut.Add("isBool", "{0} no es un dato valido", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isBool", fe.Field())
		return t
	})

	v.RegisterTranslation("is-dna", trans, func(ut ut.Translator) error {
		return ut.Add("is-dna", "{0} no es una cadena valida", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("is-dna", fe.Field())
		return t
	})

	return &trans

}
