package entities

import (
	"time"

	"github.com/google/uuid"
)

type OwnerProps struct {
	Name     string
	Document string
	Birthday time.Time
	Email    string
	Pets     *[]Pet
}

type Owner struct {
	Id string
	OwnerProps
}

func (o Owner) RegisterPet(pet Pet) {

	if o.Pets == nil {
		slice := []Pet{}
		o.Pets = &slice
	}

	*o.Pets = append(*o.Pets, pet)

}

func NewOwner(owner OwnerProps, id *string) Owner {

	if id == nil {
		newID := uuid.New().String()
		id = &newID
	}

	return Owner{
		Id:         *id,
		OwnerProps: owner,
	}
}
