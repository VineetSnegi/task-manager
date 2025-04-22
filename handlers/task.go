package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thedevsaddam/renderer"

	"task-manager/models"
)

var rnd *renderer.Render
var db *pgxpool.Pool

func Init(r *renderer.Render, database *pgxpool.Pool) {
	rnd = r
	db = database
}

func TaskRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", FetchTasks)
	r.Post("/", CreateTask)
	r.Put("/{id}", UpdateTask)
	r.Delete("/{id}", DeleteTask)
	return r
}

func FetchTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(context.Background(), `SELECT id, title, completed, created_at FROM tasks`)
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, renderer.M{"message": "Failed to fetch tasks", "error": err})
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)
		if err != nil {
			continue
		}
		tasks = append(tasks, t)
	}

	rnd.JSON(w, http.StatusOK, renderer.M{"data": tasks})
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rnd.JSON(w, http.StatusBadRequest, err)
		return
	}
	if t.Title == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{"message": "Title is required"})
		return
	}
	err := db.QueryRow(context.Background(),
		"INSERT INTO tasks (title) VALUES ($1) RETURNING id",
		t.Title).Scan(&t.ID)
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, renderer.M{"message": "Failed to create task", "error": err})
		return
	}
	rnd.JSON(w, http.StatusCreated, renderer.M{"message": "Task created", "task_id": t.ID})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{"message": "Invalid ID"})
		return
	}
	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rnd.JSON(w, http.StatusBadRequest, err)
		return
	}
	if t.Title == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{"message": "Title is required"})
		return
	}
	_, err = db.Exec(context.Background(),
		"UPDATE tasks SET title=$1, completed=$2 WHERE id=$3",
		t.Title, t.Completed, taskID)
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, renderer.M{"message": "Update failed", "error": err})
		return
	}
	rnd.JSON(w, http.StatusOK, renderer.M{"message": "Task updated"})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{"message": "Invalid ID"})
		return
	}
	_, err = db.Exec(context.Background(), "DELETE FROM tasks WHERE id=$1", taskID)
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, renderer.M{"message": "Delete failed", "error": err})
		return
	}
	rnd.JSON(w, http.StatusOK, renderer.M{"message": "Task deleted"})
}
