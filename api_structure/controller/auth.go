package controller

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(c *gin.Context, jwtKey []byte) (jwt.MapClaims, bool) {
	ck, err := c.Request.Cookie("token")
	fmt.Println(ck, "cookie")

	if err != nil {
		fmt.Print(err)
		return nil, false
	}

	tokenString := ck.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpacted signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}
	return nil, false
}
