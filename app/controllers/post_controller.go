package controllers

import (
	"go-restapi-boiltertemplate/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostControllerInterface interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
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

func (c *PostController) FindById(ctx *gin.Context) {
	// URL パラメータ "id" を取得
	postId, err := strconv.Atoi(ctx.Param("id")) // ctx.Param から "id" を取得して整数に変換
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	// サービスから投稿を取得
	post, err := c.service.FindById(postId) // 修正: postId を使って投稿を取得
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
