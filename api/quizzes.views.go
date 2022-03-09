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

func (app *application) getPublishedQuizzes(w http.ResponseWriter, r *http.Request) {
	// Get Quizzes
	quizzes, err := app.db.GetPublishedQuizzes()
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, quizzes)
}

/*
	= = = = = = = = = = = = = = = = = =
	PRIVATE VIEWS
	Requires Bearer Token
	= = = = = = = = = = = = = = = = = =
*/

func (app *application) getQuiz(w http.ResponseWriter, r *http.Request) {
	// Get UserID from JWT (to confirm ownership)
	uidToken, err := app.getUserIdFromJWT(r)
	if err != nil {
		app.writeError(w, err, http.StatusUnauthorized)
		return
	}

	// Get Quiz ID from params
	vars := mux.Vars(r)
	qzidParam, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Get Quiz
	quiz, err := app.db.GetQuiz(qzidParam)
	if err != nil || quiz.ID == 0 {
		app.writeError(w, err, http.StatusNotFound)
		return
	}

	// AUTH CHECK: Check quiz belongs to user (if not published)
	if !quiz.IsPublished && (quiz.User.ID != uidToken) {
		app.writeError(
			w,
			errors.New("not published / No Auth"),
			http.StatusUnauthorized,
		)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, quiz)
}

func (app *application) getUsersQuizzes(w http.ResponseWriter, r *http.Request) {
	// Get UserID from JWT (to get users quizzes)
	uidToken, err := app.getUserIdFromJWT(r)
	if err != nil {
		app.writeError(w, err, http.StatusUnauthorized)
		return
	}

	// Get Quizzes
	quizzes, err := app.db.GetUsersQuizzes(uidToken)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, quizzes)
}

func (app *application) getQuizzesByUser(w http.ResponseWriter, r *http.Request) {
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

	// Get Quizzes
	quizzes, err := app.db.GetQuizzesByUser(uidParams)
	if err != nil {
		app.writeError(w, err, http.StatusNotFound)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, quizzes)
}

// func (app *application) getQuizzesQuizzes(w http.ResponseWriter, r *http.Request) {

// }

func (app *application) createQuiz(w http.ResponseWriter, r *http.Request) {
	// Get UserID from JWT (to set as UserID on new Quiz)
	uidToken, err := app.getUserIdFromJWT(r)
	if err != nil {
		app.writeError(w, err, http.StatusUnauthorized)
		return
	}

	// Decode request body into QuizPayload
	var payload QuizPayload
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write new Quiz to DB, return it's ID
	rid, err := app.db.CreateQuiz(payload, uidToken)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, rid)
}
