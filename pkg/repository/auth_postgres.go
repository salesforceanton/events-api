package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/salesforceanton/events-api/domain"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domain.User) (int, error) {
	var result int

	query := fmt.Sprintf("INSERT INTO %s (email, username, password_hash) VALUES ($1, $2, $3) RETURNING id", USERS_TABLE)
	row := r.db.QueryRow(query, user.Email, user.Username, user.Password)

	if err := row.Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}

func (r *AuthPostgres) GetUser(username, password string) (domain.User, error) {
	var result domain.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", USERS_TABLE)
	err := r.db.Get(&result, query, username, password)

	return result, err
}
