package usecases

import (
	"github.com/ymoutella/king-poker-bk/internal/database"
	"github.com/ymoutella/king-poker-bk/internal/domain"
	"github.com/ymoutella/king-poker-bk/internal/repository"
)

type GetAllUseCase interface {
	Execute(args ...any) (any, error)
}

type getAllUseCase struct {
	rep repository.UserRepository
}

func NewGetAllUserUsecase() *getAllUseCase {
	db, err := database.PostgresDB()

	if err != nil {
		panic("Error connecting in db")
	}

	rep := repository.NewUserRepository(db)
	return &getAllUseCase{rep: rep}
}

func (u *getAllUseCase) Execute(page int, pageSize int) ([]domain.User, error) {

	users := u.rep.GetAll(page, pageSize)

	return users, nil
}
