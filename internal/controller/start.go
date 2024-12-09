package controller

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/tma-bot/internal/model"
	"github.com/AwesomeXjs/tma-bot/pkg/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

func (c *Controller) Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println("id: ", update.Message.Chat.ID)
	fmt.Println("username: ", update.Message.Chat.Username)
	fmt.Println("first_name: ", update.Message.Chat.FirstName)
	fmt.Println("last_name: ", update.Message.Chat.LastName)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:         update.Message.Chat.ID,
		Text:           "Hello world",
		ProtectContent: true,
	})
	
	user := &model.User{
		ID:        int(update.Message.Chat.ID),
		Username:  update.Message.Chat.Username,
		FirstName: update.Message.Chat.FirstName,
		LastName:  update.Message.Chat.LastName,
		IsPremium: 0,
	}
	err = c.svc.Registration(user)
	if err != nil {
		logger.Error("failed to register", zap.Error(err))
	}

	if err != nil {
		logger.Error("failed to send message", zap.Error(err))
	}
}
