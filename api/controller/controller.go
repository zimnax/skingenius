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

	iBySkinTypeList, err1 := gc.geniusData.IngredientBySkinType(ctx.UserContext(), qa.SkinType)
	fmt.Println(fmt.Sprintf("IngredientBySkinType: [%d], error: %v", len(iBySkinTypeList), err1))

	iBySkinSensitivityList, err2 := gc.geniusData.IngredientBySkinSensitivity(ctx.UserContext(), qa.SkinReact_Sensitivity)
	fmt.Println(fmt.Sprintf("iBySkinSensitivityList: [%d], error: %v", len(iBySkinSensitivityList), err2))

	iBySkinAcneList, err3 := gc.geniusData.IngredientByAcne(ctx.UserContext(), qa.AcneBreakouts)
	fmt.Println(fmt.Sprintf("iBySkinAcneList: [%d], error: %v", len(iBySkinAcneList), err3))

	return nil
}
