package service

import (
	"github.com/google/uuid"
	"github.com/saadrupai/order-assignment/app/consts"
	"github.com/saadrupai/order-assignment/app/entity"
	"github.com/saadrupai/order-assignment/app/models"
	"github.com/saadrupai/order-assignment/app/repository"
	"go.uber.org/zap"
	"net/http"
)

type IOrderService interface {
	Create(orderReq models.OrderReqBody) (models.OrderCreateResponse, error)
	List(queryParam models.QueryParam) (models.OrderListResponse, error)
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
		ConsignmentID:      uuid.New(),
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
		OrderType:          consts.OrderTypeDelivery,
		CODFee:             consts.CODFee,
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
func (odrSvc *orderService) List(queryParam models.QueryParam) (models.OrderListResponse, error) {
	orders, count, err := odrSvc.orderRepo.List(queryParam)
	if err != nil {
		odrSvc.logger.Error("failed to list orders", zap.Error(err))
		return models.OrderListResponse{}, err
	}
	var orderResps []models.OrderRespData
	for _, order := range orders {
		orderResp := models.OrderRespData{
			OrderConsignmentID: order.ConsignmentID,
			OrderCreatedAt:     order.CreatedAt,
			OrderDescription:   order.ItemDescription,
			MerchantOrderID:    order.MerchantOrderID,
			RecipientName:      order.RecipientName,
			RecipientAddress:   order.RecipientAddress,
			RecipientPhone:     order.RecipientPhone,
			OrderAmount:        order.AmountToCollect,
			TotalFee:           order.DeliveryFee,
			Instruction:        order.SpecialInstruction,
			OrderTypeID:        order.ItemType,
			CODFee:             order.CODFee,
			PromoDiscount:      order.PromoDiscount,
			Discount:           order.Discount,
			DeliveryFee:        order.DeliveryFee,
			OrderStatus:        order.OrderStatus,
			OrderType:          consts.OrderTypeDelivery,
		}

		if order.ItemType == consts.ItemTypeParcel {
			orderResp.ItemType = consts.ItemTypeStatusParcel
		}

		orderResps = append(orderResps, orderResp)
	}

	type orderResponse struct {
		Data []models.OrderRespData `json:"data"`
	}

	pagination := models.PaginationInfo{
		Total:       count,
		CurrentPage: int64(queryParam.Page),
		PerPage:     int64(queryParam.Limit),
		TotalInPage: int64(len(orders)),
		LastPage:    (count + int64(queryParam.Limit) - 1) / int64(queryParam.Limit),
	}

	respose := models.OrderListResponse{
		Response: models.Response{
			Message: "Orders successfully fetched",
			Type:    "success",
			Code:    http.StatusOK,
		},
		Data: orderResponse{
			Data: orderResps,
		},
		PaginationInfo: pagination,
	}
	return respose, nil
}
func (odrSvc *orderService) Cancel() {

}
