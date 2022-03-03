package main

import (
	"database/sql"
)

func serialiseQuestionsList(rows *sql.Rows) (q []QuestionListItemRes, err error) {
	var questions []QuestionListItemRes
	for rows.Next() {
		var q QuestionListItemRes
		err := q.serialise(rows)
		if err != nil {
			return []QuestionListItemRes{}, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func (q *QuestionDetailsRes) serialise(row *sql.Row) (err error) {
	err = row.Scan(
		&q.ID,
		&q.Title,
		&q.Answer,
		&q.Points,
		&q.Category,
		&q.IsPublished,
		&q.User.ID,
		&q.DateCreated,
		&q.DateUpdated,
	)
	if err != nil {
		return err
	}
	return nil
}

func (q *QuestionListItemRes) serialise(rows *sql.Rows) (err error) {
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
		return err
	}
	return nil
}

func (*TagsResponse) serialise(rows *sql.Rows) (TagsResponse, error) {
	var tags []TagResponse
	for rows.Next() {
		var qt QuestionTag
		err := rows.Scan(
			&qt.ID,
			&qt.QuestionID,
			&qt.TagID,
			&qt.Tag.Title,
		)
		if err != nil {
			return tags, err
		}
		tags = append(tags, TagResponse(qt.Tag.Title))
	}
	return TagsResponse(tags), nil
}

func (t *Tag) serialise(row *sql.Row) error {
	err := row.Scan(
		&t.ID,
		&t.Title,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserOnItem) serialise(row *sql.Row) (err error) {
	row.Scan(
		&u.ID,
		&u.Username,
	)
	return nil
}
