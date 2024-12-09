package app

import (
	"github.com/AwesomeXjs/tma-bot/internal/client/http"
	"github.com/AwesomeXjs/tma-bot/internal/config"
	"github.com/AwesomeXjs/tma-bot/internal/controller"
	"github.com/AwesomeXjs/tma-bot/internal/service"
	"github.com/AwesomeXjs/tma-bot/pkg/logger"
	"github.com/go-telegram/bot"
	"go.uber.org/zap"
)

type ServiceProvider struct {
	botConfig config.IBotConfig

	resources config.IResources

	bot        *bot.Bot
	httpClient http.IHttpClient

	controller *controller.Controller
	svc        service.IService
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) BotConfig() config.IBotConfig {
	if s.botConfig == nil {
		cfg, err := config.NewBotConfig()
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

func (s *ServiceProvider) HttpClient() http.IHttpClient {
	if s.httpClient == nil {
		s.httpClient = http.New(http.NewClient())
	}
	return s.httpClient
}

func (s *ServiceProvider) Resources() config.IResources {
	if s.resources == nil {
		res, err := config.NewResources()
		if err != nil {
			logger.Error("failed to get resources", zap.Error(err))
		}
		s.resources = res
	}
	return s.resources
}

func (s *ServiceProvider) Service() service.IService {
	if s.svc == nil {
		s.svc = service.New(s.HttpClient(), s.Resources())
	}
	return s.svc
}

func (s *ServiceProvider) Controller() *controller.Controller {
	if s.controller == nil {
		s.controller = controller.NewController(s.Bot(), s.Service())
	}
	return s.controller
}
