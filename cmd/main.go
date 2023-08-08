package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	eventsapi "github.com/salesforceanton/events-api"
	"github.com/salesforceanton/events-api/config"
	"github.com/salesforceanton/events-api/pkg/logger"
	"github.com/salesforceanton/events-api/pkg/repository"
	"github.com/salesforceanton/events-api/pkg/service"
	grpc_client "github.com/salesforceanton/events-api/pkg/transport/grpc"
	handler "github.com/salesforceanton/events-api/pkg/transport/rest"
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

	// Connect to DB
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logger.LogExecutionIssue(err)
		return
	}

	// Connect to grpc-loggerbin
	loggerbin, err := grpc_client.NewClient(cfg.LoggerbinPort)
	if err != nil {
		logger.LogExecutionIssue(err)
		return
	}

	// Init dependenties
	repos := repository.NewRepository(db)
	services := service.NewService(repos, cfg)
	handler := handler.NewHandler(services, loggerbin)

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
