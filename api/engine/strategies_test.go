package engine

import (
	"context"
	"fmt"
	"os"
	"skingenius/config"
	"skingenius/database"
	"testing"
)

func Test_FindBestProducts_matchBestStrategy(t *testing.T) {
	db, err := database.NewGormClient(config.Host, config.Port, config.User, config.Password, false)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	q1SkinTypeAnswer := "normal"
	q2SkinSensitivityAnswer := "often"
	q3AcneBreakoutsAnswer := "occasionally"
	q4PreferencesAnswer := []string{"vegetarian", "vegan"}
	q5AllergiesAnswer := []string{"nuts", "latex"}
	q6SkinConcernAnswer := []string{"rosacea", "hyperpigmentation_unevenskintone"}
	q7AgeAnswer := 20
	q8BenefitsAnswer := []string{"moisturizing", "nourishing"}

	top3 := FindBestProducts_matchBestStrategy(db, context.Background(), q1SkinTypeAnswer, q2SkinSensitivityAnswer, q3AcneBreakoutsAnswer,
		q4PreferencesAnswer, q5AllergiesAnswer, q6SkinConcernAnswer, q7AgeAnswer, q8BenefitsAnswer)

	fmt.Println(fmt.Sprintf("products len:: %v", len(top3)))
}
