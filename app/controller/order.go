package controller

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/saadrupai/order-assignment/app/models"
	"github.com/saadrupai/order-assignment/app/service"
	"go.uber.org/zap"
	"net/http"
)

type orderController struct {
	orderSvc service.IOrderService
	logger   *zap.Logger
}

func NewOrderController(orderSvc service.IOrderService, logger *zap.Logger) *orderController {
	return &orderController{
		orderSvc: orderSvc,
		logger:   logger,
	}
}

func (odrCtr *orderController) Create(ctx *gin.Context) {
	var orderReq models.OrderReqBody
	if err := ctx.ShouldBindJSON(&orderReq); err != nil {
		odrCtr.logger.Error("Invalid order request", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: "Invalid request body",
			Type:    "error",
			Code:    http.StatusBadRequest,
			Errors:  map[string]interface{}{"request": []string{err.Error()}},
		})
		return
	}

	if err := orderReq.Validate(); err != nil {
		odrCtr.logger.Warn("Order request body validation failed", zap.Error(err))

		validationErrors := make(map[string][]string)
		if fieldErrors, ok := err.(validation.Errors); ok {
			for field, fieldErr := range fieldErrors {
				validationErrors[field] = append(validationErrors[field], fieldErr.Error())
			}
		} else {
			validationErrors["general"] = []string{err.Error()}
		}

		ctx.JSON(http.StatusUnprocessableEntity, models.Response{
			Message: "Please fix the given errors",
			Type:    "error",
			Code:    http.StatusUnprocessableEntity,
			Errors:  validationErrors,
		})
		return
	}

	createdOrder, createErr := odrCtr.orderSvc.Create(orderReq)
	if createErr != nil {
		odrCtr.logger.Error("Failed to create order", zap.Error(createErr))
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "Failed to create order",
			Type:    "error",
			Code:    http.StatusInternalServerError,
			Errors:  map[string]interface{}{"order": []string{createErr.Error()}},
		})
		return

	}

	ctx.JSON(http.StatusOK, models.Response{
		Message: "Order Created Successfully",
		Type:    "success",
		Code:    http.StatusOK,
		Data:    createdOrder,
	})
}

func (odrCtr *orderController) List(ctx *gin.Context) {
	var queryParams models.QueryParam

	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		odrCtr.logger.Error("Failed to bind query parameters", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: "Invalid query parameters",
			Type:    "error",
			Code:    http.StatusBadRequest,
			Errors:  map[string]interface{}{"query": []string{"Invalid query parameters"}},
		})
		return
	}

	orders, err := odrCtr.orderSvc.List(queryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: "Failed to list orders",
			Type:    "error",
			Code:    http.StatusInternalServerError,
			Errors:  map[string]interface{}{"order": []string{err.Error()}},
		})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (odrCtr *orderController) Cancel(ctx *gin.Context) {
	conIDStr := ctx.Param("consignment-id")

	err := odrCtr.orderSvc.Cancel(conIDStr)
	if err != nil {
		if _, ok := err.(*models.NotFoundError); ok {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: "Please contact CX to cancel order",
				Type:    "error",
				Code:    http.StatusBadRequest,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: "Failed to cancel order",
				Type:    "error",
				Code:    http.StatusInternalServerError,
				Errors:  map[string]interface{}{"order": []string{err.Error()}},
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Message: "Order Cancelled Successfully",
		Type:    "success",
		Code:    http.StatusOK,
	})
}
