package dto

import (
	"ssr/internal/entity"
)

type RelationCreateReq struct {
	StudentID    int `json:"student_id"`
	SupervisorID int `json:"supervisor_id"`
	WorkID       int `json:"work_id"`
}

type RelationResp struct {
	RelationID int        `json:"relation_id"`
	Work       WorkResp   `json:"work"`
	Status     string     `json:"status"`
	Supervisor Supervisor `json:"supervisor"`
	Student    Student    `json:"student"`
}

type RelationShortResp struct {
	RelationID int             `json:"relation_id"`
	Work       WorkShortResp   `json:"work"`
	Status     string          `json:"status"`
	Supervisor SupervisorShort `json:"supervisor"`
	Student    StudentShort    `json:"student"`
}

type RelationCreateResp struct {
	RelationID int `json:"relation_id"`
}

type RelationUpdateReq struct {
	RelationID int              `json:"relation_id"`
	Status     entity.StatusSSR `json:"status"`
}

type RelationPlenty struct {
	Relations []*RelationShortResp `json:"relations"`
}
