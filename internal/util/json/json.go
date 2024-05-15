package json

import (
	jsonlib "encoding/json"
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/errors"
)

func Encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := jsonlib.NewEncoder(w).Encode(v); err != nil {
		return errors.ErrJSONEncoding.Wrap(err)
	}
	return nil
}

func Decode(r *http.Request, v interface{}) error {
	if err := jsonlib.NewDecoder(r.Body).Decode(v); err != nil {
		return errors.ErrJSONDecoding.Wrap(err)
	}
	return nil
}
