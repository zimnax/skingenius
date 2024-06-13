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
	db, err := database.NewGormClient(config.Host, config.Port, config.User, config.Password)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	geniusController, err := controller.NewGeniusController(db)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to create genius controller instance: %v", err))
		os.Exit(1)
	}

	app := fiber.New()
	middleware.FiberMiddleware(app)

	routes.GeniusRoutes(app, geniusController)
	routes.NotFoundRoute(app)

	utils.StartServerWithGracefulShutdown(app)
}
