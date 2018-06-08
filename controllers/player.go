package controllers

import (
	json "encoding/json"
	"fmt"
	http "net/http"

	config "../config"
	models "../models"
	utilities "../utilities"
	mux "github.com/gorilla/mux"
	bson "gopkg.in/mgo.v2/bson"
)

func AddPlayer(w http.ResponseWriter, r *http.Request) {

	var player models.Player
	player.Id = utilities.GenerateId()

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&player)
	config.Mongo().Insert(player)
	json.NewEncoder(w).Encode(player)

}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	var players []models.Player
	w.Header().Set("Content-Type", "application/json")
	config.Mongo().Find(nil).All(&players)
	fmt.Print(len(players))

	if len(players) == 0 {
		json.NewEncoder(w).Encode(
			utilities.Response{Message: "There are not players", Status: 404})
	} else {
		json.NewEncoder(w).Encode(players)
	}

}

func GetPlayerById(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	config.Mongo().Find(bson.M{"id": params["id"]}).One(&player)

	if player.Id == "" {
		json.NewEncoder(w).Encode(
			utilities.Response{Message: "Player not found", Status: 404})
	} else {
		json.NewEncoder(w).Encode(player)
	}
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	config.Mongo().Remove(bson.M{"id": params["id"]})
	json.NewEncoder(w).Encode(player)
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&player)

	body := bson.M{
		"id":          params["id"],
		"name":        player.Name,
		"team":        player.Team,
		"position":    player.Position,
		"nationality": player.Nationality,
		"image":       player.Image,
		"dorsal":      player.Dorsal,
	}
	config.Mongo().Update(bson.M{"id": params["id"]}, body)
	json.NewEncoder(w).Encode(body)
}
