package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	demo := router.Group("/demo")

	demo.Group("/v1").
		GET("/hello", func(c *gin.Context) {
			c.JSON(200, "Hello - v1")
		}).
		GET("/hi", func(c *gin.Context) {
			c.JSON(200, "Hi - v1")
		})

	demo.Group("/v2").
		GET("/hello", func(c *gin.Context) {
			c.JSON(200, "Hello - v2")
		}).
		GET("/hi", func(c *gin.Context) {
			c.JSON(200, "Hi - v2")
		})

	router.Run()
}
