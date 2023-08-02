package service

import "github.com/salesforceanton/events-api/pkg/repository"

type Service struct {
	Authorization
	Events
}

type Authorization interface {
}

type Events interface {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Events:        NewEventsService(repos.Events),
	}
}
