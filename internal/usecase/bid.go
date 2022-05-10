package usecase

import (
	"fmt"
	"ssr/internal/dto"
	"ssr/pkg/misc"
)

type BidUseCase struct {
	repo IBidRepo
}

func NewBidUC(r IBidRepo) *BidUseCase {
	return &BidUseCase{
		repo: r,
	}
}

func (uc *BidUseCase) GetStudentBids(studentID int) (*dto.StudentBidsDTO, error) {
	dbData, err := uc.repo.GetBidsByStudentID(studentID)
	if err != nil {
		return nil, fmt.Errorf("BidUseCase - GetStudentBids - repo.GetBidsByStudentID: %w", err)
	}

	var resp []*dto.StudentBidDTO

	for _, db := range dbData {
		resp = append(resp, &dto.StudentBidDTO{
			BidID:     db.BidID,
			Status:    db.Status,
			CreatedAt: db.CreatedAt,
			Supervisor: dto.SupervisorProfileDTO{
				SupervisorID: db.SupervisorID,
				Email:        db.Email,
				FirstName:    db.FirstName,
				LastName:     db.LastName,
				About:        db.About,
				Birthdate: misc.BirthDate{
					Time: db.Birthdate,
				},
				AvatarUrl:  misc.NullString(db.Avatar),
				Department: db.SupervisorProfile.DepartmentID,
			},
			Work: dto.WorkDTO{
				WorkID:      db.WorkID,
				Name:        db.WorkKind.Name,
				Description: db.Work.Description,
				Semester:    db.Work.Semester,
				Subject: dto.SubjectDTO{
					SubjectID:  db.SubjectID,
					Name:       db.Subject.Name,
					Department: db.Subject.DepartmentID,
				},
			},
		})
	}

	return &dto.StudentBidsDTO{Bids: resp}, nil
}

func (uc *BidUseCase) GetSupervisorBids(supervisorID int) (*dto.SupervisorBidsDTO, error) {
	dbData, err := uc.repo.GetBidsBySupervisorID(supervisorID)
	if err != nil {
		return nil, fmt.Errorf("BidUseCase - GetStudentBids - repo.GetBidsByStudentID: %w", err)
	}

	var resp []*dto.SupervisorBidDTO

	for _, db := range dbData {
		resp = append(resp, &dto.SupervisorBidDTO{
			BidID:     db.BidID,
			Status:    db.Status,
			CreatedAt: db.CreatedAt,
			Student: dto.StudentProfileDTO{
				StudentID:  db.StudentID,
				Email:      db.Email,
				FirstName:  db.FirstName,
				LastName:   db.LastName,
				Year:       db.Year,
				AvatarUrl:  misc.NullString(db.Avatar),
				Department: db.StudentProfile.DepartmentID,
			},
			Work: dto.WorkDTO{
				WorkID:      db.WorkID,
				Name:        db.WorkKind.Name,
				Description: db.Work.Description,
				Semester:    db.Work.Semester,
				Subject: dto.SubjectDTO{
					SubjectID:  db.SubjectID,
					Name:       db.Subject.Name,
					Department: db.Subject.DepartmentID,
				},
			},
		})
	}

	return &dto.SupervisorBidsDTO{Bids: resp}, nil
}

func (uc *BidUseCase) ApplyBid(data *dto.StudentApplyBidDTO) (*dto.StudentApplyBidResponseDTO, error) {
	bidID, err := uc.repo.CreateBid(data.StudentID, data.SupervisorID, data.WorkID)
	if err != nil {
		return nil, fmt.Errorf("BidUseCase - ApplyBid - repo.CreateBid %w", err)
	}

	return &dto.StudentApplyBidResponseDTO{BidID: bidID}, nil
}
