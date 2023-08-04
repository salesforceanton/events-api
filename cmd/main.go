package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	eventsapi "github.com/salesforceanton/events-api"
	"github.com/salesforceanton/events-api/config"
	"github.com/salesforceanton/events-api/pkg/handler"
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
		log.WithFields(log.Fields{
			"problem": fmt.Sprintf("Error with config initialization: %s", err.Error()),
		}).Error(err)
		return
	}

	// Connect to DB
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"problem": fmt.Sprintf("Error with database connect: %s", err.Error()),
		}).Error(err)
		return
	}

	// Init dependenties
	repos := repository.NewRepository(db)
	services := service.NewService(repos, cfg)
	handler := handler.NewHandler(services)

	// Run server
	server := new(eventsapi.Server)
	go func() {
		if err := server.Run(cfg.Port, handler.InitRoutes()); err != nil {
			log.WithFields(log.Fields{
				"problem": fmt.Sprintf("Error with server running/or server in closing: %s", err.Error()),
			}).Error(err)
			return
		}
	}()

	// Gracefull shutdown
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	if err := server.Shutdown(context.Background()); err != nil {
		log.WithFields(log.Fields{
			"problem": fmt.Sprintf("Error when server was shutting down: %s", err.Error()),
		}).Error(err)
		return
	}

	if err := db.Close(); err != nil {
		log.WithFields(log.Fields{
			"problem": fmt.Sprintf("Error with closing database connection: %s", err.Error()),
		}).Error(err)
		return
	}

	log.Info("Server shutdown successfully")
}
