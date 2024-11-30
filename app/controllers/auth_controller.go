package controllers

import (
	"go-restapi-boiltertemplate/app/dto"
	"go-restapi-boiltertemplate/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PAuthController interface {
	Signup(ctx *gin.Context)
}

type AuthController struct {
	service services.PAuthService
}

func NewAuthController(service services.PAuthService) PAuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Signup(ctx *gin.Context){
	var input dto.SignupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := c.service.Signup(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.Status(http.StatusCreated)
}
