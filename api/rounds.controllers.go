package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func (m *DBModel) GetRound(id int) (Round, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET SINGLE ROUND
	var round Round
	row := m.DB.QueryRowContext(ctx, ROUND_SELECT_QUERY, id)
	err := round.serialise(row)
	if err != nil {
		return round, err
	}

	// GET QUESTIONS FOR ROUND
	rows, err := m.DB.QueryContext(ctx, QUESTIONS_FROM_ROUNDS_QUESTIONS_QUERY, id)
	if err != nil {
		return round, err
	}
	questions, err := serialiseQuestions(rows)
	if err != nil {
		return round, err
	}
	round.Questions = questions
	return round, nil
}

func (m DBModel) GetPublishedRounds() ([]RoundItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET ROUNDS
	where := "is_published = true"
	order := "title"
	query := ROUND_LIST_QUERY(where, order)
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return serialiseRounds(rows)
}

func (m DBModel) GetUsersRounds(userId int) ([]RoundItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET ROUNDS
	where := "user_id = $1"
	order := "title"
	query := ROUND_LIST_QUERY(where, order)
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return serialiseRounds(rows)
}

func (m DBModel) GetRoundsByUser(userId int) ([]RoundItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// GET ROUNDS
	where := "user_id = $1 and is_published = true"
	order := "title"
	query := ROUND_LIST_QUERY(where, order)
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return serialiseRounds(rows)
}

func (m *DBModel) CreateRound(round RoundPayload, uid int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Create Round
	// rounds are created by providing only a title.
	// Further details provided later in the editing process
	row := m.DB.QueryRowContext(ctx, ROUND_INSERT_QUERY,
		round.Title,
		time.Now(),
		time.Now(),
		uid,
	)
	// Get ID from new Round
	var roundID int
	fmt.Println(row.Scan(
		&roundID,
	))

	return roundID, nil
}

func (m *DBModel) UpdateRound(roundID int, payload RoundPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Update Round with values from Payload
	_, err := m.DB.ExecContext(ctx, ROUND_UPDATE_QUERY,
		payload.Title,
		payload.Description,
		time.Now(),
		roundID,
	)

	updateRoundQuestions(m, ctx, payload.QIDs, roundID)

	if err != nil {
		return err
	}
	return nil

}

func (m *DBModel) DeleteRound(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, ROUND_DELETE_QUERY, id)
	if err != nil {
		return err
	}
	return nil
}

func updateRoundQuestions(m *DBModel, ctx context.Context, qids []string, roundID int) {
	m.DB.ExecContext(ctx, ROUND_QUESTION_DELETE_ALL_QUERY, roundID)
	// TODO: error handling
	for _, id := range qids {
		qid, _ := strconv.Atoi(id)
		// TODO: error handling?
		createRoundQuestion(m, ctx, roundID, qid)
	}
}

func createRoundQuestion(m *DBModel, ctx context.Context, roundID int, questionID int) {
	_, err := m.DB.ExecContext(
		ctx,
		ROUND_QUESTION_INSERT_QUERY,
		roundID,
		questionID,
	)
	if err != nil {
		// TODO: not sure what to do with this error yet
		fmt.Println(err)
	}
}
