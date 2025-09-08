package comment

import (
	abstract "ess-server/internal/core/domain/service/usecase"
	dto "ess-server/internal/core/dto/comment"
)

type TestCommentUC struct {
}

func NewTestCommentUC() *TestCommentUC {
	return &TestCommentUC{}
}

var _ abstract.ICommentUsecase = (*TestCommentUC)(nil)

func (u *TestCommentUC) GetComments(page, size int) ([]dto.CommentOutDTO, error) {
	return []dto.CommentOutDTO{
		*dto.NewCommentOutDTO(
			"Test User",
			"This is a test comment.",
			[]dto.Correction{
				{
					Original:  "teh",
					Corrected: "the",
					Reason:    "Typo correction",
				},
			},
			[]dto.NestedComment{
				{
					Nickname: "NestedUser1",
					Comment:  "This is a nested comment.",
				},
			},
		),
		*dto.NewCommentOutDTO(
			"Another User",
			"Another test comment.",
			[]dto.Correction{},
			[]dto.NestedComment{},
		),
		*dto.NewCommentOutDTO(
			"User3",
			"This is a comment from User3.",
			[]dto.Correction{
				{
					Original:  "mistke",
					Corrected: "mistake",
					Reason:    "Typo correction",
				},
			},
			[]dto.NestedComment{
				{
					Nickname: "NestedUser2",
					Comment:  "This is another nested comment.",
				},
			},
		),
	}, nil
}

func (u *TestCommentUC) PostComment(dto *dto.CommentInDTO) error {
	return nil
}
