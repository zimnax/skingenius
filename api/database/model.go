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
