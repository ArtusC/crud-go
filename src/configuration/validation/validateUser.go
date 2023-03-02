package validation

import (
	"encoding/json"
	"errors"

	restErr "github.com/ArtusC/crud-go/src/configuration/restErr"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {

		en := en.New()

		// universal translator
		unt := ut.New(en, en)

		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}

}

func ValidateUserError(validation_err error) *restErr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return restErr.NewBadRequestError("Ivalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []restErr.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := restErr.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}
		return restErr.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return restErr.NewBadRequestError("Error trying to convert fields")
	}
}
