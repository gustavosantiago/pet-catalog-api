package repository

import (
	"pet-catalog-api/app/models"
)

type PetRepository interface {
	Fetch(limit int, page int) []*models.Pet
}
