package middleware

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"task4/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var Secret = []byte("6235135")

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHandler := c.Request.Header.Get("authorization")
		fmt.Println(authHandler)
		if authHandler == "" {
			c.JSON(200, gin.H{"code": 2003, "msg": "请求头部auth为空"})
			c.Abort()
			return
		}

		// 前两部门可以直接解析出来
		jwt := strings.Split(authHandler, ".")
		cnt := 0
		for _, val := range jwt {
			cnt++
			if cnt == 3 {
				break
			}
			msg, _ := base64.StdEncoding.DecodeString(val)
			fmt.Println("val ->", string(msg))
		}

		//调用下方自己实现的token解析函数,并且在判断token是否过期
		mc, err := ParseToken(authHandler)
		if err != nil {
			fmt.Println("err = ", err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		fmt.Println(mc)
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.UserName)
		c.Set("userid", mc.UserId)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func ParseToken(tokenString string) (*model.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
