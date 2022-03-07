package main

import "database/sql"

func (u *UserOnItem) serialise(row *sql.Row) (err error) {
	row.Scan(
		&u.ID,
		&u.Username,
	)
	return nil
}
