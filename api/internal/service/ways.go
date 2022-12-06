package service

import (
	"ways/internal/repository"

	"github.com/rs/zerolog"
)

func NewWaysService(lg *zerolog.Logger, vertices *repository.WaysVerticesRepo) *WaysService {
	return &WaysService{
		logger:   lg,
		vertices: vertices,
	}
}

type WaysService struct {
	logger   *zerolog.Logger
	vertices *repository.WaysVerticesRepo
}
