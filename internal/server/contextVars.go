package server

import (
	"context"

	"github.com/arthurlee945/Docrilla/internal/logger"
	"github.com/arthurlee945/Docrilla/internal/model"
)

const ContextVarKey = "server.context.vars"

type ContextVars struct {
	CTX    context.Context
	User   *model.User
	Logger *logger.Logger
}
