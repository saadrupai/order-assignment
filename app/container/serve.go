package container

import (
	"github.com/gin-gonic/gin"
	"github.com/saadrupai/order-assignment/app/config"
	"github.com/saadrupai/order-assignment/app/controller"
	"github.com/saadrupai/order-assignment/app/middleware"
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

	userRepo := repository.NewOUserRepo(db)
	orderRepo := repository.NewOrderRepo(db)

	userSvc := service.NewUserSvc(userRepo, logger)
	orderSvc := service.NewOrderSvc(orderRepo, logger)

	orderCtr := controller.NewOrderController(orderSvc, logger)
	userCtr := controller.NewUserController(userSvc, logger)

	apiVersion.POST("/login", userCtr.Login)

	apiVersionAuth := apiVersion.Use(middleware.AuthMiddleware())

	apiVersionAuth.POST("/orders", orderCtr.Create)
	apiVersionAuth.GET("/orders/all", orderCtr.List)
	apiVersionAuth.PUT("/orders/:consignment-id/cancel", orderCtr.Cancel)

	router.Run(":" + config.LocalConfig.Port)
}
