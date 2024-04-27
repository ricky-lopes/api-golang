package entities

import "github.com/google/uuid"

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewPerson() *Person {
	person := Person{
		ID: uuid.New().String(),
	}

	return &person
}
