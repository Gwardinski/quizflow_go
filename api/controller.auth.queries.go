package main

const USER_REGISTER_QUERY string = `
	insert into users 
		(username, email, password_hash, date_created, date_updated) 
	values 
		($1, $2, $3, $4, $5)
`

const USER_LOGIN_QUERY string = `
	select 
		id, email, password_hash 
	from 
		users 
	where 
		email = $1
`

const USER_SELECT_QUERY string = `
	select
		u.id, u.username
	from
		users u
	where
		u.id = $1
`
