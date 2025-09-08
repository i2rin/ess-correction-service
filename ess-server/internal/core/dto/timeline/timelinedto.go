package timeline

type TimeLineDTO struct {
	Name         string `json:"nickname"`
	Text         string `json:"text"`
	Date         string `json:"date"`
	SubmissionId string `json:"submissionid"`
}

func NewTimeLineDTO(name, text, date, submissionId string) *TimeLineDTO {
	return &TimeLineDTO{
		Name:         name,
		Text:         text,
		Date:         date,
		SubmissionId: submissionId,
	}
}

func (t *TimeLineDTO) GetName() string {
	return t.Name
}

func (t *TimeLineDTO) GetText() string {
	return t.Text
}

func (t *TimeLineDTO) GetDate() string {
	return t.Date
}

func (t *TimeLineDTO) GetSubmissionId() string {
	return t.SubmissionId
}
