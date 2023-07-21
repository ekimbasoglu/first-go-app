package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func main() {
	// Set up routes
	http.HandleFunc("/todos", getAllTodos)
	http.HandleFunc("/todos/create", createTodo)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler to get all todos
func getAllTodos(w http.ResponseWriter, r *http.Request) {
	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Encode todos into JSON format and send the response
	json.NewEncoder(w).Encode(todos)
}

// Handler to create a new todo
func createTodo(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to Todo object
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Assign a unique ID and add the new todo to the list
	todo.ID = len(todos) + 1
	todos = append(todos, todo)

	// Send a success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}


