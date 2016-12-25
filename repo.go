package main

import (
	"github.com/jinzhu/gorm"
	"fmt"

	_ "github.com/lib/pq"
	)


type Card struct {
	ID        int       `json:"id"`
	Front 		string 		`json:"front"`
	Back 			string 		`json:"back"`
}

// Todos is a collection of Todo's
type Cards []Card


// func Init() {
// 	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=flash_cards sslmode=disable")
// 	if err != nil {
// 		panic("failed to connect to db")
// 	}
// 	defer db.Close()
// 	fmt.Println("db connected")
// 	if !db.HasTable(&Card{}) {
// 		fmt.Println("create table")
// 		db.CreateTable(&Card{})
// 	}
// }

func main() {
	// Init()
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=flash_cards sslmode=disable")
	if err != nil {
		panic("failed to connect to db")
	}
	defer db.Close()
	// test := Card{
	// 	ID: 2,
	// 	Front: "This is the front",
	// 	Back: "da booty",
	// }
	// RepoCreateCard(test, *db)
	// test := RepoFindCard(2, *db)
	// fmt.Println(test)
	RepoDestroyCard(2, *db)
}

// RepoFindTodo finds a todo
func RepoFindCard(id int, db gorm.DB) Card {
	card := Card{}
	db.Find(&card, id)
	return card
}

// RepoCreateCard creates a card
func RepoCreateCard(c Card, db gorm.DB) Card {
	db.Create(&c)
	return c
}

// RepoDestroyTodo destroys a todo
func RepoDestroyCard(id int, db gorm.DB) error {
	
}
