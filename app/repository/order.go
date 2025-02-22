package repository

import (
	"github.com/saadrupai/order-assignment/app/entity"
	"github.com/saadrupai/order-assignment/app/models"
	"gorm.io/gorm"
)

type IOrderRepo interface {
	Create(newOrder entity.Order) error
	List(queryParam models.QueryParam) ([]entity.Order, int64, error)
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
func (odrRepo *orderRepo) List(queryParam models.QueryParam) ([]entity.Order, int64, error) {
	var orders []entity.Order
	var count int64

	baseQ := odrRepo.DB.Model(entity.Order{})

	baseQ = baseQ.Where("transfer_status = ?", true).Where("archive = ?", false)

	if queryParam.Limit != 0 && queryParam.Page > 0 {
		baseQ = baseQ.Limit(queryParam.Limit).Offset((queryParam.Page - 1) * queryParam.Limit)
	}

	baseQ = baseQ.Count(&count)

	if dbErr := baseQ.Find(&orders).Error; dbErr != nil {
		return nil, 0, dbErr
	}

	return orders, count, nil
}

func (odrRepo *orderRepo) Cancel() {

}
