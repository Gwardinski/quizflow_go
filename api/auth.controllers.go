package main

import (
	"context"
	"time"
)

func (m *DBModel) RegisterUser(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, USER_REGISTER_QUERY,
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

func (m *DBModel) GetUserForLogin(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, USER_LOGIN_QUERY, email)

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
