package config

import (
	"fmt"
	"os"

	"github.com/AwesomeXjs/tma-bot/pkg/logger"
	"go.uber.org/zap"
)

const (
	EnvBotToken = "BOT_TOKEN"
)

type IBotConfig interface {
	GetToken() string
}

type BotConfig struct {
	token string
}

func NewBotConfig() (IBotConfig, error) {
	token := os.Getenv(EnvBotToken)

	if len(token) == 0 {
		logger.Error("failed to get bot token", zap.String("env", EnvBotToken))
		return nil, fmt.Errorf("env %s not set", EnvBotToken)
	}
	return &BotConfig{
		token: token,
	}, nil
}

func (c *BotConfig) GetToken() string {
	return c.token
}
