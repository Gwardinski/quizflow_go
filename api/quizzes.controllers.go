package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func (m *DBModel) GetQuiz(id int) (Quiz, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET SINGLE QUIZ
	var quiz Quiz
	row := m.DB.QueryRowContext(ctx, QUIZ_SELECT_QUERY, id)
	err := quiz.serialise(row)
	if err != nil {
		fmt.Println(1)
		return quiz, err
	}

	// GET ROUNDS FOR QUIZ
	rows, err := m.DB.QueryContext(ctx, ROUNDS_FROM_QUIZZES_ROUNDS_QUERY, id)
	if err != nil {
		fmt.Println(2)
		return quiz, err
	}
	questions, err := serialiseRounds(rows)
	if err != nil {
		fmt.Println(3)
		return quiz, err
	}
	quiz.Rounds = questions
	return quiz, nil
}

func (m DBModel) GetPublishedQuizzes() ([]QuizItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET QUIZS
	where := "is_published = true"
	order := "title"
	query := QUIZ_LIST_QUERY(where, order)
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return serialiseQuizzes(rows)
}

func (m DBModel) GetUsersQuizzes(userId int) ([]QuizItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET QUIZS
	where := "user_id = $1"
	order := "title"
	query := QUIZ_LIST_QUERY(where, order)
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return serialiseQuizzes(rows)
}

func (m DBModel) GetQuizzesByUser(userId int) ([]QuizItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET QUIZS
	where := "user_id = $1 and is_published = true"
	order := "title"
	query := QUIZ_LIST_QUERY(where, order)
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return serialiseQuizzes(rows)
}

func (m *DBModel) CreateQuiz(quiz QuizPayload, uid int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Create Quiz
	// quizs are created by providing only a title.
	// Further details provided later in the editing process
	row := m.DB.QueryRowContext(ctx, QUIZ_INSERT_QUERY,
		quiz.Title,
		time.Now(),
		time.Now(),
		uid,
	)
	// Get ID from new Quiz
	var quizID int
	fmt.Println(row.Scan(
		&quizID,
	))

	return quizID, nil
}

func (m *DBModel) UpdateQuiz(quizID int, payload QuizPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Update Quiz with values from Payload
	_, err := m.DB.ExecContext(ctx, QUIZ_UPDATE_QUERY,
		payload.Title,
		payload.Description,
		time.Now(),
		quizID,
	)

	updateQuizRounds(m, ctx, payload.RIDs, quizID)

	if err != nil {
		return err
	}
	return nil

}

func (m *DBModel) DeleteQuiz(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, QUIZ_DELETE_QUERY, id)
	if err != nil {
		return err
	}
	return nil
}

func updateQuizRounds(m *DBModel, ctx context.Context, qids []string, quizID int) {
	m.DB.ExecContext(ctx, QUIZ_ROUND_DELETE_ALL_QUERY, quizID)
	// TODO: error handling
	for _, id := range qids {
		qid, _ := strconv.Atoi(id)
		// TODO: error handling?
		createQuizRound(m, ctx, quizID, qid)
	}
}

func createQuizRound(m *DBModel, ctx context.Context, quizID int, roundID int) {
	_, err := m.DB.ExecContext(
		ctx,
		QUIZ_ROUND_INSERT_QUERY,
		quizID,
		roundID,
	)
	if err != nil {
		// TODO: not sure what to do with this error yet
		fmt.Println(err)
	}
}
