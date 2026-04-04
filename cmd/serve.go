package cmd

import (
	"github.com/labib0x9/ProjectUnsafe/config"
	"github.com/labib0x9/ProjectUnsafe/infra/db/postgres"
	"github.com/labib0x9/ProjectUnsafe/repo"
	"github.com/labib0x9/ProjectUnsafe/rest"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/admin"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/auth"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/lab"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/user"
	"github.com/labib0x9/ProjectUnsafe/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	dbConn := postgres.New().SetupAndConnection(cnf.DBConfig)
	defer dbConn.Close()

	labRepo := repo.NewLabRepo(dbConn)
	authRepo := repo.NewAuthRepository(dbConn)
	adminRepo := repo.NewAdminRepository(dbConn)
	userRepo := repo.NewUserRepository(dbConn)

	middlewares := middleware.NewMiddlewares(cnf)

	labHandler := lab.NewHandler(labRepo, middlewares)
	authHandler := auth.NewHandler(authRepo, middlewares)
	adminHandler := admin.NewHandler(adminRepo)
	userHandler := user.NewHandler(userRepo)

	server := rest.NewServer(
		labHandler,
		authHandler,
		adminHandler,
		userHandler,
	)

	server.Start(cnf)
}
