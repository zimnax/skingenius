package controller

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
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

func (gc *GeniusController) SubmitQuizV2(ctx *fiber.Ctx) error {
	logger.New().Info(ctx.Context(), packageLogPrefix+"SubmitQuizV2 route")

	userAnswers := model.QuizAnswers{}
	if err := ctx.BodyParser(&userAnswers); err != nil {
		logger.New().Error(ctx.Context(), packageLogPrefix+
			fmt.Sprintf("failed to unmarshall userAnswers req, err: %+v", err))
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

	fmt.Println("\n\n ********************  Answers V2 ********************")
	fmt.Println("Skin type: ", a1SkinType)
	fmt.Println("Sensitivity: ", a2SkinSensitivity)
	fmt.Println("Acne: ", a3Acne)
	fmt.Println("Age: ", a4Age)
	fmt.Println("Preference: ", a5Preference)
	fmt.Println("Allergy: ", a6Allergy)
	fmt.Println("Concerns: ", a7Concerns)
	fmt.Println("Benefits: ", a8Benefits)
	fmt.Println("********************  Answers  ******************** \n\n ")

	//top3 := engine.FindBestProducts_RatingStrategy(gc.geniusData, ctx.Context(),
	//	a1SkinType, a2SkinSensitivity, a3Acne, a5Preference,
	//	a6Allergy, a7Concerns, a4Age, a8Benefits)
	top3 := engine.FindBestProducts_matchBestStrategy(gc.geniusData, ctx.Context(),
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

func (gc *GeniusController) SubmitQuiz(ctx *fiber.Ctx) error {
	logger.New().Info(ctx.Context(), packageLogPrefix+"SubmitQuiz route")

	userAnswers := model.QuizAnswers{}
	if err := ctx.BodyParser(&userAnswers); err != nil {
		logger.New().Error(ctx.Context(), packageLogPrefix+
			fmt.Sprintf("failed to unmarshall userAnswers req, err: %+v", err))
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

	recommendedProducts := model.SaveRecommendationsReq{}
	if err := ctx.BodyParser(&recommendedProducts); err != nil {
		logger.New().Error(ctx.Context(), packageLogPrefix+
			fmt.Sprintf("failed to unmarshall saveRecommendations req, err: %+v", err))
		return ctx.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("failed to unmarshall saveRecommendations req, err: %v", err))
	}

	var rs []dbmodel.UserRecommendations
	for _, p := range recommendedProducts.Products {
		rs = append(rs, dbmodel.UserRecommendations{
			UserId:    userId,
			ProductId: p.Id,
			Score:     p.Score,
		})
	}

	err := gc.geniusData.SaveRecommendations(ctx.Context(), rs)
	if err != nil {
		logger.New().Error(ctx.Context(), packageLogPrefix+
			fmt.Sprintf("failed to save user recommendations, err: %+v", err))
		return ctx.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("failed to save user recommendations, err: %v", err))
	}
	return ctx.Status(http.StatusCreated).JSON(nil)
}

func (gc *GeniusController) GetRecommendation(ctx *fiber.Ctx) error {
	logger.New().Info(ctx.Context(), packageLogPrefix+"GetRecommendation route")

	userId := ctx.Params("id")
	fmt.Println(fmt.Sprintf("userID: %s", userId))

	savedRecommendation, err := gc.geniusData.GetRecommendations(ctx.Context(), userId)
	if err != nil {
		logger.New().Error(ctx.Context(), packageLogPrefix+
			fmt.Sprintf("failed to get user recommendations, err: %+v", err))
		return ctx.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("failed to get user recommendations, err: %v", err))
	}

	var recIds []int32
	for _, recommendations := range savedRecommendation {
		recIds = append(recIds, int32(recommendations.ProductId))
	}

	fmt.Println(fmt.Sprintf("product ids %v by userId %s", recIds, userId))
	fullProducts, err := gc.geniusData.FindProductsByIds(ctx.Context(), recIds)
	if err != nil {
		logger.New().Error(ctx.Context(), packageLogPrefix+
			fmt.Sprintf("failed to get full Products recommendations, err: %+v", err))
		return ctx.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("failed to get full product recommendations, err: %v", err))
	}

	for _, product := range fullProducts {
		for i, recommendations := range savedRecommendation {
			if product.ID == uint(recommendations.ProductId) {
				product.Score = savedRecommendation[i].Score
			}
		}
	}

	userQuiz, err := gc.geniusData.GetQuiz(ctx.Context(), userId)
	if err != nil {
		logger.New().Error(ctx.Context(), packageLogPrefix+
			fmt.Sprintf("failed to get user quiz, err: %+v", err))
		return ctx.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("failed to get user quiz, err: %v", err))
	}
	logger.New().Info(ctx.Context(), packageLogPrefix+fmt.Sprintf("found skin quiz %v for user %s", userQuiz, userId))

	if userQuiz.SkinConcern[0] != "" {
		logger.New().Info(ctx.Context(), packageLogPrefix+fmt.Sprintf("fetching Ingredients description for user skin concern: %s", userQuiz.SkinConcern[0]))

		for _, product := range fullProducts {
			desc, fetchDescErr := gc.geniusData.GetSkinConcernDescriptionByIngredients(ctx.Context(), database.GetIngredientsName(product.Ingredients), userQuiz.SkinConcern[0])
			if fetchDescErr != nil {
				logger.New().Error(ctx.Context(), packageLogPrefix+
					fmt.Sprintf("failed to get skin concern description, fetchDescErr: %+v", fetchDescErr))
				return ctx.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("failed to get skin concern description, err: %v", fetchDescErr))
			}

			logger.New().Info(ctx.Context(), packageLogPrefix+fmt.Sprintf("found %d ingredient description for concern %s", len(desc), userQuiz.SkinConcern[0]))

			for _, description := range desc {
				for _, ingredient := range product.Ingredients {
					if ingredient.Name == description.Ingredientname {
						ingredient.ConcernDescription = description.Description
					}
				}
			}
		}
	}

	fmt.Println(fmt.Sprintf("full products len %d by userId %s", len(fullProducts), userId))
	for _, product := range fullProducts {
		fmt.Println(fmt.Sprintf("full product %s with scores %f by userId %s", product.Name, product.Score, userId))
	}

	return ctx.Status(http.StatusOK).JSON(fullProducts)
}

func (gc *GeniusController) SaveQuiz(ctx *fiber.Ctx) error {
	logger.New().Info(ctx.Context(), packageLogPrefix+"SaveQuiz route")

	userId := ctx.Params("id")
	fmt.Println(fmt.Sprintf("userID: %s", userId))

	userAnswers := model.QuizAnswers{}
	if err := ctx.BodyParser(&userAnswers); err != nil {
		logger.New().Error(ctx.Context(), packageLogPrefix+
			fmt.Sprintf("failed to unmarshall save quiz req, err: %+v", err))
		return ctx.SendString(fmt.Sprintf("failed to unmarshall save quiz req, err: %v", err))
	}

	logger.New().Info(ctx.Context(), packageLogPrefix+fmt.Sprintf("userAnswers: %+v", userAnswers))

	// todo: extract into func -----------------------------------------------------------------------------------------
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
	//todo: ------------------------------------------------------------------------------------------------------------

	err := gc.geniusData.SaveQuiz(ctx.Context(), dbmodel.UserQuiz{
		UserId:             userId,
		SkinType:           a1SkinType,
		SkinSensitivity:    a2SkinSensitivity,
		AcneBreakouts:      a3Acne,
		ProductPreferences: a5Preference,
		FreeFromAllergens:  a6Allergy,
		SkinConcern:        a7Concerns,
		Age:                a4Age,
		ProductBenefit:     a8Benefits,
	})

	if err != nil {
		logger.New().Error(ctx.Context(), packageLogPrefix+
			fmt.Sprintf("failed to save user quiz, err: %+v", err))
		return ctx.Status(http.StatusInternalServerError).SendString("failed to save quiz")
	}

	return ctx.Status(http.StatusAccepted).JSON(nil)
}

func (gc *GeniusController) GetQuiz(ctx *fiber.Ctx) error {
	logger.New().Info(ctx.Context(), packageLogPrefix+"GetQuiz route")

	userId := ctx.Params("id")
	fmt.Println(fmt.Sprintf("userID: %s", userId))

	quiz, err := gc.geniusData.GetQuiz(ctx.Context(), userId)
	if err != nil {
		logger.New().Error(ctx.Context(), packageLogPrefix+
			fmt.Sprintf("failed to get quiz, err: %+v", err))
		return ctx.Status(http.StatusInternalServerError).SendString("failed to get quiz")
	}

	logger.New().Info(ctx.Context(), fmt.Sprintf(packageLogPrefix+"Return quiz: %#v", quiz))
	return ctx.Status(http.StatusAccepted).JSON(quiz)
}
