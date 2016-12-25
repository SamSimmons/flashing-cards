package main

// Todo is a model of a todo
type Card struct {
	ID        int       `json:"id"`
	Front 		string 		`json:"front"`
	Back 			string 		`json:"back"`
}

// Todos is a collection of Todo's
type Cards []Card
