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
	CreateUser(domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Events interface {
}

func NewService(repos *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, cfg),
		Events:        NewEventsService(repos.Events, cfg),
	}
}
