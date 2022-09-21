package usecase

import (
	"ssr/pkg/logger"
)

type Base struct {
	l logger.Interface
}

func NewBase(l logger.Interface) *Base {
	return &Base{
		l: l,
	}
}
