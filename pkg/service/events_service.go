package service

import (
	"github.com/salesforceanton/events-api/config"
	"github.com/salesforceanton/events-api/domain"
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

func (s *EventsService) GetAll(userId int) ([]domain.Event, error) {
	return s.repo.GetAll(userId)
}

func (s *EventsService) GetById(userId, eventId int) (domain.Event, error) {
	return s.repo.GetById(userId, eventId)
}

func (s *EventsService) Upsert(userId int, event domain.Event) (int, error) {
	// TODO: Add update access check
	event.OrganizerId = userId
	return s.repo.Upsert(event)
}

func (s *EventsService) Delete(userId, eventId int) error {
	return s.repo.Delete(userId, eventId)
}
