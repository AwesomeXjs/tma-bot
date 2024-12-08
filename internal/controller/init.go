package controller

import "github.com/go-telegram/bot"

type Controller struct {
	bot *bot.Bot
}

func NewController(bot *bot.Bot) *Controller {
	return &Controller{
		bot: bot,
	}
}
