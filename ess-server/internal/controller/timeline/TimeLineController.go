package timeline

import (
	abstract "ess-server/internal/core/domain/service/usecase"
	"strconv"

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
	page := ctx.Query("page")
	size := ctx.Query("size")

	if page == "" {
		page = "1"
	}
	if size == "" {
		size = "10"
	}

	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)

	timeLines, err := c.timeLineUsecase.GetTimeLine(pageInt, sizeInt)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, timeLines)
}
