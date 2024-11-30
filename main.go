package main

import (
	"go-restapi-boiltertemplate/app/infra"
	"go-restapi-boiltertemplate/app/repositories"
	"go-restapi-boiltertemplate/app/controllers"
	"go-restapi-boiltertemplate/app/services"
	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initializer()
	db := infra.SetupDB()
	// サンプルのPostデータを作成
	// posts := []models.Post{
	// 	{ID: 1, Text: "最初の投稿", UserID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	// 	{ID: 2, Text: "2番目の投稿", UserID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	// 	{ID: 3, Text: "3番目の投稿", UserID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	// }

	// PostRepositoryを作成
	postRepository := repositories.NewPostRepository(db)
	
	// PostServiceを作成
	postService := services.NewPostService(postRepository)

	// PostControllerを作成
	postController := controllers.NewPostController(postService)

	// Ginのルーターをセットアップ
	r := gin.Default()
	itemRouter := r.Group("/posts")
    r.GET("", postController.FindAll) // エンドポイントを追加
    r.GET("/:id",postController.FindById)
    r.POST("",postController.Create)
    r.PUT("/:id",postController.Update)
    r.DELETE("/:id",postController.Delete)
	r.Run("localhost:8081") // サーバーを8081番ポートで実行
}
