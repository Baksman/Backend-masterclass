package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBSetUp() (*gorm.DB, error) {
	dsn := "host=localhost user=ibrahim password=ibrahim dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}


