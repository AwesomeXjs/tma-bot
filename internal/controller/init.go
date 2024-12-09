package controller

import (
	"github.com/AwesomeXjs/tma-bot/internal/service"
	"github.com/go-telegram/bot"
)

type Controller struct {
	bot *bot.Bot

	svc service.IService
}

func NewController(bot *bot.Bot, svc service.IService) *Controller {
	return &Controller{
		bot: bot,
		svc: svc,
	}
}
