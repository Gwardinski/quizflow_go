package main

import (
	"time"
)

// Model Representation of Database Structure
// UserID is a field on the object
// Only the response models nest UserID under a 'User' field (see Quiz)
type QuizDB struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	IsPublished   bool      `json:"isPublished"`
	DateCreated   time.Time `json:"dateCreated"`
	DateUpdated   time.Time `json:"dateUpdated"`
	DatePublished time.Time `json:"datePublished"`
}

/*
	= = = = = = = = = = = = = = = = = =
	REQUEST BODY
	Used for easier decoding of recieved json on Create & Update functions.
	Each value type as a string. (no ints)
	= = = = = = = = = = = = = = = = = =
*/
type QuizPayload struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	RIDs        []string `json:"rids"`
}

/*
	= = = = = = = = = = = = = = = = = =
	RESPONSE BODYS
	Used for returning json on Get Methods
	UserID is now moved into seperate 'User' field
	= = = = = = = = = = = = = = = = = =
*/

// Quiz is returned when requesting single item
// GET - /quizzes/{ID}
type Quiz struct {
	ID            int         `json:"id"`
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	TotalPoints   int         `json:"total_points"`
	IsPublished   bool        `json:"isPublished"`
	DateCreated   time.Time   `json:"dateCreated"`
	DateUpdated   time.Time   `json:"dateUpdated"`
	DatePublished time.Time   `json:"datePublished"`
	User          UserItem    `json:"user"`
	Rounds        []RoundItem `json:"rounds"`
}

// QuizItem is returned when requesting multiple items
// GET - /quizzes/
// GET - /quizzes/user
type QuizItem struct {
	ID            int            `json:"id"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	TotalPoints   int            `json:"total_points"`
	IsPublished   bool           `json:"isPublished"`
	DateCreated   time.Time      `json:"dateCreated"`
	DateUpdated   time.Time      `json:"dateUpdated"`
	DatePublished time.Time      `json:"datePublished"`
	User          UserSubItem    `json:"user"`
	Rounds        []RoundSubItem `json:"rounds"`
}
