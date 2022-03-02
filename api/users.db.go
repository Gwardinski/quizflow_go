package main

import (
	"context"
	"time"
)

func (m *DBModel) GetUserByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, email, password_hash from users where email = $1`

	row := m.DB.QueryRowContext(ctx, query, email)

	var user User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *DBModel) RegisterUser(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into users (username, email, password_hash, date_created, date_updated) 
				values ($1, $2, $3, $4, $5)`

	_, err := m.DB.ExecContext(ctx, stmt,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.DateCreated,
		user.DateUpdated,
	)

	if err != nil {
		return err
	}
	return nil
}
