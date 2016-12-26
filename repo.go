package main

import (
	"github.com/jinzhu/gorm"
	"fmt"

	_ "github.com/lib/pq"
	)

func InitDB() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost user=postgres dbname=flash_cards sslmode=disable")
	if err != nil {
		panic("failed to connect to db")
	}
	// defer DB.Close()
	fmt.Println("db connected")
	if !DB.HasTable(&Card{}) {
		fmt.Println("create table")
		DB.CreateTable(&Card{})
	}
}

// RepoFindTodo finds a todo
func RepoFindCard(id string) Card {
	card := Card{}
	DB.Find(&card, id)
	return card
}

func RepoGetAllCards() Cards {
	cards := Cards{}
	DB.Find(&cards)
	fmt.Println(cards)
	return cards
}

// RepoCreateCard creates a card
func RepoCreateCard(c Card) Card {
	DB.Create(&c)
	return c
}

// RepoDestroyTodo destroys a todo
func RepoDestroyCard(id int) {
	card := Card{}
	DB.Find(&card, id)
	DB.Delete(&card)
	fmt.Println(card)
}
