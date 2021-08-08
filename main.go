package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.DisableConsoleColor()
	file, _ := os.Create("gin.log")                     // create log file
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout) // wirte log to file and console

	router := gin.Default() // with logger attached
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	router.Run()
}
