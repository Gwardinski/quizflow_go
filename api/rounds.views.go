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
	ridParam, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Get Round
	round, err := app.db.GetRound(ridParam)
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

func (app *application) getUsersRounds(w http.ResponseWriter, r *http.Request) {
	// Get UserID from JWT (to get users rounds)
	uidToken, err := app.getUserIdFromJWT(r)
	if err != nil {
		app.writeError(w, err, http.StatusUnauthorized)
		return
	}

	// Get Rounds
	rounds, err := app.db.GetUsersRounds(uidToken)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, rounds)
}

func (app *application) getRoundsByUser(w http.ResponseWriter, r *http.Request) {
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

	// Get Rounds
	rounds, err := app.db.GetRoundsByUser(uidParams)
	if err != nil {
		app.writeError(w, err, http.StatusNotFound)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, rounds)
}

func (app *application) getRoundsQuizzes(w http.ResponseWriter, r *http.Request) {

}

func (app *application) createRound(w http.ResponseWriter, r *http.Request) {
	// Get UserID from JWT (to set as UserID on new Round)
	uidToken, err := app.getUserIdFromJWT(r)
	if err != nil {
		app.writeError(w, err, http.StatusUnauthorized)
		return
	}

	// Decode request body into RoundPayload
	var payload RoundPayload
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write new Round to DB, return it's ID
	rid, err := app.db.CreateRound(payload, uidToken)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Write Data
	app.writeResponse(w, http.StatusOK, rid)
}
