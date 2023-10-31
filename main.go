package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			AnnotateAsHandler(NewHelloHandler),
			fx.Annotate(
				NewRouter,
				fx.ParamTags(``, `group:"handlers"`),
			),
		),
		fx.Invoke(func(router *Router) {
			router.Route()
		}),
	).Run()

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	router.Run()
}

func AnnotateAsHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Hander)),
		fx.ResultTags(`group:"handler"`),
	)
}

type Router struct {
	router   *gin.Engine
	handlers []Hander
}

func NewRouter(lc fx.Lifecycle, handlers []Hander) *Router {
	router := gin.Default()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go router.Run()
			return nil
		},
	})

	return &Router{
		router:   router,
		handlers: handlers,
	}
}

func (r Router) Route() {
	for _, h := range r.handlers {
		r.router.Handle(h.HttpMethod(), h.Pattern(), h.HandlerFunc)
	}
}

type Hander interface {
	HandlerFunc(c *gin.Context)
	HttpMethod() string
	Pattern() string
}

type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}
func (h *HelloHandler) HandlerFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func (h *HelloHandler) HttpMethod() string {
	return "GET"
}

func (h *HelloHandler) Pattern() string {
	return "/hello"
}
