package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Index route fn
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

// TodoIndex is the handler for /todos
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoShow is the handler for /todos/{todoID}
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	fmt.Fprintln(w, "Todo Show:", todoID)
}
