package main

import (
	"context"
	"time"
)

func (m *DBModel) GetRound(id int) (RoundDetailsRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET SINGLE ROUND
	var round RoundDetailsRes
	row := m.DB.QueryRowContext(ctx, ROUND_SELECT_QUERY, id)
	err := round.serialise(row)
	if err != nil {
		return round, err
	}

	// GET SINGLE USER
	var user UserOnItem
	row = m.DB.QueryRowContext(ctx, USER_SELECT_QUERY, round.User.ID)
	err = user.serialise(row)
	if err != nil {
		return round, err
	}

	// SET USER
	round.User = user

	return round, nil
}

func (m DBModel) GetPublishedRounds() ([]RoundListItemRes, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	// // GET ROUNDS
	// where := "is_published = true"
	// order := "title"
	// query := ROUND_LIST_QUERY(where, order)
	// rows, err := m.DB.QueryContext(ctx, query)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	// return serialiseRoundsList(rows)
	return []RoundListItemRes{}, nil
}
