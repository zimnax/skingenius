package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"skingenius/database"
	"skingenius/database/model"
	"skingenius/logger"
	"strings"
	"time"
)

const packageLogPrefix = "admin: "

func storeProducts(ctx context.Context, dbClient database.Connector, filepath string) {
	records := readCsvFile(filepath)

	//var logs []string
	var mp []string

	for i, record := range records {
		if i == 0 { // skip headers
			continue
		}

		logger.New().Info(context.Background(), packageLogPrefix+fmt.Sprintf("Processing product [%s]", record[ProductName]))
		currentProduct := model.Product{
			Name:            strings.ToLower(record[ProductName]),
			Brand:           record[ProductBrand],
			Ingredients:     nil,
			Link:            record[ProductLink],
			Type:            record[ProductType],
			FormulationType: record[FormulationType],
			FormulatedFor:   record[FormulatedFor],
			Price:           priceToFloat64(strings.ReplaceAll(record[ProductPrice], "$", "")),
			Image:           "",
			Description:     record[ProductDescription],
		}

		for ingredientIndex, ingredientName := range strings.Split(record[ProductIngredients], ";") {
			ingredientNameToFind := strings.ToLower(ingredientName)
			ingredientNameToFind = strings.ReplaceAll(ingredientNameToFind, "*", "")
			ingredientNameToFind = strings.TrimSpace(ingredientNameToFind)

			ingredient, err := dbClient.FindIngredientByAlias(ctx, ingredientNameToFind)
			if err != nil {
				logger.New().Warn(context.Background(), packageLogPrefix+fmt.Sprintf("FATAL, no ingredient with name [%s]", ingredientNameToFind))
				mp = append(mp, ingredientNameToFind)
				continue
			}

			//ingredient.IndexNumber = uint(ingredientIndex)

			ctx = context.WithValue(ctx, model.IngredientIndexCtxKey(ingredient.ID), ingredientIndex)
			currentProduct.Ingredients = append(currentProduct.Ingredients, *ingredient)
		}

		//imgBase64, err := image.ReadImageToBase64V2("admin/resources/product_pictures/" + strings.TrimSpace(currentProduct.Name) + ".jpg")

		//logs = append(logs, fmt.Sprintf("Product: [%s] Missing products: %#v ", currentProduct.Name, missingProducts))

		if saveErr := dbClient.SaveProduct(ctx, &currentProduct); saveErr != nil {
			logger.New().Fatal(context.Background(), packageLogPrefix+fmt.Sprintf("failed to save product [%s], error: %v", currentProduct.Name, saveErr))
			continue
		}

		logger.New().Info(context.Background(), packageLogPrefix+fmt.Sprintf("Stored a new product #%d : %s", i, currentProduct.Name))
	}

	//fmt.Println(len(logs))
	fmt.Println(fmt.Sprintf("%#v", mp))
}

//func storeProducts(ctx context.Context, dbClient database.Connector, filepath string) {
//	records := readCsvFile(filepath)
//
//	currentProduct := model.Product{
//		Name: " ",
//	}
//	first := true
//
//	var missingImages []string
//
//	for i, record := range records {
//		//if record[ProductName] != "Petitgrain Face Moisturizer" {
//		//	continue
//		//}
//
//		if i >= 1 { // skip headers
//			productName := strings.ToLower(record[ProductName])
//
//			// next product from ingredient csv table
//			if currentProduct.Name != productName {
//				fmt.Println(fmt.Sprintf("Creating a new product: %s", productName))
//
//				if productName == "" {
//					continue
//				}
//
//				if !first {
//					time.Sleep(200 * time.Millisecond) // delay before saving next
//
//					if saveErr := dbClient.SaveProduct(ctx, &currentProduct); saveErr != nil {
//						fmt.Println(fmt.Sprintf("failed to save product [%s], error: %v", currentProduct.Name, saveErr))
//						continue
//					}
//					fmt.Println(fmt.Sprintf("product [%s] saved", currentProduct.Name))
//					first = true
//				}
//
//				imgBase64, err := image.ReadImageToBase64V2("admin/resources/product_pictures/" + strings.TrimSpace(productName) + ".jpg")
//				if err != nil {
//					missingImages = append(missingImages, productName)
//				}
//
//				currentProduct = model.Product{
//					Name:  productName,
//					Brand: record[ProductBrand],
//					Link:  record[ProductLink],
//					Image: imgBase64,
//				}
//
//				first = false
//			}
//
//			// continue with same product from table, add ingredient to the product
//			var ingredient *model.Ingredient
//			var err error
//			ingredientNameToFind := strings.ToLower(record[ProductIngredientName])
//			ingredientNameToFind = strings.ReplaceAll(ingredientNameToFind, "*", "")
//			ingredientNameToFind = strings.TrimSpace(ingredientNameToFind)
//
//			ingredient, err = dbClient.FindIngredientByName(ctx, ingredientNameToFind)
//
//			fmt.Println(fmt.Sprintf("Ingredient [%s] found: %v", ingredientNameToFind, ingredient))
//			fmt.Println(fmt.Sprintf("Ingredient [%s] found error: %v", ingredientNameToFind, err))
//
//			if err != nil {
//				fmt.Println(fmt.Sprintf("failed to find igredient by name [%s], trying to find by alias", ingredientNameToFind))
//				ingredient, err = dbClient.FindIngredientByAlias(ctx, ingredientNameToFind)
//				if err != nil {
//					fmt.Println(fmt.Sprintf("failed to find igredient by alias [%s], error: %v", ingredientNameToFind, err))
//					ingredient, err = dbClient.FindIngredientByINCIName(ctx, ingredientNameToFind)
//					if err != nil {
//						fmt.Println(fmt.Sprintf("failed to find igredient by INCI name [%s], error: %v", ingredientNameToFind, err))
//
//						fmt.Println(fmt.Sprintf("FATAL, no ingredient with name [%s]", ingredientNameToFind))
//						continue
//					}
//				}
//			}
//
//			currentProduct.Ingredients = append(currentProduct.Ingredients, *ingredient)
//			fmt.Println(fmt.Sprintf("%d - product %s : added ingredient %s", i, currentProduct.Name, ingredient.Name))
//		}
//	}
//
//	fmt.Println(fmt.Sprintf("Missing images: %#v ", missingImages))
//}

func storeIngredients(ctx context.Context, dbClient database.Connector, filepath string, updateExisting bool) {
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

			name := strings.ReplaceAll(strings.TrimSpace(strings.ToLower(record[IngredientName])), "*", "")
			logger.New().Info(context.Background(), packageLogPrefix+fmt.Sprintf("Processing ingredient [%s]", name))

			//if name != "bentonite" {
			//	continue
			//}

			dbIngredient, findErr := dbClient.FindIngredientByName(context.Background(), name)
			fmt.Println(fmt.Sprintf("Ingredient [%s] found: %v", name, dbIngredient))
			fmt.Println(fmt.Sprintf("Ingredient [%s] found error: %v", name, findErr))

			if findErr == nil && !updateExisting {
				fmt.Println(fmt.Sprintf("Ingredient [%s] already exists, upadeExisting:%t ingredient: [%v]", name, updateExisting, dbIngredient))
				continue
			}

			ctx, ipref := assignPreferencesScore(ctx, record, allPreferences)
			ctx, iskintype := assignSkintypeScore(ctx, record, allskintypes)
			ctx, iskinSens := assignSkinSensitivityScore(ctx, record, allskinsensetivities)
			ctx, iacneBreakouts := assignAcneBreakoutScore(ctx, record, allAcnebreakouts)
			ctx, iallergies := assignAllergyScore(ctx, record, allAllergies)
			ctx, iskinConcerns := assignSkinConcernScore(ctx, record, allSkinconcerns)
			ctx, iages := assignAgeScore(ctx, record, allAges)
			ctx, ibenefits := assignBenefitsScore(ctx, record, allBenefits)

			aliases := strings.Split(record[Aliases], ";")
			for n, alias := range aliases {
				aliases[n] = strings.ToLower(strings.TrimSpace(alias))
			}
			aliases = append(aliases, name) // Add ingredient name to aliases for search optimization

			// in case if  ingredient does not exist, dbClient would return empty object
			if dbIngredient.Name != "" {
				logger.New().Info(context.Background(), packageLogPrefix+fmt.Sprintf("Found existing ingredient [%s], doing update", name))
			}

			dbIngredient.Name = name
			//dbIngredient.Type = strings.ToLower(record[Active_Inactive])
			//PubchemId: record[PubChemCID],
			//CasNumber: record[CASNumber],
			dbIngredient.ECNumber = ""
			dbIngredient.INCIName = record[INCIName]
			dbIngredient.Synonyms = aliases
			dbIngredient.ConcentrationRinseOffMin, dbIngredient.ConcentrationRinseOffMax = parseConcentration(record[Concentrations])
			dbIngredient.ConcentrationLeaveOnMin, dbIngredient.ConcentrationLeaveOnMax = parseConcentration(record[Concentrations])
			dbIngredient.EffectiveAtLowConcentration = assignEffectiveness(record[Effective_at_low_concentrations])

			dbIngredient.Preferences = ipref
			dbIngredient.Skintypes = iskintype
			dbIngredient.Skinsensitivities = iskinSens
			dbIngredient.Acnebreakouts = iacneBreakouts
			dbIngredient.Allergies = iallergies
			dbIngredient.Skinconcerns = iskinConcerns
			dbIngredient.Ages = iages
			dbIngredient.Benefits = ibenefits

			dbClient.SaveIngredient(ctx, dbIngredient)
			fmt.Println(fmt.Sprintf("Ingredient [%s] saved or updated", dbIngredient.Name))
			time.Sleep(100 * time.Millisecond)
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
