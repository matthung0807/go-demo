package main

import "github.com/gin-gonic/gin"

func main() {
	server := setupServer()
	server.Run(":8080")
}

func setupServer() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello world"})
	})

	return router
}
