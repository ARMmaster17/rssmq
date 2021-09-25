package pkg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func GetDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("RSSMQ_DB_HOST"),
		os.Getenv("RSSMQ_DB_USER"),
		os.Getenv("RSSMQ_DB_PASSWORD"),
		os.Getenv("RSSMQ_DB_DATABASE"))))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to DB: %w", err)
	}
	return db, nil
}
