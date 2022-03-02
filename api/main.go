package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// API Version, to be updated on each release
const version = "1.0.0"

// AppStatus struct used for collating API & development information
type AppStatus struct {
	Status      string `json:"status"`
	Enviornment string `json:"environment"`
	Version     string `json:"version"`
}

// Main control object
type application struct {
	config Config
	logger *log.Logger
	db     DBModel
}

// config, set using flags
type Config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	jwt struct {
		secret string
	}
}

// Wrapper to call controller functions in views
type DBModel struct {
	DB *sql.DB
}

// NewDB returns db pool
func NewDB(db *sql.DB) DBModel {
	return DBModel{
		DB: db,
	}
}

func main() {
	// Set config
	var cfg Config
	flag.IntVar(&cfg.port, "port", 4000, "Server port")
	flag.StringVar(&cfg.env, "env", "development", "Server env")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://gordon@localhost/go_quizflow?sslmode=disable", "Postgres connection string")
	flag.StringVar(&cfg.jwt.secret, "jwt-secret", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160", "secret")
	flag.Parse()

	// Create custom logging for Application
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Open DB Connection
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	// Set application with config and custom logging
	app := &application{
		config: cfg,
		logger: logger,
		db:     NewDB(db),
	}

	// Handle API routes
	app.routes()

	// Init http server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Println("Starting server on port", cfg.port)

	// Serving...
	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func openDB(cfg Config) (*sql.DB, error) {
	// TODO: What exactly is context? 🤔
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
