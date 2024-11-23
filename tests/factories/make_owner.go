package test_factories

import (
	"time"

	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
)

func MakeOwner() *entities.Owner {
	defaultProps := entities.OwnerProps{
		Name:     "john due",
		Document: "123132213",
		Birthday: time.Date(2003, 2, 20, 0, 0, 0, 0, time.UTC),
		Email:    "john@due.com",
	}

	owner := entities.NewOwner(defaultProps, nil)

	return &owner
}
