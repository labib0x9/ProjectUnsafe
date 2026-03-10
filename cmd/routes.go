package cmd

import (
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/handlers"
	"github.com/labib0x9/ProjectUnsafe/middleware"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// Lab APIs
	mux.Handle(
		"GET /labs",
		manager.With(
			http.HandlerFunc(handlers.GetAllLabs),
		),
	)

	mux.Handle(
		"GET /lab/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetLabByID),
		),
	)
	mux.Handle(
		"POST /lab/start",
		manager.With(
			http.HandlerFunc(handlers.StartLab),
		),
	)
	mux.Handle(
		"POST /lab/reset'",
		manager.With(
			http.HandlerFunc(handlers.ResetLab),
		),
	)
	mux.Handle(
		"POST /lab/terminate'",
		manager.With(
			http.HandlerFunc(handlers.TerminateLab),
		),
	)

	// Auth APIs
	// mux.Handle("POST /auth/signup", http.HandlerFunc())
	// mux.Handle("POST /auth/login", http.HandlerFunc())
	// mux.Handle("POST /auth/reset-password", http.HandlerFunc())
	mux.Handle(
		"POST /auth/anonymous",
		manager.With(
			http.HandlerFunc(handlers.AnonLogin),
		),
	)
	mux.Handle(
		"POST /auth/logout",
		manager.With(
			http.HandlerFunc(handlers.Logout),
		),
	)

	// Playground APIs
	// mux.Handle("GET /problems", http.HandlerFunc())
	// mux.Handle("GET /problem/{id}", http.HandlerFunc())
	// mux.Handle("POST /code/run-custom", http.HandlerFunc())
	// mux.Handle("POST /code/run", http.HandlerFunc())

	// Admin APIs
	// mux.Handle("GET /admin/users", http.HandlerFunc())
	// mux.Handle("GET /admin/containers", http.HandlerFunc())
	// mux.Handle("POST /admin/terminate", http.HandlerFunc())
}
