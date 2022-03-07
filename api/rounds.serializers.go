package main

import "database/sql"

func (q *RoundDetailsRes) serialise(row *sql.Row) (err error) {
	err = row.Scan(
		&q.ID,
		&q.Title,
		&q.Description,
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
