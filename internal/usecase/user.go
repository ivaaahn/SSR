package usecase

import (
	"fmt"
	"ssr/internal/dto"
	"ssr/internal/misc"
)

type UserUseCase struct {
	repo IUserRepo
}

func NewUserUC(r IUserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) Me(email string) (*dto.UserResponseDTO, error) {
	dbData, err := uc.repo.Get(email)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - Me- repo.Get: %w", err)
	}

	return &dto.UserResponseDTO{
		Email:     dbData.Email,
		FirstName: dbData.FirstName,
		LastName:  dbData.LastName,
		AvatarUrl: misc.NullString(dbData.Avatar),
		Role:      dbData.Role,
	}, nil
}
