package main

import (
	"log"
	"net/http"
	"github.com/jinzhu/gorm"

)

var DB *gorm.DB

func main() {
	InitDB()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
