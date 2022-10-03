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

func (service *Profile) GetStudentProfile(userID int) (*dto.Student, error) {
	dbData, err := service.stRepo.GetStudent(userID)
	if err != nil {
		return nil, err
	}

	return &dto.Student{
		Email:       dbData.User.Email,
		FirstName:   dbData.User.FirstName,
		LastName:    dbData.User.LastName,
		PhotoUrl:    dbData.User.PhotoUrl,
		Year:        dbData.Year,
		StudentCard: dbData.StudentCard,
		Department:  dbData.DepartmentID,
	}, nil
}

func (service *Profile) GetSupervisorProfile(userID int) (*dto.Supervisor, error) {
	dbData, err := service.svRepo.GetFullSupervisor(userID)
	if err != nil {
		return nil, err
	}

	return &dto.Supervisor{
		Email:      dbData.User.Email,
		FirstName:  dbData.User.FirstName,
		LastName:   dbData.User.LastName,
		PhotoUrl:   dbData.User.PhotoUrl,
		About:      dbData.About,
		Birthdate:  misc.Date{Time: dbData.Birthdate},
		Department: dbData.DepartmentID,
	}, nil
}
