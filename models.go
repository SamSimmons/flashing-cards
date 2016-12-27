package main

import 	"github.com/jinzhu/gorm"

// Card is a model of a card
type Card struct {
	gorm.Model
	Front 		string 		`json:"front"`
	Back 			string 		`json:"back"`
}

// Cards is a collection of Todo's
type Cards []Card

type ErrorMessage struct {
	Message error `json:"message"`
}
