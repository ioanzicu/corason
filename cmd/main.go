package main

import (
	"context"
	"corason/config"
	"corason/internal/application/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Booom: %v", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Error running the logger: %v", err)
	}
	sugar := logger.Sugar()
	defer logger.Sync()

	fmt.Println("Start the server...")

	app := NewApp(config, sugar)
	app.GracefulShutdown()
}

type App struct {
	Logger *zap.SugaredLogger
	Server *http.Server
}

func NewApp(config config.Config, logger *zap.SugaredLogger) *App {
	return &App{
		Logger: logger,
		Server: &http.Server{
			Addr:           fmt.Sprintf("0.0.0.0:%s", config.ApplicationPort),
			Handler:        router.NewRouter(),
			ReadTimeout:    config.HTTPServerReadTimeout,
			WriteTimeout:   config.HTTPServerWriteTimeout,
			IdleTimeout:    config.HTTPServerIdleimeout,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (app *App) GracefulShutdown() {
	// Graceful shutdown
	serverErrors := make(chan error, 1)
	// Start the listener
	go func() {
		// sugar.info(fmt.Sprintf("HTTP is running: %v", addr))
		serverErrors <- app.Server.ListenAndServe()
	}()

	// Make a channel to listen for an interrupt or terminate signal from OS
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM)

	// Block waiting for a receive on either channel
	select {
	case err := <-serverErrors:
		app.Logger.Errorf("error starting server: %v", err)

	case <-osSignal:
		// Attempt a graceful 5 second shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Attempt the graceful shutdown by closing the listener and completeing all inflight requests
		if err := app.Server.Shutdown(ctx); err != nil {
			app.Logger.Errorf("Cannot stop the server gracefully: %v", err)
			app.Logger.Info("Forcing shutdown")
			if err := app.Server.Close(); err != nil {
				app.Logger.Errorf("Could not stop the HTTP server")
			}
		}
	}

	app.Logger.Info("Shut down successful!!!")
}
