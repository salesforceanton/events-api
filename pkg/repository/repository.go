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
	CreateUser(domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type Events interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Events:        NewEventsPostgres(db),
	}
}
