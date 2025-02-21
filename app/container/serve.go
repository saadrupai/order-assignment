package container

import (
	"github.com/gin-gonic/gin"
	"github.com/saadrupai/order-assignment/app/config"
	"github.com/saadrupai/order-assignment/app/controller"
	"github.com/saadrupai/order-assignment/app/repository"
	"github.com/saadrupai/order-assignment/app/service"
	"go.uber.org/zap"
	"log"
)

func Serve(router *gin.Engine) {

	apiVersion := router.Group("/api/v1")

	logger, loggerErr := zap.NewDevelopment()
	if loggerErr != nil {
		log.Fatal("failed to initialize logger")
	}

	db := config.GetDb()

	orderRepo := repository.NewOrderRepo(db)
	orderSvc := service.NewOrderSvc(orderRepo, logger)
	orderCtr := controller.NewOrderController(orderSvc, logger)

	apiVersion.POST("/orders", orderCtr.Create)
	apiVersion.GET("/orders/all", orderCtr.List())
	apiVersion.PUT("/orders/:consignment-id/cancel", orderCtr.Cancel())

	router.Run(":" + config.LocalConfig.Port)
}
