package controllers

import (
	"encoding/json"
	"net/http"
)

func AddPlayer(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode("")
}

func GetPlayer(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode("{adsada}")
}

func GetPlayerById(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode("")
}

func DeletePlayer(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode("")
}

func UpdatePlayer(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode("")
}
