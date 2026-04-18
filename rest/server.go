package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labib0x9/ProjectUnsafe/config"
	"github.com/labib0x9/ProjectUnsafe/infra/cache/redis"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/admin"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/auth"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/lab"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/user"
	middleware "github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

type Server struct {
	LabHandler   *lab.Handler
	AuthHandler  *auth.Handler
	AdminHandler *admin.Handler
	UserHandler  *user.Handler
	dbConn       *sqlx.DB
}

func NewServer(
	LabHandler *lab.Handler,
	AuthHandler *auth.Handler,
	AdminHandler *admin.Handler,
	UserHandler *user.Handler,
) *Server {
	return &Server{
		LabHandler:   LabHandler,
		AuthHandler:  AuthHandler,
		AdminHandler: AdminHandler,
		UserHandler:  UserHandler,
	}
}

func (s *Server) Start(cnf *config.Config) {
	redisClient := redis.Setup(cnf.RedisConfig)
	defer redisClient.Close()

	rateLimiter := middleware.NewRateLimiter(redisClient, 2, 1)

	manager := middleware.NewManager()
	manager.Use(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
		rateLimiter.Limit(),
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	s.AuthHandler.RegisterRoutes(mux, manager)
	s.LabHandler.RegisterRoutes(mux, manager)
	s.AdminHandler.RegisterRoutes(mux, manager)
	s.UserHandler.RegisterRoutes(mux, manager)

	fmt.Printf("Starting Server at http://127.0.0.1:%d/\n", cnf.Port)
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
