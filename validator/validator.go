package validator

import "github.com/go-playground/validator/v10"

var validate *validator.Validate = validator.New()

// Validate validates the fields of the given struct according to the
// validation tags defined on its fields. It uses the go-playground/validator
// package under the hood. If validation fails, it returns an error describing
// the first validation issue encountered; otherwise, it returns nil.
func Validate[T any](schema *T) error {
	return validate.Struct(schema)
}
