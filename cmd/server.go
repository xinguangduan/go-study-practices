package main

import "github.com/gin-gonic/gin"

func Run() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Run(":9090")
}
