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

func formatBool(val string) bool {
	val = strings.ToLower(val)
	if val == "yes" {
		return true
	}

	return false
}

func assignPreferencesScore(ctx context.Context, record []string, allPreferences []model.Preference) (context.Context, []model.Preference) {
	ipref := allPreferences
	for _, ipreference := range ipref {
		var score bool

		switch ipreference.Name {
		case model.Paleo:
			score = formatBool(record[Paleo])
		case model.Vegetarian:
			score = formatBool(record[Vegetarian])
		case model.Vegan:
			score = formatBool(record[Vegan])
		case model.GlutenFree:
			score = formatBool(record[GlutenFree])
		}

		ctx = context.WithValue(ctx, model.PreferencesCtxKey(ipreference.ID), score)
	}

	return ctx, ipref
}

func assignSkintypeScore(ctx context.Context, record []string, allskintypes []model.Skintype) (context.Context, []model.Skintype) {
	iSkintypes := allskintypes
	for _, iSkintype := range iSkintypes {

		var score float64
		var err error

		switch iSkintype.Type {
		case model.Dry:
			//score, err = strconv.Atoi(record[Dry])
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Dry]), 32)
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skintype score: [%s] for skintype Dry", record[Dry]))
			}
		case model.Normal:
			//score, err = strconv.Atoi(record[Normal])
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Normal]), 32)
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skintype score: [%s] for skintype Normal", record[Normal]))
			}
		case model.Combination:
			//score, err = strconv.Atoi(record[Combination])
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Combination]), 32)
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skintype score: [%s] for skintype Combination", record[Combination]))
			}
		case model.Oily:
			//score, err = strconv.Atoi(record[Oily])
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Oily]), 32)
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
		var score bool
		var err error

		switch skinsen.Sensitivity {
		case model.Never:
			score = formatBool(record[NotSensitive])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skinSensetivity score: [%s] for skinSensetivity Never(NotSensitive)", record[NotSensitive]))
			}
		case model.Rarely:
			score = formatBool(record[ALittleSensitive])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skinSensetivity score: [%s] for skinSensetivity Rarely(ALittleSensitive)", record[ALittleSensitive]))
			}
		case model.Sometimes:
			score = formatBool(record[ModeratelySensitive])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skinSensetivity score: [%s] for skinSensetivity Sometimes(ModeratelySensitive)", record[ModeratelySensitive]))
			}
		case model.Often:
			score = formatBool(record[Sensitive])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skinSensetivity score: [%s] for skinSensetivity Often(Sensitive)", record[Sensitive]))
			}
		case model.Frequently:
			score = formatBool(record[VerySensitive])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast skinSensetivity score: [%s] for skinSensetivity Frequently(VerySensitive)", record[VerySensitive]))
			}
		case model.Always:
			score = formatBool(record[ExtremelySensitive])
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
		var score bool
		var err error

		switch acnebreakout.Frequency {
		case model.NeverAcne:
			score = formatBool(record[NotAcneProne])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast acneBreakout score: [%s] for acneBreakout NeverAcne(NotAcneProne)", record[NotAcneProne]))
			}
		case model.RarelyAcne:
			score = formatBool(record[ALittleAcneProne])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast acneBreakout score: [%s] for acneBreakout RarelyAcne(ALittleAcneProne)", record[ALittleAcneProne]))
			}
		case model.Occasionally:
			score = formatBool(record[ModeratelyAcneProne])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast acneBreakout score: [%s] for acneBreakout Occasionally(ModeratelyAcneProne)", record[ModeratelyAcneProne]))
			}
		case model.FrequentlyAcne:
			score = formatBool(record[AcneProne])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast acneBreakout score: [%s] for acneBreakout FrequentlyAcne(AcneProne)", record[AcneProne]))
			}
		case model.VeryFrequently:
			score = formatBool(record[VeryAcneProne])
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to cast acneBreakout score: [%s] for acneBreakout VeryFrequently(VeryAcneProne)", record[VeryAcneProne]))
			}
		case model.AlmostAlways:
			score = formatBool(record[ExtremelyAcneProne])
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
		var score bool

		switch iallergy.Name {
		case model.AllergyNuts:
			score = formatBool(record[NutFree])
		case model.AllergySoy:
			score = formatBool(record[SoyFree])
		case model.AllergyLatex:
			score = formatBool(record[LatexFree])
		case model.AllergySesame:
			score = formatBool(record[SesameFree])
		case model.AllergyCitrus:
			score = formatBool(record[CitrusFree])
		case model.AllergyDye:
			score = formatBool(record[DyeFree])
		case model.AllergyArtificialFragrance:
			score = formatBool(record[FragranceFree])
		case model.AllergyScent:
			score = formatBool(record[ScentFree])
		}

		ctx = context.WithValue(ctx, model.AllergiesCtxKey(iallergy.ID), score)
	}

	return ctx, iallergies
}

func assignSkinConcernScore(ctx context.Context, record []string, allSkinconcern []model.Skinconcern) (context.Context, []model.Skinconcern) {
	iconcerns := allSkinconcern
	for _, concern := range iconcerns {
		var score float64
		var err error

		switch concern.Name {
		case model.ConcernAcne:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Acne]), 32)
		case model.ConcernRosacea:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Rosacea]), 32)
		case model.ConcernHyperpigmentation_UnevenSkinTone:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Hyperpigmentation_UnevenSkin_tone]), 32)
		case model.ConcernDryness_Dehydration:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Dryness_Dehydration]), 32)
		case model.ConcernOiliness_Shine:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Oiliness_Shine]), 32)
		case model.ConcernFine_lines_Wrinkles:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Fine_lines_Wrinkles]), 32)
		case model.ConcernLoss_of_Elasticity_firmness:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Loss_of_Elasticity_firmness]), 32)
		case model.ConcernVisible_pores_Uneven_texture:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Visible_pores_Uneven_texture]), 32)
		case model.ConcernClogged_pores_blackheads:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Clogged_pores_blackheads]), 32)
		case model.ConcernDullness:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Dullness]), 32)
		case model.ConcernDark_circles:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Dark_circles]), 32)
		case model.ConcernBlemishes:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Blemishes]), 32)
		}

		if err != nil {
			fmt.Println(fmt.Sprintf("failed to cast skinConcern score, err: %v", err))
		}

		//Conce	rnNone                             SkinconcernValue = "no_concern"

		ctx = context.WithValue(ctx, model.SkinconcernCtxKey(concern.ID), score)
	}

	return ctx, iconcerns
}

func assignAgeScore(ctx context.Context, record []string, allAges []model.Age) (context.Context, []model.Age) {
	iAllAges := allAges
	for _, age := range iAllAges {
		var score bool

		switch age.Value {
		case model.Age10:
			score = formatBool(record[Teen])
		case model.Age20:
			score = formatBool(record[Twenties])
		case model.Age30:
			score = formatBool(record[Thirties])
		case model.Age40:
			score = formatBool(record[Forties])
		case model.Age50:
			score = formatBool(record[Fifties])
		case model.Age60:
			score = formatBool(record[SixtiesPlus])
		}

		ctx = context.WithValue(ctx, model.AgeCtxKey(age.ID), score)
	}

	return ctx, iAllAges
}

func assignBenefitsScore(ctx context.Context, record []string, allBenefits []model.Benefit) (context.Context, []model.Benefit) {
	iAllBenefits := allBenefits
	for _, benefit := range iAllBenefits {
		//var stringScore string
		var score float64
		var err error

		switch benefit.Name {
		case model.BenefitMoisturizing:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Moisturizing]), 32)
		case model.BenefitNourishing:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Nourishing]), 32)
		case model.BenefitHydrating:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Hydrating]), 32)
		case model.BenefitExfoliating:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Exfoliating]), 32)
		case model.BenefitCalming:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Calming]), 32)
		case model.BenefitSoothing:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Soothing]), 32)
		case model.BenefitUVBarrier:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[UVBarrier]), 32)
		case model.BenefitHealing:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Healing]), 32)
		case model.BenefitSmoothing:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Smoothing]), 32)
		case model.BenefitReducesAcne:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[ReducesAcne]), 32)
		case model.BenefitReducesBlemishes:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[ReducesBlemishes]), 32)
		case model.BenefitReducesWrinkles:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[ReducesWrinkles]), 32)
		case model.BenefitImprovesSymptomsOfEczema:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[ImprovesSymptomsOfEczema]), 32)
		case model.BenefitImprovesSymptomsOfPsoriasis:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[ImprovesSymptomsOfPsoriasis]), 32)
		case model.BenefitImprovesSymptomsOfDermatitis:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[ImprovesSymptomsOfDermatitis]), 32)
		case model.BenefitBrightening:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Brightening]), 32)
		case model.BenefitImprovesSkinTone:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[ImprovesSkinTone]), 32)
		case model.BenefitReducesInflammation:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[ReducesInflammation]), 32)
		case model.BenefitMinimizesPores:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[MinimizesPores]), 32)
		case model.BenefitAntiAging:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[AntiAging]), 32)
		case model.BenefitFirming:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Firming]), 32)
		case model.BenefitDetoxifying:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Detoxifying]), 32)
		case model.BenefitBalancing:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Balancing]), 32)
		case model.BenefitReducesRedness:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[ReducesRedness]), 32)
		case model.BenefitClarifying:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[Clarifying]), 32)
		case model.BenefitAntiBacterial:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[AntiBacterial]), 32)
		case model.BenefitStimulatesCollagenProduction:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[StimulatesCollagenProduction]), 32)
		case model.BenefitReducesFineLines:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[ReducesFineLines]), 32)
		case model.BenefitAntioxidantProtection:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[AntioxidantProtection]), 32)
		case model.BenefitSkinBarrierProtection:
			score, err = strconv.ParseFloat(strings.TrimSpace(record[SkinBarrierProtection]), 32)
		}

		if err != nil {
			fmt.Println(fmt.Sprintf("failed to cast benefit score, err: %v", err))
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
