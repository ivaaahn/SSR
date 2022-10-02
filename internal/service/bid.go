package service

import (
	"ssr/internal/dto"
	"ssr/pkg/logger"
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

func (service *Bid) GetSupervisorBids(supervisorID int) (*dto.SvBids, error) {
	dbData, err := service.repo.GetSupervisorBids(supervisorID)
	if err != nil {
		return nil, err
	}

	var resp []*dto.SvBid

	for _, db := range dbData {
		resp = append(resp, &dto.SvBid{
			BidID:     db.RelationID,
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
					ID:         db.SubjectID,
					Name:       db.Subject.Name,
					Department: db.Subject.DepartmentID,
				},
			},
		})
	}

	return &dto.SvBids{Bids: resp}, nil
}

//func (service *Bid) Apply(data *dto.RelationCreateReq) (*dto.RelationCreateResp, error) {
//	bidID, err := service.repo.Create(data.StudentID, data.SupervisorID, data.WorkID)
//	if err != nil {
//		return nil, err
//	}
//
//	return &dto.RelationCreateResp{BidID: bidID}, nil
//}
//
//func (service *Bid) Resolve(data *dto.ResolveBid) error {
//	var status entity.StatusSSR
//
//	if data.Accept {
//		status = "accepted"
//	} else {
//		status = "rejected"
//	}
//
//	_, err := service.repo.UpdateStatus(data.BidID, status)
//	return err
//}
