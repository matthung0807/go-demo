package main

import (
	"abc.com/demo/route"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	group := router.Group("/demo")
	route.NewDemoRoute(group).Route()
	router.Run(":8080")
}
