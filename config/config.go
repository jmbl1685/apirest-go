package config

import (
	"gopkg.in/mgo.v2"
)

func GetPort() string {
	return ":3000"
}

func Mongo() *mgo.Collection {

	const (
		Host       = "HOST_DB"
		Username   = "USERNAME_DB"
		Password   = "PASSWORD_DB"
		Database   = "DATABASE_DB"
		Collection = "COLLECTION_DB"
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

	defer session.Close()

	return session.DB(Database).C(Collection)
}
