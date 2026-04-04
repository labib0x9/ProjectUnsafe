package admin

import (
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /admin/users",
		manager.With(
			http.HandlerFunc(h.ListUsers),
			h.middlewares.Auth,
			h.middlewares.Admin,
		),
	)

	mux.Handle(
		"GET /admin/containers",
		manager.With(
			http.HandlerFunc(h.list_container),
			h.middlewares.Auth,
			h.middlewares.Admin,
		),
	)

	mux.Handle(
		"POST /admin/terminate",
		manager.With(
			http.HandlerFunc(h.terminate_container),
			h.middlewares.Auth,
			h.middlewares.Admin,
		),
	)
}
