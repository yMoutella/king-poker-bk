package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=king_poker port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, err

}

func MockDB() any {
	return nil
}
