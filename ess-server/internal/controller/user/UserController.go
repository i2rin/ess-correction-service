package user

import (
	abstract "ess-server/internal/core/domain/service/usecase"
	dto "ess-server/internal/core/dto/user"

	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	useruc abstract.IUserUsecase
}

func NewUserController(userUsecase abstract.IUserUsecase) *UserController {
	return &UserController{
		useruc: userUsecase,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var userDTO dto.UserInDTO

	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.useruc.CreateUser(&userDTO); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"status": "user created"})
}

func (c *UserController) GetUser(ctx *gin.Context) {
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

	users, err := c.useruc.GetUser(pageInt, sizeInt)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, users)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var userDTO dto.UserInDTO
	userId := ctx.Param("userId")

	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.useruc.UpdateUser(&userDTO, userId); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"status": "user updated"})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	if err := c.useruc.DeleteUser(userId); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"status": "user deleted"})
}
