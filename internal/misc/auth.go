package misc

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ExtractEmailFromContext(ctx echo.Context) string {
	token := ctx.Get("userEmail").(*jwt.Token)
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Subject
}
