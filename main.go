package main

import (
	log "log"
	http "net/http"

	config "./config"
	routes "./routes"
	cors "github.com/rs/cors"
)

func main() {
	handler := cors.Default().Handler(routes.Routers())
	log.Fatal(http.ListenAndServe(config.GetPort(), handler))
}
