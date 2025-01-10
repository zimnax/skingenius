package controller

import (
	"context"
	"fmt"
	"math"
	"skingenius/database/model"
	"skingenius/logger"
	globalModel "skingenius/model"
	"sort"
)

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
// 2. Find ingredients what are common in all answers. If answer array is empty(no matching ingredients by answer), ignoring it in merging
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

func filterProductsWithIngredients(products []model.Product, ingredients []model.Ingredient) map[string]model.Product {
	filteredProducts := make(map[string]model.Product)

	for i, product := range products {
		for _, ingredient := range ingredients {
			for _, productIngredient := range product.Ingredients {
				if productIngredient.Name == ingredient.Name || findStringInSlice(productIngredient.Name, ingredient.Synonyms) {
					filteredProducts[products[i].Name] = products[i] // can be added multiple times if a few ingredients match but map makes it unique
					logger.New().Info(context.Background(), fmt.Sprintf("Product [%s] has ingredient [%s]", product.Name, ingredient.Name))
					// TODO: mark ingredient inside product if exists
				}
			}
		}
	}

	return filteredProducts
}

func findStringInSlice(target string, slice []string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
