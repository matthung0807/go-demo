package main

import (
	"context"
	"log"
	"net/http"

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
}

func AnnotateAsHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Hander)),
		fx.ResultTags(`group:"handlers"`),
	)
}

type Router struct {
	router   *gin.Engine
	handlers []Hander
}

func NewRouter(lc fx.Lifecycle, handlers []Hander) *Router {
	router := gin.Default()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server forced to shutdown: ", err)
			}

			log.Println("Server exiting")
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
