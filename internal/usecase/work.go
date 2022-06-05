package usecase

import (
	"ssr/internal/dto"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
)

type WorkUseCase struct {
	*BaseUC
	repo IWorkRepo
}

func NewWorkUC(r IWorkRepo, l logger.Interface) *WorkUseCase {
	return &WorkUseCase{
		BaseUC: NewUC(l),
		repo:   r,
	}
}

func (uc *WorkUseCase) GetStudentWorks(studentID int) (*dto.StudentWorkPlenty, error) {
	dbData, err := uc.repo.GetWorksByStudentID(studentID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.StudentWork

	for _, db := range dbData {
		resp = append(resp, &dto.StudentWork{
			WorkID:      db.WorkID,
			Kind:        db.WorkKindName,
			Description: db.Description,
			Subject:     db.SubjectName,
		})
	}

	return &dto.StudentWorkPlenty{
		StudentID: studentID,
		Works:     resp,
	}, nil
}

func (uc *WorkUseCase) GetSupervisorWorks(supervisorID int) (*dto.SupervisorWorkPlenty, error) {
	dbData, err := uc.repo.GetWorksBySupervisorID(supervisorID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.SupervisorWork

	for _, db := range dbData {
		resp = append(resp, &dto.SupervisorWork{
			WorkID:      db.WorkID,
			Kind:        db.WorkKindName,
			Description: db.Description,
			Subject:     db.SubjectName,
			Head:        db.Head,
		})
	}

	return &dto.SupervisorWorkPlenty{
		SupervisorID: supervisorID,
		Works:        resp,
	}, nil
}

func (uc *WorkUseCase) GetWorkSupervisors(workID int) (*dto.WorkSupervisorPlenty, error) {
	dbData, err := uc.repo.GetSupervisorsByWorkID(workID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.WorkSupervisor

	for _, db := range dbData {
		resp = append(resp, &dto.WorkSupervisor{
			SupervisorProfile: dto.SupervisorProfile{
				SupervisorID: db.SupervisorID,
				Email:        db.Email,
				FirstName:    db.FirstName,
				LastName:     db.LastName,
				About:        db.About,
				Birthdate:    misc.Date{Time: db.Birthdate},
				AvatarUrl:    misc.NullString(db.Avatar),
				Department:   db.DepartmentID,
			},
			Head: db.Head,
			Full: db.Full,
		})
	}

	return &dto.WorkSupervisorPlenty{
		WorkID:      workID,
		Supervisors: resp,
	}, nil
}
