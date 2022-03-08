package main

const QUESTION_SELECT_QUERY string = `
	select 
		q.id, q.title, q.answer, q.points, q.category, q.is_published, q.date_created, q.date_updated,
		u.id, u.username
	from 
		questions q
		left join users u on (u.id = q.user_id)
	where 
		q.id = $1
`

func QUESTION_LIST_QUERY(where, order string) string {
	query := `
		select 
			q.id, q.title, q.answer, q.points, q.category, q.is_published, q.user_id
		from 
			questions q
	`
	query += ("where " + where + " ")
	query += ("order by " + order)
	return query
}

const QUESTION_INSERT_QUERY string = `
	insert into questions 
		(title, answer, points, category, date_created, date_updated, user_id) 
	values 
		($1, $2, $3, $4, $5, $6, $7)
	RETURNING 
		id
`

const QUESTION_UPDATE_QUERY string = `
	update questions set 
		title = $1,
		answer = $2,
		points = $3,
		category = $4,
		date_updated = $5
	where 
		id = $6
`

const QUESTION_DELETE_QUERY string = `
	delete from questions where id = $1
`

const TAG_SELECT_QUERY string = `
	select
		t.title
	from
		tags t
	where
		t.title = $1
`

const TAG_INSERT_QUERY string = `
	insert into tags 
		(title) 
	values 
		($1)
	RETURNING 
		id 
`

const QUESTION_TAG_SELECT_QUERY string = `
	select
		qt.id, qt.question_id, qt.tag_id
	from
		questions_tags qt
	where
		qt.question_id = $1 and qt.tag_id = $2
`

const QUESTION_TAG_INSERT_QUERY string = `
	insert into questions_tags
		(question_id, tag_id)
	values
		($1, $2)
	RETURNING 
		id
`

const QUESTION_TAG_DELETE_ALL_QUERY string = `
	delete from questions_tags where question_id = $1
`

const TAG_FROM_QUESTION_TAG_QUERY string = `
	select
		t.title
	from
		questions_tags qt
		left join tags t on (t.id = qt.tag_id)
	where
		qt.question_id = $1
`
