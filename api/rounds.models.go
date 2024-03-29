package main

import (
	"time"
)

// TODO: Add 'Image' field to all structs

// Model Representation of Database Structure
// UserID is a field on the object
// Only the response models nest UserID under a 'User' field (see Round)
type RoundDB struct {
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
type RoundPayload struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	QIDs        []string `json:"qids"`
}

/*
	= = = = = = = = = = = = = = = = = =
	RESPONSE BODYS
	Used for returning json on Get Methods
	UserID is now moved into seperate 'User' field
	= = = = = = = = = = = = = = = = = =
*/

// Round is returned when requesting single item
// GET - /rounds/{ID}
type Round struct {
	ID            int            `json:"id"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	TotalPoints   int            `json:"total_points"`
	IsPublished   bool           `json:"isPublished"`
	DateCreated   time.Time      `json:"dateCreated"`
	DateUpdated   time.Time      `json:"dateUpdated"`
	DatePublished time.Time      `json:"datePublished"`
	User          UserItem       `json:"user"`
	Questions     []QuestionItem `json:"questions"`
}

// RoundItem is returned when requesting multiple items
// GET - /rounds/
// GET - /rounds/user
// Or when requesting a single parent item
// GET - /quizzes/{ID}
type RoundItem struct {
	ID            int               `json:"id"`
	Title         string            `json:"title"`
	Description   string            `json:"description"`
	TotalPoints   int               `json:"total_points"`
	IsPublished   bool              `json:"isPublished"`
	DateCreated   time.Time         `json:"dateCreated"`
	DateUpdated   time.Time         `json:"dateUpdated"`
	DatePublished time.Time         `json:"datePublished"`
	User          UserSubItem       `json:"user"`
	Questions     []QuestionSubItem `json:"questions"`
}

// RoundSubItem is returned when requesting multiple parent items
// GET - /quizzes/
// GET - /quizzes/user
type RoundSubItem struct {
	ID int `json:"id"`
}
