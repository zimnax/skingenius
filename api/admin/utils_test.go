package main

import (
	"fmt"
	"skingenius/database/model"
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

func Test_UniqueIngredients(t *testing.T) {

	is := [][]model.Ingredient{
		{{Name: "name1", Score: 1}, {Name: "name2"}, {Name: "name3"}},
		{{Name: "name1", Score: 1}, {Name: "name2"}, {Name: "name33"}},
		{{Name: "name1", Score: 1}, {Name: "name222"}, {Name: "name333"}},
		{{Name: "name1", Score: 1}, {Name: "name222"}, {Name: "name3333"}},
	}

	actual := uniqueIngredientsNamesMap(is...)
	if len(actual) != 1 {
		t.Fatalf("expected 1, actual %d", len(actual))
	}

	fmt.Println("actual", actual)

}
