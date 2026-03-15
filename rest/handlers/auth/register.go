package auth

import (
	"net/http"

	middleware "github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"POST /auth/signup",
		manager.With(
			http.HandlerFunc(h.Signup),
		),
	)

	mux.Handle(
		"POST /auth/login",
		manager.With(
			http.HandlerFunc(h.Login),
		),
	)

	mux.Handle(
		"POST /auth/reset-password",
		manager.With(
			http.HandlerFunc(h.Reset),
		),
	)

	mux.Handle(
		"POST /auth/anonymous",
		manager.With(
			http.HandlerFunc(h.AnonLogin),
		),
	)

	mux.Handle(
		"POST /auth/logout",
		manager.With(
			http.HandlerFunc(h.Logout),
		),
	)
}
