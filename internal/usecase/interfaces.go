package usecase

import (
	"ssr/internal/dto"
	"ssr/internal/entity"
)

type (
	IAuthUseCase interface {
		Login(email, password string) (*dto.LoginResponseDTO, error)
	}
	IAuthRepo interface {
		Get(email string) (*entity.Auth, error)
	}
	IUserUseCase interface {
		Me(email string) (*dto.UserResponseDTO, error)
	}
	IUserRepo interface {
		Get(email string) (*entity.User, error)
	}
)
