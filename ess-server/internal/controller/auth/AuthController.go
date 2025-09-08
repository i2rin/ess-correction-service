package auth

import (
	"github.com/gin-gonic/gin"
	dto "ess-server/internal/core/dto/auth"
	abstract "ess-server/internal/core/domain/service/usecase"
)

type AuthController struct {
	authuc abstract.IAuthUsecase
}

func NewAuthController(authuc abstract.IAuthUsecase) *AuthController {
	return &AuthController{
		authuc: authuc,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginInDTO dto.LoginInDTO

	if err := ctx.ShouldBindJSON(&loginInDTO); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	loginOutDTO, err := c.authuc.Login(&loginInDTO)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, loginOutDTO)
}
