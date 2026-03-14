package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/config"
	"github.com/labib0x9/ProjectUnsafe/middleware"
)

func Server() {

	cnf := config.GetConfig()

	manager := middleware.NewManager()
	manager.Use(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Printf("Starting Server at http://127.0.0.1:%d/\n", cnf.Port)
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
