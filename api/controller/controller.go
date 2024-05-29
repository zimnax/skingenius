package controller

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"skingenius/database"
	"skingenius/logger"
	"skingenius/model"
)

const packageLogPrefix = "genius_controller:"

type GeniusController struct {
	geniusData database.Connector
}

func NewGeniusController(db database.Connector) (*GeniusController, error) {
	logger.New().Info(context.Background(), packageLogPrefix+"initializing new genius controller")

	return &GeniusController{geniusData: db}, nil
}

func (gc *GeniusController) SubmitQuiz(ctx *fiber.Ctx) error {
	logger.New().Info(ctx.Context(), packageLogPrefix+"SubmitQuiz route")

	qa := model.QuizAnswers{}
	if err := ctx.BodyParser(&qa); err != nil {
		logger.New().Error(ctx.Context(), packageLogPrefix+
			fmt.Sprintf("failed to unmarshall quizAnswers req, err: %+v", qa))
		return ctx.SendString(fmt.Sprintf("failed to unmarshall quizAnswers req, err: %v", err))
	}

	logger.New().Info(ctx.Context(), packageLogPrefix+fmt.Sprintf("quiz request: %+v", qa))

	return nil
}
