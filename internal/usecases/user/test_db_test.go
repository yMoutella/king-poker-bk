package usecases

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/ymoutella/king-poker-bk/internal/database"
	"github.com/ymoutella/king-poker-bk/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultTestDBName = "king_poker_test"

func TestMain(m *testing.M) {
	if os.Getenv("POSTGRES_DB") == "" {
		_ = os.Setenv("POSTGRES_DB", defaultTestDBName)
	}

	if err := ensureTestDatabase(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}

func resetTestDatabase(t *testing.T) {
	t.Helper()

	db, err := database.PostgresDB()
	if err != nil {
		t.Fatalf("Error connecting to test database: %v", err)
	}

	if err := db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error; err != nil {
		t.Fatalf("Error resetting users table: %v", err)
	}
}

func ensureTestDatabase() error {
	testDBName := os.Getenv("POSTGRES_DB")
	if testDBName == "" {
		return errors.New("POSTGRES_DB must be set for tests")
	}
	if !isValidDBName(testDBName) {
		return fmt.Errorf("invalid test database name: %s", testDBName)
	}

	adminDSN := database.PostgresDSNForDB("postgres")
	adminDB, err := gorm.Open(postgres.Open(adminDSN), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to postgres database: %w", err)
	}

	var exists bool
	if err := adminDB.Raw(
		"SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = ?)",
		testDBName,
	).Scan(&exists).Error; err != nil {
		return fmt.Errorf("error checking test database: %w", err)
	}

	if !exists {
		if err := adminDB.Exec(fmt.Sprintf("CREATE DATABASE %s", testDBName)).Error; err != nil {
			return fmt.Errorf("error creating test database: %w", err)
		}
	}

	testDB, err := database.PostgresDB()
	if err != nil {
		return fmt.Errorf("error connecting to test database: %w", err)
	}

	if err := testDB.AutoMigrate(&domain.User{}); err != nil {
		return fmt.Errorf("error migrating test database: %w", err)
	}

	return nil
}

func isValidDBName(name string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(name)
}
