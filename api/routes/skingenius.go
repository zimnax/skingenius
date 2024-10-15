package routes

import (
	"github.com/gofiber/fiber/v2"
	"skingenius/controller"
)

func GeniusRoutes(a *fiber.App, gc *controller.GeniusController) {
	route := a.Group("/api/v1/skingenius")

	route.Post("/submitQuizV2", gc.SubmitQuizV2)
	route.Post("/user/:id/recommendation", gc.SaveRecommendation)
	route.Get("/user/:id/recommendation", gc.GetRecommendation)

	route.Post("/user/:id/routine", gc.SaveUserRoutine)
	route.Get("/user/:id/routine", gc.GetUserRoutine)

	route.Post("/user/:id/quiz", gc.SaveQuiz)
	route.Get("/user/:id/quiz", gc.GetQuiz)
	route.Get("/search/:request", gc.Search)
}
