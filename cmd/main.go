package main

import (
	"hello/internal/app"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	app.InitEndPoints(r)

	r.Run(":8080")
}