package main

import (
	"context"
	"fmt"
	"time"
)

// Single Question, also queries User and Tag tables
func (m *DBModel) GetQuestion(id int) (Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET SINGLE QUESTION
	var question Question
	row := m.DB.QueryRowContext(ctx, QUESTION_SELECT_QUERY, id)
	err := question.serialise(row)
	if err != nil {
		return question, err
	}

	// GET MULTIPLE TAGS
	var tags TagsResponse
	rows, err := m.DB.QueryContext(ctx, TAG_FROM_QUESTION_TAG_QUERY, id)
	if err != nil {
		return question, err
	}
	defer rows.Close()
	tags, err = tags.serialise(rows)
	if err != nil {
		return question, err
	}

	// SET TAGS
	question.Tags = tags

	return question, nil
}

// List, does not query User or Tag tables
func (m DBModel) GetPublishedQuestions() ([]QuestionItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET QUESTIONS
	where := "is_published = true"
	order := "title"
	query := QUESTION_LIST_QUERY(where, order)
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return serialiseQuestions(rows)
}

func (m DBModel) GetUsersQuestions(userId int) ([]QuestionItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET QUESTIONS
	where := "user_id = $1"
	order := "title"
	query := QUESTION_LIST_QUERY(where, order)
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return serialiseQuestions(rows)
}

func (m DBModel) GetQuestionsByUser(userId int) ([]QuestionItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET QUESTIONS
	where := "user_id = $1 and is_published = true"
	order := "title"
	query := QUESTION_LIST_QUERY(where, order)
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return serialiseQuestions(rows)
}

func (m *DBModel) CreateQuestion(question QuestionPayload, uid int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Create Question
	row := m.DB.QueryRowContext(ctx, QUESTION_INSERT_QUERY,
		question.Title,
		question.Answer,
		question.Points,
		question.Category,
		time.Now(),
		time.Now(),
		uid,
	)
	// Get ID from new Question
	var questionID int
	fmt.Println(row.Scan(
		&questionID,
	))

	updateQuestionTags(m, ctx, question.Tags, questionID)

	return questionID, nil
}

func (m *DBModel) UpdateQuestion(questionID int, payload QuestionPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Update Question with values from Payload
	_, err := m.DB.ExecContext(ctx, QUESTION_UPDATE_QUERY,
		payload.Title,
		payload.Answer,
		payload.Points,
		payload.Category,
		time.Now(),
		questionID,
	)

	updateQuestionTags(m, ctx, payload.Tags, questionID)

	if err != nil {
		return err
	}
	return nil

}

func (m *DBModel) DeleteQuestion(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, QUESTION_DELETE_QUERY, id)
	if err != nil {
		return err
	}
	return nil
}

func updateQuestionTags(m *DBModel, ctx context.Context, tags []string, questionID int) {
	m.DB.ExecContext(ctx, QUESTION_TAG_DELETE_ALL_QUERY, questionID)
	// TODO: error handling
	for _, t := range tags {
		// Find if Tag doesn't exist
		row := m.DB.QueryRowContext(ctx, TAG_SELECT_QUERY, t)
		var tag Tag
		tag.serialise(row)
		if tag.Title == "" {
			// Tag is new, create new Tag
			r := m.DB.QueryRowContext(
				ctx,
				TAG_INSERT_QUERY,
				t,
			)
			r.Scan(
				&tag.ID,
			)
		}
		createQuestionTag(m, ctx, questionID, tag.ID)
	}
}

func createQuestionTag(m *DBModel, ctx context.Context, questionID int, tagID int) {
	_, err := m.DB.ExecContext(
		ctx,
		QUESTION_TAG_INSERT_QUERY,
		questionID,
		tagID,
	)
	if err != nil {
		// TODO: not sure what to do with this error yet
		fmt.Println(err)
	}
}
