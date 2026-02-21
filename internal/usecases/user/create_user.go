package usecases

import (
	"github.com/ymoutella/king-poker-bk/internal/database"
	"github.com/ymoutella/king-poker-bk/internal/domain"
	"github.com/ymoutella/king-poker-bk/internal/repository"
)

type CreateUserUsecase interface {
	Execute(args ...any) (any, error)
}

type createUserUsecase struct {
	rep repository.UserRepository
}

func NewCreateUserUsecase() (*createUserUsecase, error) {
	db, err := database.PostgresDB()

	if err != nil {
		panic("Error connecting in db")
	}
	rep := repository.NewUserRepository(db)
	return &createUserUsecase{rep: rep}, nil
}

func (u *createUserUsecase) Execute(user *domain.User) (*domain.User, error) {

	createdUser, err := u.rep.Create(user)

	if err != nil {
		return nil, err
	}

	return createdUser, nil

}
