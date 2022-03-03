package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
	= = = = = = = = = = = = = = = = = =
	PUBLIC VIEWS
	Do not require Bearer token
	= = = = = = = = = = = = = = = = = =
*/

func (app *application) getPublishedQuestions(w http.ResponseWriter, r *http.Request) {
	// Get Questions
	questions, err := app.db.GetPublishedQuestions()
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, questions)
}

/*
	= = = = = = = = = = = = = = = = = =
	PRIVATE VIEWS
	Requires Bearer Token
	= = = = = = = = = = = = = = = = = =
*/

func (app *application) getQuestion(w http.ResponseWriter, r *http.Request) {
	// Get UserID from JWT (to confirm ownership)
	uidToken, err := app.getUserIdFromJWT(r)
	if err != nil {
		app.writeError(w, err, http.StatusUnauthorized)
		return
	}

	// Get Question ID from params
	vars := mux.Vars(r)
	qidParam, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Get Question
	question, err := app.db.GetQuestion(qidParam)
	if err != nil || question.ID == 0 {
		app.writeError(w, err, http.StatusNotFound)
		return
	}

	// AUTH CHECK: Check question belongs to user (if not published)
	if !question.IsPublished && (question.User.ID != uidToken) {
		app.writeError(
			w,
			errors.New("not published / No Auth"),
			http.StatusUnauthorized,
		)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, question)
}

func (app *application) getUsersQuestions(w http.ResponseWriter, r *http.Request) {
	// Get UserID from JWT (to get users questions)
	uidToken, err := app.getUserIdFromJWT(r)
	if err != nil {
		app.writeError(w, err, http.StatusUnauthorized)
		return
	}

	// Get Questions
	questions, err := app.db.GetUsersQuestions(uidToken)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, questions)
}

func (app *application) getQuestionsByUser(w http.ResponseWriter, r *http.Request) {
	// Get UserID from JWT (validation check)
	_, err := app.getUserIdFromJWT(r)
	if err != nil {
		app.writeError(w, err, http.StatusUnauthorized)
		return
	}

	// Get UserID from params
	vars := mux.Vars(r)
	uidParams, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Get Questions
	questions, err := app.db.GetQuestionsByUser(uidParams)
	if err != nil {
		app.writeError(w, err, http.StatusNotFound)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, questions)
}

func (app *application) getQuestionsRounds(w http.ResponseWriter, r *http.Request) {

}

func (app *application) createQuestion(w http.ResponseWriter, r *http.Request) {
	// Get UserID from JWT (to set as UserID on new Question)
	uidToken, err := app.getUserIdFromJWT(r)
	if err != nil {
		app.writeError(w, err, http.StatusUnauthorized)
		return
	}

	// Decode request body into QuestionPayload
	var payload QuestionPayload
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write new Question to DB, return it's ID
	qid, err := app.db.CreateQuestion(payload, uidToken)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, qid)
}

func (app *application) updateQuestion(w http.ResponseWriter, r *http.Request) {
	// Get UserID from JWT (to confirm Question ownership)
	uidToken, err := app.getUserIdFromJWT(r)
	if err != nil {
		app.writeError(w, err, http.StatusUnauthorized)
		return
	}

	// Get Question ID from params
	vars := mux.Vars(r)
	qidParam, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Retrieve existing Question from DB using qid
	question, err := app.db.GetQuestion(qidParam)
	if err != nil || question.ID == 0 {
		app.writeError(w, err, http.StatusNotFound)
		return
	}

	// AUTH CHECK: Check question belongs to user
	if question.User.ID != uidToken {
		app.writeError(w, errors.New("you are not questions owner"), http.StatusUnauthorized)
		return
	}

	// Error if question is published
	if question.IsPublished {
		app.writeError(
			w,
			errors.New("this question is published and can not be edited"),
			http.StatusUnauthorized,
		)
		return
	}

	// Decode request body into QuestionPayload
	var payload QuestionPayload
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write updated Question to DB
	err = app.db.UpdateQuestion(question.ID, payload)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, nil)
}

func (app *application) deleteQuestion(w http.ResponseWriter, r *http.Request) {
	// Get UserID from JWT (to confirm Question ownership)
	uidToken, err := app.getUserIdFromJWT(r)
	if err != nil {
		app.writeError(w, err, http.StatusUnauthorized)
		return
	}

	// Get Question ID from params
	vars := mux.Vars(r)
	qidParam, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Get Question
	question, err := app.db.GetQuestion(qidParam)
	if err != nil || question.ID == 0 {
		app.writeError(w, err, http.StatusNotFound)
		return
	}

	// AUTH CHECK: Check question belongs to user
	if question.User.ID != uidToken {
		app.writeError(
			w,
			errors.New("you are not questions owner"),
			http.StatusUnauthorized,
		)
		return
	}

	err = app.db.DeleteQuestion(qidParam)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, nil)
}
