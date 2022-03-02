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
	authRoute := router.PathPrefix("/auth").Subrouter()
	authRoute.HandleFunc("/login", app.login).Methods("Post")
	authRoute.HandleFunc("/register", app.register).Methods("Post")

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

	// TODO: Round routes

	// TODO: Quiz routes

	// Handle CORS Middleware
	return app.enableCORS(router)
}
