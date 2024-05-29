package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"skingenius/config"
	"skingenius/controller"
	"skingenius/database"
	"skingenius/middleware"
	"skingenius/routes"
	"skingenius/utils"
)

func main() {
	db, err := database.NewClient(config.Host, config.Port, config.User, config.Password)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establich db connection, error: %v", err))
		os.Exit(1)
	}

	quppaController, err := controller.NewGeniusController(db)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to create genius controller instance: %v", err))
		os.Exit(1)
	}

	app := fiber.New()
	middleware.FiberMiddleware(app)

	routes.GeniusRoutes(app, quppaController)
	routes.NotFoundRoute(app)

	utils.StartServerWithGracefulShutdown(app)
}
