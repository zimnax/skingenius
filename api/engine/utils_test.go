package engine

import (
	"fmt"
	"skingenius/database/model"
	"testing"
)

func Test_UniqueIngredients(t *testing.T) {

	is := [][]model.Ingredient{
		{{Name: "name1", Score: 1}, {Name: "name2"}, {Name: "name3"}},
		{{Name: "name1", Score: 1}, {Name: "name2"}, {Name: "name33"}},
		{{Name: "name1", Score: 1}, {Name: "name222"}, {Name: "name333"}},
		{{Name: "name1", Score: 1}, {Name: "name222"}, {Name: "name3333"}},
	}

	actual := uniqueIngredientsNamesMap(is...)
	if len(actual) != 1 {
		t.Fatalf("expected 1, actual %d", len(actual))
	}

	fmt.Println("actual", actual)
}

func Test_sortProductsByScoreTop3(t *testing.T) {
	products := map[string]float64{
		"product4": 11,
		"product1": 20, // top 3
		"product5": 14,
		"product2": 40, // top 3
		"product6": 16,
		"product3": 60, // top 3
		"product7": 18,
		"product8": 19,
	}

	actual := sortProductsByScoreTop3(products)
	if len(actual) != 3 {
		t.Fatalf("expected 3, actual %d", len(actual))
	}

	fmt.Println("actual", actual) // map[product1:32.3 product2:64.7 product3:97]
}
