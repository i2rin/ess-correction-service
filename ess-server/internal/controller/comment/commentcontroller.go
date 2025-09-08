package comment

import (
	abstract "ess-server/internal/core/domain/service/usecase"
	dto "ess-server/internal/core/dto/comment"
	"github.com/gin-gonic/gin"
)


type CommentController struct {
	usecase abstract.ICommentUsecase
}

func NewCommentController(usecase abstract.ICommentUsecase) *CommentController {
	return &CommentController{
		usecase: usecase,
	}
}

func (c *CommentController) GetComments(ctx *gin.Context) {
	comments, err := c.usecase.GetComments()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, comments)
}

func (c *CommentController) PostComment(ctx *gin.Context) {
	var comment dto.CommentInDTO
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.usecase.PostComment(&comment); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{"status": "comment posted"})
}