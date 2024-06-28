package main

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"os"
	"skingenius/config"
	"skingenius/database"
	"skingenius/database/model"
	"time"
)

/*
TRUNCATE TABLE ingredients  RESTART IDENTITY CASCADE;
*/
func main() {
	dbClient, err := database.NewGormClient(config.Host, config.Port, config.User, config.Password, false)
	//dbClient, err := database.NewGormClient(config.RemoteHost, config.Port, config.User, config.Password, false)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	//storeIngredients(context.Background(), dbClient)

	//storeProducts(context.Background(), dbClient, "admin/products-to-ingredients.csv")

	ps, err := dbClient.FindAllProductsWithIngredients(context.Background(), []int{1})

	fmt.Println(fmt.Sprintf("Products #%d", len(ps)))
	fmt.Println(fmt.Sprintf("Products: %+v", ps))
	fmt.Println(err)

	// ---------------  Inventory page

	//ingredient := model.Ingredient{
	//	Name: "Testname",
	//}
	//
	//
	//ctx = context.WithValue(ctx, model.SkintypeCtxKey, 13)
	//ctx = context.WithValue(ctx, model.SkinsensetivityCtxKey, 14)
	//ctx = context.WithValue(ctx, model.AcnebreakoutsCtxKey, 15)
	//ctx = context.WithValue(ctx, model.AllergiesCtxKey, 22)
	//ctx = context.WithValue(ctx, model.SkinconcernCtxKey, 22)
	//ctx = context.WithValue(ctx, model.AgeCtxKey, 28)
	//ctx = context.WithValue(ctx, model.BenefitsCtxKey, 29)
	//
	//ingredient.Preferences = allPreferences
	//ingredient.Skintypes = allskintypes
	//ingredient.Skinsensitivities = allskinsensetivities
	//ingredient.Acnebreakouts = allAcnebreakouts
	//ingredient.Allergies = allAllergies
	//ingredient.Skinconcerns = allSkinconcerns
	//ingredient.Ages = allAges
	//ingredient.Benefits = allBenefits

	//var pref []model.Preference
	//var ctx context.Context
	//for _, preference := range allPreferences {
	//	ctx = context.WithValue(context.Background(), preference.Name, 2)
	//}

	//ctx := context.Background()
	//ingredients, err := db.GetAllIngredients(ctx)
	//iNames := database.GetIngredientsName(ingredients)
	//
	//skintypes, err := db.GetAllSkintypes(ctx)
	//skinsensetivities, err := db.GetAllSkinsensetivity(ctx)
	//acnebreakouts, err := db.GetAllAcneBreakouts(ctx)
	//preferences, err := db.GetAllPreferences(ctx)
	//
	//a := app.New()
	//containerWindow := a.NewWindow("Skingenius admin application")
	//containerWindow.SetPadded(true)
	//containerWindow.Resize(fyne.Size{
	//	Width:  1400,
	//	Height: 1000,
	//})
	//
	//addNewIngredientTab := container.NewTabItem("Add new ingredient", widget.NewLabel("Ingredient"))
	//editIngredientTab := container.NewTabItem("Edit ingredient", widget.NewLabel("Ingredient"))
	//testtab := container.NewTabItem("Test tab", widget.NewLabel("test tab"))
	//
	//tabs := container.NewAppTabs(addNewIngredientTab, editIngredientTab, testtab)
	//tabs.SetTabLocation(container.TabLocationTop)
	//
	//
	//addNewIngredientTab.Content = ingredientTabContent()
	//editIngredientTab.Content = ingredientTabContent()
	//testtab.Content = testTabContent(iNames, skintypes, skinsensetivities, acnebreakouts, preferences)
	//
	//containerWindow.SetContent(tabs)
	//containerWindow.CenterOnScreen()
	//containerWindow.ShowAndRun()
}

// ------- TABS ---->>

func productTabContent() *fyne.Container {
	newProduct := model.Product{
		ID:          0,
		Name:        "",
		Brand:       "",
		Link:        "",
		Ingredients: nil,
	}

	fmt.Println(newProduct)

	//saveProductResult := canvas.NewText("", color.Black)
	saveProductResult := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	pName := widget.NewEntry()
	pBrand := widget.NewEntry()
	pLink := widget.NewEntry()
	pIngredients := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Name", Widget: pName},
			{Text: "Brand", Widget: pBrand},
			{Text: "Link", Widget: pLink},
			{Text: "Ingredient", Widget: pIngredients},
		},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Product name top save:", pName.Text)
			saveProductResult.Text = "Product has been saved"
			saveProductResult.Show()
			time.Sleep(2 * time.Second)
			//saveProductResult.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255} // maroon
			saveProductResult.Hide()
		},
	}

	productTabLayoutContainer := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		form,
		layout.NewSpacer(),
	)

	return productTabLayoutContainer
}

func ingredientTabContent() *fyne.Container {

	//saveProductResult := canvas.NewText("", color.Black)
	saveProductResult := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	pName := widget.NewEntry()
	pBrand := widget.NewEntry()
	pLink := widget.NewEntry()
	pIngredients := widget.NewEntry()

	selecty := widget.NewSelect([]string{"dry", "normal"}, func(s string) {
		fmt.Println("selecty string: ", s)
	})
	selecty.PlaceHolder = "year"

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Name", Widget: pName},
			{Text: "Brand", Widget: pBrand},
			{Text: "Link", Widget: pLink},
			{Text: "Ingredient", Widget: pIngredients},
			{Text: "skintype", Widget: selecty},
		},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Product name top save:", pName.Text)
			saveProductResult.Text = "Product has been saved"
			saveProductResult.Show()
			time.Sleep(2 * time.Second)

			fmt.Println(selecty.Selected)
			//saveProductResult.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255} // maroon
			saveProductResult.Hide()
		},
	}

	productTabLayoutContainer := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		form,
		layout.NewSpacer(),
	)

	return productTabLayoutContainer
}

func testTabContent(ingredients []string, skintypes []model.Skintype, skinsensesivities []model.Skinsensitivity, acne []model.Acnebreakout, preferences []model.Preference) *fyne.Container {

	skintypeLayout := container.New(layout.NewVBoxLayout(), container.New(layout.NewGridLayout(2), widget.NewLabel("Skin type"), widget.NewLabel("Score")))
	for _, db_skintype := range skintypes {
		skintypeLayout.Add(container.New(layout.NewGridLayout(2), widget.NewLabel(string(db_skintype.Type)), widget.NewEntry()))
	}

	skinsensitivitiesLayout := container.New(layout.NewVBoxLayout(), container.New(layout.NewGridLayout(2), widget.NewLabel("Skin sensitivity"), widget.NewLabel("Score")))
	for _, db_skinsensesivity := range skinsensesivities {
		skinsensitivitiesLayout.Add(container.New(layout.NewGridLayout(2), widget.NewLabel(string(db_skinsensesivity.Sensitivity)), widget.NewEntry()))
	}

	acneLayout := container.New(layout.NewVBoxLayout(), container.New(layout.NewGridLayout(2), widget.NewLabel("Acne prone"), widget.NewLabel("Score")))
	for _, db_acne := range acne {
		acneLayout.Add(container.New(layout.NewGridLayout(2), widget.NewLabel(string(db_acne.Frequency)), widget.NewEntry()))
	}

	preferencesLayout := container.New(layout.NewVBoxLayout(), container.New(layout.NewGridLayout(2), widget.NewLabel("Ingredient preference"), widget.NewLabel("Score")))
	for _, db_preference := range preferences {
		preferencesLayout.Add(container.New(layout.NewGridLayout(2), widget.NewLabel(string(db_preference.Name)), widget.NewEntry()))
	}

	productTabLayoutContainer := container.New(layout.NewVBoxLayout(),
		layout.NewSpacer(),
		container.New(layout.NewVBoxLayout(),
			container.New(layout.NewGridLayout(10), widget.NewLabel("Choose the ingredient"), widget.NewSelect(ingredients, func(s string) { fmt.Println("selecty string: ", s) })),
			container.New(layout.NewGridLayout(10), widget.NewLabel("Name"), widget.NewEntry()),
			container.New(layout.NewGridLayout(10), widget.NewLabel("Pubchem Id"), widget.NewEntry()),
			container.New(layout.NewGridLayout(10), widget.NewLabel("Cas number"), widget.NewEntry()),
			container.New(layout.NewGridLayout(10), widget.NewLabel("Inci number"), widget.NewEntry()),
			container.New(layout.NewGridLayout(10), widget.NewLabel("ec number"), widget.NewEntry()),
		),
		//container.New(layout.NewGridLayout(6),
		//	container.New(layout.NewGridLayout(2), widget.NewLabel("Choose the ingredient"), widget.NewSelect([]string{"1", "2"}, func(s string) { fmt.Println("selecty string: ", s) })),
		//	container.New(layout.NewGridLayout(2), widget.NewLabel("Name"), widget.NewEntry()),
		//	container.New(layout.NewGridLayout(2), widget.NewLabel("Pubchem Id"), widget.NewEntry()),
		//	container.New(layout.NewGridLayout(2), widget.NewLabel("Cas number"), widget.NewEntry()),
		//	container.New(layout.NewGridLayout(2), widget.NewLabel("Inci number"), widget.NewEntry()),
		//	container.New(layout.NewGridLayout(2), widget.NewLabel("ec number"), widget.NewEntry()),
		//),

		layout.NewSpacer(),
		container.New(layout.NewGridLayout(4),
			//container.New(layout.NewHBoxLayout(),
			skintypeLayout,
			skinsensitivitiesLayout,
			acneLayout,
			preferencesLayout,
		),
		canvas.NewLine(color.Black),

		layout.NewSpacer(),
		container.New(layout.NewGridLayout(4),
			container.New(layout.NewVBoxLayout(),
				widget.NewLabel("Label1"),
				widget.NewLabel("Label2"),
				widget.NewLabel("Label3"),
			),
			container.New(layout.NewVBoxLayout(),
				widget.NewLabel("Label1"),
				widget.NewLabel("Label2"),
				widget.NewLabel("Label3"),
			),
			container.New(layout.NewVBoxLayout(),
				widget.NewLabel("Label1"),
				widget.NewLabel("Label2"),
				widget.NewLabel("Label3"),
			),
			container.New(layout.NewVBoxLayout(),
				widget.NewLabel("Label1"),
				widget.NewLabel("Label2"),
				widget.NewLabel("Label3"),
			),
		),
		layout.NewSpacer(),
	)

	return productTabLayoutContainer
}
