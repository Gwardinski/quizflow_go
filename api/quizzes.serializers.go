package main

import (
	"database/sql"
)

func (q *Quiz) serialise(row *sql.Row) (err error) {
	err = row.Scan(
		&q.ID,
		&q.Title,
		&q.Description,
		&q.IsPublished,
		&q.DateCreated,
		&q.DateUpdated,
		&q.User.ID,
		&q.User.Username,
	)
	if err != nil {
		return err
	}
	return nil
}

func serialiseQuizzes(rows *sql.Rows) (quizzes []QuizItem, err error) {
	for rows.Next() {
		var q QuizItem
		// TODO: array is returned as []byte 🤔
		// Needs converted into string, then into []string, then into []int 🙄
		var ridString string
		err = rows.Scan(
			&q.ID,
			&q.Title,
			&q.Description,
			&q.IsPublished,
			&q.DateCreated,
			&q.DateUpdated,
			&q.User.ID,
			&ridString,
		)
		if err != nil {
			return []QuizItem{}, err
		}
		q.Rounds = _appendRoundOnListItems(ridString)
		quizzes = append(quizzes, q)
	}
	return quizzes, nil
}

func _appendRoundOnListItems(ridString string) (rounds []RoundSubItem) {
	rids := byteStringToIntSlice(ridString)
	for _, id := range rids {
		r := RoundSubItem{
			ID: id,
		}
		rounds = append(rounds, r)
	}
	return rounds
}
