package timeline

import (
	abstract "ess-server/internal/core/domain/service/usecase"
	dto "ess-server/internal/core/dto/timeline"
)

type TestTimeLineUC struct {}

func NewTestTimeLineUC() *TestTimeLineUC {
	return &TestTimeLineUC{}
}

var _ abstract.ITimeLineUsecase = (*TestTimeLineUC)(nil)


func (u *TestTimeLineUC) GetTimeLine() ([]*dto.TimeLineDTO, error) {
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
