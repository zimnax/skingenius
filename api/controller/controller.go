package controller

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"skingenius/database"
	dbmodel "skingenius/database/model"
	"skingenius/engine"
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

	userAnswers := model.QuizAnswers{}
	if err := ctx.BodyParser(&userAnswers); err != nil {
		logger.New().Error(ctx.Context(), packageLogPrefix+
			fmt.Sprintf("failed to unmarshall userAnswers req, err: %+v", userAnswers))
		return ctx.SendString(fmt.Sprintf("failed to unmarshall userAnswers req, err: %v", err))
	}

	logger.New().Info(ctx.Context(), packageLogPrefix+fmt.Sprintf("userAnswers: %+v", userAnswers))

	a1SkinType := dbmodel.SkinTypeMapping[userAnswers.SkinType]
	a2SkinSensitivity := dbmodel.SensitivityMapping[userAnswers.SkinSensitivity]
	a3Acne := dbmodel.AcneProneMapping[userAnswers.AcneBreakouts]
	a4Age := dbmodel.AgeMapping[userAnswers.Age]

	var a5Preference []string
	for _, preference := range userAnswers.ProductPreferences {
		a5Preference = append(a5Preference, string(dbmodel.PreferenceMapping[preference]))
	}

	var a6Allergy []string

	for _, allergen := range userAnswers.FreeFromAllergens {
		a6Allergy = append(a6Allergy, string(dbmodel.AllergiesMapping[allergen]))
	}

	var a7Concerns []string
	for _, concern := range userAnswers.SkinConcern {
		a7Concerns = append(a7Concerns, string(dbmodel.SkinConcernsMapping[concern]))
	}

	var a8Benefits []string
	for _, benefit := range userAnswers.ProductBenefit {
		a8Benefits = append(a8Benefits, string(dbmodel.BenefitsMapping[benefit]))
	}

	fmt.Println("\n\n ********************  Answers  ********************")
	fmt.Println("Skin type: ", a1SkinType)
	fmt.Println("Sensitivity: ", a2SkinSensitivity)
	fmt.Println("Acne: ", a3Acne)
	fmt.Println("Age: ", a4Age)
	fmt.Println("Preference: ", a5Preference)
	fmt.Println("Allergy: ", a6Allergy)
	fmt.Println("Concerns: ", a7Concerns)
	fmt.Println("Benefits: ", a8Benefits)
	fmt.Println("********************  Answers  ******************** \n\n ")

	top3 := engine.FindBestProducts_RatingStrategy(gc.geniusData, ctx.Context(),
		a1SkinType, a2SkinSensitivity, a3Acne, a5Preference,
		a6Allergy, a7Concerns, a4Age, a8Benefits)

	fmt.Println(fmt.Sprintf("top 3: %#v", len(top3)))

	for _, topP := range top3 {
		topP.Ingredients = nil
	}

	err := ctx.JSON(top3)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to marshall top3 req, err: %+v", err))
	}

	return err
}

func (gc *GeniusController) SaveRecommendation(ctx *fiber.Ctx) error {
	logger.New().Info(ctx.Context(), packageLogPrefix+"SaveRecommendation route")

	userId := ctx.Params("id")

	fmt.Println(fmt.Sprintf("userID: %s", userId))
	fmt.Println(fmt.Sprintf("req body: %s", string(ctx.Body())))

	//userAnswers := model.QuizAnswers{}
	//if err := ctx.BodyParser(&userAnswers); err != nil {
	//	logger.New().Error(ctx.Context(), packageLogPrefix+
	//		fmt.Sprintf("failed to unmarshall userAnswers req, err: %+v", userAnswers))
	//	return ctx.SendString(fmt.Sprintf("failed to unmarshall userAnswers req, err: %v", err))
	//}

	return nil

}
