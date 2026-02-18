package repository

import (
	"github.com/ymoutella/king-poker-bk/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(page int, pageSize int) []domain.User
	GetByID(id uint) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func (repository userRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	result := repository.db.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository userRepository) GetAll(page int, pageSize int) []domain.User {
	var users []domain.User
	repository.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users)
	return users
}

func (repository userRepository) Create(user *domain.User) (*domain.User, error) {
	createdUser := *user
	result := repository.db.Create(&createdUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &createdUser, nil

}

func (repository userRepository) Update(user *domain.User) (*domain.User, error) {
	updatedUser := *user
	result := repository.db.Save(&updatedUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &updatedUser, nil
}

func (repository userRepository) Delete(id uint) error {
	result := repository.db.Delete(&domain.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
