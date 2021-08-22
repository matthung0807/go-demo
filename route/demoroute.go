package route

import (
	"abc.com/demo/handler"
	"github.com/gin-gonic/gin"
)

type DemoRoute struct {
	*gin.RouterGroup
	employeeHandler *handler.EmployeeHandler
}

func NewDemoRoute(group *gin.RouterGroup) *DemoRoute {
	return &DemoRoute{
		group,
		handler.NewEmployeeHandler(),
	}
}

func (route *DemoRoute) Route() {
	route.GET("/employee/:id", route.employeeHandler.GetEmployeeById())
}
