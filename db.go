package main

import (
  "github.com/jinzhu/gorm"
  "fmt"

  _ "github.com/lib/pq"
)

type Card struct {
  gorm.Model
  ID int
  Front string
  Back string
}

func main() {
  db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=flash_cards sslmode=disable")
  if err != nil {
    panic("failed to connect to db")
  }
  defer db.Close()

  db.AutoMigrate(&Card{})
  //
  // card := Card{
  //   ID: 1,
  //   Front: "Hi",
  //   Back: "Back of the card",
  // }
  // db.NewRecord(card)
  // db.Create(&card)
  // db.NewRecord(card)
  card := Card{}
  db.First(&card)
  fmt.Println(card)
}
