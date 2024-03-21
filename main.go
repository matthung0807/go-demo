package main

import (
	"fmt"

	"abc.com/demo/model"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/employee", demoHandler)

	router.Run(":8080")
}

func demoHandler(c *gin.Context) {
	var dep model.Department
	err := c.Bind(&dep)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", dep)
	c.JSON(200, "success")
}
