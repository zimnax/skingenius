package engine

import (
	"skingenius/database/model"
	"sort"
)

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
