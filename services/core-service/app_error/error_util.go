package app_error

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"services.core-service/i18n"
)

func translateToAppVE(i18n *i18n.I18n, lang string,
	valErrors validator.ValidationErrors, errCode string) []ValidationErrorField {

	res := make([]ValidationErrorField, len(valErrors))
	for i, valErr := range valErrors {
		res[i] = ValidationErrorField{
			Field:        valErr.Field(),
			Tag:          valErr.Tag(),
			ErrorMessage: i18n.MustLocalize(lang, fmt.Sprintf("%v.%v", errCode, valErr.Tag()), nil),
		}
	}
	return res
}

func HandleValidationErrors(language string, i18n *i18n.I18n, valErrors validator.ValidationErrors) *AppError {
	appErr := ValidationError(
		i18n.MustLocalize(language, COM0005, nil),
		"ErrValidation",
		translateToAppVE(i18n, language, valErrors, COM0005),
	)
	return appErr
}

func MustError(err error) {
	if err != nil {
		panic(err)
	}
}
