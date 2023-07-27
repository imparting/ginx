package api

import (
	"gotest/indirect/service"
	"gotest/indirect/service/interfaces/dto"

	"github.com/gin-gonic/gin"
)

func Reg(c *gin.Context) {
	in := dto.AccountCreateInDto{
		Account:  c.Query("account"),
		Password: c.Query("password"),
	}
	out := service.Account().Create(in)
	c.JSON(200, out)
}
