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
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
	// Image
}

// Simplified User object when returning Question / Round / Quiz
type UserOnItem struct {
	ID int `json:"id"`
	// following will be null on lists, only retrieved when querying item by ID
	Username string `json:"username"`
	// Image
}
