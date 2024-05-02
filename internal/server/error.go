package server

import "github.com/arthurlee945/Docrilla/internal/errors"

const (
	ErrJSONEncoding = errors.Error("server_failed_encoding: server failed encoding json.")
	ErrJSONDecoding = errors.Error("server_failed_decoding: server failed decoding json.")
)
