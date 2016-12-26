package main

// Card is a model of a card
type Card struct {
	ID        int       `json:"id"`
	Front 		string 		`json:"front"`
	Back 			string 		`json:"back"`
}

// Cards is a collection of Todo's
type Cards []Card
