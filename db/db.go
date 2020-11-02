package db

import (
	"context"

	"github.com/sergiosegrera/covlog/models"
)

type DB interface {
	SavePerson(context.Context, models.Person) error
	GetPersons(context.Context) ([]models.Person, error)
}
