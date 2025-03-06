package storage

import (
	"errors"
	"sync"
)

// Todo represents a task
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// TodoStorage manages todos in memory
type TodoStorage struct {
	mu    sync.Mutex
	todos []Todo
	nextID int
}

// NewTodoStorage initializes the storage
func NewTodoStorage() *TodoStorage {
	return &TodoStorage{
		todos:  []Todo{},
		nextID: 1,
	}
}

// AddTodo adds a new todo
func (s *TodoStorage) AddTodo(title string) Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := Todo{ID: s.nextID, Title: title, Completed: false}
	s.todos = append(s.todos, todo)
	s.nextID++
	return todo
}

// GetTodos retrieves all todos
func (s *TodoStorage) GetTodos() []Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.todos
}

// MarkAsCompleted marks a todo as completed
func (s *TodoStorage) MarkAsCompleted(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, todo := range s.todos {
		if todo.ID == id {
			s.todos[i].Completed = true
			return nil
		}
	}
	return errors.New("todo not found")
}

// DeleteTodo removes a todo by ID
func (s *TodoStorage) DeleteTodo(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, todo := range s.todos {
		if todo.ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
