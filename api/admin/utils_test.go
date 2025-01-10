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

func Test_parseConcentration(t *testing.T) {

	tt := []struct {
		value       string
		expectedMin float64
		expectedMax float64
	}{
		{
			value:       "0.01-1%",
			expectedMin: 0.01,
			expectedMax: 1,
		},
		{
			value:       "0.01-0.03%",
			expectedMin: 0.01,
			expectedMax: 0.03,
		},
	}

	for _, s := range tt {
		min, max := parseConcentration(s.value)
		if min != s.expectedMin || max != s.expectedMax {
			t.Fatalf("expected min %f, actual %f, expected max %f, actual %f, value: %s", s.expectedMin, min, s.expectedMax, max, s.value)
		}
	}
}
