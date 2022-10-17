package http

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func Handler[Response any](
	handleFunc func(_ http.ResponseWriter, _ *http.Request) (Response, error),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			httpErr        Error
			validationErrs validation.Errors
		)

		resp, err := handleFunc(w, r)

		switch {
		case err == nil:
			render.JSON(w, r, resp)
		case errors.As(err, &httpErr):
			Render(w, r, httpErr)
		case errors.As(err, &validationErrs):
			Render(w, r, NewValidationErrors(validationErrs))
		default:
			Render(w, r, NewHTTPError(err, http.StatusInternalServerError))
		}
	}
}
