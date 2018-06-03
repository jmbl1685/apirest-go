package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Team        string        `json:"team"`
	Position    string        `json:"position"`
	Dorsal      int           `json:"dorsal"`
	Nationality string        `json:"nationality"`
	ImageUrl    string        `json:"image"`
}
