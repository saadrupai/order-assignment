package repository

import (
	"github.com/saadrupai/order-assignment/app/entity"
	"gorm.io/gorm"
)

type IOrderRepo interface {
	Create(newOrder entity.Order) error
	List()
	Cancel()
}

type orderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) IOrderRepo {
	return &orderRepo{
		DB: db,
	}
}
func (odrRepo *orderRepo) Create(newOrder entity.Order) error {
	if dbErr := odrRepo.DB.Create(newOrder).Error; dbErr != nil {
		return dbErr
	}
	return nil
}
func (odrRepo *orderRepo) List() {

}

func (odrRepo *orderRepo) Cancel() {

}
