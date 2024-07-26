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
		a5Preference = append(a5Preference, preference)
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

	fmt.Println(fmt.Sprintf("top 3: %#v", top3))

	//iBySkinTypeList, err1 := gc.geniusData.IngredientBySkinType(ctx.UserContext(), qa.SkinType)
	//fmt.Println(fmt.Sprintf("IngredientBySkinType: [%d], error: %v", len(iBySkinTypeList), err1))
	//
	//iBySkinSensitivityList, err2 := gc.geniusData.IngredientBySkinSensitivity(ctx.UserContext(), qa.SkinReact_Sensitivity)
	//fmt.Println(fmt.Sprintf("iBySkinSensitivityList: [%d], error: %v", len(iBySkinSensitivityList), err2))
	//
	//iByAcneList, err3 := gc.geniusData.IngredientByAcne(ctx.UserContext(), qa.AcneBreakouts)
	//fmt.Println(fmt.Sprintf("iBySkinAcneList: [%d], error: %v", len(iByAcneList), err3))
	//
	//iByPreferences, err4 := gc.geniusData.IngredientByPreferences(ctx.UserContext(), qa.ProductPreferences)
	//fmt.Println(fmt.Sprintf("iByPreferences: [%d], error: %v", len(iByPreferences), err4))
	//
	//iByAllergens, err5 := gc.geniusData.IngredientByAllergens(ctx.UserContext(), qa.FreeFromAllergens)
	//fmt.Println(fmt.Sprintf("iByAllergens: [%d], error: %v", len(iByAllergens), err5))
	//
	//iBySkinConcern, err6 := gc.geniusData.IngredientBySkinConcern(ctx.UserContext(), qa.SkinConcern)
	//fmt.Println(fmt.Sprintf("iBySkinConcern: [%d], error: %v", len(iBySkinConcern), err6))
	//
	//iByAge, err7 := gc.geniusData.IngredientByAge(ctx.UserContext(), qa.Age)
	//fmt.Println(fmt.Sprintf("iByAge: [%d], error: %v", len(iByAge), err7))
	//
	//iByBenefit, err8 := gc.geniusData.IngredientByProductBenefit(ctx.UserContext(), qa.ProductBenefit)
	//fmt.Println(fmt.Sprintf("iByBenefit: [%d], error: %v", len(iByBenefit), err8))

	allSkinTypes, err1 := gc.geniusData.GetAllSkintypes(ctx.UserContext())
	fmt.Println(fmt.Sprintf("allSkinTypes: [%d], error: %v", allSkinTypes, err1))

	allSkinSensetivities, err2 := gc.geniusData.GetAllSkinsensetivity(ctx.UserContext())
	fmt.Println(fmt.Sprintf("allSkinSensetivities: [%d], error: %v", allSkinSensetivities, err2))

	allAcneBreakouts, err3 := gc.geniusData.GetAllSkinsensetivity(ctx.UserContext())
	fmt.Println(fmt.Sprintf("allAcneBreakouts: [%d], error: %v", allAcneBreakouts, err3))

	allPreferences, err4 := gc.geniusData.GetAllPreferences(ctx.UserContext())
	fmt.Println(fmt.Sprintf("allPreferences: [%d], error: %v", allPreferences, err4))

	allAllergies, err5 := gc.geniusData.GetAllAllergies(ctx.UserContext())
	fmt.Println(fmt.Sprintf("allAlergies: [%d], error: %v", allAllergies, err5))

	allSkinconcerns, err6 := gc.geniusData.GetAllSkinconcerns(ctx.UserContext())
	fmt.Println(fmt.Sprintf("allSkinconcerns: [%d], error: %v", allSkinconcerns, err6))

	allAges, err7 := gc.geniusData.GetAllAge(ctx.UserContext())
	fmt.Println(fmt.Sprintf("allAges: [%d], error: %v", allAges, err7))

	allBenefits, err8 := gc.geniusData.GetAllBenefits(ctx.UserContext())
	fmt.Println(fmt.Sprintf("allBenefits: [%d], error: %v", allBenefits, err8))

	return nil
}
