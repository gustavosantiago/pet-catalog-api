package handler

import (
	"encoding/json"
	"net/http"
	"pet-catalog-api/app/models"
	"pet-catalog-api/app/repository"
	"pet-catalog-api/app/repository/pet"
	"strconv"

	"github.com/go-chi/chi"
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

// GetByID get a pet data by id
func (p *Pet) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	result, err := p.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
	}

	respondwithJSON(w, http.StatusOK, result)
}

// Create a new post
func (p *Pet) Create(w http.ResponseWriter, r *http.Request) {
	pet := models.Pet{}
	json.NewDecoder(r.Body).Decode(&pet)

	err := p.repo.Create(r.Context(), &pet)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
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

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
