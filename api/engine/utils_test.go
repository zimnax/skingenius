package engine

import (
	"fmt"
	"skingenius/database/model"
	"testing"
)

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
