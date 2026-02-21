package usecases

import (
	"testing"

	"github.com/ymoutella/king-poker-bk/internal/domain"
)

func TestGetAll(t *testing.T) {
	resetTestDatabase(t)

	createUserUsecase, err := NewCreateUserUsecase()

	if err != nil {
		t.Fatalf("Error creating createUserUsecase")
	}

	user := domain.User{
		Email:     "ymoutella@ymoutella.com.br",
		Password:  "TestingPassword",
		FirstName: "Yure",
		LastName:  "Moutella",
	}

	createUserUsecase.Execute(&user)

	u := NewGetAllUserUsecase()

	returnedUser, err := u.Execute(1, 1)

	if err != nil {
		t.Fatalf("Error executing getAllUsecase")
	}

	if len(returnedUser) == 0 {
		t.Errorf("The user list should return filled")
	}

}
