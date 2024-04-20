package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/leonardonicola/tickethub/config"
	_ "github.com/leonardonicola/tickethub/docs"
	"github.com/leonardonicola/tickethub/pkg/router"
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
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("env variables: %v", err)
	}
	config.InitDB()
	r, err := router.InitRoutes()
	// Load godot to retrieve variables from .env
	if err != nil {
		log.Fatalf("%v", err)
	}
	if err := r.Run(":3000"); err != nil {
		panic(err.Error())
	}
}
