package controller

import (
	"context"

	"github.com/AwesomeXjs/tma-bot/internal/controller/filters"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (c *Controller) RegisterHandlers() {
	arr := []struct {
		filter  func(*models.Update) bool
		handler func(ctx context.Context, b *bot.Bot, update *models.Update)
	}{
		{filters.IsStart, c.Start},
		{filters.IsHelp, c.Help},
		{filters.IsPhoto, c.Photo},
		{filters.IsVideo, c.Video},
		{filters.IsMyID, c.MyID},
	}

	for _, v := range arr {
		c.bot.RegisterHandlerMatchFunc(v.filter, v.handler)
	}
}
