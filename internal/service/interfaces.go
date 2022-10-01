package service

import (
	"ssr/internal/entity"
)

type (
	AuthRepo interface {
		GetUserByEmail(email string) (*entity.Auth, error)
	}
	ProfileRepo interface {
		GetStProfile(email string) (*entity.StProfile, error)
		GetSvProfile(email string) (*entity.SvProfile, error)
	}
	RelationRepo interface {
		Create(studentID, supervisorID, workID int) (int, error)
		GetStudentBids(studentID int) ([]*entity.StRelation, error)
		GetSupervisorBids(studentID int) ([]*entity.SvRelation, error)
		GetStudentRelations(studentID int) ([]*entity.StRelation, error)
		GetStudentRelation(studentID, ssrID int) (*entity.StRelation, error)
		UpdateStatus(id int, newStatus entity.StatusSSR) (int, error)
	}
	WorkRepo interface {
		GetWorksByStudentID(studentID int) ([]*entity.Work, error)
		GetWorksBySupervisorID(supervisorID int) ([]*entity.SvWork, error)
		GetSupervisorsByWorkID(workID int) ([]*entity.WorkSv, error)
	}
	FeedbackRepo interface {
		Create(studentID, supervisorID, workID int, content string) (int, error)
		GetBySupervisorID(supervisorID int) ([]*entity.Feedback, error)
	}
)