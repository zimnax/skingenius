package model

import "fmt"

func PreferencesCtxKey(prefId uint) string {
	return fmt.Sprintf("ingredient_pref_score_%d", prefId)
}

func SkintypeCtxKey(skintypeId uint) string {
	return fmt.Sprintf("skintype_score_%d", skintypeId)
}

func SkinsensetivityCtxKey(skinsensId uint) string {
	return fmt.Sprintf("skinsens_score_%d", skinsensId)
}

func AcnebreakoutsCtxKey(acnebreakoutId uint) string {
	return fmt.Sprintf("acnebreakouts_score_%d", acnebreakoutId)
}

func AllergiesCtxKey(allergyId uint) string {
	return fmt.Sprintf("allergies_score_%d", allergyId)
}

func SkinconcernCtxKey(concernId uint) string {
	return fmt.Sprintf("skinconcern_score_%d", concernId)
}

func AgeCtxKey(ageId uint) string {
	return fmt.Sprintf("age_score_%d", ageId)
}

func BenefitsCtxKey(benefitId uint) string {
	return fmt.Sprintf("benefits_score_%d", benefitId)
}

var SkinSensitivityToDbValue = map[string]string{
	"never":      "No",
	"rarely":     "No",
	"sometimes":  "No",
	"often":      "Yes",
	"frequently": "Yes",
}

var SkinAcneToDbValue = map[string]string{
	"rarely":          "No",
	"occasionally":    "No",
	"frequently":      "Yes",
	"very_frequently": "Yes",
	"almost_always":   "Yes",
}

var SkinConcernToDbValue = map[string][]string{
	"df":   {"dryness"},
	"ov":   {"oiliness", "visible_pores"},
	"ud":   {"uneven_skin_tone", "dark_spots"},
	"swf":  {"signs_of_aging", "wrinkles", "fine_lines"},
	"ab":   {"acne", "blemishes"},
	"repd": {"redness", "eczema", "psoriasis", "dermatitis"},
	"du":   {"dullness", "uneven_texture"},
	"ds":   {"damaged_skin", "sunburned_skin"},
	"dc":   {"dark_circles"},
	"ss":   {"sensitive_skin"},
}

var AgeToDbValue = map[string]string{
	"10s": "teen",
	"20s": "age20s",
	"30s": "age30s",
	"40s": "age40s",
	"50s": "age50s",
	"60s": "age60s",
}

var SkinBenefitsToDbValue = map[string][]string{
	"mnh":  {"moisturizing", "nourishing", "hydrating"},
	"ebi":  {"exfoliating", "brightening", "improves_skin_tone"},
	"sm":   {"smoothing", "minimize_pores"},
	"srr":  {"smoothing", "reduces_wrinkles", "reduces_signs_of_aging"},
	"rr":   {"reduces_acne", "reduces_blemishes"},
	"cshr": {"calming", "soothing", "healing", "reduces_inflammation"},
	"ides": {"improves_symptoms_of_psoriasis", "improves_symptoms_of_dermatitis", "improves_symptoms_of_eczema", "moisturizing", "cooling"},
}
