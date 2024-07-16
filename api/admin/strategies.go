package main

import (
	"context"
	"fmt"
	"skingenius/admin/csv"
	"skingenius/database"
	"skingenius/database/model"
	"strconv"
)

func scoreByQuestionReport(dbClient database.Connector,
	skintypeIng, skinSensIng, acneIng, prefIng, allergiesIng, skinConcernIng, ageIng, benefitsIng []model.Ingredient,
	skintypeAns, skinSensAns, acneAns, prefAns, allergyAns, concernAns string, ageAns int, benefitsAns string) {

	skintypeIngMap := IngredientsSliceToMap(skintypeIng)
	skinSensIngMap := IngredientsSliceToMap(skinSensIng)
	acneIngMap := IngredientsSliceToMap(acneIng)
	prefIngMap := IngredientsSliceToMap(prefIng)
	allergiesIngMap := IngredientsSliceToMap(allergiesIng)
	skinConcernIngMap := IngredientsSliceToMap(skinConcernIng)
	ageIngMap := IngredientsSliceToMap(ageIng)
	benefitsIngMap := IngredientsSliceToMap(benefitsIng)

	headers := []string{"Ingredient name", skintypeAns, skinSensAns, acneAns, prefAns, allergyAns, concernAns, strconv.Itoa(ageAns), benefitsAns, "Total score"}

	allIng, err := dbClient.GetAllIngredients(context.Background())
	if err != nil {
		panic(err)
	}

	var scores [][]string

	for _, ingredient := range allIng {
		skintypeScore := skintypeIngMap[ingredient.Name]
		skinSensScore := skinSensIngMap[ingredient.Name]
		acneScore := acneIngMap[ingredient.Name]
		prefScore := prefIngMap[ingredient.Name]
		allergiesScore := allergiesIngMap[ingredient.Name]
		skinConcernScore := skinConcernIngMap[ingredient.Name]
		ageScore := ageIngMap[ingredient.Name]
		benefitsScore := benefitsIngMap[ingredient.Name]

		totalScore := skintypeScore + skinSensScore + acneScore + prefScore + allergiesScore + skinConcernScore + ageScore + benefitsScore

		vals := []string{ingredient.Name, strconv.Itoa(skintypeScore), strconv.Itoa(skinSensScore), strconv.Itoa(acneScore), strconv.Itoa(prefScore),
			strconv.Itoa(allergiesScore), strconv.Itoa(skinConcernScore), strconv.Itoa(ageScore), strconv.Itoa(benefitsScore), strconv.Itoa(totalScore)}

		scores = append(scores, vals)
	}

	if csverr := csv.WriteToCsv("scores_per_ingredient", headers, scores); csverr != nil {
		panic(csverr)
	}
}

func findBestProducts_RatingStrategy(dbClient database.Connector, ctx context.Context,
	q1SkinTypeAnswer string, q2SkinSensitivityAnswer string, q3AcneBreakoutsAnswer string, q4PreferencesAnswer []string,
	q5AllergiesAnswer []string, q6SkinConcernAnswer []string, q7AgeAnswer int, q8BenefitsAnswer []string) {

	q1Ing, q2Ing, q3Ing, q4Ing, q5Ing, q6Ing, q7Ing, q8Ing := findIngredientsByQuestion(dbClient, ctx, q1SkinTypeAnswer, q2SkinSensitivityAnswer, q3AcneBreakoutsAnswer, q4PreferencesAnswer, q5AllergiesAnswer, q6SkinConcernAnswer, q7AgeAnswer, q8BenefitsAnswer)

	scoreByQuestionReport(dbClient, q1Ing, q2Ing, q3Ing, q4Ing, q5Ing, q6Ing, q7Ing, q8Ing,
		q1SkinTypeAnswer, q2SkinSensitivityAnswer, q3AcneBreakoutsAnswer, q4PreferencesAnswer[0], q5AllergiesAnswer[0], q6SkinConcernAnswer[0], q7AgeAnswer, q8BenefitsAnswer[0])

	fullIngredients := mergeIngredientsWithScores(q1Ing, q2Ing, q3Ing, q4Ing, q5Ing, q6Ing, q7Ing, q8Ing)
	fmt.Println(fmt.Sprintf("\n total ingredients after score [sum]: %d \n\n ", len(fullIngredients)))

	allProducts, err := dbClient.FindAllProducts(ctx)
	if err != nil {
		panic(err)
	}

	productScoreMap := make(map[string]int)

	for _, singleProduct := range allProducts {
		if singleProduct.Name == "" {
			continue
		}

		fullProduct, err := dbClient.FindProductByName(context.Background(), singleProduct.Name)
		if err != nil {
			panic(err)
		}

		singleProduct = *fullProduct

		productIngredientsWithScore := make(map[string]int)
		for _, ingredientFromProduct := range singleProduct.Ingredients {
			if score, ok := fullIngredients[ingredientFromProduct.Name]; ok {
				productIngredientsWithScore[ingredientFromProduct.Name] = score
			}
		}
		if csverr := csv.SingleProductExtendedReport("product_extended_report", []string{"Product", "Ingredient", "Score"}, singleProduct.Name, productIngredientsWithScore); csverr != nil {
			panic(csverr)
		}

		productWithAllIngredients := make(map[string]int)
		for name, score := range fullIngredients {
			if _, ok := productIngredientsWithScore[name]; ok {
				productWithAllIngredients[name] = score
			} else {
				productWithAllIngredients[name] = 0
			}
		}
		if csverr := csv.SingleProductExtendedReport("productWithAllIngredients_extended_report", []string{"Product", "Ingredient", "Score"}, singleProduct.Name, productWithAllIngredients); csverr != nil {
			panic(csverr)
		}

		if len(productIngredientsWithScore) != 0 {
			fmt.Println(fmt.Sprintf("[%s] __ %d __ %+v", singleProduct.Name, len(productIngredientsWithScore), productIngredientsWithScore))
			totalScore := 0
			for _, i := range productIngredientsWithScore {
				totalScore = totalScore + i
			}
			productScoreMap[singleProduct.Name] = totalScore
		}
	}

	//fmt.Println(fmt.Sprintf("product score map: %#v", productScoreMap))

	if csverr := csv.WriteToFile("product_total_score", []string{"Product", "Score"}, productScoreMap); csverr != nil {
		panic(csverr)
	}

}

/*

	unique ingredients + find products with N ingredients from unoque list
*/

func findBestProducts_matchBestStrategy(dbClient database.Connector, ctx context.Context,
	q1SkinTypeAnswer string, q2SkinSensitivityAnswer string, q3AcneBreakoutsAnswer string, q4PreferencesAnswer []string,
	q5AllergiesAnswer []string, q6SkinConcernAnswer []string, q7AgeAnswer int, q8BenefitsAnswer []string) {

	q1Ing, q2Ing, q3Ing, q4Ing, q5Ing, q6Ing, q7Ing, q8Ing := findIngredientsByQuestion(dbClient, ctx, q1SkinTypeAnswer, q2SkinSensitivityAnswer, q3AcneBreakoutsAnswer,
		q4PreferencesAnswer, q5AllergiesAnswer, q6SkinConcernAnswer, q7AgeAnswer, q8BenefitsAnswer)

	//mergedIngredientsList := mergeIngredients(q1Ing, q2Ing, q3Ing, q4Ing, q5Ing, q6Ing, q7Ing, q8Ing)
	//fmt.Println(fmt.Sprintf("merged ingredients: %v", len(mergedIngredientsList)))

	//fmt.Println(fmt.Sprintf("-->> q1 ingredients: %v", getIngredientsNames(q1Ing)))
	//fmt.Println(fmt.Sprintf("-->> q2 ingredients: %v", getIngredientsNames(q2Ing)))
	//fmt.Println(fmt.Sprintf("-->> q3 ingredients: %v", getIngredientsNames(q3Ing)))
	//fmt.Println(fmt.Sprintf("-->> q4 ingredients: %v", getIngredientsNames(q4Ing)))
	//fmt.Println(fmt.Sprintf("-->> q5 ingredients: %v", getIngredientsNames(q5Ing)))
	//fmt.Println(fmt.Sprintf("-->> q6 ingredients: %v", getIngredientsNames(q6Ing)))
	//fmt.Println(fmt.Sprintf("-->> q7 ingredients: %v", getIngredientsNames(q7Ing)))
	//fmt.Println(fmt.Sprintf("-->> q8 ingredients: %v", getIngredientsNames(q8Ing)))

	iNames := uniqueIngredientsNamesList(q1Ing, q2Ing, q3Ing, q4Ing, q5Ing, q6Ing, q7Ing, q8Ing)
	fmt.Println(fmt.Sprintf("unuqie ingredients: %#v", len(iNames)))
	fmt.Println(fmt.Sprintf("unuqie ingredients: %#v", iNames))

	ps, err := dbClient.FindAllProductsWithIngredients(context.Background(), iNames, uint(3)) // len(iNames)
	fmt.Println(fmt.Sprintf("Products #%d", len(ps)))
	fmt.Println(fmt.Sprintf("Products: %+v", ps))
	fmt.Println(err)
}
