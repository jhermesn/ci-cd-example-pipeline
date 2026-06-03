package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jhermesn/ci-cd-example-pipeline/handler"
)

func main() {
	r := gin.Default()

	r.GET("/health", handler.Health)
	r.POST("/words", handler.Words)
	r.GET("/count", handler.Count)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
