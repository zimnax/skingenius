package model

type QuizAnswers struct {
	SkinType           string   `json:"q1"`
	SkinSensitivity    string   `json:"q2"`
	AcneBreakouts      string   `json:"q3"`
	ProductPreferences []string `json:"q4"`
	FreeFromAllergens  []string `json:"q5"`
	SkinConcern        []string `json:"q6"`
	Age                string   `json:"q7"`
	ProductBenefit     []string `json:"q8"`
}
