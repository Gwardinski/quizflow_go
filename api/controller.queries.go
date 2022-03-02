package main

func (*QuestionDetailsRes) selectQuery() string {
	return `
		select 
			id, title, answer, points, category, is_published, user_id, date_created, date_updated
		from 
			questions
		where 
			id = $1
	`
}
func (*QuestionDetailsRes) insertQuery() string {
	return `
		update questions set 
			title = $1,
			answer = $2,
			points = $3,
			category = $4,
			date_updated = $5
		where 
			id = $6
	`
}

func (*QuestionPayload) insertQuery() string {
	return `
		insert into questions 
			(title, answer, points, category, date_created, date_updated, user_id) 
		values 
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING 
			id
	`
}

func (*TagResponse) selectQuery() string {
	return `
		select
			t.id, t.title
		from
			tags t
		where
			t.title = $1
	`
}

func (*TagsResponse) selectQuery() string {
	return `
		select
			qt.id, qt.question_id, qt.tag_id, t.title
		from
			questions_tags qt
			left join tags t on (t.id = qt.tag_id)
		where
			qt.question_id = $1
	`
}

func (*UserOnItem) selectQuery() string {
	return `
		select
			u.id, u.username
		from
			users u
		where
			u.id = $1
	`
}

func GET_QUESTION_LIST_QUERY(where, order string) string {
	query := `
		select 
			id, title, answer, points, category, is_published, user_id
		from questions
	`
	query += ("where " + where + " ")
	query += ("order by " + order)
	return query
}
