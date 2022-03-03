package main

const ROUND_SELECT_QUERY string = `
	select 
		id, title, description, is_published, user_id, date_created, date_updated
	from 
		questions
	where 
		id = $1
`
