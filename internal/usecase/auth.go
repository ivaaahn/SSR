package usecase

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"ssr/internal/dto"
	"time"
)

type AuthUseCase struct {
	repo       IAuthRepo
	tokenExp   time.Duration
	signingKey []byte
}

func NewAuthUC(r IAuthRepo, tokenExpMinutes int, signingKey []byte) *AuthUseCase {
	return &AuthUseCase{
		repo:       r,
		tokenExp:   time.Duration(tokenExpMinutes) * time.Minute,
		signingKey: signingKey,
	}
}

func (uc *AuthUseCase) Login(email, password string) (*dto.LoginResponseDTO, error) {
	dbData, err := uc.repo.Get(email)
	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - Login - repo.Get: %w", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(password)); err != nil {
		return nil, echo.ErrUnauthorized
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(uc.tokenExp).Unix(),
		Subject:   dbData.Email,
	})

	tokenStr, err := token.SignedString(uc.signingKey)

	if err != nil {
		return nil, fmt.Errorf("AuthUseCase - token.SignedString: %w", err)
	}

	return &dto.LoginResponseDTO{Token: tokenStr}, nil
}
