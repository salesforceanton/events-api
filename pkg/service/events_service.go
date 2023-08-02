package service

import "github.com/salesforceanton/events-api/pkg/repository"

type EventsService struct {
	repo repository.Events
}

func NewEventsService(repo repository.Events) *EventsService {
	return &EventsService{repo: repo}
}
