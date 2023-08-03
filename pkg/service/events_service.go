package service

import (
	"github.com/salesforceanton/events-api/config"
	"github.com/salesforceanton/events-api/pkg/repository"
)

type EventsService struct {
	repo repository.Events
	cfg  *config.Config
}

func NewEventsService(repo repository.Events, cfg *config.Config) *EventsService {
	return &EventsService{
		repo: repo,
		cfg:  cfg,
	}
}
