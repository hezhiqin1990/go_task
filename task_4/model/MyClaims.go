package model

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	UserId   uint   `json:"userid"`
	UserName string `json:"username"`
	jwt.StandardClaims
}
