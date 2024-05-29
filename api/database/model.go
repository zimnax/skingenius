package database

var skinSensitivityToDbValue = map[string]string{
	"never":      "No",
	"rarely":     "No",
	"sometimes":  "No",
	"often":      "Yes",
	"frequently": "Yes",
}

var skinAcneToDbValue = map[string]string{
	"rarely":          "No",
	"occasionally":    "No",
	"frequently":      "Yes",
	"very_frequently": "Yes",
	"almost_always":   "Yes",
}

var skinConcernToDbValue = map[string][]string{
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

var ageToDbValue = map[string]string{
	"10s": "teen",
	"20s": "age20s",
	"30s": "age30s",
	"40s": "age40s",
	"50s": "age50s",
	"60s": "age60s",
}

var skinBenefitsToDbValue = map[string][]string{
	"mnh":  {"moisturizing", "nourishing", "hydrating"},
	"ebi":  {"exfoliating", "brightening", "improves_skin_tone"},
	"sm":   {"smoothing", "minimize_pores"},
	"srr":  {"smoothing", "reduces_wrinkles", "reduces_signs_of_aging"},
	"rr":   {"reduces_acne", "deduces_blemishes"},
	"cshr": {"calming", "soothing", "healing", "reduces_inflammation"},
	"ides": {"improves_symptoms_of_psoriasis", "improves_symptoms_of_dermatitis", "improves_symptoms_of_eczema", "sunburn"},
}
