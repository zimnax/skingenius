package database

import (
	"context"
	"fmt"
	"os"
	"skingenius/config"
	"skingenius/database/model"
	"testing"
	"time"
)

var env = "test"

func Test_FindTop3ByIds(t *testing.T) {
	db, err := NewGormClient(config.LocalHost, config.Port, config.User, config.Password, false, env)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	ps, err := db.FindProductsByIds(context.Background(), []int32{100, 89, 85})

	fmt.Println(err)
	fmt.Println(fmt.Sprintf("products len:: %d", len(ps)))
	fmt.Println(fmt.Sprintf("0: %d", len(ps[0].Ingredients)))
}

func Test_FindIngredientByAlias(t *testing.T) {
	db, err := NewGormClient(config.LocalHost, config.Port, config.User, config.Password, false, env)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	ing, err := db.FindIngredientByAlias(context.Background(), "ACV")

	fmt.Println(err)
	fmt.Println(fmt.Sprintf("ingredient:: %v", ing))
}

func Test_FindExistingIngredient(t *testing.T) {
	db, err := NewGormClient(config.LocalHost, config.Port, config.User, config.Password, false, env)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}
	//butter butyrospermum parkii (shea) butter

	//if saveErr := db.SaveIngredient(context.Background(), &model.Ingredient{Name: "testIngredient"}); saveErr != nil {
	//	t.Fatalf("failed to save ingredient, error:%v", saveErr)
	//}

	ing, findErr := db.FindIngredientByName(context.Background(), "butter butyrospermum parkii (shea) butter")
	if ing != nil {
		fmt.Println(fmt.Sprintf("ingredient:: %#v", ing))
	}

	if findErr != nil {
		t.Fatalf("error should be nil, error: %v", findErr)
	}
}

func Test_SaveExistingIngredient(t *testing.T) {
	db, err := NewGormClient(config.LocalHost, config.Port, config.User, config.Password, false, env)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	i := model.Ingredient{Name: "testIngredient", INCIName: "inciName_1"}

	if saveErr := db.SaveIngredient(context.Background(), &i); saveErr != nil {
		t.Fatalf("failed to save ingredient, error:%v", saveErr)
	}

	i.INCIName = "inciName_2"
	if saveErr := db.SaveIngredient(context.Background(), &i); saveErr != nil {
		t.Fatalf("failed to save ingredient, error:%v", saveErr)
	}

	ing, findErr := db.FindIngredientByName(context.Background(), "testIngredient")
	if ing != nil {
		fmt.Println(fmt.Sprintf("ingredient:: %v", ing))
	}

	if findErr != nil {
		t.Fatalf("error should be nil")
	}

	if ing.INCIName != "inciName_2" {
		t.Fatalf("inci name should be updated")
	}

	if deleteErr := db.DeleteIngredientByName(context.Background(), "testIngredient"); deleteErr != nil {
		t.Fatalf("failed to delete ingredient, err: %v", deleteErr)
	}

	removedIng, findErrAfterDelete := db.FindIngredientByName(context.Background(), "testIngredient")
	if removedIng.Name != "" {
		t.Fatalf("ingredient should be nil, current: %v", removedIng)
	}

	if findErrAfterDelete == nil {
		t.Fatalf("error should not be nil")
	}
}

//func Test_SaveIngredient_LowConcentrationTest(t *testing.T) {
//	db, err := NewGormClient(config.LocalHost, config.Port, config.User, config.Password, false, env)
//	if err != nil {
//		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
//		os.Exit(1)
//	}
//
//	i := model.Ingredient{Name: "testIngredient", EffectiveConcentrations: "0.5-10%", EffectiveAtLowConcentration: model.EffectiveYes}
//
//	if saveErr := db.SaveIngredient(context.Background(), &i); saveErr != nil {
//		t.Fatalf("failed to save ingredient, error:%v", saveErr)
//	}
//
//	ing, findErr := db.FindIngredientByName(context.Background(), "testIngredient")
//	if ing != nil {
//		fmt.Println(fmt.Sprintf("ingredient:: %v", ing))
//	}
//
//	if findErr != nil {
//		t.Fatalf("error should be nil")
//	}
//
//	if ing.EffectiveAtLowConcentration != model.EffectiveYes {
//		t.Fatal("EffectiveAtLowConcentration should be Yes")
//	}
//
//	if ing.EffectiveConcentrations != "0.5-10%" {
//		t.Fatal("EffectiveAtLowConcentration should be 0.5-10%")
//	}
//
//	fmt.Println(fmt.Sprintf("ingredient:: %v", ing))
//}

/*
intention of this test is to check is the ingredients are being retrieved with skinconcerns
*/
func TestFindProductsByIds(t *testing.T) {
	db, err := NewGormClient(config.LocalHost, config.Port, config.User, config.Password, false, env)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	ps, err := db.FindProductsByIds(context.Background(), []int32{5})

	fmt.Println(err)
	fmt.Println(fmt.Sprintf("products len:: %d", len(ps)))
	fmt.Println(fmt.Sprintf("ps[0].Ingredients: %#v", ps[0].Ingredients))
}

func Test_GetSkinconcernDescriptionByIngredients(t *testing.T) {
	db, err := NewGormClient(config.RemoteHost, config.Port, config.User, config.Password, false, env)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	ing, err := db.GetSkinConcernDescriptionByIngredients(context.Background(), []string{"algae extract", "glycerin"}, "oiliness_shine")

	fmt.Println(err)
	fmt.Println(fmt.Sprintf("ingredient:: %#v", ing))
}

func TestGormConnector_SaveRecommendations(t *testing.T) {
	db, err := NewGormClient(config.LocalHost, config.Port, config.User, config.Password, false, env)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	ur := []model.UserRecommendations{
		{UserId: "1", ProductId: 11, Score: 11.1},
		{UserId: "1", ProductId: 22, Score: 22.2},
		{UserId: "1", ProductId: 33, Score: 33.3},
	}

	err = db.SaveRecommendations(context.Background(), ur)
	if err != nil {
		t.Fatal(err)
	}

	rec, err := db.GetRecommendations(context.Background(), "1")
	if err != nil {
		t.Fatal(err)
	}

	if len(rec) != 3 {
		t.Fatalf("expected 3 recommendations, got %d", len(rec))
	}

	fmt.Println(rec)
}

func TestGormConnector_FindIngredientByINCIName(t *testing.T) {
	db, err := NewGormClient(config.LocalHost, config.Port, config.User, config.Password, false, env)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	i, err := db.FindIngredientByINCIName(context.Background(), "3-O-Ethyl Ascorbic Acid")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(i)
}

func Test_FuzzySearchWithLike(t *testing.T) {
	db, err := NewGormClient(config.LocalHost, config.Port, config.User, config.Password, false, env)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	products, err := db.LiveSearch(context.Background(), "mois")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(products)
}

func Test_SaveUserRoutine(t *testing.T) {
	db, err := NewGormClient(config.RemoteHost, config.Port, config.User, config.Password, true, env)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	ps := []model.Product{
		{Name: "product01"},
		{Name: "product02"},
		//{Name: "product03"},
		//{Name: "product04"},
	}

	var savedProducts []model.Product
	for _, p := range ps {
		if err := db.SaveProduct(context.Background(), &p); err != nil {
			t.Fatal(err)
		}

		dbProduct, err := db.FindProductByName(context.Background(), p.Name)
		if err != nil {
			t.Fatal(err)
		}

		savedProducts = append(savedProducts, *dbProduct)
	}

	defer func() {
		for _, sp := range savedProducts {
			fmt.Printf("Delete user routine by id: %d \n", sp.ID)
			if delErr := db.DeleteUserRoutine(context.Background(), "1", int(sp.ID)); delErr != nil {
				t.Fatalf("failed to remove userRoutine, error: %v ", delErr)
			}
		}

		for _, p := range ps {
			fmt.Println("running cleanup from product: ", p.Name)
			if err := db.DeleteProductByName(context.Background(), p.Name); err != nil {
				t.Fatal(err)
			}
		}
	}()

	// start of the test

	for _, savedProduct := range savedProducts {
		r := model.UserRoutine{
			UserId:    "1",
			ProductID: savedProduct.ID,
			//Product:     savedProduct,
			TimeOfDay:   "Day",
			TimesPerDay: "once",
			HowLong:     "1 month",
			Note:        "user note",
		}

		fmt.Println(fmt.Sprintf("saving routine: %#v", r))

		err = db.SaveUserRoutine(context.Background(), r)
		if err != nil {
			t.Fatal(err)
		}

	}

	ur, err := db.GetUserRoutine(context.Background(), "1")
	if len(ur) != 2 {
		for _, routine := range ur {
			fmt.Println(fmt.Sprintf("routine %#v", routine))
		}

		t.Fatalf("expected 2 routine products, got %d", len(ur))
	}

	for _, routine := range ur {
		fmt.Println(fmt.Sprintf("routine %#v", routine))
	}

	time.Sleep(4 * time.Second)

}

func Test_FindAllProductsHavingIngredients(t *testing.T) {
	db, err := NewGormClient(config.LocalHost, config.Port, config.User, config.Password, true, "test")
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	defer func() {
		//for _, p := range ps {
		//	fmt.Println("running cleanup from product: ", p.Name)
		//	if err := db.DeleteProductByName(context.Background(), p.Name); err != nil { // delete doesn't work because it just marks it as deleted
		//		t.Fatal(err)
		//	}
		//}

		if err := db.DROPDB(); err != nil {
			t.Fatal(err)
		}
	}()

	ps := []model.Product{
		{ID: 1, Name: "product01", Ingredients: []model.Ingredient{{ID: 6, Name: "Ing1"}, {ID: 2, Name: "Ing2"}, {ID: 3, Name: "Ing3"}, {ID: 4, Name: "Ing4"}, {ID: 5, Name: "Ing5"}}}, // has all ingredients + 1
		{ID: 2, Name: "product02", Ingredients: []model.Ingredient{{ID: 6, Name: "Ing1"}, {ID: 2, Name: "Ing2"}, {ID: 3, Name: "Ing3"}, {ID: 4, Name: "Ing4"}}},                        // has all ingredients
		{ID: 3, Name: "product002", Ingredients: []model.Ingredient{{ID: 6, Name: "Ing1"}, {ID: 2, Name: "Ing2"}, {ID: 3, Name: "Ing3"}, {ID: 7, Name: "Ing7"}}},                       // has all ingredients
		{ID: 4, Name: "product03", Ingredients: []model.Ingredient{{ID: 2, Name: "Ing2"}, {ID: 3, Name: "Ing3"}, {ID: 4, Name: "Ing4"}, {ID: 5, Name: "Ing5"}}},                        // missing 1
		{ID: 5, Name: "product04", Ingredients: []model.Ingredient{{ID: 6, Name: "Ing1"}, {ID: 3, Name: "Ing3"}, {ID: 4, Name: "Ing4"}, {ID: 5, Name: "Ing5"}}},                        // missing 1
	}

	var savedProducts []model.Product
	for _, p := range ps {
		if err := db.SaveProduct(context.Background(), &p); err != nil {
			t.Fatal(err)
		}

		dbProduct, err := db.FindProductByName(context.Background(), p.Name)
		if err != nil {
			t.Fatal(err)
		}

		savedProducts = append(savedProducts, *dbProduct)
	}

	productsWithCommonIngredients, err := db.FindAllProductsHavingIngredients(context.Background(), []string{"Ing1", "Ing2", "Ing3", "Ing4", "IngEXSTRA", "Ing7"})

	if err != nil {
		t.Fatal(err)
	}

	for _, ingredient := range productsWithCommonIngredients {
		fmt.Println(fmt.Sprintf("AllProductsHavingIngredients: %#v", ingredient))
	}

	if len(productsWithCommonIngredients) != 2 {
		t.Fatalf("Expecting 2, actual %d", len(productsWithCommonIngredients))
	}
}

func TestGormConnector_GetIngredientsBySkinconcerns(t *testing.T) {
	db, err := NewGormClient(config.LocalHost, config.Port, config.User, config.Password, true, "test")
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	ings, err := db.GetIngredientsBySkinconcerns(context.Background(), []string{"acne"})

	for _, ing := range ings {
		fmt.Println(fmt.Sprintf("ingredient:: %#v", ing))
	}
	fmt.Println(err)

}

func TestGormConnector_SaveIngredient(t *testing.T) {
	db, err := NewGormClient(config.LocalHost, config.Port, config.User, config.Password, true, "test")
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	//defer func() {
	//	if err := db.DROPDB(); err != nil {
	//		t.Fatal(err)
	//	}
	//}()

	ing := model.Ingredient{
		Name:                        "testIngredient",
		INCIName:                    "inciName_1",
		EffectiveAtLowConcentration: model.EffectiveYes,
		Roleinformulations: []model.Roleinformulation{
			{ID: 5, Name: model.Emollient},
			{ID: 1, Name: model.Active},
		},
	}

	if saveErr := db.SaveIngredient(context.Background(), &ing); saveErr != nil {
		//t.Fatalf("failed to save ingredient, error:%v", saveErr)
	}

	ing2, findErr := db.FindIngredientByName(context.Background(), "testIngredient")
	if ing2 != nil {
		fmt.Println(fmt.Sprintf("ingredient:: %v", ing2))
	}

	if findErr != nil {
		t.Fatalf("error should be nil")
	}

	fmt.Println(fmt.Sprintf("ingredient:: %#v", ing2))

	if ing2.EffectiveAtLowConcentration != model.EffectiveYes {
		t.Fatal("EffectiveAtLowConcentration should be Yes")
	}

	if len(ing2.Roleinformulations) != 2 {
		t.Fatalf("expected 2 roles, got %d", len(ing2.Roleinformulations))
	}
}
