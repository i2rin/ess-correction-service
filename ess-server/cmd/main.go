package main

import (
	timelinectl "ess-server/internal/controller/timeline"
	timelineuc "ess-server/internal/usecase/timeline"
	"ess-server/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()
	timeLineUsecase := timelineuc.NewTestTimeLineUC()
	timeLineController := timelinectl.NewTimeLineController(timeLineUsecase)

	r.GET("/timeline", timeLineController.GetTimeLine)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
