package controller

import "github.com/gin-gonic/gin"

type AdminController struct {
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
