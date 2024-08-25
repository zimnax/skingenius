package model

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

var PreferenceMapping = map[string]IngredientPreferencesValue{
	"Vegetarian":    Vegetarian,
	"Vegan":         Vegan,
	"Gluten-free":   GlutenFree,
	"Paleo":         Paleo,
	"No Preference": NoPreference,
}

var AllergiesMapping = map[string]AllergyValue{
	"Nut-free":           AllergyNuts,
	"Soy-free":           AllergySoy,
	"Latex-free":         AllergyLatex,
	"Sesame-free":        AllergySesame,
	"Citrus-free":        AllergyCitrus,
	"Dye-free":           AllergyDye,
	"Fragrance-free":     AllergyArtificialFragrance,
	"Scent-free":         AllergyScent,
	"No Known Allergies": AllergyNone,
}

var SkinConcernsMapping = map[string]SkinconcernValue{
	"Rosacea":                     ConcernRosacea,
	"Hyperpigmentation":           ConcernHyperpigmentation_UnevenSkinTone,
	"Acne":                        ConcernAcne,
	"Dryness":                     ConcernDryness_Dehydration,
	"Oiliness":                    ConcernOiliness_Shine,
	"Fine lines":                  ConcernFine_lines_Wrinkles,
	"Loss of Elasticity/firmness": ConcernLoss_of_Elasticity_firmness,
	"Visible pores":               ConcernVisible_pores_Uneven_texture,
	"Clogged pores, blackheads":   ConcernClogged_pores_blackheads,
	"Dullness":                    ConcernDullness,
	"Blemishes":                   ConcernBlemishes,
	"No Concerns":                 ConcernNone,


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

var BenefitsMapping = map[string]BenefitValue{
	"Moisturizing":                    BenefitMoisturizing,
	"Nourishing":                      BenefitNourishing,
	"Hydrating":                       BenefitHydrating,
	"Exfoliating":                     BenefitExfoliating,
	"Calming":                         BenefitCalming,
	"Soothing":                        BenefitSoothing,
	"UV-barrier":                      BenefitUVBarrier,
	"Healing":                         BenefitHealing,
	"Smoothing":                       BenefitSmoothing,
	"Reduces acne":                    BenefitReducesAcne,
	"Reduces blemishes":               BenefitReducesBlemishes,
	"Reduces wrinkles":                BenefitReducesWrinkles,
	"Improves symptoms of eczema":     BenefitImprovesSymptomsOfEczema,
	"Improves symptoms of psoriasis":  BenefitImprovesSymptomsOfPsoriasis,
	"Improves symptoms of dermatitis": BenefitImprovesSymptomsOfDermatitis,
	"Brightening":                     BenefitBrightening,
	"Improves skin tone":              BenefitImprovesSkinTone,
	"Reduces inflammation":            BenefitReducesInflammation,
	"Minimizes pores":                 BenefitMinimizesPores,
	"Anti-aging":                      BenefitAntiAging,
	"Firming":                         BenefitFirming,
	"Detoxifying":                     BenefitDetoxifying,
	"Balancing":                       BenefitBalancing,
	"Reduces redness":                 BenefitReducesRedness,
	"Clarifying":                      BenefitClarifying,
	"Anti-bacterial":                  BenefitAntiBacterial,
	"Stimulates collagen production":  BenefitStimulatesCollagenProduction,
	"Reduces fine lines":              BenefitReducesFineLines,
	"Antioxidant Protection":          BenefitAntioxidantProtection,
	"Skin Barrier Protection":         BenefitSkinBarrierProtection,
}
