package controller

import "github.com/gin-gonic/gin"

type StudentController struct {
}

func (a StudentController) DebtInfo(c *gin.Context) {
	c.JSON(200, "hello")
	return
}

func (a StudentController) StudentRetakeRequest(c *gin.Context) {
	c.JSON(200, "hello")
	return
}
