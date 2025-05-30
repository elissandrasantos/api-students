package main

// @title Student API
// @version 1.0
// @description This is a sample server Student API. You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@yourdomain.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
import (
	"github.com/elissandrasantos/api-students/api"
	"github.com/rs/zerolog/log"

	_ "github.com/elissandrasantos/api-students/docs" // Mantenha esta!
)

func main() {
	server := api.NewServer()
	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
