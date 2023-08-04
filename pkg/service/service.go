package service

import (
	"github.com/salesforceanton/events-api/config"
	"github.com/salesforceanton/events-api/domain"
	"github.com/salesforceanton/events-api/pkg/repository"
)

type Service struct {
	Authorization
	Events
}

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Events interface {
	GetAll(userId int) ([]domain.Event, error)
	GetById(userId, eventId int) (domain.Event, error)
	Create(userId int, event domain.SaveEventRequest) (int, error)
	Update(userId, eventId int, event domain.SaveEventRequest) (domain.Event, error)
	Delete(userId, eventId int) error
}

func NewService(repos *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, cfg),
		Events:        NewEventsService(repos.Events, cfg),
	}
}
