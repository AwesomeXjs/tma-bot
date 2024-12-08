package app

import (
	"github.com/AwesomeXjs/tma-bot/internal/controller"
	"github.com/AwesomeXjs/tma-bot/pkg/logger"
	"github.com/go-telegram/bot"
	"go.uber.org/zap"
)

type ServiceProvider struct {
	botConfig IBotConfig

	bot *bot.Bot

	controller *controller.Controller
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) BotConfig() IBotConfig {
	if s.botConfig == nil {
		cfg, err := NewBotConfig()
		if err != nil {
			logger.Error("failed to get bot config", zap.Error(err))
		}
		s.botConfig = cfg
	}
	return s.botConfig
}

func (s *ServiceProvider) Bot() *bot.Bot {
	if s.bot == nil {
		opts := []bot.Option{
			// bot.WithDefaultHandler(s.Controller(s.bot).DefaultHandler),
		}

		b, err := bot.New(s.BotConfig().GetToken(), opts...)
		if err != nil {
			logger.Error("failed to create bot", zap.Error(err))
		}
		s.bot = b
	}
	return s.bot
}

func (s *ServiceProvider) Controller(bot *bot.Bot) *controller.Controller {
	if s.controller == nil {
		s.controller = controller.NewController(bot)
	}
	return s.controller
}
