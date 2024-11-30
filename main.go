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

	postRepository := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepository)
	postController := controllers.NewPostController(postService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	// Ginのルーターをセットアップ
	r := gin.Default()

	postRouter := r.Group("/posts")
	{
		postRouter.GET("", postController.FindAll)
		postRouter.GET("/:id", postController.FindByID)
		postRouter.POST("", postController.Create)
		postRouter.PUT("/:id", postController.Update)
		postRouter.DELETE("/:id", postController.Delete)
	}

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/signup", authController.Signup)
	}

	r.Run("localhost:8081")// サーバーを8081番ポートで実行
}
