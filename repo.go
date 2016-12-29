package main

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"errors"

	_ "github.com/lib/pq"
	)

func InitDB() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost user=postgres dbname=flash_cards sslmode=disable")
	if err != nil {
		panic("failed to connect to db")
	}

	fmt.Println("db connected")
	if !DB.HasTable(&Card{}) {
		fmt.Println("create table")
		DB.CreateTable(&Card{})
	}
	DB.AutoMigrate(&Card{})
}

// RepoFindTodo finds a todo
func RepoFindCard(id string) (Card, error) {
	card := Card{}
	DB.Find(&card, id)
	if card.Model.ID == 0 {
		return card, errors.New("Can't find card")
	}
	return card, nil
}

func RepoGetAllCards() Cards {
	cards := Cards{}
	DB.Find(&cards)
	fmt.Println(cards)
	return cards
}

// RepoCreateCard creates a card
func RepoCreateCard(c Card) Card {
	fmt.Println("coming in to db", c)
	DB.Create(&c)
	return c
}

// RepoDestroyTodo destroys a todo
func RepoDestroyCard(id string) (Card, error) {
	card := Card{}
	DB.Find(&card, id)
	if card.Model.ID == 0 {
		return card, errors.New("Can't find card")
	}
	DB.Delete(&card)
	return card, nil
}
