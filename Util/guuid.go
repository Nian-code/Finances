package main

import 	guid "github.com/google/uuid"

func GenerateUID() string {
	id := guid.New()
	return id.String()
}
