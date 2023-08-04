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

func (s *EventsService) Create(userId int, request domain.SaveEventRequest) (int, error) {
	return s.repo.Create(userId, request)
}

func (s *EventsService) Update(userId, eventId int, request domain.SaveEventRequest) (domain.Event, error) {
	return s.repo.Update(userId, eventId, request)
}

func (s *EventsService) Delete(userId, eventId int) error {
	return s.repo.Delete(userId, eventId)
}
