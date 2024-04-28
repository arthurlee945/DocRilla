package server

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/model"
	"go.uber.org/zap"
)

const ContextVarKey = "server.context.vars"

type ContextVars struct {
	CTX    context.Context
	User   *model.User
	Logger *zap.Logger
}
