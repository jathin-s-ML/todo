package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jathin-s-ML/todo/internal/handlers"
	"github.com/jathin-s-ML/todo/internal/storage"
)

func main() {
	fmt.Println("Starting TODO Application...")

	// Initialize Todo storage
	store := storage.NewTodoStorage()

	// Setup router
	r := mux.NewRouter()
	handlers.RegisterTodoRoutes(r, store)

	// Start HTTP Server
	port := ":8080"
	fmt.Println("Server running on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
