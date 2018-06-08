package utiities

import (
	shortid "github.com/jasonsoft/go-short-id"
)

type Response struct {
	Message string
	Status  int
}

func GenerateId() string {

	options := shortid.Options{
		Number:        14,
		StartWithYear: true,
		EndWithHost:   false,
	}

	return shortid.Generate(options)
}
