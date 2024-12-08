package controller

import "github.com/go-telegram/bot"

func (c *Controller) RegisterHandlers() {
	c.bot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, c.Start)
}
