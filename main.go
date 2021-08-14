package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/employee", demoHandler)

	router.Run(":8080")
}

func demoHandler(c *gin.Context) {

	var m map[string]interface{}
	err := c.Bind(&m)
	if err != nil {
		return
	}

	fmt.Printf("%v\n", m)

	c.JSON(200, "success")
}
