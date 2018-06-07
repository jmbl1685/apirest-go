package routes

import (
	controllers "../controllers"
	mux "github.com/gorilla/mux"
)

func Routers() (router *mux.Router) {

	router = mux.NewRouter()

	/* Player routes */
	router.HandleFunc("/player", controllers.AddPlayer).Methods("POST")
	router.HandleFunc("/player", controllers.GetPlayer).Methods("GET")
	router.HandleFunc("/player/{id}", controllers.GetPlayerById).Methods("GET")
	router.HandleFunc("/player/{id}", controllers.DeletePlayer).Methods("DELETE")
	router.HandleFunc("/player/{id}", controllers.UpdatePlayer).Methods("PUT")
	return

}
