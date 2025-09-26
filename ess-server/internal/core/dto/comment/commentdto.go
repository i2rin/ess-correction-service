package comment

// Correction 修正情報を表すDTO
type Correction struct {
	Original  string `json:"original"`
	Corrected string `json:"corrected"`
	Reason    string `json:"reason"`
}

// NestedComment 入れ子になったコメント情報
type NestedComment struct {
	Nickname string `json:"nickname"`
	Comment  string `json:"comment"`
}

// CommentInDTO 入力用DTO
type CommentOutDTO struct {
	Name        string          `json:"name"`
	Text        string          `json:"text"`
	Corrections []Correction    `json:"corrections"`
	Comments    []NestedComment `json:"comments"`
}

func NewCommentOutDTO(name, text string, corrections []Correction, comments []NestedComment) *CommentOutDTO {
	return &CommentOutDTO{
		Name:        name,
		Text:        text,
		Corrections: corrections,
		Comments:    comments,
	}
}

func (c *CommentOutDTO) GetName() string {
	return c.Name
}

func (c *CommentOutDTO) GetText() string {
	return c.Text
}

func (c *CommentOutDTO) GetCorrections() []Correction {
	return c.Corrections
}

func (c *CommentOutDTO) GetComments() []NestedComment {
	return c.Comments
}

// CommentOutDTO 出力用DTO
type CommentInDTO struct {
	Comment []string `json:"comment"`
}

func NewCommentInDTO() *CommentInDTO {
	return &CommentInDTO{
		Comment: []string{},
	}
}

func (c *CommentInDTO) GetComment() []string {
	return c.Comment
}
