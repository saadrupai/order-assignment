package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/saadrupai/order-assignment/app/models"
	"github.com/saadrupai/order-assignment/app/service"
	"go.uber.org/zap"
	"net/http"
)

type userController struct {
	userService service.IUserService
	logger      *zap.Logger
}

func NewUserController(userService service.IUserService, logger *zap.Logger) *userController {
	return &userController{
		userService: userService,
		logger:      logger,
	}
}

func (usrCtr *userController) Login(ctx *gin.Context) {
	var userReq models.UserReq
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		usrCtr.logger.Error("Invalid user request", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: "Invalid request body",
			Type:    "error",
			Code:    http.StatusBadRequest,
			Errors:  map[string]interface{}{"request": []string{err.Error()}},
		})
		return
	}

	loginResp, loginErr := usrCtr.userService.Login(userReq)
	if loginErr != nil {
		if _, ok := loginErr.(*models.NotFoundError); ok {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Message: "The user credentials were incorrect.",
				Type:    "error",
				Code:    http.StatusBadRequest,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Message: "Failed to login user",
				Type:    "error",
				Code:    http.StatusInternalServerError,
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, loginResp)
}
