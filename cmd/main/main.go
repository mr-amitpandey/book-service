package main

import (
	"log"

	"github.com/book-service/api/app/api"
	"github.com/book-service/api/cmd/setup"
)

func main() {
	log.Println(" Starting Book Service...")

	// App setup (configuration, DB, etc.)
	router, configData := setup.AppSetup()
	log.Println("✅ Setup completed")

	router.Static("/files", "./public")

	// Initialize all API routes
	MyRouters := api.Routers{
		Router: router,
	}
	MyRouters.Init()
	log.Println("✅ Routes initialized")

	// Start HTTP server
	serverAddress := configData.App.Host + ":" + configData.App.Port
	log.Printf(" Starting Book Service on %s", serverAddress)

	// Add this debug line
	log.Printf(" Server will listen on: %s", serverAddress)

	if err := router.Run(serverAddress); err != nil {
		log.Fatal("❌ Unable to start the server:", err)
	}

	// This line should never be reached if server starts successfully
	log.Println("❌ Server stopped unexpectedly")
}
