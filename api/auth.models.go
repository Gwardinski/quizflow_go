package main

import "time"

// Model Representation of Database Structure
// Full object only returned on request to view profile information
// Simplified represenations used on item requests
type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"-"`
	PasswordHash string    `json:"-"`
	Image        string    `json:"image"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"-"`
}

// Simplified User object when returning Question / Round / Quiz
type UserItem struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Image    string `json:"image"`
}

// Simplified User object when returning list of Questions / Rounds / Quizs
type UserSubItem struct {
	ID int `json:"id"`
}
