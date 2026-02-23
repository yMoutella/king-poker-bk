package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymoutella/king-poker-bk/internal/domain"
)

func TestGetByID(t *testing.T) {
	resetTestDatabase(t)

	createUserUsecase, err := FactoryCreateUserUS()
	if err != nil {
		t.Fatalf("Error creating createUserUsecase")
	}

	user := &domain.User{
		Email:     "getbyid@ymoutella.com.br",
		Password:  "TestingPassword",
		FirstName: "Yure",
		LastName:  "Moutella",
	}

	createdUser, err := createUserUsecase.Execute(user)
	if err != nil {
		t.Fatalf("Error executing createUserUsecase")
	}

	u := FactoryGetByIdUS()
	returnedUser, err := u.Execute(createdUser.ID)
	if err != nil {
		t.Fatalf("Error executing getByIdUsecase")
	}

	assert.NotEmpty(t, returnedUser)
	assert.Equal(t, createdUser.ID, returnedUser.ID)
	assert.Equal(t, createdUser.Email, returnedUser.Email)
}
