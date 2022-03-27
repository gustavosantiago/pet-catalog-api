package main

import (
	"log"
	"net/http"
	"pet-catalog-api/app/database"
	"pet-catalog-api/app/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	var conn *gorm.DB = database.ConnectDB()

	pHandler := handler.NewPetHandler(conn)
	r.Get("/pets", pHandler.Fetch)
	r.Post("/pets", pHandler.Create)
	r.Get("/pets/{id}", pHandler.GetByID)
	r.Put("/pets/{id}", pHandler.Update)
	r.Delete("/pets/{id}", pHandler.Delete)

	log.Println("Server listen at :3030")
	http.ListenAndServe(":3030", r)
}
