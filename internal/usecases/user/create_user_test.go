package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymoutella/king-poker-bk/internal/domain"
)

func TestCreateUser(t *testing.T) {
	resetTestDatabase(t)

	u, err := NewCreateUserUsecase()

	if err != nil {
		t.Fatalf("Error creating createUserUsecase")
	}

	user, err := u.Execute(&domain.User{
		Email:     "yure@ymoutella.com.br",
		Password:  "TestingPassword",
		FirstName: "Yure",
		LastName:  "Moutella",
	})

	if err != nil {
		t.Fatalf("Error executing createUserUsecase")
	}

	assert.NotEmpty(t, user)

}
