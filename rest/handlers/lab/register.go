package lab

import (
	"net/http"

	middleware "github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /labs",
		manager.With(
			http.HandlerFunc(h.GetAllLabs),
		),
	)

	mux.Handle(
		"GET /lab/{id}",
		manager.With(
			http.HandlerFunc(h.GetLabByID),
		),
	)

	mux.Handle(
		"POST /lab/create",
		manager.With(
			http.HandlerFunc(h.CreateLab),
			h.middlewares.Auth,
		),
	)

	mux.Handle(
		"POST /lab/delete",
		manager.With(
			http.HandlerFunc(h.GetLabByID),
			h.middlewares.Auth,
		),
	)

	mux.Handle(
		"POST /lab/start",
		manager.With(
			http.HandlerFunc(h.StartLab),
		),
	)

	mux.Handle(
		"POST /lab/reset'",
		manager.With(
			http.HandlerFunc(h.ResetLab),
		),
	)

	mux.Handle(
		"POST /lab/terminate'",
		manager.With(
			http.HandlerFunc(h.TerminateLab),
		),
	)
}
