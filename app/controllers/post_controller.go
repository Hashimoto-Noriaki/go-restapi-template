package controllers

import (
	"go-restapi-boiltertemplate/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostControllerInterface interface {
	FindAll(ctx *gin.Context)
}

type PostController struct {
	service services.PostService
}

func NewPostController(service services.PostService) PostController {
	return PostController{service: service}
}

func (c *PostController) FindAll(ctx *gin.Context) {
	posts, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": posts})
}
