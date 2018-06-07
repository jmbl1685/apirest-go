
# API REST Example using Go [Golang]

Used modules
```
go get github.com/gorilla/mux
go get github.com/rs/cors
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson
go get github.com/jasonsoft/go-short-id
```

## CRUD Football player


Player Model:

```go
type Player struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Team        string `json:"team"`
	Position    string `json:"position"`
	Dorsal      int    `json:"dorsal"`
	Nationality string `json:"nationality"`
	ImageUrl    string `json:"image"`
}
```
Player Controllers:

```go
func AddPlayer(w http.ResponseWriter, r *http.Request){...}
func GetPlayer(w http.ResponseWriter, r *http.Request){...}
func GetPlayerById(w http.ResponseWriter, r *http.Request){...}
func DeletePlayer(w http.ResponseWriter, r *http.Request){...}
func UpdatePlayer(w http.ResponseWriter, r *http.Request){...}
```
Player Routes:
```go
router = mux.NewRouter()

router.HandleFunc("/player", controllers.AddPlayer).Methods("POST")
router.HandleFunc("/player", controllers.GetPlayer).Methods("GET")
router.HandleFunc("/player/{id}", controllers.GetPlayerById).Methods("GET")
router.HandleFunc("/player/{id}", controllers.DeletePlayer).Methods("DELETE")
router.HandleFunc("/player/{id}", controllers.UpdatePlayer).Methods("PUT")
```
