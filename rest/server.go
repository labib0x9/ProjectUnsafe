package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/config"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/auth"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/lab"
	middleware "github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

type Server struct {
	LabHandler  *lab.Handler
	AuthHandler *auth.Handler
	dbConn      *sqlx.DB
}

func NewServer(
	LabHandler *lab.Handler,
	AuthHandler *auth.Handler,
) *Server {
	return &Server{
		LabHandler:  LabHandler,
		AuthHandler: AuthHandler,
	}
}

func (s *Server) Start(cnf *config.Config) {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	s.AuthHandler.RegisterRoutes(mux, manager)
	s.LabHandler.RegisterRoutes(mux, manager)

	fmt.Printf("Starting Server at http://127.0.0.1:%d/\n", cnf.Port)
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
