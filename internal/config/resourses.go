package config

import (
	"os"

	"github.com/AwesomeXjs/tma-bot/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	BaseURL = "BASE_URL"

	registerPrefix = "/api/v1/registration"
)

type Resources struct {
	registerURL string
}

type IResources interface {
	RegisterURL() string
}

func NewResources() (*Resources, error) {
	baseURL := os.Getenv(BaseURL)
	if len(baseURL) == 0 {
		logger.Error("failed to get base url", zap.String("env", BaseURL))
		return nil, errors.New("env " + BaseURL + " not set")
	}

	return &Resources{
		registerURL: baseURL + registerPrefix,
	}, nil
}

func (r *Resources) RegisterURL() string {
	return r.registerURL
}
