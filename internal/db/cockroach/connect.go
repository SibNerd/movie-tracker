package cockroach

import (
	"fmt"
	"movietracker/internal/db"
	entity "movietracker/internal/entities"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func ConnectDB() (db.DB, error) {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &Database{}, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Show{},
		&entity.UserShow{})
	if err != nil {
		return &Database{}, err
	}

	fmt.Printf("Connected to Database %s\n", db.Migrator().CurrentDatabase())

	return &Database{DB: db}, err
}
