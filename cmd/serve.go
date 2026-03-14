package cmd

import (
	"github.com/labib0x9/ProjectUnsafe/config"
	"github.com/labib0x9/ProjectUnsafe/rest"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/auth"
	"github.com/labib0x9/ProjectUnsafe/rest/handlers/lab"
)

func Serve() {
	cnf := config.GetConfig()

	labHandler := lab.NewHandler()
	authHandler := auth.NewHandler()

	server := rest.NewServer(
		labHandler,
		authHandler,
	)

	server.Start(cnf)
}
