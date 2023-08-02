package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/salesforceanton/events-api/config"
)

const (
	POSTGRESS_DB_TYPE = "postgres"
	USERS_TABLE       = "users"
	EVENTS_TABLE      = "events"
)

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(fmt.Sprintf("%s://%s:%s@%s/%s", POSTGRESS_DB_TYPE, cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBName))
	db, err := sqlx.Open(POSTGRESS_DB_TYPE, pgUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
