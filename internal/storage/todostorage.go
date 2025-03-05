package storage

import (
	"errors"
	"sync"

	"github.com/jathin-s-ML/todo/internal/models"
)

// TodoStorage manages all todo operations
type TodoStorage struct {
	mu     sync.Mutex
	todos  map[int]*models.Todo
	nextID int
}

// NewTodoStorage initializes an empty todo storage
func NewTodoStorage() *TodoStorage {
	return &TodoStorage{
		todos:  make(map[int]*models.Todo),
		nextID: 1,
	}
}

// AddTodo adds a new task to the list
func (s *TodoStorage) AddTodo(title string) *models.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := &models.Todo{
		ID:        s.nextID,
		Title:     title,
		Completed: false,
	}
	s.todos[s.nextID] = todo
	s.nextID++
	return todo
}

// GetTodos retrieves all tasks
func (s *TodoStorage) GetTodos() []*models.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	var todos []*models.Todo
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}
	return todos
}

// MarkAsCompleted marks a task as completed
func (s *TodoStorage) MarkAsCompleted(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if todo, exists := s.todos[id]; exists {
		todo.Completed = true
		return nil
	}
	return errors.New("task not found")
}

// DeleteTodo removes a task from the list
func (s *TodoStorage) DeleteTodo(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.todos[id]; exists {
		delete(s.todos, id)
		return nil
	}
	return errors.New("task not found")
}
