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

type DBAnswerModel struct {
	SkinType        string
	SkinSensitivity string
	AcneProne       string
	Age             int
	Preferences     []string
	Allergies       []string
	Concerns        []string
	Benefits        []string
}

type SaveRecommendationsReq struct {
	Products []Recommendation `json:"recommendations"`
}
type Recommendation struct {
	Id    int     `json:"id"`
	Score float64 `json:"score"`
}

type UserRoutine struct {
	Products    []int32
	TimeOfDay   string
	TimesPerDay int
	HowLong     string
	Note        string
}
