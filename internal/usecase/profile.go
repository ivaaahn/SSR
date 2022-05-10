package usecase

import (
	"fmt"
	"ssr/internal/dto"
	"ssr/pkg/misc"
)

type ProfileUseCase struct {
	repo IProfileRepo
}

func NewProfileUC(r IProfileRepo) *ProfileUseCase {
	return &ProfileUseCase{
		repo: r,
	}
}

func (uc *ProfileUseCase) GetStudentProfile(email string) (*dto.StudentProfileDTO, error) {
	dbData, err := uc.repo.GetStudentProfile(email)
	if err != nil {
		return nil, fmt.Errorf("ProfileUseCase - GetProfile - repo.GetStudentProfile: %w", err)
	}

	return &dto.StudentProfileDTO{
		StudentID:   dbData.StudentID,
		Email:       dbData.Email,
		FirstName:   dbData.FirstName,
		LastName:    dbData.LastName,
		AvatarUrl:   misc.NullString(dbData.Avatar),
		Year:        dbData.Year,
		StudentCard: dbData.StudentCard,
		Department:  dbData.DepartmentID,
	}, nil
}

func (uc *ProfileUseCase) GetSupervisorProfile(email string) (*dto.SupervisorProfileDTO, error) {
	dbData, err := uc.repo.GetSupervisorProfile(email)
	if err != nil {
		return nil, fmt.Errorf("SupervisorUseCase - GetProfile - repo.GetSupervisorProfile: %w", err)
	}

	return &dto.SupervisorProfileDTO{
		SupervisorID: dbData.SupervisorID,
		Email:        dbData.Email,
		FirstName:    dbData.FirstName,
		LastName:     dbData.LastName,
		AvatarUrl:    misc.NullString(dbData.Avatar),
		About:        dbData.About,
		Birthdate:    misc.BirthDate{Time: dbData.Birthdate},
		Department:   dbData.DepartmentID,
	}, nil
}
