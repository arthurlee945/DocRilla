package config

import (
	"fmt"
	"os"

	"github.com/arthurlee945/Docrilla/internal/errors"
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

var (
	configPath = "config/%s.config.yaml"
)

type Config struct {
	DSN string `yaml:"dsn"`
}

func Load() {

}

func loadFromFiels(cfg interface{}) error {
	wd, err := os.Getwd()
	if err != nil {
		return ErrNoWorkingDir.Wrap(err)
	}
	data, err := os.ReadFile(fmt.Sprintf("%s/%s", wd, fmt.Sprintf(configPath, getAppEnv())))
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
