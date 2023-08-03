package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/salesforceanton/events-api/domain"
)

type EventsPostgres struct {
	db *sqlx.DB
}

func NewEventsPostgres(db *sqlx.DB) *EventsPostgres {
	return &EventsPostgres{db: db}
}

func (r *EventsPostgres) GetAll(userId int) ([]domain.Event, error) {
	var result []domain.Event

	query := fmt.Sprintf(
		"SELECT id, title, timezoneId, startDatetime, description FROM %s WHERE organizerId=$1",
		EVENTS_TABLE,
	)
	err := r.db.Select(&result, query, userId)

	return result, err
}

func (r *EventsPostgres) GetById(userId, eventId int) (domain.Event, error) {
	var result domain.Event

	query := fmt.Sprintf(
		"SELECT id, title, timezoneId, startDatetime, description FROM %s WHERE organizerId=$1 AND id=$2",
		EVENTS_TABLE,
	)
	err := r.db.Get(&result, query, userId, eventId)

	return result, err
}

func (r *EventsPostgres) Upsert(event domain.Event) (int, error) {
	var result int

	query := fmt.Sprintf(
		`INSERT INTO %s (id, title, timezoneId, startDatetime, description, organizerId) 
		VALUES ($1 $2 $3 $4 $5 $6)
		ON CONFLICT (id) DO UPDATE
		RETURNING id`,
		EVENTS_TABLE,
	)
	row := r.db.QueryRow(query, event.Id, event.Title, event.StartDatetime, event.Description, event.OrganizerId)
	if err := row.Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}

func (r *EventsPostgres) Delete(userId, eventId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE organizerId=$1 AND id=$2", EVENTS_TABLE)

	r.db.Exec(query, eventId)
	_, err := r.db.Exec(query, eventId)

	return err
}
