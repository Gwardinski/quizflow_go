package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func (m *DBModel) GetQuestion(id int) (QuestionDetailsRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET SINGLE QUESTION
	var question QuestionDetailsRes
	row := m.DB.QueryRowContext(ctx, question.selectQuery(), id)
	err := question.serialise(row)
	if err != nil {
		return question, err
	}

	// GET SINGLE USER
	var user UserOnItem
	row = m.DB.QueryRowContext(ctx, user.selectQuery(), question.User.ID)
	err = user.serialise(row)
	if err != nil {
		return question, err
	}

	// GET MULTIPLE TAGS
	var tags TagsResponse
	rows, err := m.DB.QueryContext(ctx, tags.selectQuery(), id)
	if err != nil {
		return question, err
	}
	defer rows.Close()
	tags, err = tags.serialise(rows)
	if err != nil {
		return question, err
	}

	// SET TAGS / USER
	question.Tags = tags
	question.User = user

	return question, nil
}

func (m DBModel) GetPublishedQuestions() ([]QuestionListItemRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET QUESTIONS
	where := "is_published = true"
	order := "title"
	query := GET_QUESTION_LIST_QUERY(where, order)
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return serialiseQuestionsList(rows)
}

func (m DBModel) GetUsersQuestions(userId int) ([]QuestionListItemRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET QUESTIONS
	where := "user_id = $1"
	order := "title"
	query := GET_QUESTION_LIST_QUERY(where, order)
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return serialiseQuestionsList(rows)
}

func (m DBModel) GetQuestionsByUser(userId int) ([]QuestionListItemRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET QUESTIONS
	where := "user_id = $1 and is_published = true"
	order := "title"
	query := GET_QUESTION_LIST_QUERY(where, order)
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return serialiseQuestionsList(rows)
}

func (m *DBModel) CreateQuestion(question QuestionPayload, uid int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Convert int values from string
	points, err := strconv.Atoi(question.Points)
	if err != nil {
		return 0, err
	}
	// Create Question
	row := m.DB.QueryRowContext(ctx, question.insertQuery(),
		question.Title,
		question.Answer,
		points,
		question.Category,
		time.Now(),
		time.Now(),
		uid,
	)
	// Get ID from new Question
	var newId int
	fmt.Println(row.Scan(
		&newId,
	))

	// Save Tags
	createOrUpdateTags(m, ctx, question.Tags)

	return newId, nil
}

func (m *DBModel) UpdateQuestion(question QuestionDetailsRes, payload QuestionPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Update Question with values from Payload
	question.updateFromPayload(payload)
	_, err := m.DB.ExecContext(ctx, question.insertQuery(),
		question.Title,
		question.Answer,
		question.Points,
		question.Category,
		question.DateUpdated,
		question.ID,
	)

	//TODO: Create / Update Tags

	if err != nil {
		return err
	}
	return nil

}

func (m *DBModel) DeleteQuestion(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "delete from questions where id = $1"
	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func createOrUpdateTags(m *DBModel, ctx context.Context, questionTags []string) {
	for _, t := range questionTags {
		tr := TagResponse(t)
		row := m.DB.QueryRowContext(ctx, tr.selectQuery(), t)
		if row != nil {
			// tag already exists
			// check QuestionsTags
			// if nil == true, create new entry in QuestionsTags
		} else {
			// create new entry in Tags
			// create new entry in QuestionsTags
		}
	}
}
