package main

import (
	"fmt"
	"os"

	eventsapi "github.com/salesforceanton/events-api"
	"github.com/salesforceanton/events-api/config"
	"github.com/salesforceanton/events-api/pkg/handler"
	"github.com/salesforceanton/events-api/pkg/repository"
	"github.com/salesforceanton/events-api/pkg/service"
	log "github.com/sirupsen/logrus"
)

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
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	server := new(eventsapi.Server)
	if err := server.Run(cfg.Port, handler.InitRoutes()); err != nil {
		log.WithFields(log.Fields{
			"problem": fmt.Sprintf("Error with server running: %s", err.Error()),
		}).Error(err)
		return
	}
}
