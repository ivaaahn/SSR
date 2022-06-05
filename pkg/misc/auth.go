package misc

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
)

type AppJWTClaims struct {
	jwt.StandardClaims
	Role string
}

func NewAppJWTClaims(exp time.Duration, sub, role string) *AppJWTClaims {
	return &AppJWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exp).Unix(),
			Subject:   sub,
		},
		Role: role,
	}
}

func ExtractCtx(ctx echo.Context) (string, string) {
	token := ctx.Get("ctx").(*jwt.Token)
	claims := token.Claims.(*AppJWTClaims)
	return claims.Subject, claims.Role
}
