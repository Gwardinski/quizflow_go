package main

import "time"

// Model Representation of Database Structure
// Full object only returned on request to view profile information
// Simplified represenations used on item requests
type UserDB struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"-"`
	PasswordHash string    `json:"-"`
	Image        string    `json:"image"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"-"`
}

// UserItem is returned when requesting single parent item
// GET - /quizzes/{ID}
// GET - /rounds/{ID}
// GET - /questions/{ID}
type UserItem struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Image    string `json:"image"`
}

// UserSubItem is returned when requesting multiple parent items
// GET - /quizzes/
// GET - /rounds/
// GET - /questions/
type UserSubItem struct {
	ID int `json:"id"`
}
