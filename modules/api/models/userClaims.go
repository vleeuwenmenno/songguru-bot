package apimodels

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/oauth2"
)

type UserClaims struct {
	Token       oauth2.Token
	DiscordUser DiscordUser
	jwt.StandardClaims
}
