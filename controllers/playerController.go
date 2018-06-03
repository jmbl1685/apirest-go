package controllers

import (
	"encoding/json"
	"net/http"

	"../config"
	"../models"
	"github.com/gorilla/mux"
)

func AddPlayer(w http.ResponseWriter, r *http.Request) {

	//------------------------ //
	var player models.Player
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&player)
	//------------------------ //

	config.Mongo().Insert(player)

	json.NewEncoder(w).Encode(player)

}

func GetPlayer(w http.ResponseWriter, r *http.Request) {

	player := models.Player{
		Name:        "James Rodriguez",
		Team:        "Bayern Munich",
		Position:    "Forward",
		Dorsal:      11,
		Nationality: "Colombian",
		ImageUrl:    "james.png",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}

func GetPlayerById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["id"])
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("")
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("")
}
