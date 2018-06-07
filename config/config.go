package config

import (
	"gopkg.in/mgo.v2"
)

func GetPort() string {
	return ":3000"
}

func Mongo() *mgo.Collection {

	const (
		Host       = "ds046367.mlab.com:46367"
		Username   = "root"
		Password   = "12345"
		Database   = "db-social-mean5"
		Collection = "player"
	)

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{Host},
		Username: Username,
		Password: Password,
		Database: Database,
	})

	if err != nil {
		panic(err)
	}

	return session.DB(Database).C(Collection)
}
