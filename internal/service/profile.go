package service

import (
	"ssr/internal/dto"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
)

type Profile struct {
	*Base
	repo ProfileRepo
}

func NewProfile(r ProfileRepo, l logger.Interface) *Profile {
	return &Profile{
		Base: NewBase(l),
		repo: r,
	}
}

func (uc *Profile) GetStudentProfile(email string) (*dto.StProfile, error) {
	dbData, err := uc.repo.GetStProfile(email)
	if err != nil {
		return nil, err
	}

	return &dto.StProfile{
		StudentID:   dbData.StudentID,
		Email:       dbData.Email,
		FirstName:   dbData.FirstName,
		LastName:    dbData.LastName,
		AvatarUrl:   misc.NullString(dbData.PhotoUrl),
		Year:        dbData.Year,
		StudentCard: dbData.StudentCard,
		Department:  dbData.DepartmentID,
	}, nil
}

func (uc *Profile) GetSupervisorProfile(email string) (*dto.SvProfile, error) {
	dbData, err := uc.repo.GetSvProfile(email)
	if err != nil {
		return nil, err
	}

	return &dto.SvProfile{
		SupervisorID: dbData.SupervisorID,
		Email:        dbData.Email,
		FirstName:    dbData.FirstName,
		LastName:     dbData.LastName,
		AvatarUrl:    misc.NullString(dbData.PhotoUrl),
		About:        dbData.About,
		Birthdate:    misc.Date{Time: dbData.Birthdate},
		Department:   dbData.DepartmentID,
	}, nil
}
