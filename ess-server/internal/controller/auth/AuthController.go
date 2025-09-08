package auth

import (
	abstract "ess-server/internal/core/domain/service/usecase"
	dto "ess-server/internal/core/dto/auth"

	"github.com/gin-gonic/gin"
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

func (c *AuthController) Logout(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Logged out successfully"})
}
