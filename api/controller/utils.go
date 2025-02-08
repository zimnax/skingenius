package controller

import (
	"context"
	"fmt"
	"math"
	"skingenius/database/model"
	globalModel "skingenius/frontModel"
	"skingenius/logger"
	"sort"
)

var marginErrors = []float64{17.03, 7.13, 3.86, 2.92, 2.15, 1.25, 1.03, 0.64, 0.61, 0.63, 0.68, 0.44, 0.42, 0.33, 0.32, 0.27, 0.33, 0.27, 0.31, 0.24, 0.11, 0.12}

// merge scores of same ingredient if occurs in multiple answers
func mergeIngredientsWithScores(ingredients ...[]model.Ingredient) map[string]float64 {
	imap := make(map[string]float64)

	for _, ingredient := range ingredients {
		for _, i := range ingredient {
			imap[i.Name] = imap[i.Name] + i.Score
		}
	}

	return imap
}

func getIngredientsNames(is []model.Ingredient) []string {
	var names []string
	for _, i := range is {
		names = append(names, i.Name)
	}
	return names
}

func getIngredientsIds(is []model.Ingredient) []uint {
	var ids []uint
	for _, i := range is {
		ids = append(ids, i.ID)
	}
	return ids
}

func IngredientsSliceToMap(is []model.Ingredient) map[string]float64 {
	m := make(map[string]float64)

	for _, i := range is {
		m[i.Name] = i.Score
	}

	return m
}

func FindTop3Products(products map[string]int) []string {
	keys := make([]string, 0, len(products))
	for k := range products {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return products[keys[i]] > products[keys[j]]
	})

	return keys[:3]
}

// 1. Merge scores of same ingredients if occurs in multiple answers
// 2. Find ingredients what are common in all answers. If answer array is empty(no matching ingredients by answer), ignoring it in merging - should NOT happen :)
func uniqueIngredientsNamesMap(ingredients ...[]model.Ingredient) (map[string]float64, []string) {
	countMap := make(map[string]int)
	scoreMap := make(map[string]float64)
	uniqueFullIngredients := make(map[string]model.Ingredient)
	var uniqueIngredientNames []string

	var ingredientArrays int

	for _, answerIngredients := range ingredients {
		if len(answerIngredients) > 0 { // skip empty arrays if no Allergies, no Concerns, no Benefits etc...
			ingredientArrays++
		}

		for _, ingredient := range answerIngredients {
			uniqueFullIngredients[ingredient.Name] = ingredient // collect all ingredients (we need only unique ones, can be optimized later)

			countMap[ingredient.Name]++
			scoreMap[ingredient.Name] = scoreMap[ingredient.Name] + ingredient.Score
		}
	}

	// Find strings that appear in all slices (count == len(ingredients))
	result := make(map[string]float64)
	for str, count := range countMap {
		if count == ingredientArrays {
			result[str] = scoreMap[str]
			uniqueIngredientNames = append(uniqueIngredientNames, str)
		}
	}

	// If ingredient has effective at low concentration YES: *1, moderate effective *0,5, no * 0
	for _, fullIngredient := range uniqueFullIngredients {
		if _, ok := result[fullIngredient.Name]; ok {
			switch fullIngredient.EffectiveAtLowConcentration {
			case model.EffectiveYes:
				result[fullIngredient.Name] = result[fullIngredient.Name] * float64(1)
			case model.EffectiveModerate:
				result[fullIngredient.Name] = result[fullIngredient.Name] * float64(0.5)
			case model.EffectiveNo:
				result[fullIngredient.Name] = result[fullIngredient.Name] * float64(0)
			}
		}
	}

	return result, uniqueIngredientNames
}

/*
*
If ingredient is in the product:

1. add the score to the product
2. multiply the score by 6 if the ingredient is the first in the product, 5 if the second, 4 if the third, 3 if the fourth, 2 if the fifth, 1 if the sixth
3. divide by the number of ingredients in the product
*/
func matchProductsAndIngredients(ingredients map[string]float64, allProducts []model.Product) map[string]float64 {
	productScores := make(map[string]float64)
	for productIngredientPlace, product := range allProducts {
		for _, productIngredient := range product.Ingredients {
			if _, ok := ingredients[productIngredient.Name]; ok {
				switch productIngredientPlace {
				case 0:
					productScores[product.Name] = productScores[product.Name] + ingredients[productIngredient.Name]*6
				case 1:
					productScores[product.Name] = productScores[product.Name] + ingredients[productIngredient.Name]*5
				case 2:
					productScores[product.Name] = productScores[product.Name] + ingredients[productIngredient.Name]*4
				case 3:
					productScores[product.Name] = productScores[product.Name] + ingredients[productIngredient.Name]*3
				case 4:
					productScores[product.Name] = productScores[product.Name] + ingredients[productIngredient.Name]*2
				default:
					productScores[product.Name] = productScores[product.Name] + ingredients[productIngredient.Name]
				}
			}
		}

		//fmt.Println(fmt.Sprintf("-> %f , %f, %v", productScores[product.Name], float64(len(product.Ingredients)), productScores[product.Name]/float64(len(product.Ingredients))))

		productScores[product.Name] = productScores[product.Name] / float64(len(product.Ingredients))
	}
	return productScores
}

// TODO Rename from top# to Top(N) - returning first 20 products
func sortProductsByScoreTop3(products map[string]float64) map[string]float64 {

	N := 20 // magic number - number of products to return and display at - SEE ALL PRODUCTS - page

	type kv struct {
		Key   string
		Value float64
	}

	var sortedPairs []kv
	for k, v := range products {
		sortedPairs = append(sortedPairs, kv{k, v})
	}

	// Sort the slice based on the value
	sort.Slice(sortedPairs, func(i, j int) bool {
		return sortedPairs[i].Value > sortedPairs[j].Value // Descending order
	})

	// Print the sorted map
	//for _, pair := range sortedPairs {
	//	fmt.Printf(">>  %s: %f\n", pair.Key, pair.Value)
	//}

	if len(sortedPairs) == 0 { //|| len(sortedPairs) < 3
		fmt.Println(fmt.Sprintf("Not enogh products to return, total sorted products: %d", len(sortedPairs)))
		return map[string]float64{}
	}

	// score adjustments to ScalingRepresentation value being #1 in recommendations
	maxScore := sortedPairs[0].Value
	if maxScore == 0 {
		maxScore = 1 // inf NAN issue fix
	}
	scalingFactor := globalModel.ScalingRepresentation / maxScore

	mapToReturn := make(map[string]float64)
	for i, _ := range sortedPairs {
		if i == N {
			break
		}

		mapToReturn[sortedPairs[i].Key] = roundToFirstDecimal(sortedPairs[i].Value * scalingFactor)
	}

	//return map[string]float64{
	//	sortedPairs[0].Key: roundToFirstDecimal(sortedPairs[0].Value * scalingFactor),
	//	sortedPairs[1].Key: roundToFirstDecimal(sortedPairs[1].Value * scalingFactor),
	//	sortedPairs[2].Key: roundToFirstDecimal(sortedPairs[2].Value * scalingFactor),
	//}
	return mapToReturn
}

func roundToFirstDecimal(value float64) float64 {
	const places = 1
	pow := math.Pow(10, float64(places))
	return math.Round(value*pow) / pow
}

func determineSkinSensitivity(answ []string) string {
	var weight int

	for i, _ := range answ {
		switch answ[i] {
		case "dry":
			weight += 1
		case "normal":
			weight += 2
		case "combination":
			weight += 4
		case "oily":
			weight += 8
		}
	}

	avgWeight := float64(weight)/float64(len(answ)) + 0.5
	fmt.Println("avgWeight:", avgWeight)

	if avgWeight < 1.5 {
		return "dry"
	} else if avgWeight >= 1.5 && avgWeight <= 2.5 {
		return "normal"
	} else if avgWeight > 2.5 && avgWeight < 6.5 {
		return "combination"
	} else if avgWeight >= 6.5 {
		return "oily"
	}

	return "normal" // todo
}

func calculateConcentrations(p model.Product) map[string]string {
	concentrations := make(map[string]string)

	for i, _ := range p.Ingredients {
		concenrationByRole := model.ConcentrationMap[string(p.Ingredients[i].Roleinformulations[0].Name)] //TODO: assume role is the only one in the array

		c := concenrationByRole.Min + (concenrationByRole.Max-concenrationByRole.Min)/float64(i)
		logger.New().Info(context.Background(), fmt.Sprintf("Product %s - Ingredient %s calculated concentration: %f", p.Name, p.Ingredients[i].Name, c))
		//logger.New().Info(context.Background(), fmt.Sprintf("Ingredient %s has effective concentration in range", p.Ingredients[i].Name))

		if p.Type == model.MoisturizerProductType {
			if p.Ingredients[i].ConcentrationRinseOffMin < c && p.Ingredients[i].ConcentrationRinseOffMax > c {
				logger.New().Info(context.Background(), fmt.Sprintf("Ingredient %s has effective concentration in range", p.Ingredients[i].Name))
			}
		} else {
			if p.Ingredients[i].ConcentrationLeaveOnMin < c && p.Ingredients[i].ConcentrationLeaveOnMax > c {
				logger.New().Info(context.Background(), fmt.Sprintf("Ingredient %s has effective concentration in range", p.Ingredients[i].Name))
			}
		}

		//concentrations[i.Name] = i.EffectiveConcentrations // a+(b-a)/nT
	}

	return concentrations
}

// Checking the effective concentration of the active ingredients in the product, and if its below the minimum, adjust the score (set ot to 0.2)
func validateConcentrations(product model.Product) model.Product {
	count := 0 // number of active ingredients from 1 product with concentrations in the range

	for _, activeName := range product.ActiveIngredients {
		c := product.Concentrations[activeName]
		//fmt.Println(fmt.Sprintf("product [%s] active ingredient [%s] concentration: %v", product.Name, activeName, c))

		for j, productIngredient := range product.Ingredients {
			if productIngredient.Name == activeName {

				margin := 0.12 // default margin for ingredients with index > 21 //todo: improve in future
				if productIngredient.Index < len(marginErrors)-1 {
					margin = marginErrors[productIngredient.Index]
				}

				if product.Type == model.MoisturizerProductType {
					if productIngredient.ConcentrationRinseOffMin-margin < c {
						logger.New().Debug(context.Background(), fmt.Sprintf("Ingredient [%s] effective concentration [%f] "+
							"is in the range  [%f- %f]", productIngredient.Name, c, productIngredient.ConcentrationRinseOffMin, productIngredient.ConcentrationRinseOffMax))
						count++
						continue
					}
					logger.New().Debug(context.Background(), fmt.Sprintf("Ingredient [%s] effective concentration [%f] is NOT in the "+
						"range [%f- %f]", productIngredient.Name, c, productIngredient.ConcentrationRinseOffMin, productIngredient.ConcentrationRinseOffMax))
					product.Ingredients[j].Score = 0.2
				} else {
					if productIngredient.ConcentrationLeaveOnMin-margin < c {
						logger.New().Debug(context.Background(), fmt.Sprintf("Ingredient [%s] effective concentration [%f] "+
							"is in the range  [%f- %f]", productIngredient.Name, c, productIngredient.ConcentrationRinseOffMin, productIngredient.ConcentrationRinseOffMax))
						count++
						continue
					}
					logger.New().Debug(context.Background(), fmt.Sprintf("Ingredient [%s] effective concentration [%f] is NOT in the "+
						"range [%f- %f]", productIngredient.Name, c, productIngredient.ConcentrationRinseOffMin, productIngredient.ConcentrationRinseOffMax))
					product.Ingredients[j].Score = 0.2
				}
			}
		}
	}

	logger.New().Info(context.Background(), fmt.Sprintf("Product [%s] has %d/%d effective concentrations  in the range", product.Name, count, len(product.ActiveIngredients)))

	fmt.Println(fmt.Sprintf("%#v", product))
	return product
}

func calculateWeightedAverageScorePassive(p model.Product) float64 {
	var sum float64
	for _, passiveIngredient := range p.PassiveIngredients {
		for _, ingredient := range p.Ingredients {

			if passiveIngredient == ingredient.Name {
				was := ingredient.Score * p.Concentrations[ingredient.Name]
				logger.New().Info(context.Background(), fmt.Sprintf("Product [%s] PASSIVE ingredient [%s] score: [%f], concentration: [%f], WAS: [%f]", p.Name, ingredient.Name, ingredient.Score, p.Concentrations[ingredient.Name], was))

				sum = sum + was
			}

		}
	}
	return sum
}

func calculateWeightedAverageScoreActive(p model.Product) float64 {
	var sum float64
	for _, activeIngredient := range p.ActiveIngredients {
		for i, ingredient := range p.Ingredients {

			if activeIngredient == ingredient.Name {
				was := ingredient.Score * p.Concentrations[ingredient.Name]
				p.Ingredients[i].WAS = was
				logger.New().Info(context.Background(), fmt.Sprintf("Product [%s] ACTIVE ingredient [%s] score: [%f], concentration: [%f], WAS: [%f]", p.Name, ingredient.Name, ingredient.Score, p.Concentrations[ingredient.Name], was))

				sum = sum + was
			}

		}
	}
	return sum
}

// 1.2(AC)*4.46 * 10^(-3) * W49^(-1.79)
func calculateConcentrationsBogdanFormula(p model.Product, adjustmentCoefficients float64) map[string]float64 {
	concentrations := make(map[string]float64)
	sum := 0.0

	for i, ingredient := range p.Ingredients {
		it := float64(ingredient.Index+1) / float64(len(p.Ingredients))
		concentration := (adjustmentCoefficients * 4.46 * math.Pow(10, -3) * math.Pow(it, -1.79)) * 100

		// Checking if we have known concentration for this ingredient, if yes, we use it instead of calculated
		if p.Type == model.MoisturizerProductType && ingredient.KnownConcentrationRinseOffMin != 0 && ingredient.KnownConcentrationRinseOffMax != 0 {
			if concentration < ingredient.KnownConcentrationRinseOffMin {
				concentration = ingredient.KnownConcentrationRinseOffMin
				logger.New().Debug(context.Background(), fmt.Sprintf("product [%s] Ingredient [%s] concentration is less than min, set to min: %f", p.Name, ingredient.Name, concentration))
			}

			if concentration > ingredient.KnownConcentrationRinseOffMax {
				concentration = ingredient.KnownConcentrationRinseOffMax
				logger.New().Debug(context.Background(), fmt.Sprintf("product [%s] Ingredient [%s] concentration is greater than max, set to max: %f", p.Name, ingredient.Name, concentration))
			}
		} else if p.Type != model.MoisturizerProductType && ingredient.KnownConcentrationLeaveOnMin != 0 && ingredient.KnownConcentrationLeaveOnMax != 0 {
			if concentration < ingredient.KnownConcentrationLeaveOnMin {
				concentration = ingredient.KnownConcentrationLeaveOnMin
				logger.New().Debug(context.Background(), fmt.Sprintf("product [%s] Ingredient [%s] concentration is less than min, set to min: %f", p.Name, ingredient.Name, concentration))

			}

			if concentration > ingredient.KnownConcentrationLeaveOnMax {
				concentration = ingredient.KnownConcentrationLeaveOnMax
				logger.New().Debug(context.Background(), fmt.Sprintf("product [%s] Ingredient [%s] concentration is greater than max, set to max: %f", p.Name, ingredient.Name, concentration))
			}
		}

		concentrations[p.Ingredients[i].Name] = math.Round(concentration*1000) / 1000

		sum = sum + concentration
	}

	//fmt.Printf("coef: %.2f sum: %.4f \n", adjustmentCoefficients, sum)
	if sum > 99 { // magic number (precision)
		logger.New().Info(context.Background(), fmt.Sprintf("product [%s] Sum of concentrations is greater than 99, sum: %f", p.Name, sum))
		return concentrations
	}
	return calculateConcentrationsBogdanFormula(p, adjustmentCoefficients+0.01)
}

func findProductsWithActiveIngredients(products []model.Product, ingredients []model.Ingredient) map[string]model.Product {
	filteredProducts := make(map[string]model.Product)

	productHasActiveIngredients := false
	for i, product := range products {

		for _, productIngredient := range product.Ingredients {
			isIngredientActive := false
			for _, ingredient := range ingredients {
				if ok, _ := findStringInSlice(productIngredient.Synonyms, ingredient.Name); ok {
					products[i].ActiveIngredients = append(products[i].ActiveIngredients, productIngredient.Name)
					logger.New().Debug(context.Background(), fmt.Sprintf("Product [%s] has active ingredient [%s]", product.Name, ingredient.Name))

					productHasActiveIngredients = true
					isIngredientActive = true
					continue
				}
			}
			if !isIngredientActive {
				products[i].PassiveIngredients = append(products[i].PassiveIngredients, productIngredient.Name)
			}
		}

		if productHasActiveIngredients {
			filteredProducts[products[i].Name] = products[i] // can be added multiple times if a few ingredients match but map makes it unique
		}
		productHasActiveIngredients = false
	}

	return filteredProducts
}

func findStringInSlice(slice []string, target string) (bool, string) {
	for _, s := range slice {
		if s == target {
			return true, target
		}
	}
	return false, ""
}

func findProductIntersection(productsByConcern map[string][]model.Product) []model.Product {
	productCount := make(map[uint]int)
	totalConcerns := len(productsByConcern)

	// Count occurrences of each product ID across all concerns
	for _, products := range productsByConcern {
		seen := make(map[uint]bool) // To avoid double counting in the same concern

		for _, product := range products {
			if !seen[product.ID] {
				productCount[product.ID]++
				seen[product.ID] = true
			}
		}
	}

	// Find products that appear in all concerns
	var commonProducts []model.Product
	for _, products := range productsByConcern {
		for _, product := range products {
			if productCount[product.ID] == totalConcerns {
				commonProducts = append(commonProducts, product)
			}
		}
		break // Only need one list to extract unique products
	}

	return commonProducts
}
