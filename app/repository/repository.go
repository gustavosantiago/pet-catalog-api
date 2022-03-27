package repository

import (
	"pet-catalog-api/app/models"
)

type PetRepository interface {
	Fetch(limit int, page int) []*models.Pet
	GetByID(id int64) (models.Pet, error)
	Create(p *models.Pet) error
	Update(id int64, p models.UpdatePetInput) (models.Pet, error)
	Delete(id int64) (bool, error)
}
