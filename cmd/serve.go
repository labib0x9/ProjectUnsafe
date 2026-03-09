package cmd

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/labib0x9/ProjectUnsafe/global_router"
	"github.com/labib0x9/ProjectUnsafe/handlers"
	"github.com/labib0x9/ProjectUnsafe/model"
)

func init() {
	lab1 := model.Lab{
		Id:          "lab-ssh-basics",
		Title:       "SSH Basics",
		Difficulty:  "easy",
		Category:    "networking",
		Description: "Learn to connect to remote machines, manage keys, and tunnel ports using SSH.",
		LongDescription: `In this lab you will practice the fundamentals of SSH. You will: 
- Connect to a remote host using password and key-based auth
- Generate and manage SSH key pairs
- Use SSH tunneling to forward local ports
- Explore the ~/.ssh/config file

The container runs a minimal Debian image with sshd configured and ready.`,
		Hints: []string{
			"Use ssh-keygen -t ed25519 to generate a modern key pair",
			"Copy your public key with ssh-copy-id user@host",
			"Check /etc/ssh/sshd_config for server settings",
		},
		Completions:   0,
		EstimatedTime: "45 min",
		Tags:          []string{"ssh", "networking", "linux"},
	}

	model.LabList = append(model.LabList, lab1)
}

func Server() {

	mux := http.NewServeMux()

	// Lab APIs
	mux.Handle("GET /labs", http.HandlerFunc(handlers.GetAllLabs))
	// mux.Handle("GET /lab/${id}", http.HandlerFunc())
	// mux.Handle("POST /lab/start", http.HandlerFunc())
	// mux.Handle("POST /lab/reset'", http.HandlerFunc())
	// mux.Handle("POST /lab/terminate'", http.HandlerFunc())
	// mux.Handle("GET /lab/hints?labId=${labId}", http.HandlerFunc())

	// Auth APIs
	// mux.Handle("POST /auth/login", http.HandlerFunc())
	// mux.Handle("POST /auth/signup", http.HandlerFunc())
	mux.Handle("POST /auth/anonymous", http.HandlerFunc(handlers.AnonLogin))
	// mux.Handle("POST /auth/reset-password", http.HandlerFunc())

	// Playground APIs
	// mux.Handle("GET /problems", http.HandlerFunc())
	// mux.Handle("GET /problem/${id}", http.HandlerFunc())
	// mux.Handle("POST /code/run-custom", http.HandlerFunc())
	// mux.Handle("POST /code/run", http.HandlerFunc())

	// Admin APIs
	// mux.Handle("GET /admin/containers", http.HandlerFunc())
	// mux.Handle("GET /admin/users", http.HandlerFunc())
	// mux.Handle("POST /admin/terminate", http.HandlerFunc())

	globalRouter := global_router.NewGlobalRouter(mux)

	slog.Info("Starting Server at http://127.0.0.1:8080/")
	log.Fatal(http.ListenAndServe(":8080", globalRouter))
}
