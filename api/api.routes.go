package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	// Util routes
	statusRoute := router.PathPrefix("/status").Subrouter()
	statusRoute.HandleFunc("/", app.getAppStatus).Methods("GET")

	// Auth routes
	authRoutes := router.PathPrefix("/auth").Subrouter()
	authRoutes.HandleFunc("/login", app.login).Methods("Post")
	authRoutes.HandleFunc("/register", app.register).Methods("Post")

	// Question routes
	questionsRoute := router.PathPrefix("/questions").Subrouter()
	questionsRoute.HandleFunc("/", app.getPublishedQuestions).Methods("GET")
	questionsRoute.HandleFunc("/user", app.getUsersQuestions).Methods("GET")
	questionsRoute.HandleFunc("/", app.createQuestion).Methods("POST")
	questionsRoute.HandleFunc("/{id}", app.getQuestion).Methods("GET")
	questionsRoute.HandleFunc("/{id}", app.updateQuestion).Methods("PUT")
	questionsRoute.HandleFunc("/{id}", app.deleteQuestion).Methods("DELETE")
	questionsRoute.HandleFunc("/{id}/rounds", app.getQuestionsRounds).Methods("GET")
	questionsRoute.HandleFunc("/user/{id}", app.getQuestionsByUser).Methods("GET")

	// Round routes
	roundsRoute := router.PathPrefix("/rounds").Subrouter()
	roundsRoute.HandleFunc("/", app.getPublishedRounds).Methods("GET")
	// roundsRoute.HandleFunc("/user", app.getUsersRounds).Methods("GET")
	// roundsRoute.HandleFunc("/", app.createRound).Methods("POST")
	roundsRoute.HandleFunc("/{id}", app.getRound).Methods("GET")
	// roundsRoute.HandleFunc("/{id}", app.updateRound).Methods("PUT")
	// roundsRoute.HandleFunc("/{id}", app.deleteRound).Methods("DELETE")
	// roundsRoute.HandleFunc("/{id}/rounds", app.getRoundsRounds).Methods("GET")
	// roundsRoute.HandleFunc("/user/{id}", app.getRoundsByUser).Methods("GET")

	// TODO: Quiz routes

	// Handle CORS Middleware
	return app.enableCORS(router)
}
