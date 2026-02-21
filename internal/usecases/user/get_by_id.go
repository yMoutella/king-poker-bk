package usecases

import (
	"github.com/google/uuid"
	"github.com/ymoutella/king-poker-bk/internal/database"
	"github.com/ymoutella/king-poker-bk/internal/domain"
	"github.com/ymoutella/king-poker-bk/internal/repository"
)

type GetByIdUsecase interface {
	Execute(args ...any) (any, error)
}

type getByIdUsecase struct {
	rep repository.UserRepository
}

func CreateNewGetByIdUsecase() *getByIdUsecase {
	db, err := database.PostgresDB()

	if err != nil {
		panic("Error connecting in db")
	}

	rep := repository.NewUserRepository(db)

	return &getByIdUsecase{rep: rep}
}

func (u *getByIdUsecase) Execute(id uuid.UUID) (*domain.User, error) {

	user, err := u.rep.GetByID(id)

	if err != nil {
		return nil, err
	}

	return user, nil

}
