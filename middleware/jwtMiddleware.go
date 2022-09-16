package middleware

import (
	"fmt"
	"main/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix() //token will be expire after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	keys, err := token.SignedString([]byte(constants.SecretKey()))
	if err != nil {
		fmt.Println(err)
	}
	return keys, err
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(int)
		return userId
	}
	return 0
}
