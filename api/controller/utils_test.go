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

	actual := findProductsWithActiveIngredients(products, ingredients)
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
		//Ingredients: []frontModel.Ingredient{
		//	{Name: "Water", Index: 0},
		//	{Name: "Caprylic Capric Triglyceride", Index: 1},
		//	{Name: "Cetearyl Glucoside, Glyceryl Stearate, Coffee Arabica Seed Oil (CreamMaker® Green Coffee)", Index: 2},
		//	{Name: "Persea Gratissima (Avocado) Oil, Glycine Soja (Soybean) Lipids, Beeswax (Avocado Butter)", Index: 3},
		//	{Name: "Glycerin", Index: 4},
		//	{Name: "Coffea Arabica (Coffee) Seed Extract", Index: 5},
		//	{Name: "Musa Sapientum (Banana) Leaf/Trunk Extract (AntiMicro Banana)", Index: 6},
		//	{Name: "Hydroxypropyl Guar", Index: 7},
		//	{Name: "Caprylyl Glycol, Ethylhexylglycerin", Index: 8},
		//	{Name: "Papaya Banana Fragrance", Index: 9},
		//},
		Ingredients: []model.Ingredient{
			{Name: "Water", Index: 0},
			{Name: "Caprylic Capric Triglyceride", Index: 1, KnownConcentrationLeaveOnMin: 10, KnownConcentrationLeaveOnMax: 15},
			{Name: "Cetearyl Glucoside, Glyceryl Stearate, Coffee Arabica Seed Oil (CreamMaker® Green Coffee)", Index: 2},
			{Name: "Persea Gratissima (Avocado) Oil, Glycine Soja (Soybean) Lipids, Beeswax (Avocado Butter)", Index: 3},
			{Name: "Glycerin", Index: 4},
			{Name: "Coffea Arabica (Coffee) Seed Extract", Index: 5, KnownConcentrationLeaveOnMin: 4.5, KnownConcentrationLeaveOnMax: 10.5},
			{Name: "Musa Sapientum (Banana) Leaf/Trunk Extract (AntiMicro Banana)", Index: 6},
			{Name: "Hydroxypropyl Guar", Index: 7},
			{Name: "Caprylyl Glycol, Ethylhexylglycerin", Index: 8},
			{Name: "Papaya Banana Fragrance", Index: 9},
		},
	}

	//p2 := frontModel.Product{
	//	Name: "product1",
	//	Ingredients: []frontModel.Ingredient{
	//		{Name: "Water"},
	//		{Name: "Coco - caprylate/caprate"},
	//		{Name: "CreamMaker® Stearate"},
	//		{Name: "Glycerin"},
	//		{Name: "Meadowfoam Seed Oil"},
	//		{Name: "PolyGel Emollient"},
	//		{Name: "Sheabutter Glycerides"},
	//		{Name: "Raspberry Seed Oil"},
	//		{Name: "Ceramide Complex"},
	//		{Name: "72h Moisture"},
	//		{Name: "Capryly Glycol EHG"},
	//		{Name: "Phenoxyethanol"},
	//		{Name: "GelMaker® Hydro"},
	//	},
	//}
	//}

	tt := []struct {
		products []model.Product
	}{
		{
			products: []model.Product{p1},
		},
	}

	for _, tst := range tt {
		for _, product := range tst.products {
			fmt.Println("product ->>> ", product.Name)
			actual := calculateConcentrationsBogdanFormula(product, 0.1)

			if product.Name == "product1" {
				for iName, cVal := range actual {
					if iName == "Coffea Arabica (Coffee) Seed Extract" && cVal != 4.5 {
						t.Fatalf("expected 4.5, actual %f", cVal)
					}
					if iName == "Caprylic Capric Triglyceride" && cVal != 15 {
						t.Fatalf("expected 15, actual %f", cVal)
					}

					fmt.Println(iName, cVal)
				}
			}

		}
	}
}

func Test_findProductIntersection(t *testing.T) {
	productsByConcern := map[string][]model.Product{
		"acne": {
			{ID: 1, Name: "Product A"},
			{ID: 2, Name: "Product B"},
			{ID: 3, Name: "Product C"},
		},
		"wrinkles": {
			{ID: 2, Name: "Product B"},
			{ID: 3, Name: "Product C"},
			{ID: 4, Name: "Product D"},
		},
		"dryness": {
			{ID: 1, Name: "Product A"},
			{ID: 2, Name: "Product B"},
			{ID: 3, Name: "Product C"},
			{ID: 3, Name: "Product D"},
			{ID: 3, Name: "Product E"},
		},
	}

	commonProducts := findProductIntersection(productsByConcern)

	if len(commonProducts) != 2 {
		t.Fatalf("expected 2, actual %d", len(commonProducts))
	}

	fmt.Println("Common Products:")
	for _, product := range commonProducts {
		fmt.Printf("ID: %d, Name: %s\n", product.ID, product.Name)
	}
}

func Test_ValidateConcentrations(t *testing.T) {
	p := model.Product{
		Type: model.MoisturizerProductType,
		Ingredients: []model.Ingredient{
			{Name: "name1", ConcentrationRinseOffMin: 62, Score: 1, Index: 0},   // margin 17.03
			{Name: "name2", ConcentrationRinseOffMin: 32, Score: 1, Index: 1},   // margin  7.13
			{Name: "name3", ConcentrationRinseOffMin: 22, Score: 1, Index: 2},   // margin  3.86
			{Name: "name4", ConcentrationRinseOffMin: 12, Score: 0.6, Index: 3}, // margin  2.92
		},
		Concentrations: map[string]float64{
			"name1": 44, // less than min
			"name2": 24, // less than min
			"name3": 20, // more than min
			"name4": 12, // equal to min
		},
		ActiveIngredients: []string{"name1", "name2", "name3", "name4"},
	}

	adjustedScoreProducts := validateConcentrations(p)

	if len(adjustedScoreProducts.Ingredients) != 4 {
		t.Fatalf("expected 4, actual %d", len(adjustedScoreProducts.Ingredients))
	}

	for _, ing := range adjustedScoreProducts.Ingredients {
		if ing.Name == "name1" && ing.Score != 0.2 {
			t.Fatalf("expected 0.6, actual %f", ing.Score)
		}
		if ing.Name == "name2" && ing.Score != 0.2 {
			t.Fatalf("expected 0.6, actual %f", ing.Score)
		}
		if ing.Name == "name3" && ing.Score != 1 {
			t.Fatalf("expected 1, actual %f", ing.Score)
		}

	}
}
