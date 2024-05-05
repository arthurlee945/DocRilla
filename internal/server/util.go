package server

import (
	"encoding/json"
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/errors"
)

// do i just wamt to use marshall and unmarshall?

func Encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return errors.ErrJSONEncoding.Wrap(err)
	}
	return nil
}

func Decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return v, errors.ErrJSONDecoding.Wrap(err)
	}
	return v, nil
}
