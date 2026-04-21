package cmd

import (
	"github.com/go-playground/validator/v10"
	"github.com/labib0x9/ProjectUnsafe/config"
	"github.com/labib0x9/ProjectUnsafe/infra/cache/redis"
	"github.com/labib0x9/ProjectUnsafe/infra/db/postgres"
	"github.com/labib0x9/ProjectUnsafe/repo"
	"github.com/labib0x9/ProjectUnsafe/rest"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/admin"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/auth"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/user"
	"github.com/labib0x9/ProjectUnsafe/rest/middleware"
	"github.com/labib0x9/ProjectUnsafe/utils/mailer"
)

func Serve() {
	cnf := config.GetConfig()

	postgresConn := postgres.New()
	dbConn := postgresConn.SetupAndConnection(cnf.DBConfig)
	defer dbConn.Close()

	redisClient := redis.Setup(cnf.RedisConfig)
	defer redisClient.Close()

	authRepo := repo.NewAuthRepository(dbConn, redisClient)
	adminRepo := repo.NewAdminRepository(dbConn)
	userRepo := repo.NewUserRepository(dbConn)

	middlewares := middleware.NewMiddlewares(cnf, authRepo)
	validate := validator.New()
	mailer := mailer.NewMailer(cnf)

	authHandler := auth.NewHandler(authRepo, middlewares, validate, mailer)
	adminHandler := admin.NewHandler(adminRepo, middlewares)
	userHandler := user.NewHandler(userRepo, middlewares)

	server := rest.NewServer(
		authHandler,
		adminHandler,
		userHandler,
	)

	server.Start(redisClient, cnf)
}
