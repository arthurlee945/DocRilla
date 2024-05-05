package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/logger"
	"go.uber.org/zap"
)

const (
	ErrJSONEncoding = errors.Error("server_failed_encoding: server failed encoding json.")
	ErrJSONDecoding = errors.Error("server_failed_decoding: server failed decoding json.")
)

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
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	errJsonObj := struct {
		Error string `json:"error"`
	}{
		Error: strings.Split(err.Error(), errors.ErrSeperator)[0],
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
