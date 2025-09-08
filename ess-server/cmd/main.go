package main

import (
	"ess-server/pkg/db"
	authctl "ess-server/internal/controller/auth"
	authuc "ess-server/internal/usecase/auth"
	userctl "ess-server/internal/controller/user"
	useruc "ess-server/internal/usecase/user"
	timelineuc "ess-server/internal/usecase/timeline"
	timelinectl "ess-server/internal/controller/timeline"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()


	r := gin.Default()
	authUsecase := authuc.NewTestAuthUC()
	userusecase := useruc.NewTestUserUC()
	timeLineUsecase := timelineuc.NewTestTimeLineUC()

	authcontroller := authctl.NewAuthController(authUsecase)
	usercontroller := userctl.NewUserController(userusecase)
	timeLineController := timelinectl.NewTimeLineController(timeLineUsecase)

	r.GET("/timeline", timeLineController.GetTimeLine)
	r.POST("/user/create", usercontroller.CreateUser)
	r.GET("/user/get/:userId", usercontroller.GetUser)
	r.PUT("/user/update", usercontroller.UpdateUser)
	r.DELETE("/user/delete/:userId", usercontroller.DeleteUser)
	r.POST("/auth/login", authcontroller.Login)
	r.Run(":8080")
}
