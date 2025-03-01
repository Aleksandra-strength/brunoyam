package app

import (
	"hello/internal/app/handlers"
	"github.com/gin-gonic/gin"
)

func InitEndPoints(r *gin.Engine) {
	r.POST("/task", handlers.PostTask)
	r.GET("/task", handlers.GetTask)
	r.PUT("/task/:id", handlers.PutTask)
	r.DELETE("/task/:id", handlers.DelTask)
	r.POST("/tasks/save", handlers.SaveTasks)
	r.GET("/tasks/load", handlers.LoadTasks)

}
