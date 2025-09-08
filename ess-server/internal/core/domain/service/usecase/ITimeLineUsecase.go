package usecase

import (
	dto "ess-server/internal/core/dto/timeline"
)

type ITimeLineUsecase interface {
	GetTimeLine() ([]*dto.TimeLineDTO, error)
}