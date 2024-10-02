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

func SkinconcernDescCtxKey(concernId uint) string {
	return fmt.Sprintf("skinconcern_description_%d", concernId)
}

func AgeCtxKey(ageId uint) string {
	return fmt.Sprintf("age_score_%d", ageId)
}

func BenefitsCtxKey(benefitId uint) string {
	return fmt.Sprintf("benefits_score_%d", benefitId)
}

var SkinConcernToUserFriendlyDescription = map[string][]string{
	string(ConcernAcne):                             []string{"Clear skin", "Reduced inflammation", "Consistent skin tone"},
	string(ConcernRosacea):                          []string{"Reduced redness", "Even skin tone", "Healthy glow"},
	string(ConcernDryness_Dehydration):              []string{"Soft and smooth texture", "Improved skin barrier", "Enhanced hydration"},
	string(ConcernHyperpigmentation_UnevenSkinTone): []string{"Even skin tone", "Reduced appearance of dark spots", "Clear, healthy skin"},
	string(ConcernOiliness_Shine):                   []string{"Matte & shine-free skin", "Balanced complexion", "Smooth & even texture"},
	string(ConcernFine_lines_Wrinkles):              []string{"Firm and toned skin", "Reduced appearance of fine lines & wrinkles", "Radiance and hydration"},
	string(ConcernLoss_of_Elasticity_firmness):      []string{"Tight and toned skin", "Smooth and event texture", "Younger look"},
	string(ConcernVisible_pores_Uneven_texture):     []string{"Even and uniform skin texture", "Reduced appearance of pores", "Tight & glowing skin"},
	string(ConcernClogged_pores_blackheads):         []string{"Clear and smooth skin", "Refined pores", "Balanced and hydrated skin"},
	string(ConcernDullness):                         []string{"Even skin tone", "Healthy glow", "Hydrated and nourished skin"},
	string(ConcernDark_circles):                     []string{"Reduced appearance of dark circles", "Even texture", "Healthy & glowing skin"},
	string(ConcernBlemishes):                        []string{"Healed, healthy skin", "Even texture & tone", "Balanced complexion"},
}

var BenefitsToUserFriendlyDescription = map[string][]string{
	string(BenefitMoisturizing):                 []string{"Well-moisturized skin"},
	string(BenefitNourishing):                   []string{"Enhanced skin health & strong barrier"},
	string(BenefitHydrating):                    []string{"Enhanced hydration"},
	string(BenefitExfoliating):                  []string{"Enhanced skin renewal & better texture"},
	string(BenefitCalming):                      []string{"Reduced redness & stronger barrier"},
	string(BenefitSoothing):                     []string{"Reduced sensitivity & faster healing"},
	string(BenefitUVBarrier):                    []string{"Healthy and youthful skin"},
	string(BenefitHealing):                      []string{"Regenerated & healthy skin"},
	string(BenefitSmoothing):                    []string{"Better texture and radiance"},
	string(BenefitReducesAcne):                  []string{"Healing & reduction of breakouts"},
	string(BenefitReducesBlemishes):             []string{"Healing & reduced redness"},
	string(BenefitReducesWrinkles):              []string{"Reduce appearance of wrinkles"},
	string(BenefitImprovesSymptomsOfEczema):     []string{"??"},
	string(BenefitImprovesSymptomsOfPsoriasis):  []string{"??"},
	string(BenefitImprovesSymptomsOfDermatitis): []string{"??"},
	string(BenefitBrightening):                  []string{"Enhanced radiance"},
	string(BenefitImprovesSkinTone):             []string{"Better skin tone"},
	string(BenefitReducesInflammation):          []string{"??"},
	string(BenefitMinimizesPores):               []string{"Smooth skin texture & reduction of pores"},
	string(BenefitAntiAging):                    []string{"??"},
	string(BenefitFirming):                      []string{"Improved elasticity"},
	string(BenefitDetoxifying):                  []string{"??"},
	string(BenefitBalancing):                    []string{"??"},
	string(BenefitReducesRedness):               []string{"??"},
	string(BenefitClarifying):                   []string{"??"},
	string(BenefitAntiBacterial):                []string{"??"},
	string(BenefitStimulatesCollagenProduction): []string{"??"},
	string(BenefitReducesFineLines):             []string{"Smoother, younger skin"},
	string(BenefitAntioxidantProtection):        []string{"??"},
	string(BenefitSkinBarrierProtection):        []string{"??"},
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
