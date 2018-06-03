package main

import (
	"log"
	"net/http"

	"./config"
	"./routes"
	"github.com/rs/cors"
)

func main() {
	handler := cors.Default().Handler(routes.Routers())
	log.Fatal(http.ListenAndServe(config.GetPort(), handler))
}
