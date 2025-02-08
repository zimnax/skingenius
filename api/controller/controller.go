package controller

import (
	"context"
	"fmt"
	"skingenius/database"
	dbmodel "skingenius/database/model"
	"skingenius/frontModel"
	"skingenius/logger"
)

const packageLogPrefix = "genius_controller : "
const adjustmentCoefficients = 0.1

type GeniusController struct {
	geniusData database.Connector
}

func NewGeniusController(db database.Connector) *GeniusController {
	logger.New().Info(context.Background(), packageLogPrefix+"initializing new genius handler")

	return &GeniusController{geniusData: db}
}

/*
	1. Find ingredients by Allergies
	2. Find ingredients by Preferences
	3. Find ingredients by SkinSensitivity
	4. Find ingredients by SkinConcerns (for each separately)
*/

func (gc *GeniusController) AlgorithmV3(ctx context.Context, quizAnswers frontModel.DBAnswerModel) error {
	logger.New().Info(context.Background(), packageLogPrefix+"AlgorithmV3")

	// 1.
	allergensIng, err := gc.geniusData.GetIngredientsByAllergies(ctx, quizAnswers.Allergies)
	if err != nil {
		logger.New().Warn(context.Background(), fmt.Sprintf("failed to get ingredients by [allergies], error: %v", err))
		return frontModel.ProcessingServerError
	}

	// 2.
	preferencesIng, err := gc.geniusData.GetIngredientsByPreferences(ctx, quizAnswers.Preferences)
	if err != nil {
		logger.New().Warn(context.Background(), fmt.Sprintf("failed to get ingredients by [preferences], error: %v", err))
		return frontModel.ProcessingServerError
	}

	// 3.
	skinSensitivityIng, err := gc.geniusData.GetIngredientsBySkinsensitivity(ctx, quizAnswers.SkinSensitivity)
	if err != nil {
		logger.New().Warn(context.Background(), fmt.Sprintf("failed to get ingredients by [skinSensitivity], error: %v", err))
		return frontModel.ProcessingServerError
	}

	// 4.
	_, uniqueIngredients := uniqueIngredientsNamesMap(allergensIng, preferencesIng, skinSensitivityIng) // todo: do we need merge scores?
	logger.New().Info(context.Background(), fmt.Sprintf("unique ingredients len: %v", len(uniqueIngredients)))

	concernRecommendedProducts, err := gc.concernFlow(ctx, quizAnswers.Concerns, uniqueIngredients)
	benefitsRecommendedProducts, err := gc.benefitsFlow(ctx, quizAnswers.Benefits, uniqueIngredients)

	// 10. Find common products between concerns
	commonProducts := findProductIntersection(concernRecommendedProducts)
	fmt.Println(fmt.Sprintf("common products: %v", len(commonProducts)))

	return nil
}

func (gc *GeniusController) benefitsFlow(ctx context.Context, benefits []string, ingredients []string) (map[string][]dbmodel.Product, error) {
	return nil, nil
}

func (gc *GeniusController) concernFlow(ctx context.Context, concerns []string, ingredients []string) (map[string][]dbmodel.Product, error) {
	productsByConcern := make(map[string][]dbmodel.Product)

	for _, concern := range concerns {
		allProducts, allPErr := gc.geniusData.FindAllProductsHavingIngredients(ctx, ingredients)
		if allPErr != nil {
			logger.New().Warn(context.Background(), fmt.Sprintf("failed to find all products having ingredients, error: %v", allPErr))
			return productsByConcern, frontModel.ProcessingServerError
		}
		logger.New().Info(context.Background(), fmt.Sprintf("found %d products from db for given ingredients count[%d]", len(allProducts), len(ingredients)))

		// 5. Find ingredients that address skin concern
		skinConcernIng, ibscErr := gc.geniusData.GetIngredientsBySkinconcerns(ctx, []string{concern})
		if ibscErr != nil {
			logger.New().Warn(context.Background(), fmt.Sprintf("failed to get ingredients by [skinConcern] [%s], error: %v", concern, ibscErr))
			return productsByConcern, frontModel.ProcessingServerError
		}
		logger.New().Info(context.Background(), fmt.Sprintf("skin concern [%s] ingredients: %v, %v", concern, len(skinConcernIng), skinConcernIng))

		// 6. Find products with ingredients that address skin concern
		concernProducts := findProductsWithActiveIngredients(allProducts, skinConcernIng)
		logger.New().Info(context.Background(), fmt.Sprintf("found %d products with concern [%s]", len(concernProducts), concern))

		var activeProducts []dbmodel.Product
		for i, product := range concernProducts {
			p := concernProducts[i]

			// 7. Calculate concentrations
			p.Concentrations = calculateConcentrationsBogdanFormula(product, adjustmentCoefficients) // cannot assign to concernProducts[i] directly
			concernProducts[i] = p

			// 8. Look at concentration based on product category and check if calculated concentration within the range in DB + adjust score
			adjustedScoreProduct := validateConcentrations(p)
			activeProducts = append(activeProducts, adjustedScoreProduct)
		}

		// 9. Calculate weighted average score passive ingredients
		for i, product := range activeProducts {
			passiveWas := calculateWeightedAverageScorePassive(product)
			activeWas := calculateWeightedAverageScoreActive(product)

			activeProducts[i].ActiveWASTotal = activeWas
			activeProducts[i].PassiveWASTotal = passiveWas
		}

		//todo DEBUG ONLY: remove
		fmt.Println(fmt.Sprintf("Products with concern [%s], len: %d", concern, len(activeProducts)))
		for i, product := range activeProducts {
			fmt.Println(fmt.Sprintf("%d ->>, Product: %s, activeWAS: %f, passiveWAS: %f,  ", i, product.Name, product.ActiveWASTotal, product.PassiveWASTotal))
		}
		//todo DEBUG ONLY: remove

		productsByConcern[concern] = activeProducts
	}

	return productsByConcern, nil
}
