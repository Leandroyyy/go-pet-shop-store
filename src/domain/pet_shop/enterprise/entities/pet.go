package entities

import (
	"time"

	"github.com/google/uuid"
)

type PetKind string

type PetGender string

const (
	Male   PetGender = "male"
	Female PetGender = "female"
)

const (
	Dog    PetKind = "dog"
	Cat    PetKind = "cat"
	Turtle PetKind = "turtle"
)

type PetProps struct {
	Name     string
	Birthday time.Time
	Breed    string
	Gender   PetGender
	Kind     PetKind
}

type Pet struct {
	id string
	PetProps
}

func NewPet(pet PetProps, id *string) Pet {

	if id == nil {
		newID := uuid.New().String()
		id = &newID
	}

	return Pet{
		id:       *id,
		PetProps: pet,
	}
}
