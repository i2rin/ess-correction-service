package main

import (
	"ess-server/pkg/db"
	authctl "ess-server/internal/controller/auth"
	authuc "ess-server/internal/usecase/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()


	r := gin.Default()
	authUsecase := authuc.NewTestAuthUC()
	authcontroller := authctl.NewAuthController(authUsecase)
	r.POST("/auth/login", authcontroller.Login)
	r.Run(":8080")
}
