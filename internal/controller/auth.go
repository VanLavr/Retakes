package controller

import "github.com/gin-gonic/gin"

type AuthService interface{}

type AuthController struct {
	s AuthService
}

func NewAuthController(s AuthService) *AuthController {
	return &AuthController{s: s}
}

func (a AuthController) Login(c *gin.Context) {
	c.JSON(200, "hello")
	return
}
