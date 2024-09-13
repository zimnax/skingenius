package controller

import (
	"fmt"
	dbmodel "skingenius/database/model"
	"skingenius/model"
)

func quizAnswersToDbModel(quizAnswers model.QuizAnswers) model.DBAnswerModel {
	a1SkinType := dbmodel.SkinTypeMapping[quizAnswers.SkinType]
	a2SkinSensitivity := dbmodel.SensitivityMapping[quizAnswers.SkinSensitivity]
	a3Acne := dbmodel.AcneProneMapping[quizAnswers.AcneBreakouts]
	a4Age := dbmodel.AgeMapping[quizAnswers.Age]

	var a5Preference []string
	for _, preference := range quizAnswers.ProductPreferences {
		a5Preference = append(a5Preference, string(dbmodel.PreferenceMapping[preference]))
	}

	var a6Allergy []string

	for _, allergen := range quizAnswers.FreeFromAllergens {
		a6Allergy = append(a6Allergy, string(dbmodel.AllergiesMapping[allergen]))
	}

	var a7Concerns []string
	for _, concern := range quizAnswers.SkinConcern {
		a7Concerns = append(a7Concerns, string(dbmodel.SkinConcernsMapping[concern]))
	}

	var a8Benefits []string
	for _, benefit := range quizAnswers.ProductBenefit {
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

	return model.DBAnswerModel{
		SkinType:        a1SkinType,
		SkinSensitivity: a2SkinSensitivity,
		AcneProne:       a3Acne,
		Age:             a4Age,
		Preferences:     a5Preference,
		Allergies:       a6Allergy,
		Concerns:        a7Concerns,
		Benefits:        a8Benefits,
	}
}
