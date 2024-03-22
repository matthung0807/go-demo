package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type QueryString struct {
	Name      string     `form:"name" binding:"required"`
	Age       int        `form:"age" binding:"required"`
	StartTime *time.Time `form:"startTime" binding:"required" time_format:"2006-01-02T15:04:05" time_utc:"1"`
	EndTime   *time.Time `form:"endTime" binding:"required" time_format:"2006-01-02T15:04:05" time_utc:"1"`
}

func (q QueryString) String() string {
	return fmt.Sprintf("{Name=%s, Age=%d, StartTime=%s, EndTime=%s}", q.Name, q.Age, q.StartTime, q.EndTime)
}

func main() {
	router := gin.Default()
	router.GET("/find", func(c *gin.Context) {
		var qs QueryString
		err := c.BindQuery(&qs)
		if err != nil {
			return
		}
		fmt.Println(qs)
		c.Status(http.StatusOK)
	})
	router.Run()
}
