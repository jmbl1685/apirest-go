package config

import (
	mongodb "gopkg.in/mgo.v2"
)

func GetPort() string {
	return ":3000"
}

func Mongo() *mongodb.Collection {

	const (
		Host       = "ds046367.mlab.com:46367"
		Username   = "root"
		Password   = "12345"
		Database   = "db-social-mean5"
		Collection = "player"
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
