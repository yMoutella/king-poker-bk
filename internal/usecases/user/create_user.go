package usecases

import (
	"github.com/ymoutella/king-poker-bk/internal/database"
	"github.com/ymoutella/king-poker-bk/internal/domain"
	"github.com/ymoutella/king-poker-bk/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserUsecase interface {
	Execute(args ...any) (any, error)
}

type createUserUsecase struct {
	rep repository.UserRepository
}

func FactoryCreateUserUS() (*createUserUsecase, error) {
	db, err := database.PostgresDB()

	if err != nil {
		panic("Error connecting in db")
	}
	rep := repository.NewUserRepository(db)
	return &createUserUsecase{rep: rep}, nil
}

func (u *createUserUsecase) Execute(user *domain.User) (*domain.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

	if err != nil {
		return nil, err
	}

	user.Password = string(hash)

	createdUser, err := u.rep.Create(user)

	if err != nil {
		return nil, err
	}

	createdUser.Password = ""

	return createdUser, nil

}
