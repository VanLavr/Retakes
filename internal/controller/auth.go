package controller

import "github.com/gin-gonic/gin"

type AuthController struct {
}

func (a AuthController) Login(c *gin.Context) {
	c.JSON(200, "hello")
	return
}
