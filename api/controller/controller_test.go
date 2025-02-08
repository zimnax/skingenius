package controller

import (
	"context"
	"fmt"
	"os"
	"skingenius/config"
	"skingenius/database"
	"skingenius/frontModel"
	"testing"
)

func TestGeniusController_AlgorithmV3(t *testing.T) {
	db, err := database.NewGormClient(config.LocalHost, config.Port, config.User, config.Password, false, "test")
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	q1SkinTypeAnswer := "normal"
	q2SkinSensitivityAnswer := "often"
	q3AcneBreakoutsAnswer := "occasionally"
	q4PreferencesAnswer := []string{"vegetarian", "vegan"}
	q5AllergiesAnswer := []string{"nuts", "latex"}
	q6SkinConcernAnswer := []string{"acne"} //"hyperpigmentation_unevenskintone"
	q7AgeAnswer := 20
	q8BenefitsAnswer := []string{"moisturizing", "nourishing"}

	geniusController := NewGeniusController(db)

	geniusController.AlgorithmV3(context.Background(), frontModel.DBAnswerModel{
		SkinType:        q1SkinTypeAnswer,
		SkinSensitivity: q2SkinSensitivityAnswer,
		AcneProne:       q3AcneBreakoutsAnswer,
		Age:             q7AgeAnswer,
		Preferences:     q4PreferencesAnswer,
		Allergies:       q5AllergiesAnswer,
		Concerns:        q6SkinConcernAnswer,
		Benefits:        q8BenefitsAnswer,
	})

}
