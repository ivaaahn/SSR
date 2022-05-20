package usecase

import (
	"database/sql"
	"github.com/golang/mock/gomock"
	"reflect"
	"ssr/internal/dto"
	"ssr/internal/entity"
	"ssr/internal/usecase/mocks"
	"ssr/pkg/misc"
	"testing"
	"time"
)

func TestBidUseCase_Apply(t *testing.T) {
	type args struct {
		bid *dto.ApplyBid
	}
	tests := []struct {
		name       string
		repository func(ctrl *gomock.Controller) *mocks.MockIRelRepo
		args       args
		want       *dto.ApplyBidResponse
		wantErr    bool
	}{
		{
			name: "success apply",
			repository: func(ctrl *gomock.Controller) *mocks.MockIRelRepo {
				m := mocks.NewMockIRelRepo(ctrl)
				m.EXPECT().Create(1, 1, 1).Return(1, nil)

				return m
			},
			args: args{
				bid: &dto.ApplyBid{
					StudentID:    1,
					SupervisorID: 1,
					WorkID:       1,
				},
			},
			want:    &dto.ApplyBidResponse{BidID: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := NewBidUC(tt.repository(ctrl))

			received, err := u.Apply(tt.args.bid)

			if (err != nil) != tt.wantErr {
				t.Errorf("Bid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(received, tt.want) {
				t.Errorf("Bid() got = %v, want %v", received, tt.want)
			}
		})
	}
}

func TestBidUseCase_Resolve(t *testing.T) {
	type args struct {
		bid *dto.ResolveBid
	}
	tests := []struct {
		name       string
		repository func(ctrl *gomock.Controller) *mocks.MockIRelRepo
		args       args
		wantErr    bool
	}{
		{
			name: "accept success",
			repository: func(ctrl *gomock.Controller) *mocks.MockIRelRepo {
				m := mocks.NewMockIRelRepo(ctrl)
				m.EXPECT().UpdateStatus(1, entity.StatusSSR("accepted")).Return(1, nil)

				return m
			},
			args: args{
				bid: &dto.ResolveBid{
					SupervisorID: 1,
					BidID:        1,
					Accept:       true,
				},
			},
			wantErr: false,
		},
		{
			name: "reject success",
			repository: func(ctrl *gomock.Controller) *mocks.MockIRelRepo {
				m := mocks.NewMockIRelRepo(ctrl)
				m.EXPECT().UpdateStatus(1, entity.StatusSSR("rejected")).Return(1, nil)

				return m
			},
			args: args{
				bid: &dto.ResolveBid{
					SupervisorID: 1,
					BidID:        1,
					Accept:       false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := NewBidUC(tt.repository(ctrl))

			err := u.Resolve(tt.args.bid)

			if (err != nil) != tt.wantErr {
				t.Errorf("Bid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSSRUseCase_Create(t *testing.T) {
	type args struct {
		req *dto.CreateSSR
	}
	tests := []struct {
		name       string
		repository func(ctrl *gomock.Controller) *mocks.MockIRelRepo
		args       args
		want       *dto.StudentViewSSR
		wantErr    bool
	}{
		{
			name: "success ssr create",
			repository: func(ctrl *gomock.Controller) *mocks.MockIRelRepo {
				m := mocks.NewMockIRelRepo(ctrl)
				m.EXPECT().UpdateStatus(1, entity.StatusSSR("wip")).Return(1, nil)
				m.EXPECT().GetStudentViewSSR(1, 1).Return(&entity.StudentViewSSR{
					BidID:     1,
					CreatedAt: time.Time{},
					Status:    "wip",
					SupervisorProfile: &entity.SupervisorProfile{
						User: &entity.User{
							UserID:    2,
							Email:     "kek@kek.com",
							FirstName: "Иван",
							LastName:  "Иванов",
							Avatar:    sql.NullString{},
						},
						SupervisorID: 1,
						Birthdate:    time.Time{},
						About:        "Обо мне",
						DepartmentID: "ИУ7",
					},
					Work: &entity.Work{
						WorkKind: &entity.WorkKind{
							WorkKindID:   1,
							WorkKindName: "Курсовая работа",
						},
						Subject: &entity.Subject{
							SubjectID:    1,
							SubjectName:  "Операционные системы",
							DepartmentID: "ИУ7",
						},
						WorkID:      1,
						Description: "Работа для профи!",
						Semester:    7,
					},
				}, nil)

				return m
			},
			args: args{
				req: &dto.CreateSSR{
					StudentID: 1,
					BidID:     1,
				},
			},
			want: &dto.StudentViewSSR{
				RelID:     1,
				Status:    "wip",
				CreatedAt: time.Time{},
				Supervisor: dto.SupervisorProfile{
					SupervisorID: 1,
					Email:        "kek@kek.com",
					FirstName:    "Иван",
					LastName:     "Иванов",
					About:        "Обо мне",
					Birthdate:    misc.Date{},
					AvatarUrl:    misc.NullString{},
					Department:   "ИУ7",
				},
				Work: dto.Work{
					WorkID:      1,
					Name:        "Курсовая работа",
					Description: "Работа для профи!",
					Semester:    7,
					Subject: dto.SubjectResp{
						SubjectID:  1,
						Name:       "Операционные системы",
						Department: "ИУ7",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success ssr create",
			repository: func(ctrl *gomock.Controller) *mocks.MockIRelRepo {
				m := mocks.NewMockIRelRepo(ctrl)
				m.EXPECT().UpdateStatus(1, entity.StatusSSR("wip")).Return(1, nil)
				m.EXPECT().GetStudentViewSSR(1, 1).Return(&entity.StudentViewSSR{
					BidID:     1,
					CreatedAt: time.Time{},
					Status:    "wip",
					SupervisorProfile: &entity.SupervisorProfile{
						User: &entity.User{
							UserID:    2,
							Email:     "kek@kek.com",
							FirstName: "Иван",
							LastName:  "Иванов",
							Avatar:    sql.NullString{},
						},
						SupervisorID: 1,
						Birthdate:    time.Time{},
						About:        "Обо мне",
						DepartmentID: "ИУ7",
					},
					Work: &entity.Work{
						WorkKind: &entity.WorkKind{
							WorkKindID:   1,
							WorkKindName: "Курсовая работа",
						},
						Subject: &entity.Subject{
							SubjectID:    1,
							SubjectName:  "Операционные системы",
							DepartmentID: "ИУ7",
						},
						WorkID:      1,
						Description: "Работа для профи!",
						Semester:    7,
					},
				}, nil)

				return m
			},
			args: args{
				req: &dto.CreateSSR{
					StudentID: 1,
					BidID:     1,
				},
			},
			want: &dto.StudentViewSSR{
				RelID:     1,
				Status:    "wip",
				CreatedAt: time.Time{},
				Supervisor: dto.SupervisorProfile{
					SupervisorID: 1,
					Email:        "kek@kek.com",
					FirstName:    "Иван",
					LastName:     "Иванов",
					About:        "Обо мне",
					Birthdate:    misc.Date{},
					AvatarUrl:    misc.NullString{},
					Department:   "ИУ7",
				},
				Work: dto.Work{
					WorkID:      1,
					Name:        "Курсовая работа",
					Description: "Работа для профи!",
					Semester:    7,
					Subject: dto.SubjectResp{
						SubjectID:  1,
						Name:       "Операционные системы",
						Department: "ИУ7",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := NewSsrUC(tt.repository(ctrl))

			received, err := u.Create(tt.args.req)

			if (err != nil) != tt.wantErr {
				t.Errorf("Bid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(received, tt.want) {
				t.Errorf("Bid() got = %v, want %v", received, tt.want)
			}
		})
	}
}
