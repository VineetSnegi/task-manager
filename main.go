package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thedevsaddam/renderer"

	"task-manager/handlers"
)

const (
	port        = ":9000"
	postgresURL = "postgres://postgres:Admin@123@localhost:5432/taskdb"
)

var rnd *renderer.Render
var db *pgxpool.Pool

func init() {
	rnd = renderer.New()
	var err error
	db, err = pgxpool.New(context.Background(), postgresURL)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := rnd.Template(w, http.StatusOK, []string{"static/home.tpl"}, nil)
	if err != nil {
		log.Println("Template render error:", err)
	}
}

func main() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	handlers.Init(rnd, db)

	r.Get("/", homeHandler)
	r.Mount("/task", handlers.TaskRouter())

	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("Server started at", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Server gracefully stopped.")
}
