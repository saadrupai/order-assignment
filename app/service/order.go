package service

import (
	"github.com/saadrupai/order-assignment/app/consts"
	"github.com/saadrupai/order-assignment/app/entity"
	"github.com/saadrupai/order-assignment/app/models"
	"github.com/saadrupai/order-assignment/app/repository"
	"go.uber.org/zap"
)

type IOrderService interface {
	Create(orderReq models.OrderReqBody) (models.OrderCreateResponse, error)
	List()
	Cancel()
}

type orderService struct {
	orderRepo repository.IOrderRepo
	logger    *zap.Logger
}

func NewOrderSvc(orderRepo repository.IOrderRepo, logger *zap.Logger) IOrderService {
	return &orderService{
		orderRepo: orderRepo,
		logger:    logger,
	}
}

func (odrSvc *orderService) Create(orderReq models.OrderReqBody) (models.OrderCreateResponse, error) {
	newOrder := entity.Order{
		ConsignmentID:      2,
		StoreID:            orderReq.StoreID,
		MerchantOrderID:    orderReq.MerchantOrderID,
		RecipientName:      orderReq.RecipientName,
		RecipientAddress:   orderReq.RecipientAddress,
		RecipientPhone:     orderReq.RecipientPhone,
		RecipientArea:      orderReq.RecipientArea,
		RecipientCity:      orderReq.RecipientCity,
		RecipientZone:      orderReq.RecipientZone,
		DeliveryType:       orderReq.DeliveryType,
		ItemType:           orderReq.ItemType,
		SpecialInstruction: orderReq.SpecialInstruction,
		ItemQuantity:       orderReq.ItemQuantity,
		ItemWeight:         orderReq.ItemWeight,
		AmountToCollect:    orderReq.AmountToCollect,
		ItemDescription:    orderReq.ItemDescription,
		OrderStatus:        consts.OrderStatusPending,
		DeliveryFee:        consts.DeliveryFee,
	}
	createErr := odrSvc.orderRepo.Create(newOrder)
	if createErr != nil {
		odrSvc.logger.Error("failed to create order", zap.Error(createErr))
		return models.OrderCreateResponse{}, createErr
	}

	return models.OrderCreateResponse{
		ConsignmentID:   newOrder.ConsignmentID,
		MerchantOrderID: newOrder.MerchantOrderID,
		OrderStatus:     newOrder.OrderStatus,
		DeliveryFee:     newOrder.DeliveryFee,
	}, nil
}
func (odrSvc *orderService) List() {

}
func (odrSvc *orderService) Cancel() {

}
