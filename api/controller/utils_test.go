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

	actual, _ := uniqueIngredientsNamesMap(is...)
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

func Test_filterProductsWithIngredients(t *testing.T) {
	products := []model.Product{
		{
			Name: "product1",
			Ingredients: []model.Ingredient{
				{Name: "ingredient1"},
				{Name: "ingredient2"},
			},
		},

		{
			Name: "product2",
			Ingredients: []model.Ingredient{
				{Name: "ingredient2"},
				{Name: "ingredient4"},
			},
		},
		{
			Name: "product3",
			Ingredients: []model.Ingredient{
				{Name: "ingredient5"},
				{Name: "ingredient6"},
			},
		},
		{
			Name: "product4",
			Ingredients: []model.Ingredient{
				{Name: "ingredient2"},
				{Name: "ingredient7"},
			},
		},
	}

	ingredients := []model.Ingredient{
		{Name: "ingredient1"},
		{Name: "ingredient2"},
	}

	actual := filterProductsWithIngredients(products, ingredients)
	//if len(actual) != 0 {
	//	t.Fatalf("expected 0, actual %d", len(actual))
	//}

	fmt.Println("actual len", len(actual))
}

/*
Water
Caprylic Capric Triglyceride
Cetearyl Glucoside, Glyceryl Stearate, Coffee Arabica Seed Oil (CreamMaker® Green Coffee)
Persea Gratissima (Avocado) Oil, Glycine Soja (Soybean) Lipids, Beeswax (Avocado Butter)
Glycerin
Coffea Arabica (Coffee) Seed Extract
Musa Sapientum (Banana) Leaf/Trunk Extract (AntiMicro Banana)
Hydroxypropyl Guar
Caprylyl Glycol, Ethylhexylglycerin
Papaya Banana Fragrance
*/
func Test_CalculateConcentrations(t *testing.T) {
	p1 := model.Product{
		Name: "product1",
		Ingredients: []model.Ingredient{
			{Name: "Water"},
			{Name: "Caprylic Capric Triglyceride"},
			{Name: "Cetearyl Glucoside, Glyceryl Stearate, Coffee Arabica Seed Oil (CreamMaker® Green Coffee)"},
			{Name: "Persea Gratissima (Avocado) Oil, Glycine Soja (Soybean) Lipids, Beeswax (Avocado Butter)"},
			{Name: "Glycerin"},
			{Name: "Coffea Arabica (Coffee) Seed Extract"},
			{Name: "Musa Sapientum (Banana) Leaf/Trunk Extract (AntiMicro Banana)"},
			{Name: "Hydroxypropyl Guar"},
			{Name: "Caprylyl Glycol, Ethylhexylglycerin"},
			{Name: "Papaya Banana Fragrance"},
		},
	}

	p2 := model.Product{
		Name: "product1",
		Ingredients: []model.Ingredient{
			{Name: "Water"},
			{Name: "Coco - caprylate/caprate"},
			{Name: "CreamMaker® Stearate"},
			{Name: "Glycerin"},
			{Name: "Meadowfoam Seed Oil"},
			{Name: "PolyGel Emollient"},
			{Name: "Sheabutter Glycerides"},
			{Name: "Raspberry Seed Oil"},
			{Name: "Ceramide Complex"},
			{Name: "72h Moisture"},
			{Name: "Capryly Glycol EHG"},
			{Name: "Phenoxyethanol"},
			{Name: "GelMaker® Hydro"},
		},
	}

	tt := []struct {
		products []model.Product
	}{
		{
			products: []model.Product{p1, p2},
		},
	}

	for _, tst := range tt {
		for _, product := range tst.products {
			fmt.Println("product ->>> ", product.Name)
			actual := calculateConcentrationsBogdanFormula(product, 0.1)
			for s, s2 := range actual {
				fmt.Println(s, s2)
			}
		}
	}
}
