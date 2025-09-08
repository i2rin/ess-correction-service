package comment

type CommentInDTO struct {
	Name       string
	Text       string
	Corrections []struct {
		Original string
		Corrected string
		Reason   string
	}
}

func NewCommentInDTO(name, text string, corrections []struct {
	Original string
	Corrected string
	Reason   string
}) *CommentInDTO {
	return &CommentInDTO{
		Name:       name,
		Text:       text,
		Corrections: corrections,
	}
}

func (c *CommentInDTO) GetName() string {
	return c.Name
}

func (c *CommentInDTO) GetText() string {
	return c.Text
}

func (c *CommentInDTO) GetCorrections() []struct {
	Original string
	Corrected string
	Reason   string
} {
	return c.Corrections
}


type CommentOutDTO struct {
	Submissionid string
	Message      []string
	Date        string
}

func NewCommentOutDTO(submissionid string, message []string, date string) *CommentOutDTO {
	return &CommentOutDTO{
		Submissionid: submissionid,
		Message:      message,
		Date:        date,
	}
}

func (c *CommentOutDTO) GetSubmissionID() string {
	return c.Submissionid
}

func (c *CommentOutDTO) GetMessage() []string {
	return c.Message
}

func (c *CommentOutDTO) GetDate() string {
	return c.Date
}