package app

import (
	"context"
	"flag"
	"fmt"
	"sync"

	"github.com/AwesomeXjs/tma-bot/pkg/closer"
	"github.com/AwesomeXjs/tma-bot/pkg/logger"
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
}

func New() (*App, error) {
	app := &App{}
	err := app.InitDeps()
	if err != nil {
		logger.Fatal("failed to initialize dependencies", zap.Error(err))
		return nil, err
	}
	return app, nil
}

func (app *App) InitDeps() error {
	inits := []func() error{
		app.InitConfig,
		app.initServiceProvider,
	}
	for _, fun := range inits {
		if err := fun(); err != nil {
			logger.Error("failed to init deps", zap.Error(err))
			return err
		}
	}

	return nil
}

func (app *App) InitConfig() error {
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

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		app.runBot(ctx)
	}()
	wg.Wait()
	return nil
}

// initServiceProvider initializes the service provider.
func (app *App) initServiceProvider() error {
	app.serviceProvider = NewServiceProvider() // Create a new service provider
	return nil
}

func (app *App) runBot(ctx context.Context) {
	flag.Parse()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(logLevel)))

	app.serviceProvider.Controller().RegisterHandlers()
	logger.Info("bot started")
	app.serviceProvider.Bot().Start(ctx)
}
