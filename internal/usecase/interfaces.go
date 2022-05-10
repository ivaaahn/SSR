package usecase

import (
	"ssr/internal/dto"
	"ssr/internal/entity"
)

type (
	IAuthRepo interface {
		Get(email string) (*entity.Auth, error)
	}
	IAuthUC interface {
		Login(email, password string) (*dto.LoginResponseDTO, error)
	}

	IProfileRepo interface {
		GetStudentProfile(email string) (*entity.StudentProfile, error)
		GetSupervisorProfile(email string) (*entity.SupervisorProfile, error)
	}
	IProfileUC interface {
		GetStudentProfile(email string) (*dto.StudentProfileDTO, error)
		GetSupervisorProfile(email string) (*dto.SupervisorProfileDTO, error)
	}

	IBidRepo interface {
		GetBidsByStudentID(studentID int) ([]*entity.StudentBid, error)
		GetBidsBySupervisorID(studentID int) ([]*entity.SupervisorBid, error)
		CreateBid(studentID, supervisorID, workID int) (int, error)
	}
	IStudentBidUC interface {
		GetStudentBids(studentID int) (*dto.StudentBidsDTO, error)
		ApplyBid(data *dto.StudentApplyBidDTO) (*dto.StudentApplyBidResponseDTO, error)
	}
	ISupervisorBidUC interface {
		GetSupervisorBids(supervisorID int) (*dto.SupervisorBidsDTO, error)
	}
)
