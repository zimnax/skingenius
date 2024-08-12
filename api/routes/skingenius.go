package routes

import (
	"github.com/gofiber/fiber/v2"
	"skingenius/controller"
)

func GeniusRoutes(a *fiber.App, gc *controller.GeniusController) {
	route := a.Group("/api/v1/skingenius")

	route.Post("/submitQuiz", gc.SubmitQuiz)
	route.Post("/user/:id/recommendation", gc.SaveRecommendation)
	route.Get("/user/:id/recommendation", gc.GetRecommendation)
	route.Post("/user/:id/quiz", gc.SaveQuiz)
	route.Get("/user/:id/quiz", gc.GetQuiz)

}
