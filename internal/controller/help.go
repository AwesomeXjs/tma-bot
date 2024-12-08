package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/AwesomeXjs/tma-bot/pkg/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

func (c *Controller) Help(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Я тебе помогу через 3 сек!",
	})

	go func() {
		time.Sleep(time.Second * 3)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Я тебе помог!",
		})
	}()
	fmt.Println(update.Message.Chat.ID)

	if err != nil {
		logger.Error("failed to send message", zap.Error(err))
	}
}
