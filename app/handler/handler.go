package handler

import (
	"encoding/json"
	"net/http"
	"pet-catalog-api/app/repository"
	"pet-catalog-api/app/repository/pet"
	"strconv"

	"gorm.io/gorm"
)

// NewPetHandler ...
func NewPetHandler(db *gorm.DB) *Pet {
	return &Pet{
		repo: pet.NewPetRepo(db),
	}
}

// Pet ...
type Pet struct {
	repo repository.PetRepository
}

// Fetch all pet data
func (p *Pet) Fetch(w http.ResponseWriter, r *http.Request) {
	var pageParams string = r.URL.Query().Get("page")
	var limitParams string = r.URL.Query().Get("limit")

	page, limit := setPageAndLimit(pageParams, limitParams)

	response := p.repo.Fetch(limit, page)

	respondwithJSON(w, http.StatusOK, response)
}

func setPageAndLimit(pageString string, limitString string) (int, int) {
	page, err := strconv.ParseInt(pageString, 10, 64)

	if err != nil {
		page = 0
	}

	limit, err := strconv.ParseInt(limitString, 10, 64)

	if err != nil {
		limit = 15
	}

	return int(page), int(limit)
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
