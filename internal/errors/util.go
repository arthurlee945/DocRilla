package errors

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/arthurlee945/Docrilla/internal/logger"
	"go.uber.org/zap"
)

func RepoError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case Is(err, sql.ErrNoRows):
		return ErrNotFound.Wrap(err)
	case Is(err, ErrDecodeCursor):
		return ErrInvalidRequest.Wrap(err)
	default:
		return ErrUnknown.Wrap(err)
	}
}

func AuthError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case Is(err, sql.ErrNoRows) || Is(err, ErrInvalidToken) || Is(err, ErrUnauthorized):
		return ErrUnauthorized
	default:
		return ErrUnknown.Wrap(err)
	}
}

func ServerError(ctx context.Context, w http.ResponseWriter, err error) {
	if err == nil {
		w.WriteHeader(http.StatusOK)
	}

	switch {
	case Is(err, ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
	case Is(err, ErrInvalidRequest):
		w.WriteHeader(http.StatusBadRequest)
	case Is(err, ErrValidation):
		w.WriteHeader(http.StatusBadRequest)
	case Is(err, ErrUnauthorized):
		w.WriteHeader(http.StatusUnauthorized)
	case Is(err, ErrJSONEncoding):
		fallthrough
	case Is(err, ErrJSONDecoding):
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
