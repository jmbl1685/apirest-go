package routes

import (
	"fmt"
	http "net/http"

	controllers "../controllers"
	mux "github.com/gorilla/mux"
)

func Routers() (router *mux.Router) {

	method := map[string]string{
		"HTTP_GET":    "GET",
		"HTTP_POST":   "POST",
		"HTTP_PUT":    "PUT",
		"HTTP_DELETE": "DELETE",
	}

	router = mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>API RUNNING</h1>")
	})

	router.HandleFunc("/player", controllers.AddPlayer).Methods(method["HTTP_POST"])
	router.HandleFunc("/player", controllers.GetPlayer).Methods(method["HTTP_GET"])
	router.HandleFunc("/player/{id}", controllers.GetPlayerById).Methods(method["HTTP_GET"])
	router.HandleFunc("/player/{id}", controllers.DeletePlayer).Methods(method["HTTP_DELETE"])
	router.HandleFunc("/player/{id}", controllers.UpdatePlayer).Methods(method["HTTP_PUT"])
	return

}
