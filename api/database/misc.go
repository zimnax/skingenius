package database

import "skingenius/database/model"

func GetIngredientsName(is []model.Ingredient) []string {
	var names []string

	for _, i := range is {
		names = append(names, i.Name)
	}
	return names
}
