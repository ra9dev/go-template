package api

import validation "github.com/go-ozzo/ozzo-validation/v4"

type (
	Sanitizer interface {
		Sanitize()
	}

	SafeRequest interface {
		Sanitizer
		validation.Validatable
	}
)

// IsSafeRequest to check if your request can be processed
func IsSafeRequest(req SafeRequest) error {
	req.Sanitize()

	return req.Validate() // nolint:wrapcheck
}
