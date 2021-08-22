package handler

import (
	"log"
	"strconv"

	"abc.com/demo/model"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{}
}

func (h *EmployeeHandler) GetEmployeeById() gin.HandlerFunc {
	m := map[int]model.Employee{
		1: {Id: 1, Name: "John", Age: 33},
		2: {Id: 2, Name: "Mary", Age: 28},
	}

	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		log.Printf("id=%d\n", id)

		name := m[id].Name
		c.JSON(200, gin.H{
			"name": name,
		})
	}
}
