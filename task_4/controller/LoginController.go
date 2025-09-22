package controller

import (
	"net/http"
	"task4/core"
	"task4/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	BaseController
}

func (l LoginController) Login(ctx *gin.Context) {

	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser model.User
	if err := core.DB.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	cla := model.MyClaims{
		storedUser.Model.ID,
		storedUser.Username,
		jwt.StandardClaims{
			Id:        string(storedUser.Model.ID),
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)

	tokenString, err := token.SignedString([]byte("6235135"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	//ctx.SetCookie("token", tokenString)
	ctx.SetCookie("token", tokenString, int(time.Now().Add(time.Hour*24).Unix()), "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString, "data": storedUser})
}
