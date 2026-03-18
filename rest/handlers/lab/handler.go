package lab

import "github.com/labib0x9/ProjectUnsafe/rest/middleware"

type Handler struct {
	middlewares middleware.Middlewares
}

func NewHandler() *Handler {
	return &Handler{}
}
