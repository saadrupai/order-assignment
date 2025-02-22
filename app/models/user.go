package models

import "github.com/dgrijalva/jwt-go"

type UserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	UserEmail string `json:"email"`
	jwt.StandardClaims
}
