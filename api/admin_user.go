package api

import (
	logging "github.com/sirupsen/logrus"
	"CarDemo1/service"
	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	service := service.ListUsersService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

