package usecase

import (
	"ssr/internal/dto"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
)

type WorkUseCase struct {
	*BaseUC
	repoWork IRepoWork
	repoSsr  IRepoSSR
}

func NewWorkUC(rWork IRepoWork, rSsr IRepoSSR, l logger.Interface) *WorkUseCase {
	return &WorkUseCase{
		BaseUC:   NewUC(l),
		repoWork: rWork,
		repoSsr:  rSsr,
	}
}

func checkIfBegin(relations []*entity.StudentSsr, workID int) bool {
	for _, rel := range relations {
		if rel.Work.WorkID == workID {
			return true
		}
	}

	return false
}

func (uc *WorkUseCase) GetStudentWorks(studentID int) (*dto.StudentWorks, error) {
	dbData, err := uc.repoWork.GetWorksByStudentID(studentID)
	if err != nil {
		return nil, err
	}

	relations, err := uc.repoSsr.GetStudentRelations(studentID)
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
			IsStarted:   checkIfBegin(relations, db.WorkID),
		})
	}

	return &dto.StudentWorks{
		StudentID: studentID,
		Works:     resp,
	}, nil
}

func (uc *WorkUseCase) GetSupervisorWorks(supervisorID int) (*dto.SupervisorWorkPlenty, error) {
	dbData, err := uc.repoWork.GetWorksBySupervisorID(supervisorID)
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
	dbData, err := uc.repoWork.GetSupervisorsByWorkID(workID)
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
