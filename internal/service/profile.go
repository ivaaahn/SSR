package service

import (
	"ssr/internal/dto"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
)

type Profile struct {
	*Base
	stRepo StudentRepo
	svRepo SupervisorRepo
}

func NewProfile(stRepo StudentRepo, svRepo SupervisorRepo, l logger.Interface) *Profile {
	return &Profile{
		Base:   NewBase(l),
		stRepo: stRepo,
		svRepo: svRepo,
	}
}

func (uc *Profile) GetStudentProfile(userID int) (*dto.StProfile, error) {
	dbData, err := uc.stRepo.GetFullStudent(userID)
	if err != nil {
		return nil, err
	}

	return &dto.StProfile{
		Email:       dbData.User.Email,
		FirstName:   dbData.User.FirstName,
		LastName:    dbData.User.LastName,
		PhotoUrl:    dbData.User.PhotoUrl,
		Year:        dbData.Year,
		StudentCard: dbData.StudentCard,
		Department:  dbData.DepartmentID,
	}, nil
}

func (uc *Profile) GetSupervisorProfile(userID int) (*dto.SvProfile, error) {
	dbData, err := uc.svRepo.GetFullSupervisor(userID)
	if err != nil {
		return nil, err
	}

	return &dto.SvProfile{
		Email:      dbData.User.Email,
		FirstName:  dbData.User.FirstName,
		LastName:   dbData.User.LastName,
		PhotoUrl:   dbData.User.PhotoUrl,
		About:      dbData.About,
		Birthdate:  misc.Date{Time: dbData.Birthdate},
		Department: dbData.DepartmentID,
	}, nil
}
