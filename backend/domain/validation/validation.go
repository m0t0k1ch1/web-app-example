package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	pgvalidator "github.com/go-playground/validator/v10"
	validatortranslationsen "github.com/go-playground/validator/v10/translations/en"
	"github.com/pkg/errors"
)

var (
	vldtr *validator
)

func init() {
	if err := initValidator(); err != nil {
		panic(err)
	}
}

type validator struct {
	validate   *pgvalidator.Validate
	translator ut.Translator
}

func initValidator() error {
	translate := func(translator ut.Translator, fldErr pgvalidator.FieldError) string {
		msg, err := translator.T(fldErr.Tag(), fldErr.Field())
		if err != nil {
			return ""
		}

		return msg
	}

	vldtr := &validator{
		validate: pgvalidator.New(pgvalidator.WithRequiredStructEnabled()),
	}
	{
		vldtr.translator, _ = ut.New(en.New()).GetTranslator("en")

		if err := validatortranslationsen.RegisterDefaultTranslations(vldtr.validate, vldtr.translator); err != nil {
			return errors.Wrap(err, "failed to register default translations")
		}
	}
	{
		vldtr.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("en")
		})
	}
	{
		tag := "hostname_rfc1123"

		vldtr.validate.RegisterTranslation(tag, vldtr.translator, func(translator ut.Translator) error {
			return translator.Add(tag, "{0} must be a valid hostname", false)
		}, translate)
	}
	{
		tag := "required"

		vldtr.validate.RegisterTranslation(tag, vldtr.translator, func(translator ut.Translator) error {
			return translator.Add(tag, "{0} is required", true)
		}, translate)
	}

	return nil
}

func Var(v any, tag string) error {
	return handleError(vldtr.validate.Var(v, tag))
}

func Struct(v any) error {
	return handleError(vldtr.validate.Struct(v))
}

func handleError(err error) error {
	if err != nil {
		vldtnErrs, ok := err.(pgvalidator.ValidationErrors)
		if ok {
			msgs := make([]string, len(vldtnErrs))
			for idx, vldtnErr := range vldtnErrs {
				msgs[idx] = vldtnErr.Translate(vldtr.translator)
			}

			return errors.New(strings.Join(msgs, "; "))
		}
	}

	return err
}
