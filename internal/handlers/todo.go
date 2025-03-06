package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jathin-s-ML/todo/internal/storage"
)

// RegisterTodoRoutes sets up routes
func RegisterTodoRoutes(router *mux.Router, store *storage.TodoStorage) {
	router.HandleFunc("", AddTodoHandler(store)).Methods(http.MethodPost)
	router.HandleFunc("", GetTodosHandler(store)).Methods(http.MethodGet)
	router.HandleFunc("/{id}", MarkTodoCompletedHandler(store)).Methods(http.MethodPut)
	router.HandleFunc("/{id}", DeleteTodoHandler(store)).Methods(http.MethodDelete)
}

func AddTodoHandler(store *storage.TodoStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Title string `json:"title"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Title == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		todo := store.AddTodo(req.Title)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todo)
	}
}

func GetTodosHandler(store *storage.TodoStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := store.GetTodos()
		sort.Slice(todos, func(i, j int) bool { return todos[i].ID < todos[j].ID })

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	}
}

func MarkTodoCompletedHandler(store *storage.TodoStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil || store.MarkAsCompleted(id) != nil {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteTodoHandler(store *storage.TodoStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil || store.DeleteTodo(id) != nil {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
