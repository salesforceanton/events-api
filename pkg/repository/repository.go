package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Authorization
	Events
}

type Authorization interface {
}

type Events interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Events:        NewEventsPostgres(db),
	}
}
