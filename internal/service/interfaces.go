package service

import (
	"ssr/internal/entity"
)

type (
	UserRepo interface {
		GetUserByEmail(email string) (*entity.UserFull, error)
	}
	StudentRepo interface {
		GetStudent(userID int) (*entity.Student, error)
		GetStudentShort(userID int) (*entity.StudentShort, error)
	}
	SupervisorRepo interface {
		GetSupervisor(userID int) (*entity.Supervisor, error)
		GetSupervisorsByWorkID(workID int) ([]*entity.WorkSupervisor, error)
	}
	RelationRepo interface {
		Create(studentID, supervisorID, workID int) (int, error)
		GetRelationsByStudentID(studentID int) ([]*entity.RelationShort, error)
		GetRelationsBySupervisorID(supervisorID int) ([]*entity.RelationShort, error)
		Get(id int) (*entity.Relation, error)
		Update(id int, status entity.StatusSSR) (int, error)
		GetRelationStatus(studentID, workID int) (entity.StatusSSR, error)
	}
	WorkRepo interface {
		Get(workID int) (*entity.Work, error)
		GetStudentWorks(departmentID string, semester int) ([]*entity.Work, error)
		GetSupervisorWorks(supervisorID int) ([]*entity.SupervisorViewWork, error)
	}
	WaypointRepo interface {
		GetPlenty(workID int) ([]*entity.Waypoint, error)
	}
)
