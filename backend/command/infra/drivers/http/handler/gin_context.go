package handler

import (
	"github.com/gin-gonic/gin"
)

type GinContext struct {
	C *gin.Context
}

func (g *GinContext) JSON(code int, obj interface{}) {
	g.C.JSON(code, obj)
}

func (g *GinContext) Bind(obj interface{}) error {
	return g.C.Bind(obj)
}

func (g *GinContext) ShouldBindJSON(obj interface{}) error {
	return g.C.ShouldBindJSON(obj)
}

func (g *GinContext) Param(key string) string {
	return g.C.Param(key)
}

func (g *GinContext) Query(key string) string {
	return g.C.Query(key)
}

func (g *GinContext) Status(code int) {
	g.C.Status(code)
}
