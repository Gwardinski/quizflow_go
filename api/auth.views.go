package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type RegisterPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	// Manually passing in pre-hashed password for local testing
	// TODO: add salt & hash, or go for oauth only approach
	Password string `json:"password"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	var payload RegisterPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.writeError(w, errors.New("json decode err"), http.StatusBadRequest)
		return
	}

	// Create new User from RegisterPayload values
	// TODO: Storing plain text passwords in DB is highly secure and always encouraged ✅
	user := User{
		Username:     payload.Username,
		Email:        payload.Email,
		PasswordHash: payload.Password,
		DateCreated:  time.Now(),
		DateUpdated:  time.Now(),
	}

	// Write new User to DB
	err = app.db.RegisterUser(user)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	// Return new User
	app.writeResponse(w, http.StatusOK, user)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	// Decode LoginPayload from request body
	var payload LoginPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.writeError(w, errors.New("json decode err"), http.StatusBadRequest)
		return
	}

	// Get user from DB where user.email == LoginPayload.email
	user, e := app.db.GetUserForLogin(payload.Email)
	if e != nil {
		// User shouldn't know if email exists or not, just us
		fmt.Println("User not found")
		app.writeError(w, errors.New(err.Error()), http.StatusBadRequest) // non-specific error text
		return
	}

	// Compare stored passwordHash with provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(payload.Password))
	if err != nil {
		// User shouldn't know it was the password specifically that failed, just us
		fmt.Println("Password is incorrect")
		app.writeError(w, errors.New("invalid Username / Password combination"), http.StatusBadRequest) // non-specific error text
		return
	}

	// Create and return JWT
	jwt, e := createJWT(user.ID, app.config.jwt.secret)

	if e != nil {
		app.writeError(w, errors.New("token signing error"), http.StatusBadRequest)
		return
	}
	app.writeResponse(w, http.StatusOK, string(jwt))
}
