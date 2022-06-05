package usecase

import (
	"ssr/internal/dto"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
)

type ProfileUseCase struct {
	*BaseUC
	repo IProfileRepo
}

func NewProfileUC(r IProfileRepo, l logger.Interface) *ProfileUseCase {
	return &ProfileUseCase{
		BaseUC: NewUC(l),
		repo:   r,
	}
}

func (uc *ProfileUseCase) GetStudentProfile(email string) (*dto.StudentProfile, error) {
	dbData, err := uc.repo.GetStudentProfile(email)
	if err != nil {
		return nil, err
	}

	return &dto.StudentProfile{
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

func (uc *ProfileUseCase) GetSupervisorProfile(email string) (*dto.SupervisorProfile, error) {
	dbData, err := uc.repo.GetSupervisorProfile(email)
	if err != nil {
		return nil, err
	}

	return &dto.SupervisorProfile{
		SupervisorID: dbData.SupervisorID,
		Email:        dbData.Email,
		FirstName:    dbData.FirstName,
		LastName:     dbData.LastName,
		AvatarUrl:    misc.NullString(dbData.Avatar),
		About:        dbData.About,
		Birthdate:    misc.Date{Time: dbData.Birthdate},
		Department:   dbData.DepartmentID,
	}, nil
}
