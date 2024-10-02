package database

import (
	"context"
	"fmt"
	"os"
	"skingenius/config"
	"skingenius/database/model"
	"testing"
)

func Test_FindTop3ByIds(t *testing.T) {
	db, err := NewGormClient(config.Host, config.Port, config.User, config.Password, false)
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
	db, err := NewGormClient(config.Host, config.Port, config.User, config.Password, false)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	ing, err := db.FindIngredientByAlias(context.Background(), "ACV")

	fmt.Println(err)
	fmt.Println(fmt.Sprintf("ingredient:: %v", ing))
}

func Test_FindExistingIngredient(t *testing.T) {
	db, err := NewGormClient(config.Host, config.Port, config.User, config.Password, false)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	if saveErr := db.SaveIngredient(context.Background(), &model.Ingredient{Name: "testIngredient"}); saveErr != nil {
		t.Fatalf("failed to save ingredient, error:%v", saveErr)
	}

	ing, findErr := db.FindIngredientByName(context.Background(), "testIngredient")
	if ing != nil {
		fmt.Println(fmt.Sprintf("ingredient:: %v", ing))
	}

	if findErr != nil {
		t.Fatalf("error should be nil")
	}
}

func Test_SaveExistingIngredient(t *testing.T) {
	db, err := NewGormClient(config.Host, config.Port, config.User, config.Password, false)
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

func Test_SaveIngredient_LowConcentrationTest(t *testing.T) {
	db, err := NewGormClient(config.Host, config.Port, config.User, config.Password, false)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	i := model.Ingredient{Name: "testIngredient", Concentrations: "0.5-10%", EffectiveAtLowConcentration: model.EffectiveYes}

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

	if ing.EffectiveAtLowConcentration != model.EffectiveYes {
		t.Fatal("EffectiveAtLowConcentration should be Yes")
	}

	if ing.Concentrations != "0.5-10%" {
		t.Fatal("EffectiveAtLowConcentration should be 0.5-10%")
	}

	fmt.Println(fmt.Sprintf("ingredient:: %v", ing))
}

/*
intention of this test is to check is the ingredients are being retrieved with skinconcerns
*/
func TestFindProductsByIds(t *testing.T) {
	db, err := NewGormClient(config.Host, config.Port, config.User, config.Password, false)
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
	db, err := NewGormClient(config.RemoteHost, config.Port, config.User, config.Password, false)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	ing, err := db.GetSkinConcernDescriptionByIngredients(context.Background(), []string{"algae extract", "glycerin"}, "oiliness_shine")

	fmt.Println(err)
	fmt.Println(fmt.Sprintf("ingredient:: %#v", ing))
}

func TestGormConnector_SaveRecommendations(t *testing.T) {
	db, err := NewGormClient(config.Host, config.Port, config.User, config.Password, false)
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
	db, err := NewGormClient(config.Host, config.Port, config.User, config.Password, false)
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
	db, err := NewGormClient(config.Host, config.Port, config.User, config.Password, false)
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
	db, err := NewGormClient(config.RemoteHost, config.Port, config.User, config.Password, false)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	ps := []model.Product{
		{Name: "product01"},
		//{Name: "product02"},
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
		//for _, sp := range savedProducts {
		//	if delErr := db.DeleteUserRoutine(context.Background(), int(sp.ID), "1"); delErr != nil {
		//		t.Fatal(delErr)
		//	}
		//}

		for _, p := range ps {
			if err := db.DeleteProductByName(context.Background(), p.Name); err != nil {
				t.Fatal(err)
			}
		}
	}()

	// start of the test

	for _, savedProduct := range savedProducts {
		r := model.UserRoutine{
			UserId:      "1",
			Product:     savedProduct,
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
	if len(ur) != 4 {
		t.Fatalf("expected 4 routine products, got %d", len(ur))
	}

	//ur.Products = savedProducts[2:]
	//
	//for _, product := range ur.Products {
	//	fmt.Println(fmt.Sprintf("product to save second batch: %#v", product))
	//}
	//
	//err = db.SaveUserRoutine(context.Background(), ur)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//ur2, err := db.GetUserRoutine(context.Background(), "1")
	//if len(ur2.Products) != 2 {
	//	t.Fatalf("expected 2 products, got %d, %#v", len(ur2.Products), ur2.Products)
	//}
	//
	//fmt.Println(fmt.Sprintf("user routine: %#v", ur2))
}
