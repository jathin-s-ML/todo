# Todo API

A simple Todo List API built with Golang using Gorilla Mux for routing. It supports adding, retrieving, updating, and deleting tasks, with authentication and logging middleware.

## 🚀 Features

- Add, retrieve, update, and delete todos.
- Middleware for authentication and logging.
- RESTful API using Gorilla Mux.
- Thread-safe in-memory storage.

## 📁 Project Structure

```
.
├── internal
│   ├── handlers
│   │   ├── todo.go         # Todo handlers (API endpoints)
│   ├── middleware
│   │   ├── auth.go         # Authentication middleware
│   │   ├── logging.go      # Logging middleware
│   ├── storage
│   │   ├── todostorage.go  # In-memory storage for todos
│   │   ├── todostorage_test.go  # Unit tests for storage
├── main.go                 # Entry point
├── go.mod                  # Go module file
├── go.sum                  # Dependencies checksum
```

## 🛠 Setup & Installation

### Prerequisites

- Install [Go](https://go.dev/dl/)
- Install [Gorilla Mux](https://github.com/gorilla/mux)

### Clone Repository

```sh
git clone https://github.com/jathin-s-ML/todo.git
cd todo
```

### Install Dependencies

```sh
go mod tidy
```

### Run the Application

```sh
go run main.go
```

Server will start on `http://localhost:8080`

## 📌 API Endpoints

| Method | Endpoint      | Description          | Auth Required |
| ------ | ------------- | -------------------- | ------------- |
| GET    | `/`           | Welcome Message      | No            |
| GET    | `/todos`      | Get all todos        | Yes           |
| POST   | `/todos`      | Add a new todo       | Yes           |
| PUT    | `/todos/{id}` | Mark a todo complete | Yes           |
| DELETE | `/todos/{id}` | Delete a todo        | Yes           |

## 🔑 Authentication

- Uses **Basic Authentication**.
- Default credentials:
  - **Username:** `admin`
  - **Password:** `password`

## 🧪 Running Tests

Tests are written using Go's testing package.
Run tests with:

```sh
go test ./internal/storage -v
```
