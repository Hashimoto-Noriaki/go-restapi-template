package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/template", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ボイラーテンプレート",
		})
	})
	r.Run("localhost:8080")
}
