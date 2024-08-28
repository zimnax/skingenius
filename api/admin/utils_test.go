package main

import (
	"testing"
)

func Test_yesNoTo01(t *testing.T) {

	tt := []struct {
		value         string
		expectedScore int
	}{
		{
			value:         "Yes",
			expectedScore: 1,
		},
		{
			value:         "yes",
			expectedScore: 1,
		},
		{
			value:         "No",
			expectedScore: 0,
		},
		{
			value:         "no",
			expectedScore: 0,
		},
		{
			value:         "Yes",
			expectedScore: 1,
		},
		{
			value:         "Yes/No",
			expectedScore: 0,
		},
	}

	for _, s := range tt {
		if actualScore := yesNoTo01(s.value); actualScore != s.expectedScore {
			t.Fatalf("expected score %d, actual %d, value: %s", s.expectedScore, actualScore, s.value)
		}
	}
}
