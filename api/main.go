package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"skingenius/api"
	"skingenius/config"
	"skingenius/controller"
	"skingenius/database"
	"skingenius/handler"
	"skingenius/middleware"
	"skingenius/utils"
)

/*
build on windows -  $Env:GOOS = "linux"; $Env:GOARCH = "amd64"; go build -o skingv1 .\main.go
nohup ./skingv7 &
*/
func main() {
	db, err := database.NewGormClient(config.LocalHost, config.Port, config.User, config.Password, false, "prod")
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	geniusController := controller.NewGeniusController(db)
	geniusHandler := handler.NewGeniusHandler(geniusController)

	app := fiber.New()
	middleware.FiberMiddleware(app)

	api.GeniusRoutes(app, geniusHandler)
	api.NotFoundRoute(app)

	utils.StartServerWithGracefulShutdown(app)
}
