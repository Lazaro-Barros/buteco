package router

import (
	"github.com/Lazaro-Barros/buteco/command/infra/drivers/http/handler"
	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	engine *gin.Engine
}

func NewGinRouter() Router {
	return &GinRouter{
		engine: gin.Default(),
	}
}

func (r *GinRouter) GET(path string, handler func(ctx handler.Context)) {
	r.engine.GET(path, func(c *gin.Context) {
		handler(c)
	})
}

func (r *GinRouter) POST(path string, handler func(ctx handler.Context)) {
	r.engine.POST(path, func(c *gin.Context) {
		handler(c)
	})
}

func (r *GinRouter) PUT(path string, handler func(ctx handler.Context)) {
	r.engine.PUT(path, func(c *gin.Context) {
		handler(c)
	})
}

func (r *GinRouter) DELETE(path string, handler func(ctx handler.Context)) {
	r.engine.DELETE(path, func(c *gin.Context) {
		handler(c)
	})
}

func (r *GinRouter) Run(addr string) error {
	return r.engine.Run(addr)
}
