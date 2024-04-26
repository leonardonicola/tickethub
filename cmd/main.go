package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/leonardonicola/tickethub/config"
	_ "github.com/leonardonicola/tickethub/docs"
	"github.com/leonardonicola/tickethub/internal/pkg/router"
)

//	@title			Tickethub
//	@version		1.0
//	@description	Serviço de ingressaria e tickets.

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/api/v1

//	@securityDefinitions.apikey JWT
//  @in header
//  @name Authorization

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
var (
	logger *config.Logger
)

func main() {
	logger = config.NewLogger()
	if err := godotenv.Load(".env"); err != nil {
		logger.Fatalf("env variables: %v", err)
		os.Exit(1)
	}
	// TODO: return error on intializations and give os.Exit(1)
	config.InitDB()
	config.InitS3Client()
	r, err := router.InitRoutes()
	// Load godot to retrieve variables from .env
	if err != nil {
		logger.Fatalf("ROUTER error: %v", err)
		os.Exit(1)
	}
	if err := r.Run(":3000"); err != nil {
		logger.Fatalf("HTTP ROUTER: %v", err)
		os.Exit(1)
	}
}
