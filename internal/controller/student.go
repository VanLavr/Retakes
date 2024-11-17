package controller

import "github.com/gin-gonic/gin"

type StudentService interface{}

type StudentController struct {
	s StudentService
}

func NewStudentService(s StudentService) *StudentController {
	return &StudentController{
		s: s,
	}
}

func (a StudentController) DebtInfo(c *gin.Context) {
	c.JSON(200, "hello")
	return
}

func (a StudentController) StudentRetakeRequest(c *gin.Context) {
	c.JSON(200, "hello")
	return
}
