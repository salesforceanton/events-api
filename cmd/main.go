package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/sessions"
	eventsapi "github.com/salesforceanton/events-api"
	"github.com/salesforceanton/events-api/config"
	"github.com/salesforceanton/events-api/pkg/handler"
	"github.com/salesforceanton/events-api/pkg/logger"
	"github.com/salesforceanton/events-api/pkg/repository"
	"github.com/salesforceanton/events-api/pkg/service"
	log "github.com/sirupsen/logrus"
)

// @title Events API
// @version 1.0
// @description API Server for booking Events

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	// Initialize app configuration
	cfg, err := config.InitConfig()
	if err != nil {
		logger.LogExecutionIssue(err)
		return
	}
	// Init session store
	// TODO: move params to yaml config
	sessionStore := sessions.NewCookieStore([]byte(cfg.SessionKey))
	sessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15, // Expiration time
		HttpOnly: true,
	}

	// Connect to DB
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logger.LogExecutionIssue(err)
		return
	}

	// Init dependenties
	repos := repository.NewRepository(db)
	services := service.NewService(repos, cfg)
	handler := handler.NewHandler(services, sessionStore)

	// Run server
	server := new(eventsapi.Server)
	go func() {
		if err := server.Run(cfg.Port, handler.InitRoutes()); err != nil {
			logger.LogExecutionIssue(err)
			return
		}
	}()

	// Gracefull shutdown
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	if err := server.Shutdown(context.Background()); err != nil {
		logger.LogExecutionIssue(err)
		return
	}

	if err := db.Close(); err != nil {
		logger.LogExecutionIssue(err)
		return
	}

	log.Info("Server shutdown successfully")
}
