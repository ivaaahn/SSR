package service

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"ssr/internal/dto"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
	"time"
)

type Auth struct {
	*Base
	repo       UserRepo
	tokenExp   time.Duration
	signingKey []byte
}

func NewAuth(r UserRepo, l logger.Interface, tokenExpMinutes int, signingKey []byte) *Auth {
	return &Auth{
		Base:       NewBase(l),
		repo:       r,
		tokenExp:   time.Duration(tokenExpMinutes) * time.Minute,
		signingKey: signingKey,
	}
}

func (service *Auth) Login(email, password string) (*dto.LoginResponse, error) {
	dbData, err := service.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(password)); err != nil {
		service.l.Error(err)
		return nil, err
	}

	tokenClaims := misc.NewAppJWTClaims(service.tokenExp, dbData.UserID, string(dbData.Role))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	tokenStr, err := token.SignedString(service.signingKey)
	if err != nil {
		service.l.Error(err)
		return nil, err
	}

	return &dto.LoginResponse{
		Token:     tokenStr,
		TokenType: "Bearer",
		UserID:    dbData.UserID,
		Role:      string(dbData.Role),
	}, nil
}
