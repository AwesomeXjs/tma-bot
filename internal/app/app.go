package app

import (
	"context"
	"flag"
	"fmt"

	"github.com/AwesomeXjs/tma-bot/pkg/closer"
	"github.com/AwesomeXjs/tma-bot/pkg/logger"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const (
	envPath = ".env"
)

// logLevel is a command-line flag for specifying the log level.
var logLevel = flag.String("l", "info", "log level")

type App struct {
	serviceProvider *ServiceProvider
	bot             *bot.Bot
}

func New(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.InitDeps(ctx)
	if err != nil {
		logger.Fatal("failed to initialize dependencies", zap.Error(err))
		return nil, err
	}
	return app, nil
}

func (app *App) InitDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		app.InitConfig,
		app.initServiceProvider,
		app.initBot,
	}
	for _, fun := range inits {
		if err := fun(ctx); err != nil {
			// Log fatal error if any dependency initialization fails
			logger.Fatal("failed to init deps", zap.Error(err))
		}
	}
	return nil
}

func (app *App) InitConfig(_ context.Context) error {
	err := godotenv.Load(envPath)
	if err != nil {
		logger.Error("Error loading .env file", zap.String("path", envPath))
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return err
}

func (app *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll() // Close all services/resources
		closer.Wait()     // Wait for all services to close
	}()
	err := app.runBot(ctx)
	if err != nil {
		return err
	}
	return nil
}

// initServiceProvider initializes the service provider.
func (app *App) initServiceProvider(_ context.Context) error {
	app.serviceProvider = NewServiceProvider() // Create a new service provider
	return nil
}

func (app *App) initBot(_ context.Context) error {
	b, err := bot.New(app.serviceProvider.BotConfig().GetToken())
	if err != nil {
		return fmt.Errorf("failed to create bot: %v", err)
	}
	app.serviceProvider.Controller(b).RegisterHandlers()

	app.bot = b

	return nil
}

func (app *App) runBot(ctx context.Context) error {
	flag.Parse()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(logLevel)))

	fmt.Println("Bot started")
	app.bot.Start(ctx)

	return nil
}
