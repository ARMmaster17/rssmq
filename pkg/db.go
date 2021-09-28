package pkg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"net/url"
)

func GetDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable",
		os.Getenv("RSSMQ_DB_USER"),
		url.QueryEscape(os.Getenv("RSSMQ_DB_PASSWORD")),
		os.Getenv("RSSMQ_DB_HOST"),
		os.Getenv("RSSMQ_DB_DATABASE"))))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to DB: %w", err)
	}
	return db, nil
}