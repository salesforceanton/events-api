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
		"SELECT id, title, timezoneId, startDatetime, organizerId, description FROM %s WHERE organizerId=$1",
		EVENTS_TABLE,
	)
	err := r.db.Select(&result, query, userId)

	return result, err
}

func (r *EventsPostgres) GetById(userId, eventId int) (domain.Event, error) {
	var result domain.Event

	query := fmt.Sprintf(
		"SELECT id, title, timezoneId, startDatetime, description, organizerId FROM %s WHERE organizerId=$1 AND id=$2",
		EVENTS_TABLE,
	)
	err := r.db.Get(&result, query, userId, eventId)

	return result, err
}

func (r *EventsPostgres) Create(event domain.Event) (int, error) {
	var result int

	query := fmt.Sprintf(
		`INSERT INTO %s (title, timezoneId, startDatetime, description, organizerId) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
		EVENTS_TABLE,
	)
	row := r.db.QueryRow(query, event.Title, event.TimezoneId, event.StartDatetime, event.Description, event.OrganizerId)
	if err := row.Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}

func (r *EventsPostgres) Update(userId int, event domain.Event) (domain.Event, error) {
	var result domain.Event
	fmt.Println(userId)

	query := fmt.Sprintf(
		`UPDATE %s SET title='%s', timezoneid='%s', startdatetime='%s', description='%s' 
		 WHERE id=$1 AND organizerid=$2
		 RETURNING id, title, description, organizerid, startdatetime, timezoneid`,
		EVENTS_TABLE, event.Title, event.TimezoneId, event.StartDatetime, event.Description,
	)
	err := r.db.Get(
		&result,
		query,
		event.Id, userId,
	)

	return result, err
}

func (r *EventsPostgres) Delete(userId, eventId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE organizerId=$1 AND id=$2", EVENTS_TABLE)

	r.db.Exec(query, eventId)
	_, err := r.db.Exec(query, userId, eventId)

	return err
}
