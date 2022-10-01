package service

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"ssr/internal/dto"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
	"time"
)

type auth struct {
	*Base
	repo       AuthRepo
	tokenExp   time.Duration
	signingKey []byte
}

func NewAuth(r AuthRepo, l logger.Interface, tokenExpMinutes int, signingKey []byte) *auth {
	return &auth{
		Base:       NewBase(l),
		repo:       r,
		tokenExp:   time.Duration(tokenExpMinutes) * time.Minute,
		signingKey: signingKey,
	}
}

func (service *auth) Login(email, password string) (*dto.LoginResponse, error) {
	dbData, err := service.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(password)); err != nil {
		service.l.Error(err)
		return nil, err
	}

	tokenClaims := misc.NewAppJWTClaims(service.tokenExp, dbData.Email, string(dbData.Role))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	tokenStr, err := token.SignedString(service.signingKey)
	if err != nil {
		service.l.Error(err)
		return nil, err
	}

	return &dto.LoginResponse{
		Token: tokenStr,
		Email: dbData.Email,
		Role:  string(dbData.Role),
	}, nil
}
