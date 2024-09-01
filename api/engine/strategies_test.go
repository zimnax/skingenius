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
	db, err := database.NewGormClient(config.RemoteHost, config.Port, config.User, config.Password, false)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	q1SkinTypeAnswer := "normal"
	q2SkinSensitivityAnswer := "often"
	q3AcneBreakoutsAnswer := "occasionally"
	q4PreferencesAnswer := []string{"vegetarian", "vegan"}
	q5AllergiesAnswer := []string{"nuts", "latex"}
	q6SkinConcernAnswer := []string{"acne", "hyperpigmentation_unevenskintone"}
	q7AgeAnswer := 20
	q8BenefitsAnswer := []string{"moisturizing", "nourishing"}

	top3 := FindBestProducts_matchBestStrategy(db, context.Background(), q1SkinTypeAnswer, q2SkinSensitivityAnswer, q3AcneBreakoutsAnswer,
		q4PreferencesAnswer, q5AllergiesAnswer, q6SkinConcernAnswer, q7AgeAnswer, q8BenefitsAnswer)

	fmt.Println(fmt.Sprintf("products len:: %v", len(top3)))
}

//found 50 products from db
//Product: hyaluronic acid moisturizer, Score: 29.000000
//Product: moisturizing oat & calendula miracle face cream, Score: 28.800000
//Product: perfect facial hydrating cream, Score: 25.200000

//Product: moisturizing oat & calendula miracle face cream, Score: 288.000000
//Product: perfect facial hydrating cream, Score: 252.000000
//Product: hyaluronic acid moisturizer, Score: 290.000000

/*

Skin type:  normal
Sensitivity:  never
Acne:  veryfrequently
Age:  20
Preference:  [paleo]
Allergy:  [artificial_fragrance]
Concerns:  []
Benefits:  [improves_skin_tone]


gives scores 0
*/
