package validator

import "errors"

var (
	errValidation = errors.New("request validation failed")
)

// Validator checks request data, ensuring it conforms to the expected format and standards.
type Validator[T any] interface {
	Validate() error
}
