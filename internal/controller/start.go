package controller

import (
	"context"

	"github.com/AwesomeXjs/tma-bot/pkg/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

func (c *Controller) Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:         update.Message.Chat.ID,
		Text:           "Hello world",
		ProtectContent: true,
	})
	if err != nil {
		logger.Error("failed to send message", zap.Error(err))
	}
}
