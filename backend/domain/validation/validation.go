package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validatortranslationsen "github.com/go-playground/validator/v10/translations/en"
	"github.com/pkg/errors"
)

type Validator struct {
	validate   *validator.Validate
	translator ut.Translator
}

func NewValidator() (*Validator, error) {
	vldtr := &Validator{
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
	{
		vldtr.translator, _ = ut.New(en.New()).GetTranslator("en")

		if err := validatortranslationsen.RegisterDefaultTranslations(vldtr.validate, vldtr.translator); err != nil {
			return nil, errors.Wrap(err, "failed to register default translations")
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

	return vldtr, nil
}

func (vldtr *Validator) Var(v any, tag string) error {
	return vldtr.handleError(vldtr.validate.Var(v, tag))
}

func (vldtr *Validator) Struct(v any) error {
	return vldtr.handleError(vldtr.validate.Struct(v))
}

func (vldtr *Validator) handleError(err error) error {
	if err != nil {
		vldtnErrs, ok := err.(validator.ValidationErrors)
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

func translate(translator ut.Translator, fldErr validator.FieldError) string {
	msg, err := translator.T(fldErr.Tag(), fldErr.Field())
	if err != nil {
		return ""
	}

	return msg
}
