package models

type Player struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Team        string `json:"team"`
	Position    string `json:"position"`
	Dorsal      int    `json:"dorsal"`
	Nationality string `json:"nationality"`
	ImageUrl    string `json:"image"`
}
