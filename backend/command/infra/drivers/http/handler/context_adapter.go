package handler

type Context interface {
	JSON(code int, obj interface{})
	Bind(obj interface{}) error
	ShouldBindJSON(obj interface{}) error
	Param(key string) string
	Query(key string) string
	Status(code int)
}
