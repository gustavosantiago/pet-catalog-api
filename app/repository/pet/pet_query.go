package pet

import (
	"context"
	"pet-catalog-api/app/models"
	"pet-catalog-api/app/repository"
	"time"

	"gorm.io/gorm"
)

type PetRepo struct {
	Conn *gorm.DB
}

func NewPetRepo(Conn *gorm.DB) repository.PetRepository {
	return &PetRepo{
		Conn: Conn,
	}
}

func (p *PetRepo) Fetch(limit int, page int) []*models.Pet {
	var pets []*models.Pet
	p.Conn.Limit(limit).Offset(page).Select("id", "name", "description", "breed", "url").Find(&pets)

	return pets
}

func (pr *PetRepo) Create(ctx context.Context, p *models.Pet) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	result := pr.Conn.Create(&p)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
