package usecase

import (
	dto "ess-server/internal/core/dto/comment"
)

type ICommentUsecase interface {
	GetComments() ([]dto.CommentOutDTO, error)
	PostComment(dto *dto.CommentInDTO) error
}