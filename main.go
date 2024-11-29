package main

import (
	"time"
	"go-restapi-boiltertemplate/app/models"
	"go-restapi-boiltertemplate/app/repositories"
	"go-restapi-boiltertemplate/app/controllers"
	"go-restapi-boiltertemplate/app/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// サンプルのPostデータを作成
	posts := []models.Post{
		{ID: 1, Text: "最初の投稿", UserID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Text: "2番目の投稿", UserID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 3, Text: "3番目の投稿", UserID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	// PostRepositoryを作成
	postRepository := repositories.NewPostMemoryRepository(posts)
	
	// PostServiceを作成
	postService := services.NewPostService(postRepository)  // 変数名修正

	// PostControllerを作成
	postController := controllers.NewPostController(postService)  // 変数名修正

	// Ginのルーターをセットアップ
	r := gin.Default()
	r.GET("/posts", postController.FindAll)
	r.GET("/posts/:id", postController.FindById)
	r.POST("/posts", postController.Create)
	r.Run("localhost:8081") // サーバーを8081番ポートで実行
}
