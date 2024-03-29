package middlewares

import (
	"MiniProject/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateAdminToken(adminId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["adminId"] = adminId
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix() // Token expires after 5 hours

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
