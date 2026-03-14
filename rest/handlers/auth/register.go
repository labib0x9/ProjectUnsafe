package auth

import (
	"net/http"

	middleware "github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// mux.Handle("POST /auth/signup", http.HandlerFunc())
	// mux.Handle("POST /auth/login", http.HandlerFunc())
	// mux.Handle("POST /auth/reset-password", http.HandlerFunc())
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

	// Admin APIs
	// mux.Handle("GET /admin/users", http.HandlerFunc())
	// mux.Handle("GET /admin/containers", http.HandlerFunc())
	// mux.Handle("POST /admin/terminate", http.HandlerFunc())
}
