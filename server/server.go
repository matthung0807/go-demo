package server

import (
	"net/http"

	"abc.com/demo/internal/handler"
	"abc.com/demo/internal/infra/mq"
	"abc.com/demo/internal/infra/postgres/repo"
	"abc.com/demo/internal/proxy"
	"abc.com/demo/internal/usecase/order"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	ginEngine *gin.Engine
	gormDB    *gorm.DB
	rabbitmq  *mq.RabbitMQ
}

func NewServer(ginEngine *gin.Engine, gormDB *gorm.DB, rabbitMQ *mq.RabbitMQ) *Server {
	return &Server{
		ginEngine: ginEngine,
		gormDB:    gormDB,
		rabbitmq:  rabbitMQ,
	}
}

func (s *Server) Route(db *gorm.DB) {
	r := s.ginEngine
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	sagaRepo := repo.NewSagaRepo(db)
	orderProxyService := proxy.NewOrderProxyService(s.rabbitmq)
	inventoryProxyService := proxy.NewInventoryProxyService(s.rabbitmq)
	orderUseCase := order.NewCreateUseCase(orderProxyService, inventoryProxyService, sagaRepo)
	orderHandler := handler.NewCreateOrderHandler(orderUseCase)
	r.POST("/order", orderHandler.Exec)
}

func (s *Server) Run() {
	s.ginEngine.Run()
}
