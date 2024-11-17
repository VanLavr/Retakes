package controller

import "github.com/gin-gonic/gin"

type TeacherService interface{}

type TeacherController struct {
	s TeacherService
}

func NewTeacherController(s TeacherService) *TeacherController {
	return &TeacherController{
		s: s,
	}
}

func (a TeacherController) TeacherRetakeRequest(c *gin.Context) {
	c.JSON(200, "hello")
	return
}

func (a TeacherController) TeacherUpdateRetake(c *gin.Context) {
	c.JSON(200, "hello")
	return
}
