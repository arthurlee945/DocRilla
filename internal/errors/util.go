package errors

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

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
	default:
		return ErrUnknown.Wrap(err)
	}
}

func ServerHandleError(ctx context.Context, w http.ResponseWriter, err error) {
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
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	errJsonObj := struct {
		Error string `json:"error"`
	}{
		Error: strings.Split(err.Error(), ErrSeperator)[0],
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
