package model

const ScalingRepresentation = 97

type QuizAnswers struct {
	SkinType           string   `json:"skintype"`
	SkinSensitivity    string   `json:"skinsensitivity"`
	AcneBreakouts      string   `json:"acnebreakouts"`
	ProductPreferences []string `json:"preferences"`
	FreeFromAllergens  []string `json:"allergens"`
	SkinConcern        []string `json:"concerns"`
	Age                string   `json:"age"`
	ProductBenefit     []string `json:"benefits"`
}

type SaveRecommendationsReq struct {
	Products []Recommendation `json:"recommendations"`
}
type Recommendation struct {
	Id    int     `json:"id"`
	Score float64 `json:"score"`
}
