package controller

import "github.com/gin-gonic/gin"

type TeacherController struct {
}

func (a TeacherController) TeacherRetakeRequest(c *gin.Context) {
	c.JSON(200, "hello")
	return
}

func (a TeacherController) TeacherUpdateRetake(c *gin.Context) {
	c.JSON(200, "hello")
	return
}
