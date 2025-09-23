package main

import (
	authctl "ess-server/internal/controller/auth"
	authuc "ess-server/internal/usecase/auth"
	"ess-server/pkg/db"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()
	authUsecase := authuc.NewTestAuthUC()
	authcontroller := authctl.NewAuthController(authUsecase)
	r.POST("/login", authcontroller.Login)
	r.GET("/logout", authcontroller.Logout)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
