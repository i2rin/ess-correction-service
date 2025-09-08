package main

import (
	userctl "ess-server/internal/controller/user"
	useruc "ess-server/internal/usecase/user"
	"ess-server/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	userusecase := useruc.NewTestUserUC()
	usercontroller := userctl.NewUserController(userusecase)

	r.POST("/users", usercontroller.CreateUser)
	r.GET("/users", usercontroller.GetUser)
	r.PUT("/users/:userId", usercontroller.UpdateUser)
	r.DELETE("/users/:userId", usercontroller.DeleteUser)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
