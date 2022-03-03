package main

import (
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

func (app *application) getPublishedRounds(w http.ResponseWriter, r *http.Request) {
	// Get Rounds
	rounds, err := app.db.GetPublishedRounds()
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, rounds)
}

/*
	= = = = = = = = = = = = = = = = = =
	PRIVATE VIEWS
	Requires Bearer Token
	= = = = = = = = = = = = = = = = = =
*/

func (app *application) getRound(w http.ResponseWriter, r *http.Request) {
	// Get UserID from JWT (to confirm ownership)
	uidToken, err := app.getUserIdFromJWT(r)
	if err != nil {
		app.writeError(w, err, http.StatusUnauthorized)
		return
	}

	// Get Round ID from params
	vars := mux.Vars(r)
	qidParam, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Get Round
	round, err := app.db.GetRound(qidParam)
	if err != nil || round.ID == 0 {
		app.writeError(w, err, http.StatusNotFound)
		return
	}

	// AUTH CHECK: Check round belongs to user (if not published)
	if !round.IsPublished && (round.User.ID != uidToken) {
		app.writeError(
			w,
			errors.New("not published / No Auth"),
			http.StatusUnauthorized,
		)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, round)
}
