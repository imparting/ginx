package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Global(c *gin.Context) {
	fmt.Println("------1-----")
	c.Next()
	fmt.Println("------2-----")
}
