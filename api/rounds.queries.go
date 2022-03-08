package main

const ROUND_SELECT_QUERY string = `
	select 
		r.id, r.title, r.description, r.is_published, r.date_created, r.date_updated,
		u.id, u.username
	from 
		rounds r
		left join users u on (u.id = r.user_id)
	where 
		r.id = $1
`

func ROUND_LIST_QUERY(where, order string) string {
	query := `
		select 
			r.id, r.title, r.description, r.is_published, r.date_created, r.date_updated, r.user_id,
			array(
				select question_id
				from rounds_questions rq 
				where rq.round_id = r.id
			) as qids
		from 
			rounds r
	`
	query += ("where " + where + " ")
	query += ("order by " + order)
	return query
}

const ROUND_INSERT_QUERY string = `
	insert into rounds 
		(title, date_created, date_updated, user_id) 
	values 
		($1, $2, $3, $4)
	RETURNING 
		id
`

const ROUND_UPDATE_QUERY string = `
	update questions set 
		title = $1,
		description = $2,
		date_updated = $3
	where 
		id = $4
`

const ROUND_DELETE_QUERY string = `
	delete from rounds where id = $1
`

const ROUND_QUESTIONS_SELECT string = `
	select
		id
	from
		rounds_questions
	where
		round_id = $1
`

const ROUND_QUESTION_DELETE_ALL_QUERY string = `
	delete from rounds_questions where round_id = $1
`

const ROUND_QUESTION_INSERT_QUERY string = `
	insert into rounds_questions
		(round_id, question_id)
	values
		($1, $2)
	RETURNING 
		id
`

const QUESTIONS_FROM_ROUNDS_QUESTIONS_QUERY string = `
	select
		q.id, q.title, q.answer, q.points, q.category, q.is_published, q.user_id
	from
		rounds_questions rq
		left join questions q on (q.id = rq.question_id)
	where
		rq.round_id = $1
`
