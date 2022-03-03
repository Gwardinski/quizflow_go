package main

import (
	"time"
)

// TODO: Add 'Image' field to all structs

// Model Representation of Database Structure
// UserID is a field on the object
// Only the response models nest UserID under a 'User' field (see QuestionDetailsRes)
type Question struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Title         string    `json:"title"`
	Answer        string    `json:"answer"`
	Points        int       `json:"points"`
	Category      string    `json:"category"`
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
type QuestionPayload struct {
	Title    string   `json:"title"`
	Answer   string   `json:"answer"`
	Points   string   `json:"points"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

/*
	= = = = = = = = = = = = = = = = = =
	RESPONSE BODYS
	Used for returning json on Get Methods
	UserID is now moved into seperate 'User' field
	= = = = = = = = = = = = = = = = = =
*/
// 'Full' response object for queries on a specific item
// /questions/{:id}
// Multiple Look ups made in controller to get Tag strings and User struct from DB
type QuestionDetailsRes struct {
	ID            int          `json:"id"`
	Title         string       `json:"title"`
	Answer        string       `json:"answer"`
	Points        int          `json:"points"`
	Category      string       `json:"category"`
	IsPublished   bool         `json:"isPublished"`
	DateCreated   time.Time    `json:"dateCreated"`
	DateUpdated   time.Time    `json:"dateUpdated"`
	DatePublished time.Time    `json:"datePublished"`
	Tags          TagsResponse `json:"tags"`
	User          UserOnItem   `json:"user"`
	Rounds        []string     `json:"rounds"`
}

// 'Lite' response object for queries on multiple items
// /questions/
// /questions/user
// /questions/user/{:id}
type QuestionListItemRes struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Answer      string     `json:"answer"`
	Points      int        `json:"points"`
	Category    string     `json:"category"`
	IsPublished bool       `json:"isPublished"`
	User        UserOnItem `json:"user"`
}

/*
	= = = = = = = = = = = = = = = = = =
	TAGS
	= = = = = = = = = = = = = = = = = =
*/
// Main Tag Model
type Tag struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	DateCreated time.Time `json:"-"`
}

// Many-to-Many Questions/Tags
type QuestionTag struct {
	ID          int       `json:"-"`
	QuestionID  int       `json:"-"`
	TagID       int       `json:"-"`
	Tag         Tag       `json:"tag"`
	DateCreated time.Time `json:"-"`
}

// Response Object
type TagsResponse []TagResponse
type TagResponse string
