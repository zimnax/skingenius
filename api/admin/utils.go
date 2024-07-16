package main

import (
	"context"
	"fmt"
	"skingenius/database/model"
	"strconv"
	"strings"
)

func yesNoTo01(val string) int {
	val = strings.ToLower(val)

	if val == "yes" {
		return 1
	}

	return 0 // in any other case - Default value
}

func assignPreferencesScore(ctx context.Context, record []string, allPreferences []model.Preference) (context.Context, []model.Preference) {
	ipref := allPreferences
	for _, ipreference := range ipref {
		var score int

		switch ipreference.Name {
		case model.Paleo:
			score = yesNoTo01(record[Paleo])
		case model.Vegetarian:
			score = yesNoTo01(record[Vegetarian])
		case model.Vegan:
			score = yesNoTo01(record[Vegan])
		case model.GlutenFree:
			score = yesNoTo01(record[GlutenFree])
		}

		ctx = context.WithValue(ctx, model.PreferencesCtxKey(ipreference.ID), score)
	}

	return ctx, ipref
}

func assignSkintypeScore(ctx context.Context, record []string, allskintypes []model.Skintype) (context.Context, []model.Skintype) {
	iSkintypes := allskintypes
	for _, iSkintype := range iSkintypes {

		var score int
		var err error

		switch iSkintype.Type {
		case model.Dry:
			score, err = strconv.Atoi(record[Dry])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skintype score: [%s] for skintype Dry", record[Dry]))
			}
		case model.Normal:
			score, err = strconv.Atoi(record[Normal])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skintype score: [%s] for skintype Normal", record[Normal]))
			}
		case model.Combination:
			score, err = strconv.Atoi(record[Combination])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skintype score: [%s] for skintype Combination", record[Combination]))
			}
		case model.Oily:
			score, err = strconv.Atoi(record[Oily])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skintype score: [%s] for skintype Oily", record[Oily]))
			}
		}

		ctx = context.WithValue(ctx, model.SkintypeCtxKey(iSkintype.ID), score)
	}

	return ctx, iSkintypes
}

func assignSkinSensitivityScore(ctx context.Context, record []string, allSkinSens []model.Skinsensitivity) (context.Context, []model.Skinsensitivity) {
	iSkinsens := allSkinSens

	for _, skinsen := range iSkinsens {
		var score int
		var err error

		switch skinsen.Sensitivity {
		case model.Never:
			score, err = strconv.Atoi(record[NotSensitive])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skinSensetivity score: [%s] for skinSensetivity Never(NotSensitive)", record[NotSensitive]))
			}
		case model.Rarely:
			score, err = strconv.Atoi(record[ALittleSensitive])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skinSensetivity score: [%s] for skinSensetivity Rarely(ALittleSensitive)", record[ALittleSensitive]))
			}
		case model.Sometimes:
			score, err = strconv.Atoi(record[ModeratelySensitive])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skinSensetivity score: [%s] for skinSensetivity Sometimes(ModeratelySensitive)", record[ModeratelySensitive]))
			}
		case model.Often:
			score, err = strconv.Atoi(record[Sensitive])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skinSensetivity score: [%s] for skinSensetivity Often(Sensitive)", record[Sensitive]))
			}
		case model.Frequently:
			score, err = strconv.Atoi(record[VerySensitive])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skinSensetivity score: [%s] for skinSensetivity Frequently(VerySensitive)", record[VerySensitive]))
			}
		case model.Always:
			score, err = strconv.Atoi(record[ExtremelySensitive])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skinSensetivity score: [%s] for skinSensetivity Always(ExtremelySensitive)", record[ExtremelySensitive]))
			}
		}

		ctx = context.WithValue(ctx, model.SkinsensetivityCtxKey(skinsen.ID), score)

	}
	return ctx, iSkinsens
}

func assignAcneBreakoutScore(ctx context.Context, record []string, allAcneBreakouts []model.Acnebreakout) (context.Context, []model.Acnebreakout) {
	iAcneBreakout := allAcneBreakouts

	for _, acnebreakout := range iAcneBreakout {
		var score int
		var err error

		switch acnebreakout.Frequency {
		case model.NeverAcne:
			score, err = strconv.Atoi(record[NotAcneProne])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast acneBreakout score: [%s] for acneBreakout NeverAcne(NotAcneProne)", record[NotAcneProne]))
			}
		case model.RarelyAcne:
			score, err = strconv.Atoi(record[ALittleAcneProne])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast acneBreakout score: [%s] for acneBreakout RarelyAcne(ALittleAcneProne)", record[ALittleAcneProne]))
			}
		case model.Occasionally:
			score, err = strconv.Atoi(record[ModeratelyAcneProne])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast acneBreakout score: [%s] for acneBreakout Occasionally(ModeratelyAcneProne)", record[ModeratelyAcneProne]))
			}
		case model.FrequentlyAcne:
			score, err = strconv.Atoi(record[AcneProne])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast acneBreakout score: [%s] for acneBreakout FrequentlyAcne(AcneProne)", record[AcneProne]))
			}
		case model.VeryFrequently:
			score, err = strconv.Atoi(record[VeryAcneProne])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast acneBreakout score: [%s] for acneBreakout VeryFrequently(VeryAcneProne)", record[VeryAcneProne]))
			}
		case model.AlmostAlways:
			score, err = strconv.Atoi(record[ExtremelyAcneProne])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast acneBreakout score: [%s] for acneBreakout AlmostAlways(ExtremelyAcneProne)", record[ExtremelyAcneProne]))
			}
		}

		ctx = context.WithValue(ctx, model.AcnebreakoutsCtxKey(acnebreakout.ID), score)
	}
	return ctx, iAcneBreakout
}

func assignAllergyScore(ctx context.Context, record []string, allAllergy []model.Allergy) (context.Context, []model.Allergy) {
	iallergies := allAllergy
	for _, iallergy := range iallergies {
		var score int

		switch iallergy.Name {
		case model.AllergyNuts:
			score = yesNoTo01(record[NutFree])
		case model.AllergySoy:
			score = yesNoTo01(record[SoyFree])
		case model.AllergyLatex:
			score = yesNoTo01(record[LatexFree])
		case model.AllergySesame:
			score = yesNoTo01(record[SesameFree])
		case model.AllergyCitrus:
			score = yesNoTo01(record[CitrusFree])
		case model.AllergyDye:
			score = yesNoTo01(record[DyeFree])
		case model.AllergyArtificialFragrance:
			score = yesNoTo01(record[FragranceFree])
		case model.AllergyScent:
			score = yesNoTo01(record[ScentFree])
		}

		ctx = context.WithValue(ctx, model.AllergiesCtxKey(iallergy.ID), score)
	}

	return ctx, iallergies
}

func assignSkinConcernScore(ctx context.Context, record []string, allSkinconcern []model.Skinconcern) (context.Context, []model.Skinconcern) {
	iconcerns := allSkinconcern
	for _, concern := range iconcerns {
		var stringScore string

		switch concern.Name {
		case model.ConcernAcne:
			stringScore = record[Acne]
		case model.ConcernRosacea:
			stringScore = record[Rosacea]
		case model.ConcernCysticAcne:
			stringScore = record[CysticAcne]
		case model.ConcernHyperpigmentation:
			stringScore = record[Hyperpigmentation]
		case model.ConcernMelasma:
			stringScore = record[Melasma]
		case model.ConcernXerosis:
			stringScore = record[Xerosis]
		case model.ConcernDryness:
			stringScore = record[Dryness]
		case model.ConcernRedness:
			stringScore = record[Redness]
		case model.ConcernOiliness:
			stringScore = record[Oiliness]
		case model.ConcernUnevenSkinTone:
			stringScore = record[UnevenSkinTone]
		case model.ConcernSignsOfAging:
			stringScore = record[SignsOfAging]
		case model.ConcernFineLines:
			stringScore = record[FineLines]
		case model.ConcernWrinkles:
			stringScore = record[Wrinkles]
		case model.ConcernDarkSpots:
			stringScore = record[DarkSpots]
		case model.ConcernLostOfElasticityFirmness:
			stringScore = record[LossOfElasticityFirmness]
		case model.ConcernVisiblePores:
			stringScore = record[VisiblePores]
		case model.ConcernCloggedPoresBlackheads:
			stringScore = record[CloggedPoresBlackheads]
		case model.ConcernDullness:
			stringScore = record[Dullness]
		case model.ConcernDamagedSkin:
			stringScore = record[DamagedSkin]
		case model.ConcernUnevenTexture:
			stringScore = record[UnevenTexture]
		case model.ConcernEczema:
			stringScore = record[Eczema]
		case model.ConcernPsoriasis:
			stringScore = record[Psoriasis]
		case model.ConcernDermatitis:
			stringScore = record[Dermatitis]
		case model.ConcernSunburnedSkin:
			stringScore = record[SunburnedSkin]
		case model.ConcernDarkCircles:
			stringScore = record[DarkCircles]
		case model.ConcernBlemishes:
			stringScore = record[Blemishes]
		case model.ConcernSensitiveSkin:
			stringScore = record[SensitiveSkin]
		}

		stringScore = strings.ReplaceAll(stringScore, " ", "")
		if stringScore == "" {
			stringScore = "0" // Default velue
		}

		score, err := strconv.Atoi(stringScore)
		if err != nil {
			fmt.Println(fmt.Sprintf("failed to cast skinConcern score: [%s] for skinConcern %s", stringScore, concern.Name))
		}

		ctx = context.WithValue(ctx, model.SkinconcernCtxKey(concern.ID), score)
	}

	return ctx, iconcerns
}

func assignAgeScore(ctx context.Context, record []string, allAges []model.Age) (context.Context, []model.Age) {
	iAllAges := allAges
	for _, age := range iAllAges {
		var score int

		switch age.Value {
		case model.Age10:
			score = yesNoTo01(record[Teen])
		case model.Age20:
			score = yesNoTo01(record[Twenties])
		case model.Age30:
			score = yesNoTo01(record[Thirties])
		case model.Age40:
			score = yesNoTo01(record[Forties])
		case model.Age50:
			score = yesNoTo01(record[Fifties])
		case model.Age60:
			score = yesNoTo01(record[SixtiesPlus])
		}

		ctx = context.WithValue(ctx, model.AgeCtxKey(age.ID), score)
	}

	return ctx, iAllAges
}

func assignBenefitsScore(ctx context.Context, record []string, allBenefits []model.Benefit) (context.Context, []model.Benefit) {
	iAllBenefits := allBenefits
	for _, benefit := range iAllBenefits {
		var stringScore string

		switch benefit.Name {
		case model.BenefitMoisturizing:
			stringScore = record[Moisturizing]
		case model.BenefitNourishing:
			stringScore = record[Nourishing]
		case model.BenefitHydrating:
			stringScore = record[Hydrating]
		case model.BenefitExfoliating:
			stringScore = record[Exfoliating]
		case model.BenefitCalming:
			stringScore = record[Calming]
		case model.BenefitSoothing:
			stringScore = record[Soothing]
		case model.BenefitUVBarrier:
			stringScore = record[UVBarrier]
		case model.BenefitHealing:
			stringScore = record[Healing]
		case model.BenefitSmoothing:
			stringScore = record[Smoothing]
		case model.BenefitReducesAcne:
			stringScore = record[ReducesAcne]
		case model.BenefitReducesBlemishes:
			stringScore = record[ReducesBlemishes]
		case model.BenefitReducesWrinkles:
			stringScore = record[ReducesWrinkles]
		case model.BenefitImprovesSymptomsOfEczema:
			stringScore = record[ImprovesSymptomsOfEczema]
		case model.BenefitImprovesSymptomsOfPsoriasis:
			stringScore = record[ImprovesSymptomsOfPsoriasis]
		case model.BenefitImprovesSymptomsOfDermatitis:
			stringScore = record[ImprovesSymptomsOfDermatitis]
		case model.BenefitBrightening:
			stringScore = record[Brightening]
		case model.BenefitImprovesSkinTone:
			stringScore = record[ImprovesSkinTone]
		case model.BenefitReducesInflammation:
			stringScore = record[ReducesInflammation]
		case model.BenefitMinimizesPores:
			stringScore = record[MinimizesPores]
		case model.BenefitAntiAging:
			stringScore = record[AntiAging]
		case model.BenefitFirming:
			stringScore = record[Firming]
		case model.BenefitDetoxifying:
			stringScore = record[Detoxifying]
		case model.BenefitBalancing:
			stringScore = record[Balancing]
		case model.BenefitReducesRedness:
			stringScore = record[ReducesRedness]
		case model.BenefitClarifying:
			stringScore = record[Clarifying]
		case model.BenefitAntiBacterial:
			stringScore = record[AntiBacterial]
		case model.BenefitStimulatesCollagenProduction:
			stringScore = record[StimulatesCollagenProduction]
		case model.BenefitReducesFineLines:
			stringScore = record[ReducesFineLines]
		case model.BenefitAntioxidantProtection:
			stringScore = record[AntioxidantProtection]
		case model.BenefitSkinBarrierProtection:
			stringScore = record[SkinBarrierProtection]

		}

		stringScore = strings.ReplaceAll(stringScore, " ", "")
		if stringScore == "" {
			stringScore = "0" // Default velue
		}

		score, err := strconv.Atoi(stringScore)
		if err != nil {
			fmt.Println(fmt.Sprintf("failed to cast skinConcern score: [%s] for skinConcern %s", stringScore, benefit.Name))
		}

		ctx = context.WithValue(ctx, model.BenefitsCtxKey(benefit.ID), score)
	}

	return ctx, iAllBenefits
}

func mergeIngredients(ingredients ...[]model.Ingredient) []string {

	// Initialize a map to track unique elements
	unique := make(map[string]bool)

	// Function to append unique elements from a slice to the map
	appendUnique := func(slice []model.Ingredient) {
		for _, ing := range slice {
			unique[ing.Name] = true
		}
	}

	for _, ingredientList := range ingredients {
		appendUnique(ingredientList)
	}

	// Create a slice from the unique map keys
	var merged []string
	for str := range unique {
		merged = append(merged, str)
	}

	return merged
}

func uniqueIngredientsNamesList(ingredients ...[]model.Ingredient) []string {
	// Create a map to count occurrences of each string
	countMap := make(map[string]int)

	// Populate the count map
	for _, slice := range ingredients {
		unique := make(map[string]bool)
		for _, ing := range slice {
			if !unique[ing.Name] {
				countMap[ing.Name]++
				unique[ing.Name] = true
			}
		}
	}

	// Find strings that appear in all slices (count == len(ingredients))
	var result []string
	for str, count := range countMap {
		if count == len(ingredients) {
			result = append(result, str)
		}
	}

	return result
}

func uniqueIngredientsNamesMap(ingredients ...[]model.Ingredient) map[string]int {
	countMap := make(map[string]int)
	scoreMap := make(map[string]int)

	for _, answerIngredients := range ingredients {
		for _, ingredient := range answerIngredients {
			countMap[ingredient.Name]++
			scoreMap[ingredient.Name] = scoreMap[ingredient.Name] + ingredient.Score
		}
	}

	// Find strings that appear in all slices (count == len(ingredients))
	result := make(map[string]int)
	for str, count := range countMap {
		if count == len(ingredients) {
			result[str] = scoreMap[str]
		}
	}

	return result
}

func mergeIngredientsWithScores(ingredients ...[]model.Ingredient) map[string]int {
	imap := make(map[string]int)

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

func IngredientsSliceToMap(is []model.Ingredient) map[string]int {
	m := make(map[string]int)

	for _, i := range is {
		m[i.Name] = i.Score
	}

	return m
}
