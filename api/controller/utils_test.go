package controller

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

func Test_determineSkinSensitivity(t *testing.T) {

	tests := []struct {
		Input []string
		Want  string
	}{
		{
			Input: []string{"dry", "dry", "normal", "normal", "combination"},
			Want:  "normal",
		},
		{
			Input: []string{"dry", "dry", "combination", "oily", "combination"},
			Want:  "combination", // Comb
		},
		{
			Input: []string{"dry", "normal", "combination", "combination", "combination"},
			Want:  "combination", // comb
		},
		{
			Input: []string{"normal", "combination", "oily", "oily", "combination"},
			Want:  "combination",
		},
		{
			Input: []string{"normal", "combination", "oily", "oily", "oily"},
			Want:  "oily",
		},
	}

	for _, test := range tests {
		actual := determineSkinSensitivity(test.Input)
		if actual != test.Want {
			t.Errorf("expected %s, actual %s", test.Want, actual)
		}
	}
}
