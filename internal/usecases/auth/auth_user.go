package usecases

import (
	"github.com/ymoutella/king-poker-bk/internal/database"
	"github.com/ymoutella/king-poker-bk/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Execute(args ...any) (any, error)
}

type authUseCase struct {
	rep repository.UserRepository
}

func FactoryAuthUseCase() *authUseCase {
	db, err := database.PostgresDB()

	if err != nil {
		panic("Error connecting in db")
	}
	rep := repository.NewUserRepository(db)
	return &authUseCase{rep: rep}
}

func (us *authUseCase) Execute(email string, password string) (string, error) {

	user, err := us.rep.GetByEmail(email)

	if err != nil {
		return "", err
	}

	hashSuccess := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if hashSuccess != nil {
		return "", hashSuccess
	}

	token, err := GenerateToken(user.Email, "PLAYER")

	if err != nil {
		return "", err
	}

	return token, nil

}
