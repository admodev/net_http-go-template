package main

import (
	"github.com/create-go-app/net_http-go-template/pkg/configs"
	"github.com/create-go-app/net_http-go-template/pkg/routes"
	"github.com/create-go-app/net_http-go-template/pkg/utils"
	"github.com/gorilla/mux"

	_ "github.com/create-go-app/net_http-go-template/docs" //
	_ "github.com/joho/godotenv/autoload"                  // load .env file automatically
)

// @title Golang net/http Template API
// @version 1.0
// @description This is an auto-generated API Docs for Golang net/http Template.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Initialize a new router.
	router := mux.NewRouter()

	// List of app routes:
	routes.PublicRoutes(router)
	routes.PrivateRoutes(router)
	routes.SwaggerRoutes(router)

	// Register API routes.
	server := configs.ServerConfig(router)

	// Start API server.
	utils.StartServerWithGracefulShutdown(server)
}
