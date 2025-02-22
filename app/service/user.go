package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/saadrupai/order-assignment/app/config"
	"github.com/saadrupai/order-assignment/app/models"
	"github.com/saadrupai/order-assignment/app/repository"
	"go.uber.org/zap"
	"time"
)

type IUserService interface {
	Login(userReq models.UserReq) (models.UserLoginResp, error)
}

type userService struct {
	userRepo repository.IUserRepo
	logger   *zap.Logger
}

func NewUserSvc(userRepo repository.IUserRepo, logger *zap.Logger) IUserService {
	return &userService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (usrSvc *userService) Login(userReq models.UserReq) (models.UserLoginResp, error) {
	user, err := usrSvc.userRepo.GetUser(userReq.Email, userReq.Password)
	if err != nil {
		usrSvc.logger.Error("failed to get user", zap.Error(err))
		return models.UserLoginResp{}, err
	}

	expirationTime := time.Now().Add(2 * 24 * time.Hour)
	claims := &models.Claims{
		UserEmail: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	config := config.LoadConfig()
	jwtKey := []byte(config.SecretKey)
	tokenString, tokenErr := token.SignedString(jwtKey)
	if tokenErr != nil {
		usrSvc.logger.Error("Failed to generate token", zap.Error(tokenErr))
		return models.UserLoginResp{}, tokenErr
	}

	return models.UserLoginResp{
		TokenType:    "Bearer",
		ExpiresIn:    172800,
		AccessToken:  tokenString,
		RefreshToken: "refresh_token",
	}, nil
}
