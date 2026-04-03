package lab

import (
	"net/http"

	middleware "github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /labs",
		manager.With(
			http.HandlerFunc(h.List),
		),
	)

	mux.Handle(
		"GET /labs/{id}",
		manager.With(
			http.HandlerFunc(h.GetLabID),
			h.middlewares.Auth,
		),
	)

	mux.Handle(
		"POST /labs/create",
		manager.With(
			http.HandlerFunc(h.Create),
			h.middlewares.Auth,
			h.middlewares.Admin,
		),
	)

	mux.Handle(
		"POST /labs/update",
		manager.With(
			http.HandlerFunc(h.Update),
			h.middlewares.Auth,
			h.middlewares.Admin,
		),
	)

	mux.Handle(
		"POST /labs/delete",
		manager.With(
			http.HandlerFunc(h.Delete),
			h.middlewares.Auth,
			h.middlewares.Admin,
		),
	)

	mux.Handle(
		"POST /labs/start",
		manager.With(
			http.HandlerFunc(h.Start),
			h.middlewares.Auth,
		),
	)

	mux.Handle(
		"POST /labs/reset'",
		manager.With(
			http.HandlerFunc(h.Reset),
			h.middlewares.Auth,
		),
	)

	mux.Handle(
		"POST /labs/terminate'",
		manager.With(
			http.HandlerFunc(h.Terminate),
			h.middlewares.Auth,
		),
	)
}
