package server

import (
	"github.com/arthurlee945/Docrilla/internal/model"
	"go.uber.org/zap"
)

type ContextVars struct {
	User   *model.User
	Logger *zap.Logger
}
