package http

import (
	"net/http"

	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/ra9dev/go-template/pkg/sre/log"
)

type Error struct {
	err     error
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func NewHTTPError(err error, code int) Error {
	message := err.Error()
	if code >= http.StatusInternalServerError {
		message = http.StatusText(code)
	}

	return Error{
		err:     err,
		Code:    code,
		Message: message,
	}
}

func (e Error) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Code)

	if e.Code >= http.StatusInternalServerError {
		log.Errorf(r.Context(), "[HTTP] internal error: %v", e.err)
	}

	return nil
}

type ValidationErrors map[string]string

func NewValidationErrors(source validation.Errors) ValidationErrors {
	errs := make(map[string]string, len(source))

	for k, v := range source {
		errs[k] = v.Error()
	}

	return errs
}

func (e ValidationErrors) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusUnprocessableEntity)

	return nil
}

func Render(w http.ResponseWriter, r *http.Request, v render.Renderer) {
	if err := render.Render(w, r, v); err != nil {
		log.Errorf(r.Context(), "render error: %v", err)

		return
	}
}
