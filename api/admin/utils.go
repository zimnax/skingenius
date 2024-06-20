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
