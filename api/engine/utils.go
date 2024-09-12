package engine

import (
	"fmt"
	"math"
	"skingenius/database/model"
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
// 2. Find ingredients what are common in all answers. If answer array is empty, ignoring it in merging
func uniqueIngredientsNamesMap(ingredients ...[]model.Ingredient) map[string]float64 {
	countMap := make(map[string]int)
	scoreMap := make(map[string]float64)

	var ingredientArrays int

	for _, answerIngredients := range ingredients {

		if len(answerIngredients) > 0 { // skip empty arrays in no Allergies, No concerns, no Benefits ets...
			ingredientArrays++
		}

		for _, ingredient := range answerIngredients {
			countMap[ingredient.Name]++
			scoreMap[ingredient.Name] = scoreMap[ingredient.Name] + ingredient.Score
		}
	}

	// Find strings that appear in all slices (count == len(ingredients))
	result := make(map[string]float64)
	for str, count := range countMap {
		if count == ingredientArrays {
			result[str] = scoreMap[str]
		}
	}

	return result
}

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

func sortProductsByScoreTop3(products map[string]float64) map[string]float64 {
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

	if len(sortedPairs) == 0 || len(sortedPairs) < 3 {
		fmt.Println(fmt.Sprintf("Not enogh products to return, total sorted products: %d", len(sortedPairs)))
		return map[string]float64{}
	}

	// score adjustments to ScalingRepresentation value being #1 in recommendations
	maxScore := sortedPairs[0].Value
	if maxScore == 0 {
		maxScore = 1 // inf NAN issue fix
	}
	scalingFactor := globalModel.ScalingRepresentation / maxScore

	return map[string]float64{
		sortedPairs[0].Key: roundToFirstDecimal(sortedPairs[0].Value * scalingFactor),
		sortedPairs[1].Key: roundToFirstDecimal(sortedPairs[1].Value * scalingFactor),
		sortedPairs[2].Key: roundToFirstDecimal(sortedPairs[2].Value * scalingFactor),
	}
}

func roundToFirstDecimal(value float64) float64 {
	const places = 1
	pow := math.Pow(10, float64(places))
	return math.Round(value*pow) / pow
}
