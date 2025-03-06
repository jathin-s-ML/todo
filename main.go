package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jathin-s-ML/todo/internal/handlers"
	"github.com/jathin-s-ML/todo/internal/middleware"
	"github.com/jathin-s-ML/todo/internal/storage"
)

func main() {
	store := storage.NewTodoStorage()
	router := mux.NewRouter()

	// Apply Logging Middleware globally
	router.Use(middleware.LoggingMiddleware)

	// Public Routes (No Auth)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Todo API!"))
	}).Methods("GET")

	// Protected Routes (Require Auth)
	protectedRoutes := router.PathPrefix("/todos").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)
	handlers.RegisterTodoRoutes(protectedRoutes, store)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)
}
