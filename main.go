package main

import (
	"fmt"

	"abc.com/demo/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(
		middleware.PrintHello("1"),
		middleware.PrintHello("2"),
		middleware.PrintHello("3"),
	)

	router.GET("/demo", func(c *gin.Context) {
		msg := "demo"
		fmt.Println(msg)
		c.JSON(200, msg)
	})
	router.Run()
}
