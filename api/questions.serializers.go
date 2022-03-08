package main

import (
	"database/sql"
)

func (q *Question) serialise(row *sql.Row) (err error) {
	err = row.Scan(
		&q.ID,
		&q.Title,
		&q.Answer,
		&q.Points,
		&q.Category,
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

func serialiseQuestions(rows *sql.Rows) (q []QuestionItem, err error) {
	var questions []QuestionItem
	for rows.Next() {
		var q QuestionItem
		err = rows.Scan(
			&q.ID,
			&q.Title,
			&q.Answer,
			&q.Points,
			&q.Category,
			&q.IsPublished,
			&q.User.ID,
		)
		if err != nil {
			return questions, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func (*TagsResponse) serialise(rows *sql.Rows) (TagsResponse, error) {
	var tags []TagResponse
	for rows.Next() {
		var t TagResponse
		err := rows.Scan(
			&t,
		)
		if err != nil {
			return tags, err
		}
		tags = append(tags, t)
	}
	return TagsResponse(tags), nil
}

func (t *Tag) serialise(row *sql.Row) error {
	err := row.Scan(
		&t.Title,
	)
	if err != nil {
		return err
	}
	return nil
}
