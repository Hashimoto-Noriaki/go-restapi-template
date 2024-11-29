package controllers

import (
	"go-restapi-boiltertemplate/app/services"
	"go-restapi-boiltertemplate/app/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostControllerInterface interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type PostController struct {
	service services.PostService
}

func NewPostController(service services.PostService) *PostController {
	return &PostController{service: service} // ポインタを返すように変更
}

func (c *PostController) FindAll(ctx *gin.Context) {
	posts, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": posts})
}

func (c *PostController) FindByID(ctx *gin.Context) {
	// URL パラメータ "id" を取得
	postID, err := strconv.Atoi(ctx.Param("id")) // ctx.Param から "id" を取得して整数に変換
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

    postIDUint := uint(postID)

	post, err := c.service.FindByID(postIDUint)
	if err != nil {
		if err.Error() == "投稿が見つかりません" { // エラーメッセージを修正
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	// 投稿が見つかった場合、データを返す
	ctx.JSON(http.StatusOK, gin.H{"data": post}) // 修正: post を返す
}

func (c *PostController) Create(ctx *gin.Context){
	var input dto.CreatePostInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newPost, err := c.service.Create(input)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"data": newPost})
}

func (c *PostController) Update(ctx *gin.Context) {
    postID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
        return 
    }

    var input dto.UpdatePostInput
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedPost, err := c.service.Update(uint(postID), input)
    if err != nil {
        if err.Error() == "Post not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"data": updatedPost})
}

func (c *PostController) Delete(ctx *gin.Context) {
    // URL パラメータ "id" を取得
    postID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
        return 
    }

    err = c.service.Delete(uint(postID))
    if err != nil {
        if err.Error() == "Post not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        return
    }
    ctx.Status(http.StatusOK)
}