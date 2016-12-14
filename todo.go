package main

import "time"

// Todo is a model of a todo
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos is a collection of Todo's
type Todos []Todo
