package helpers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	apimodels "songguru_bot/modules/api/models"
)

func ParseAccessToken(accessToken string, secret string) *apimodels.UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &apimodels.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return parsedAccessToken.Claims.(*apimodels.UserClaims)
}

func NewAccessToken(claims apimodels.UserClaims, secret string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(secret))
}

func GetUserClaims(c *gin.Context, secret string) (*apimodels.UserClaims, error) {
	jwt, err := c.Cookie("jwt")
	if err != nil {
		return nil, err
	}

	if jwt == "" {
		return nil, errors.New("jwt token was empty")
	}

	userClaims := ParseAccessToken(jwt, secret)
	if userClaims.StandardClaims.Valid() != nil {
		c.Status(http.StatusUnauthorized)
		return nil, errors.New("invalid payload")
	}
	return userClaims, nil
}
