package controller

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (c *Controller) Help(ctx context.Context, b *bot.Bot, update *models.Update) {
	replyMenu := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: "/start"},
				{Text: "Меню 2"},
			},
			{
				{Text: "Помощь"},
			},
		},
		ResizeKeyboard: true,
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Добро пожаловать! Выберите пункт меню:",
		ReplyMarkup: replyMenu,
	})

	inlineMenu := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Информация", CallbackData: "info"},
				{Text: "Помощь", CallbackData: "help"},
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Выберите опцию:",
		ReplyMarkup: inlineMenu,
	})
}
