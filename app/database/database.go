package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "127.0.0.1"
	user     = "postgres"
	password = "postgres"
	dbName   = "pet_catalog_development"
	port     = 5432
	sslMode  = "disable"
)

// DB ...
type DB struct {
	*gorm.DB
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectParams()), &gorm.Config{})

	if err != nil {
		log.Println(
			err.Error(),
		)

		panic("failed to connect database")
	}

	log.Println("Connection Opened to Database")

	return db
}

func connectParams() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		host, user, password, dbName, port, sslMode,
	)
}
