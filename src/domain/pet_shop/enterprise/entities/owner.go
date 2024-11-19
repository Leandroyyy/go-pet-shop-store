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
}

type Owner struct {
	Id string
	OwnerProps
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
