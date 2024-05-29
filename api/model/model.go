package model

type QuizAnswers struct {
	Q1 string   `json:"q1"`
	Q2 string   `json:"q2"`
	Q3 string   `json:"q3"`
	Q4 []string `json:"q4"`
	Q5 []string `json:"q5"`
	Q6 string   `json:"q6"`
	Q7 string   `json:"q7"`
	Q8 string   `json:"q8"`
}
