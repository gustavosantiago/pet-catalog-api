package repository

import (
	"context"
	"pet-catalog-api/app/models"
)

type PetRepository interface {
	Fetch(limit int, page int) []*models.Pet
	Create(ctx context.Context, p *models.Pet) error
}
