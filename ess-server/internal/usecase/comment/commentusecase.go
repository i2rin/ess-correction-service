package comment

import (
	dto "ess-server/internal/core/dto/comment"
	abstract "ess-server/internal/core/domain/service/usecase"
)

type CommentUsecase struct {
}

func NewCommentUsecase() *CommentUsecase {
	return &CommentUsecase{}
}

var _ abstract.ICommentUsecase = (*CommentUsecase)(nil)


func (u *CommentUsecase) GetComments() ([]dto.CommentOutDTO, error) {
	return []dto.CommentOutDTO{
		*dto.NewCommentOutDTO("submission1", []string{"Great job!", "Needs improvement in section 2."}, "2023-10-01"),
		*dto.NewCommentOutDTO("submission2", []string{"Well done!", "Consider adding more examples."}, "2023-10-02"),
	}, nil
}

func (u *CommentUsecase) PostComment(dto *dto.CommentInDTO) error {
	return nil
}