package main

import (
	"log"
	"net/http"

	"./config"
	"./routes"
)

func main() {
	log.Fatal(http.ListenAndServe(config.GetPort(), routes.Routers()))
}
