package frontModel

import "skingenius/database/model"

var SkinTypeMapping = map[string]string{
	"Normal":      "normal",
	"Dry":         "dry",
	"Oily":        "oily",
	"Combination": "combination",
}

var SensitivityMapping = map[string]string{
	"Not sensitive":        "never",
	"A little sensitive":   "rarely",
	"Moderately sensitive": "sometimes",
	"Sensitive":            "often",
	"Very sensitive":       "frequently",
	"Extremely sensitive":  "always",
}

var AcneProneMapping = map[string]string{
	"Not acne prone":        "never",
	"A little acne-prone":   "rarely",
	"Moderately acne-prone": "occasionally",
	"Acne-prone":            "frequently",
	"Very acne-prone":       "veryfrequently",
	"Extremely acne-prone":  "always",
}

var AgeMapping = map[string]int{
	"Teen": 10,
	"20s":  20,
	"30s":  30,
	"40s":  40,
	"50s":  50,
	"60s":  60,
}

var PreferenceMapping = map[string]model.IngredientPreferencesValue{
	"Vegetarian":    model.Vegetarian,
	"Vegan":         model.Vegan,
	"Gluten-free":   model.GlutenFree,
	"Paleo":         model.Paleo,
	"No Preference": model.NoPreference,
}

var AllergiesMapping = map[string]model.AllergyValue{
	"Nut-free":           model.AllergyNuts,
	"Soy-free":           model.AllergySoy,
	"Latex-free":         model.AllergyLatex,
	"Sesame-free":        model.AllergySesame,
	"Citrus-free":        model.AllergyCitrus,
	"Dye-free":           model.AllergyDye,
	"Fragrance-free":     model.AllergyArtificialFragrance,
	"Scent-free":         model.AllergyScent,
	"No Known Allergies": model.AllergyNone,
}

var SkinConcernsMapping = map[string]model.SkinconcernValue{
	"Rosacea":                     model.ConcernRosacea,
	"Hyperpigmentation":           model.ConcernHyperpigmentation_UnevenSkinTone,
	"Acne":                        model.ConcernAcne,
	"Dryness":                     model.ConcernDryness_Dehydration,
	"Oiliness":                    model.ConcernOiliness_Shine,
	"Fine lines":                  model.ConcernFine_lines_Wrinkles,
	"Loss of Elasticity/firmness": model.ConcernLoss_of_Elasticity_firmness,
	"Visible pores":               model.ConcernVisible_pores_Uneven_texture,
	"Clogged pores, blackheads":   model.ConcernClogged_pores_blackheads,
	"Dullness":                    model.ConcernDullness,
	"Dark circles":                model.ConcernDark_circles,
	"Blemishes":                   model.ConcernBlemishes,
	"No Concerns":                 model.ConcernNone,

	//"Rosacea":                     ConcernRosacea,
	//"Hyperpigmentation":           ConcernHyperpigmentation,
	//"Melasma":                     ConcernMelasma,
	//"Cystic Acne":                 ConcernCysticAcne,
	//"Xerosis":                     ConcernXerosis,
	//"Dryness":                     ConcernDryness,
	//"Oiliness":                    ConcernOiliness,
	//"Uneven skin tone":            ConcernUnevenSkinTone,
	//"Signs of Aging":              ConcernSignsOfAging,
	//"Fine lines":                  ConcernFineLines,
	//"Wrinkles":                    ConcernWrinkles,
	//"Dark Spots":                  ConcernDarkSpots,
	//"Loss of Elasticity/firmness": ConcernLostOfElasticityFirmness,
	//"Visible pores":               ConcernVisiblePores,
	//"Clogged pores, blackheads":   ConcernCloggedPoresBlackheads,
	//"Redness":                     ConcernRedness,
	//"Dullness":                    ConcernDullness,
	//"Damaged skin":                ConcernDamagedSkin,
	//"Uneven texture":              ConcernUnevenTexture,
	//"Eczema":                      ConcernEczema,
	//"Psoriasis":                   ConcernPsoriasis,
	//"Dermatitis":                  ConcernDermatitis,
	//"Sunburned skin":              ConcernSunburnedSkin,
	//"Dark circles":                ConcernDarkCircles,
	//"Blemishes":                   ConcernBlemishes,
	//"Sensitive skin":              ConcernSensitiveSkin,
	//"No Concerns":                 ConcernNone,
}

var BenefitsMapping = map[string]model.BenefitValue{
	"Moisturizing":                    model.BenefitMoisturizing,
	"Nourishing":                      model.BenefitNourishing,
	"Hydrating":                       model.BenefitHydrating,
	"Exfoliating":                     model.BenefitExfoliating,
	"Calming":                         model.BenefitCalming,
	"Soothing":                        model.BenefitSoothing,
	"UV-barrier":                      model.BenefitUVBarrier,
	"Healing":                         model.BenefitHealing,
	"Smoothing":                       model.BenefitSmoothing,
	"Reduces acne":                    model.BenefitReducesAcne,
	"Reduces blemishes":               model.BenefitReducesBlemishes,
	"Reduces wrinkles":                model.BenefitReducesWrinkles,
	"Improves symptoms of eczema":     model.BenefitImprovesSymptomsOfEczema,
	"Improves symptoms of psoriasis":  model.BenefitImprovesSymptomsOfPsoriasis,
	"Improves symptoms of dermatitis": model.BenefitImprovesSymptomsOfDermatitis,
	"Brightening":                     model.BenefitBrightening,
	"Improves skin tone":              model.BenefitImprovesSkinTone,
	"Reduces inflammation":            model.BenefitReducesInflammation,
	"Minimizes pores":                 model.BenefitMinimizesPores,
	"Anti-aging":                      model.BenefitAntiAging,
	"Firming":                         model.BenefitFirming,
	"Detoxifying":                     model.BenefitDetoxifying,
	"Balancing":                       model.BenefitBalancing,
	"Reduces redness":                 model.BenefitReducesRedness,
	"Clarifying":                      model.BenefitClarifying,
	"Anti-bacterial":                  model.BenefitAntiBacterial,
	"Stimulates collagen production":  model.BenefitStimulatesCollagenProduction,
	"Reduces fine lines":              model.BenefitReducesFineLines,
	"Antioxidant Protection":          model.BenefitAntioxidantProtection,
	"Skin Barrier Protection":         model.BenefitSkinBarrierProtection,
}
