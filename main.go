package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/jmbl1685/apirest-go/models"
	"github.com/gorilla/mux"
)

var playerList []Player

func main() {

	player := Player {
		Id: "1", 
		Name: "Lionel Messi", 
		Team: "Barcelona FC",
		Position: "Forward",
		Dorsal: 10,
		Nationality: "Argentinian",
		ImageUrl: "https://pbs.twimg.com/profile_images/478353923992854528/iqLeQfIT.png",
	}

	playerList = append(playerList, player)

	router := mux.NewRouter()
	router.HandleFunc("/player", GetPlayer).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func GetPlayer(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(playerList)
}
