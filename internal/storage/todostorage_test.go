package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTodo(t *testing.T) {
	store := NewTodoStorage()
	todo := store.AddTodo("Test Task")

	assert.Equal(t, 1, todo.ID)
	assert.Equal(t, "Test Task", todo.Title)
	assert.False(t, todo.Completed)

	// Check if the todo is stored
	todos := store.GetTodos()
	assert.Len(t, todos, 1)
	assert.Equal(t, todo, todos[0])
}

func TestGetTodos(t *testing.T) {
	store := NewTodoStorage()
	store.AddTodo("Task 1")
	store.AddTodo("Task 2")

	todos := store.GetTodos()

	assert.Len(t, todos, 2)
	assert.Equal(t, "Task 1", todos[0].Title)
	assert.Equal(t, "Task 2", todos[1].Title)
}

func TestMarkAsCompleted(t *testing.T) {
	store := NewTodoStorage()
	todo := store.AddTodo("Test Task")

	// Mark as completed
	err := store.MarkAsCompleted(todo.ID)
	assert.Nil(t, err)

	// Verify completion status
	todos := store.GetTodos()
	assert.True(t, todos[0].Completed)
}

func TestMarkAsCompletedInvalidID(t *testing.T) {
	store := NewTodoStorage()

	// Try marking a non-existent task
	err := store.MarkAsCompleted(999)
	assert.NotNil(t, err)
	assert.Equal(t, "todo not found", err.Error())
}

func TestDeleteTodo(t *testing.T) {
	store := NewTodoStorage()
	todo := store.AddTodo("Task to delete")

	// Delete the todo
	err := store.DeleteTodo(todo.ID)
	assert.Nil(t, err)

	// Ensure it's removed
	todos := store.GetTodos()
	assert.Len(t, todos, 0)
}

func TestDeleteTodoInvalidID(t *testing.T) {
	store := NewTodoStorage()

	// Try deleting a non-existent task
	err := store.DeleteTodo(999)
	assert.NotNil(t, err)
	assert.Equal(t, "todo not found", err.Error())
}
