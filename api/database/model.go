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
