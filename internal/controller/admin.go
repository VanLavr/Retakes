package controller

import "github.com/gin-gonic/gin"

type AdminService interface {
}

type AdminController struct {
	s AdminService
}

func NewAdminController(s AdminController) *AdminController {
	return &AdminController{
		s: s,
	}
}

func (a AdminController) UploadRetakeData(c *gin.Context) {
	c.JSON(200, "hello")
	return
}

func (a AdminController) UploadStudentRetakes(c *gin.Context) {
	c.JSON(200, "hello")
	return
}

func (a AdminController) UploadTeacherRetakes(c *gin.Context) {
	c.JSON(200, "hello")
	return
}
