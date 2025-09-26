package main

import (
	commentctl "ess-server/internal/controller/comment"
	commentuc "ess-server/internal/usecase/comment"
	"ess-server/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()
	commentUsecase := commentuc.NewTestCommentUC()
	commentController := commentctl.NewCommentController(commentUsecase)

	r.GET("/comments", commentController.GetComments)
	r.POST("/comments", commentController.PostComment)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
