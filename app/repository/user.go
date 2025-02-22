package repository

import (
	"errors"
	"github.com/saadrupai/order-assignment/app/entity"
	"github.com/saadrupai/order-assignment/app/models"
	"gorm.io/gorm"
)

type IUserRepo interface {
	GetUser(email string, password string) (entity.User, error)
}

type userRepo struct {
	DB *gorm.DB
}

func NewOUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{
		DB: db,
	}
}

func (usrRepo *userRepo) GetUser(email string, password string) (entity.User, error) {
	var user entity.User

	result := usrRepo.DB.Model(entity.User{}).
		Where("email = ? AND password = ?", email, password).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, &models.NotFoundError{Message: "Incorrect credentials"}
		}

		return entity.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.User{}, &models.NotFoundError{Message: "No user found with the given ccredentials"}
	}
	return user, nil
}
