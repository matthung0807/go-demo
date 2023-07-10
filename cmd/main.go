package main

import (
	"context"
	"fmt"

	"abc.com/demo/internal/event/router"
	"abc.com/demo/internal/infra/mq"
	"abc.com/demo/server"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	_ "abc.com/demo/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Demo
// @version 1.0
// @description Swagger API.
// @host localhost:8080
// @contact.name matthung0807
// @contact.email matthung0807blog@gmail.com
func main() {
	ctx := context.Background()
	gormDB, err := initGormDB()
	if err != nil {
		panic(err)
	}

	mqConn, err := initMQ()
	if err != nil {
		panic(err)
	}

	r := gin.Default() // get gin engine
	rabbitmq := mq.NewRabbitMQ(mqConn)
	s := server.NewServer(r, gormDB, rabbitmq)

	initSwaggerDoc(r)

	router.InitServerEventRouter(ctx, gormDB, rabbitmq)
	s.Route(gormDB)
	s.Run()
}

const (
	HOST     = "localhost"
	PORT     = "5432"
	DATABASE = "postgres"
	USER     = "admin"
	PASSWORD = "12345"
	SSL      = "disable"

	MQ_URL = "amqp://guest:guest@localhost:5672/"
)

// postgresql
func initGormDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		HOST, PORT, USER, PASSWORD, DATABASE, SSL)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}

// rabbitmq
func initMQ() (*amqp.Connection, error) {
	mqConn, err := amqp.Dial(MQ_URL)
	if err != nil {
		return nil, err
	}
	return mqConn, nil
}

func initSwaggerDoc(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DocExpansion("list")))
}
