package repository

import "github.com/jmoiron/sqlx"

type EventsPostgres struct {
	db *sqlx.DB
}

func NewEventsPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
