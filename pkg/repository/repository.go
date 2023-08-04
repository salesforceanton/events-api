package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/salesforceanton/events-api/domain"
)

type Repository struct {
	Authorization
	Events
}

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type Events interface {
	GetAll(userId int) ([]domain.Event, error)
	GetById(userId, eventId int) (domain.Event, error)
	Create(userId int, request domain.SaveEventRequest) (int, error)
	Update(userId, eventId int, request domain.SaveEventRequest) (domain.Event, error)
	Delete(userId, eventId int) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Events:        NewEventsPostgres(db),
	}
}
