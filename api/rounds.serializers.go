package main

import (
	"database/sql"
)

func (r *Round) serialise(row *sql.Row) (err error) {
	err = row.Scan(
		&r.ID,
		&r.Title,
		&r.Description,
		&r.IsPublished,
		&r.DateCreated,
		&r.DateUpdated,
		&r.User.ID,
		&r.User.Username,
	)
	if err != nil {
		return err
	}
	return nil
}

func serialiseRounds(rows *sql.Rows) (rounds []RoundItem, err error) {
	for rows.Next() {
		var r RoundItem
		// TODO: array is returned as []byte 🤔
		// Needs converted into string, then into []string, then into []int 🙄
		var qidString string
		err = rows.Scan(
			&r.ID,
			&r.Title,
			&r.Description,
			&r.IsPublished,
			&r.DateCreated,
			&r.DateUpdated,
			&r.User.ID,
			&qidString,
		)
		if err != nil {
			return []RoundItem{}, err
		}
		r.Questions = _appendQuestionOnListItems(qidString)
		rounds = append(rounds, r)
	}
	return rounds, nil
}

func _appendQuestionOnListItems(qidString string) (questions []QuestionSubItem) {
	qids := byteStringToIntSlice(qidString)
	for _, id := range qids {
		q := QuestionSubItem{
			ID: id,
		}
		questions = append(questions, q)
	}
	return questions
}
