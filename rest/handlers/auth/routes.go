package auth

import (
	"net/http"

	middleware "github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

func (h *Handler) RegisterRoutes(
	mux *http.ServeMux,
	manager *middleware.Manager,
) {
	mux.Handle(
		"POST /auth/signup",
		manager.With(
			http.HandlerFunc(h.Signup),
			h.middlewares.MaxBody1MB,
		),
	)

	mux.Handle(
		"GET /auth/verify",
		manager.With(
			http.HandlerFunc(h.Verify),
		),
	)

	mux.Handle(
		"POST /auth/verify/resend",
		manager.With(
			http.HandlerFunc(h.ResendVerify),
			h.middlewares.MaxBody1MB,
			h.middlewares.OneTimePerEmail,
			h.middlewares.BlockIP,
		),
	)

	mux.Handle(
		"POST /auth/login",
		manager.With(
			http.HandlerFunc(h.Login),
			h.middlewares.MaxBody1MB,
		),
	)

	mux.Handle(
		"POST /auth/reset-password",
		manager.With(
			http.HandlerFunc(h.Reset),
		),
	)

	mux.Handle(
		"GET /auth/logout",
		manager.With(
			http.HandlerFunc(h.Logout),
			h.middlewares.Auth,
		),
	)
}
