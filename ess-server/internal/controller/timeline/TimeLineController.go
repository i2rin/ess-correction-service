package timeline

import (
	abstract "ess-server/internal/core/domain/service/usecase"
	"github.com/gin-gonic/gin"
)

type TimeLineController struct {
	timeLineUsecase abstract.ITimeLineUsecase
}

func NewTimeLineController(timeLineUsecase abstract.ITimeLineUsecase) *TimeLineController {
	return &TimeLineController{
		timeLineUsecase: timeLineUsecase,
	}
}

func (c *TimeLineController) GetTimeLine(ctx *gin.Context) {
	timeLines, err := c.timeLineUsecase.GetTimeLine()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, timeLines)
}