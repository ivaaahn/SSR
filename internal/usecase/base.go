package usecase

import (
	"ssr/pkg/logger"
)

type BaseUC struct {
	l logger.Interface
}

func NewUC(l logger.Interface) *BaseUC {
	return &BaseUC{
		l: l,
	}
}
