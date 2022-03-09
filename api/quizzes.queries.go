package main

const QUIZ_SELECT_QUERY string = `
	select 
		q.id, q.title, q.description, q.is_published, q.date_created, q.date_updated,
		u.id, u.username
	from 
		quizzes q
		left join users u on (u.id = q.user_id)
	where 
		q.id = $1
`

func QUIZ_LIST_QUERY(where, order string) string {
	query := `
		select 
			q.id, q.title, q.description, q.is_published, q.date_created, q.date_updated, q.user_id,
			array(
				select round_id
				from quizzes_rounds qr 
				where qr.quiz_id = q.id
			) as qids
		from 
			quizzes q
	`
	query += ("where " + where + " ")
	query += ("order by " + order)
	return query
}

const QUIZ_INSERT_QUERY string = `
	insert into quizzes 
		(title, date_created, date_updated, user_id) 
	values 
		($1, $2, $3, $4)
	RETURNING 
		id
`

const QUIZ_UPDATE_QUERY string = `
	update rounds set 
		title = $1,
		description = $2,
		date_updated = $3
	where 
		id = $4
`

const QUIZ_DELETE_QUERY string = `
	delete from quizzes where id = $1
`

const QUIZ_ROUNDS_SELECT string = `
	select
		id
	from
		quizzes_rounds
	where
		quiz_id = $1
`

const QUIZ_ROUND_DELETE_ALL_QUERY string = `
	delete from quizzes_rounds where quiz_id = $1
`

const QUIZ_ROUND_INSERT_QUERY string = `
	insert into quizzes_rounds
		(quiz_id, round_id)
	values
		($1, $2)
	RETURNING 
		id
`

const ROUNDS_FROM_QUIZZES_ROUNDS_QUERY string = `
	select
		r.id, r.title, r.description, r.is_published
	from
		quizzes_rounds qr
		left join rounds r on (r.id = qr.round_id)
	where
		qr.quiz_id = $1
`
