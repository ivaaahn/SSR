package usecase

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"ssr/internal/dto"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
	"time"
)

type AuthUseCase struct {
	*BaseUC
	repo       IRepoAuth
	tokenExp   time.Duration
	signingKey []byte
}

func NewAuthUC(r IRepoAuth, l logger.Interface, tokenExpMinutes int, signingKey []byte) *AuthUseCase {
	return &AuthUseCase{
		BaseUC:     NewUC(l),
		repo:       r,
		tokenExp:   time.Duration(tokenExpMinutes) * time.Minute,
		signingKey: signingKey,
	}
}

func (uc *AuthUseCase) Login(email, password string) (*dto.LoginResponse, error) {
	dbData, err := uc.repo.GetUserInfo(email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(password)); err != nil {
		uc.l.Error(err)
		return nil, err
	}

	tokenClaims := misc.NewAppJWTClaims(uc.tokenExp, dbData.Email, string(dbData.Role))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	tokenStr, err := token.SignedString(uc.signingKey)
	if err != nil {
		uc.l.Error(err)
		return nil, err
	}

	return &dto.LoginResponse{
		Token: tokenStr,
		Email: dbData.Email,
		Role:  string(dbData.Role),
	}, nil
}
