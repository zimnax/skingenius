package handler

import (
	"fmt"
	"skingenius/frontModel"
)

func quizAnswersToDbModel(quizAnswers frontModel.QuizAnswers) frontModel.DBAnswerModel {
	a1SkinType := frontModel.SkinTypeMapping[quizAnswers.SkinType]
	a2SkinSensitivity := frontModel.SensitivityMapping[quizAnswers.SkinSensitivity]
	a3Acne := frontModel.AcneProneMapping[quizAnswers.AcneBreakouts]
	a4Age := frontModel.AgeMapping[quizAnswers.Age]

	var a5Preference []string
	for _, preference := range quizAnswers.ProductPreferences {
		a5Preference = append(a5Preference, string(frontModel.PreferenceMapping[preference]))
	}

	var a6Allergy []string

	for _, allergen := range quizAnswers.FreeFromAllergens {
		a6Allergy = append(a6Allergy, string(frontModel.AllergiesMapping[allergen]))
	}

	var a7Concerns []string
	for _, concern := range quizAnswers.SkinConcern {
		a7Concerns = append(a7Concerns, string(frontModel.SkinConcernsMapping[concern]))
	}

	var a8Benefits []string
	for _, benefit := range quizAnswers.ProductBenefit {
		a8Benefits = append(a8Benefits, string(frontModel.BenefitsMapping[benefit]))
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

	return frontModel.DBAnswerModel{
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
