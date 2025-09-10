package timeline

import (
	abstract "ess-server/internal/core/domain/service/usecase"
	dto "ess-server/internal/core/dto/timeline"
)

type TimeLineUsecase struct {}

func NewTimeLineUsecase() *TimeLineUsecase {
	return &TimeLineUsecase{}
}

var _ abstract.ITimeLineUsecase = (*TimeLineUsecase)(nil)


func (u *TimeLineUsecase) GetTimeLine() ([]*dto.TimeLineDTO, error) {
	return []*dto.TimeLineDTO{
		{
			Name:        "Sample Name",
			Text:       "Sample Text",
			Date:       "2023-01-01",
			SubmissionId: "12345",
		},
		{
			Name:        "Another Name",
			Text:       "Another Text",
			Date:       "2023-02-01",
			SubmissionId: "67890",
		},
	}, nil
}
