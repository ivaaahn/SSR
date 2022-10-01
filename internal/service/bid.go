package service

import (
	"ssr/internal/dto"
	"ssr/internal/entity"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
)

type Bid struct {
	*Base
	repo RelationRepo
}

func NewBid(r RelationRepo, l logger.Interface) *Bid {
	return &Bid{
		Base: NewBase(l),
		repo: r,
	}
}

func (service *Bid) GetStudentBids(studentID int) (*dto.StBids, error) {
	dbData, err := service.repo.GetStudentBids(studentID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.StBid

	for _, db := range dbData {
		resp = append(resp, &dto.StBid{
			BidID:     db.BidID,
			Status:    db.Status,
			CreatedAt: db.CreatedAt,
			Supervisor: dto.SvProfile{
				Email:     db.Email,
				FirstName: db.FirstName,
				LastName:  db.LastName,
				About:     db.About,
				Birthdate: misc.Date{
					Time: db.Birthdate,
				},
				PhotoUrl:   db.PhotoUrl,
				Department: db.SvProfile.DepartmentID,
			},
			Work: dto.WorkResp{
				WorkID:      db.WorkID,
				Description: db.Work.Description,
				Semester:    db.Work.Semester,
				Kind: dto.WorkKindResp{
					ID:   db.WorkKind.WorkKindID,
					Name: db.WorkKind.Name,
				},
				Subject: dto.SubjectResp{
					SubjectID:  db.SubjectID,
					Name:       db.Subject.Name,
					Department: db.Subject.DepartmentID,
				},
			},
		})
	}

	return &dto.StBids{Bids: resp}, nil
}

func (service *Bid) GetSupervisorBids(supervisorID int) (*dto.SvBids, error) {
	dbData, err := service.repo.GetSupervisorBids(supervisorID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.SvBid

	for _, db := range dbData {
		resp = append(resp, &dto.SvBid{
			BidID:     db.BidID,
			Status:    db.Status,
			CreatedAt: db.CreatedAt,
			Student: dto.StProfile{
				Email:      db.Email,
				FirstName:  db.FirstName,
				LastName:   db.LastName,
				Year:       db.Year,
				PhotoUrl:   db.PhotoUrl,
				Department: db.StProfile.DepartmentID,
			},
			Work: dto.WorkResp{
				Description: db.Work.Description,
				Semester:    db.Work.Semester,
				Kind: dto.WorkKindResp{
					ID:   db.WorkKind.WorkKindID,
					Name: db.WorkKind.Name,
				},
				Subject: dto.SubjectResp{
					SubjectID:  db.SubjectID,
					Name:       db.Subject.Name,
					Department: db.Subject.DepartmentID,
				},
			},
		})
	}

	return &dto.SvBids{Bids: resp}, nil
}

func (service *Bid) Apply(data *dto.ApplyBid) (*dto.ApplyBidResp, error) {
	bidID, err := service.repo.Create(data.StudentID, data.SupervisorID, data.WorkID)
	if err != nil {
		return nil, err
	}

	return &dto.ApplyBidResp{BidID: bidID}, nil
}

func (service *Bid) Resolve(data *dto.ResolveBid) error {
	var status entity.StatusSSR

	if data.Accept {
		status = "accepted"
	} else {
		status = "rejected"
	}

	_, err := service.repo.UpdateStatus(data.BidID, status)
	return err
}
