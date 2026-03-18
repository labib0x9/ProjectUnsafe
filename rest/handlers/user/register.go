package user

import (
	"net/http"

	middleware "github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /profiles/{id}",
		manager.With(
			http.HandlerFunc(h.GetProfile),
			h.middlewares.Auth,
		),
	)
}
