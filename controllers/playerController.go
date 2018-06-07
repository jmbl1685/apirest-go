package controllers

import (
	json "encoding/json"
	http "net/http"

	config "../config"
	models "../models"
	"github.com/gorilla/mux"
	shortid "github.com/jasonsoft/go-short-id"
	bson "gopkg.in/mgo.v2/bson"
)

func AddPlayer(w http.ResponseWriter, r *http.Request) {

	var player models.Player

	opt := shortid.Options{
		Number:        14,
		StartWithYear: true,
		EndWithHost:   false,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&player)

	player.Id = shortid.Generate(opt)
	config.Mongo().Insert(player)
	json.NewEncoder(w).Encode(player)

}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	var players []models.Player
	w.Header().Set("Content-Type", "application/json")
	config.Mongo().Find(nil).All(&players)
	json.NewEncoder(w).Encode(players)
}

func GetPlayerById(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	config.Mongo().Find(bson.M{"id": params["id"]}).One(&player)
	json.NewEncoder(w).Encode(player)
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
	config.Mongo().Update(bson.M{"id": params["id"]}, bson.M{})
	json.NewEncoder(w).Encode(player)
}
