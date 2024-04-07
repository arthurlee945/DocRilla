package config

import (
	"context"
	"fmt"
	"os"

	"github.com/arthurlee945/Docrilla/internal/errors"
	"github.com/arthurlee945/Docrilla/internal/logger"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

const (
	ErrNoWorkingDir = errors.Error("failed getting working directory of Docrilla")
	ErrInvalidEnv   = errors.Error("Docrilla Environment is not set")
	ErrValidation   = errors.Error("invalid configuration")
	ErrEnvVars      = errors.Error("env parsing failed")
	ErrRead         = errors.Error("failed reading config file")
	ErrUnmarshal    = errors.Error("failed to unmarshall yaml file")
)

const configPath = "config/%s.config.yaml"

type Config struct {
	DSN string `yaml:"dsn"`
}

func Load(ctx context.Context) (*Config, error) {
	cfg := &Config{}
	if err := loadEnvironment(ctx, cfg); err != nil {
		return nil, err
	}

	v := validator.New()

	if err := v.Struct(cfg); err != nil {
		return nil, ErrValidation.Wrap(err)
	}
	return cfg, nil
}

func loadEnvironment(ctx context.Context, cfg interface{}) error {
	wd, err := os.Getwd()
	if err != nil {
		return ErrNoWorkingDir.Wrap(err)
	}
	path := fmt.Sprintf("%s/%s", wd, fmt.Sprintf(configPath, getAppEnv()))

	logger.From(ctx).Info("Loading Configuration", zap.String("path", path))

	data, err := os.ReadFile(path)
	if err != nil {
		return ErrRead.Wrap(err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return ErrUnmarshal.Wrap(err)
	}

	return nil
}

func getAppEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return "local"
	}
	return env
}
