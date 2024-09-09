package router

import "github.com/Lazaro-Barros/buteco/command/infra/drivers/http/handler"

type Router interface {
	Run(addr string) error
	GET(path string, handler func(ctx handler.Context))
	POST(path string, handler func(ctx handler.Context))
	PUT(path string, handler func(ctx handler.Context))
	DELETE(path string, handler func(ctx handler.Context))
}
