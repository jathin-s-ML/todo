package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jathin-s-ML/todo/internal/storage"
)

// RegisterTodoRoutes sets up the todo routes
func RegisterTodoRoutes(router *mux.Router, store *storage.TodoStorage) {
	router.HandleFunc("/todos", AddTodoHandler(store)).Methods(http.MethodPost)
	router.HandleFunc("/todos", GetTodosHandler(store)).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", MarkTodoCompletedHandler(store)).Methods(http.MethodPut)
	router.HandleFunc("/todos/{id}", DeleteTodoHandler(store)).Methods(http.MethodDelete)
}

func AddTodoHandler(store *storage.TodoStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Title string `json:"title"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Debugging
		fmt.Println("Received Title:", req.Title)

		if req.Title == "" {
			http.Error(w, "Task title cannot be empty", http.StatusBadRequest)
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
        
        // Sort todos by ID
        sort.Slice(todos, func(i, j int) bool {
            return todos[i].ID < todos[j].ID
        })

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(todos)
    }
}


// MarkTodoCompletedHandler handles marking a todo as completed
func MarkTodoCompletedHandler(store *storage.TodoStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := store.MarkAsCompleted(id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// DeleteTodoHandler handles deleting a todo
func DeleteTodoHandler(store *storage.TodoStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := store.DeleteTodo(id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
