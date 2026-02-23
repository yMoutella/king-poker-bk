package database

import (
	"fmt"
	"os"

	"github.com/ymoutella/king-poker-bk/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	TimeZone string
}

func PostgresConfigFromEnv() PostgresConfig {
	return PostgresConfig{
		Host:     getEnv("POSTGRES_HOST", "localhost"),
		User:     getEnv("POSTGRES_USER", "postgres"),
		Password: getEnv("POSTGRES_PASSWORD", "postgres"),
		DBName:   getEnv("POSTGRES_DB", "king_poker"),
		Port:     getEnv("POSTGRES_PORT", "5432"),
		TimeZone: getEnv("POSTGRES_TIMEZONE", "America/Sao_Paulo"),
	}
}

func PostgresDSNForDB(dbName string) string {
	cfg := PostgresConfigFromEnv()
	if dbName == "" {
		dbName = cfg.DBName
	}
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		dbName,
		cfg.Port,
		cfg.TimeZone,
	)
}

func PostgresDB() (*gorm.DB, error) {
	dsn := PostgresDSNForDB("")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&domain.User{})

	return db, err

}

func MockDB() any {
	return nil
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
