package misc

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
)

type JWTClaimsSSR struct {
	jwt.StandardClaims
	Role string
}

func NewJWTClaimsSSR(exp time.Duration, sub, role string) *JWTClaimsSSR {
	return &JWTClaimsSSR{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exp).Unix(),
			Subject:   sub,
		},
		Role: role,
	}
}

func ExtractInfoFromContext(ctx echo.Context) (string, string) {
	token := ctx.Get("userEmail").(*jwt.Token)
	claims := token.Claims.(*JWTClaimsSSR)
	return claims.Subject, claims.Role
}
