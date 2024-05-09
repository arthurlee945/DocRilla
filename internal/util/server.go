package util

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/logger"
	"go.uber.org/zap"
)

func Encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return errors.ErrJSONEncoding.Wrap(err)
	}
	return nil
}

func Decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return errors.ErrJSONDecoding.Wrap(err)
	}
	return nil
}

func HandleServerError(ctx context.Context, w http.ResponseWriter, err error) {
	if err == nil {
		w.WriteHeader(http.StatusOK)
	}

	switch {
	case errors.Is(err, errors.ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, errors.ErrInvalidRequest):
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, errors.ErrValidation):
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, errors.ErrUnauthorized):
		w.WriteHeader(http.StatusUnauthorized)
	case errors.Is(err, errors.ErrJSONEncoding):
		fallthrough
	case errors.Is(err, errors.ErrJSONDecoding):
		fallthrough
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	errJsonObj := struct {
		Error string `json:"error"`
	}{
		// Error: strings.Split(err.Error(), errors.ErrSeperator)[0],
		Error: err.Error(),
	}

	data, err := json.Marshal(errJsonObj)
	if err != nil {
		logger.From(ctx).Error("failed to serialize error response", zap.Error(err))
		data = []byte(`{"error": "internal server error"}`)
	}
	if _, err := w.Write(data); err != nil {
		logger.From(ctx).Error("failed to write error response", zap.Error(err))
	}
}
