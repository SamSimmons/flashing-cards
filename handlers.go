package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Index route fn
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

// CardIndex returns all cards
func CardIndex(w http.ResponseWriter, r *http.Request) {
	cards := RepoGetAllCards()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cards); err != nil {
		panic(err)
	}
}

// CardShow is the handler for /card/{cardID}
func CardShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardID := vars["cardID"]
	c, err := RepoFindCard(cardID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
	} else if err = json.NewEncoder(w).Encode(c); err != nil {
		panic(err)
	}
}

// CardCreate is the handler for creating a card
func CardCreate(w http.ResponseWriter, r *http.Request) {
	var card Card
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &card); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// unprocessable entitiy
		w.WriteHeader(422)
		errMsg := ErrorMessage{Message: err}
		if err := json.NewEncoder(w).Encode(errMsg); err != nil {
			panic(err)
		}
	}
	fmt.Println("json", json.Unmarshal(body, &card))
	t := RepoCreateCard(card)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func CardDestroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardID := vars["cardID"]
	c, err := RepoDestroyCard(cardID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		errMsg := ErrorMessage{Message: err}
		json.NewEncoder(w).Encode(errMsg)
	} else {
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(c); err != nil {
			panic(err)
		}
	}
}

func HandleCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}
