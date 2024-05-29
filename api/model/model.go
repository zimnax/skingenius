package model

type QuizAnswers struct {
	SkinType    string   `json:"q1"`
	SkinReact   string   `json:"q2"`
	Breakouts   string   `json:"q3"`
	ProductPref []string `json:"q4"`
	FreeFrom    []string `json:"q5"`
	SkinConcern string   `json:"q6"`
	Age         string   `json:"q7"`
	ProductGoal string   `json:"q8"`
}
