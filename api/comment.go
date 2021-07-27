package api

import (
	"CarDemo1/pkg/logging"
	"CarDemo1/service"
	"github.com/gin-gonic/gin"
	"strconv"
)


//新增评论
func CreateComment(c *gin.Context) {
	service := service.CreateNewComment{}
	if err := c.ShouldBind(&service);err==nil{
		res:=service.Create()
		c.JSON(200,res)
	}else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//删除评论
func DeleteComment(c *gin.Context) {
	service := service.DeleteCommentService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}


//ShowSingleComm 查找单条评论
func ShowSingleComm(c *gin.Context) {
	service := service.Page{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Find(c.Query("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}


// ShowSingleChildren 查找单条子评论
func ShowSingleChildren(c *gin.Context) {
	service := service.Page{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.FindChildren(c.Query("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func ShowAllComment(c *gin.Context) {
	service := service.Page{}
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c.Param("id"),pageSize, pageNum)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func ShowAllComChildren(c *gin.Context) {
	service := service.Page{}
	if err := c.ShouldBindQuery(&service); err == nil {
		res := service.ListChild()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}




