package config

import (
	mongodb "gopkg.in/mgo.v2"
)

func GetPort() string {
	return ":3000"
}

func Mongo() *mongodb.Collection {

	const (
		Host       = "HOST_DB"
		Username   = "USERNAME_DB"
		Password   = "PASSWORD_DB"
		Database   = "DATABASE_DB"
		Collection = "COLLECTION_DB"
	)

	session, err := mongodb.DialWithInfo(&mongodb.DialInfo{
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
