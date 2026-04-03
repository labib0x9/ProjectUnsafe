package cmd

import (
	"github.com/labib0x9/ProjectUnsafe/config"
	"github.com/labib0x9/ProjectUnsafe/infra/db"
	"github.com/labib0x9/ProjectUnsafe/repo"
	"github.com/labib0x9/ProjectUnsafe/rest"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/auth"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/lab"
)

func Serve() {
	cnf := config.GetConfig()

	dbConn, err := db.NewConnection(cnf)
	if err != nil {
		panic(err)
	}

	labRepo := repo.NewLabRepo(dbConn)
	authRepo := repo.NewAuthRepository(dbConn)

	labHandler := lab.NewHandler(labRepo)
	authHandler := auth.NewHandler(authRepo)

	server := rest.NewServer(
		labHandler,
		authHandler,
	)

	server.Start(cnf)
}
