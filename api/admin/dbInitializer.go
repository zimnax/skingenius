package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"skingenius/database"
	"skingenius/database/model"
	"time"
)

func storeProducts(ctx context.Context, dbClient database.Connector, filepath string) {
	records := readCsvFile(filepath)

	currentProduct := model.Product{
		Name: " ",
	}
	first := true

	for i, record := range records {
		//if record[ProductName] == "Perfect Facial Hydrating Cream" {

		if i >= 1 { // skip headers
			if currentProduct.Name != record[ProductName] {
				fmt.Println(fmt.Sprintf("Creating a new product: %s", record[ProductName]))

				if !first {
					time.Sleep(200 * time.Millisecond) // delay before saving next

					if saveErr := dbClient.SaveProduct(ctx, &currentProduct); saveErr != nil {
						fmt.Println(fmt.Sprintf("failed to save product [%s], error: %v", currentProduct.Name, saveErr))
						continue
					}
					fmt.Println(fmt.Sprintf("product [%s] saved", currentProduct.Name))
					first = true
				}

				currentProduct = model.Product{
					Name:  record[ProductName],
					Brand: record[ProductBrand],
					Link:  record[ProductLink],
				}

				first = false
			}

			ingredient, err := dbClient.FindIngredientByName(ctx, record[ProductIngredientName])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to find igredient by name [%s]", record[ProductIngredientName]))
				fmt.Println(err)
				continue
			}
			currentProduct.Ingredients = append(currentProduct.Ingredients, *ingredient)
			fmt.Println(fmt.Sprintf("%d - product %s : added ingredient %s", i, currentProduct.Name, ingredient.Name))
		}
	}
}

func storeIngredients(ctx context.Context, dbClient database.Connector, filepath string) {
	var err error

	if err = dbClient.SetupJoinTables(); err != nil {
		os.Exit(1)
		return
	}

	allPreferences, err := dbClient.GetAllPreferences(ctx)
	allskintypes, err := dbClient.GetAllSkintypes(ctx)
	allskinsensetivities, err := dbClient.GetAllSkinsensetivity(ctx)
	allAcnebreakouts, err := dbClient.GetAllAcneBreakouts(ctx)
	allAllergies, err := dbClient.GetAllAllergies(ctx)
	allSkinconcerns, err := dbClient.GetAllSkinconcerns(ctx)
	allAges, err := dbClient.GetAllAge(ctx)
	allBenefits, err := dbClient.GetAllBenefits(ctx)

	records := readCsvFile(filepath)
	for i, record := range records {

		if i >= 2 { // skip headers
			fmt.Println(record)
			ctx := context.WithValue(context.Background(), "key", "val")

			ctx, ipref := assignPreferencesScore(ctx, record, allPreferences)
			ctx, iskintype := assignSkintypeScore(ctx, record, allskintypes)
			ctx, iskinSens := assignSkinSensitivityScore(ctx, record, allskinsensetivities)
			ctx, iacneBreakouts := assignAcneBreakoutScore(ctx, record, allAcnebreakouts)
			ctx, iallergies := assignAllergyScore(ctx, record, allAllergies)
			ctx, iskinConcerns := assignSkinConcernScore(ctx, record, allSkinconcerns)
			ctx, iages := assignAgeScore(ctx, record, allAges)
			ctx, ibenefits := assignBenefitsScore(ctx, record, allBenefits)

			ingredient := model.Ingredient{
				Name:      record[IngredientName],
				PubchemId: record[PubChemCID],
				CasNumber: record[CASNumber],
				ECNumber:  "",
				Synonyms:  []string{},

				Preferences:       ipref,
				Skintypes:         iskintype,
				Skinsensitivities: iskinSens,
				Acnebreakouts:     iacneBreakouts,
				Allergies:         iallergies,
				Skinconcerns:      iskinConcerns,
				Ages:              iages,
				Benefits:          ibenefits,
			}

			dbClient.SaveIngredient(ctx, &ingredient)
			fmt.Println(fmt.Sprintf("Ingredient [%s] saved", ingredient.Name))
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func readCsvFile(filePath string) [][]string {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current directory: ", dir)

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, " ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
