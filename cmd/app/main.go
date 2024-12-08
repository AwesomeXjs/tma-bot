package main

import (
	"context"

	"github.com/AwesomeXjs/tma-bot/internal/app"
	"github.com/AwesomeXjs/tma-bot/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	bot, err := app.New()
	if err != nil {
		logger.Fatal("failed to create app", zap.Error(err))
	}

	err = bot.Run(ctx)
	if err != nil {
		logger.Fatal("failed to run bot", zap.Error(err))
	}
}
