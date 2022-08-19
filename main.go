package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", Hello)

	router.Run(":3000")
}

func Hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello")
}
