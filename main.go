package main

import (
	"gotest/indirect/api"
	"gotest/indirect/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
)

func main() {
	ginS.Use(gin.Logger(), gin.Recovery())
	rg := ginS.Group("/account")
	rg.Use(middleware.Global)
	{
		rg.GET("/reg", api.Reg)
		rg.GET("/login", api.Login)
	}
	ginS.Run("127.0.0.1:8000")
}
