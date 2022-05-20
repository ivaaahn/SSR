package usecase

import (
	"fmt"
	"ssr/internal/dto"
	"ssr/internal/entity"
	"ssr/pkg/misc"
)

type BidUseCase struct {
	repo IRelRepo
}

func NewBidUC(r IRelRepo) *BidUseCase {
	return &BidUseCase{
		repo: r,
	}
}

func (uc *BidUseCase) GetStudentBids(studentID int) (*dto.StudentBids, error) {
	dbData, err := uc.repo.GetStudentViewBidPlenty(studentID)
	if err != nil {
		return nil, fmt.Errorf("BidUseCase - GetStudentBids - repo.GetStudentViewBidPlenty: %w", err)
	}

	var resp []*dto.StudentBid

	for _, db := range dbData {
		resp = append(resp, &dto.StudentBid{
			BidID:     db.BidID,
			Status:    db.Status,
			CreatedAt: db.CreatedAt,
			Supervisor: dto.SupervisorProfile{
				SupervisorID: db.SupervisorID,
				Email:        db.Email,
				FirstName:    db.FirstName,
				LastName:     db.LastName,
				About:        db.About,
				Birthdate: misc.Date{
					Time: db.Birthdate,
				},
				AvatarUrl:  misc.NullString(db.Avatar),
				Department: db.SupervisorProfile.DepartmentID,
			},
			Work: dto.Work{
				WorkID:      db.WorkID,
				Name:        db.WorkKind.WorkKindName,
				Description: db.Work.Description,
				Semester:    db.Work.Semester,
				Subject: dto.SubjectResp{
					SubjectID:  db.SubjectID,
					Name:       db.Subject.SubjectName,
					Department: db.Subject.DepartmentID,
				},
			},
		})
	}

	return &dto.StudentBids{Bids: resp}, nil
}

func (uc *BidUseCase) GetSupervisorBids(supervisorID int) (*dto.SupervisorBids, error) {
	dbData, err := uc.repo.GetSupervisorViewBidPlenty(supervisorID)
	if err != nil {
		return nil, fmt.Errorf("BidUseCase - GetStudentBids - repo.GetStudentViewBidPlenty: %w", err)
	}

	var resp []*dto.SupervisorBid

	for _, db := range dbData {
		resp = append(resp, &dto.SupervisorBid{
			BidID:     db.BidID,
			Status:    db.Status,
			CreatedAt: db.CreatedAt,
			Student: dto.StudentProfile{
				StudentID:  db.StudentID,
				Email:      db.Email,
				FirstName:  db.FirstName,
				LastName:   db.LastName,
				Year:       db.Year,
				AvatarUrl:  misc.NullString(db.Avatar),
				Department: db.StudentProfile.DepartmentID,
			},
			Work: dto.Work{
				WorkID:      db.WorkID,
				Name:        db.WorkKind.WorkKindName,
				Description: db.Work.Description,
				Semester:    db.Work.Semester,
				Subject: dto.SubjectResp{
					SubjectID:  db.SubjectID,
					Name:       db.Subject.SubjectName,
					Department: db.Subject.DepartmentID,
				},
			},
		})
	}

	return &dto.SupervisorBids{Bids: resp}, nil
}

func (uc *BidUseCase) Apply(data *dto.ApplyBid) (*dto.ApplyBidResponse, error) {
	bidID, err := uc.repo.Create(data.StudentID, data.SupervisorID, data.WorkID)
	if err != nil {
		return nil, fmt.Errorf("BidUseCase - Apply - repo.Create %w", err)
	}

	return &dto.ApplyBidResponse{BidID: bidID}, nil
}

func (uc *BidUseCase) Resolve(data *dto.ResolveBid) error {
	var status entity.StatusSSR

	if data.Accept {
		status = "accepted"
	} else {
		status = "rejected"
	}

	_, err := uc.repo.UpdateStatus(data.BidID, status)
	if err != nil {
		return fmt.Errorf("BidUseCase - Resolve - repo.UpdateStatus %w", err)
	}

	return nil
}
