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
	"skingenius/xlsx"
	"time"
)

/*
TRUNCATE TABLE ingredients  RESTART IDENTITY CASCADE;
*/
func main() {
	//dbClient, err := database.NewGormClient(config.Host, config.Port, config.User, config.Password, false)
	fmt.Println("Skingenious 2024")

	dbClient, err := database.NewGormClient(config.RemoteHost, config.Port, config.User, config.Password, false) // REMOTE
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}
	time.Sleep(100 * time.Millisecond)

	ctx := context.Background()

	answers := xlsx.ReadAnswers("admin/input.xlsx")

	a1SkinType := model.SkinTypeMapping[answers[0]]
	a2SkinSensitivity := model.SensitivityMapping[answers[1]]
	a3Acne := model.AcneProneMapping[answers[2]]
	a4Age := model.AgeMapping[answers[3]]
	a5Preference := string(model.PreferenceMapping[answers[4]])
	a6Allergy := string(model.AllergiesMapping[answers[5]])
	a7Concerns := string(model.SkinConcernsMapping[answers[6]])
	a8Benefits := string(model.BenefitsMapping[answers[7]])

	fmt.Println("\n\n ********************  Answers  ********************")
	fmt.Println("Skin type: ", a1SkinType)
	fmt.Println("Sensitivity: ", a2SkinSensitivity)
	fmt.Println("Acne: ", a3Acne)
	fmt.Println("Age: ", a4Age)
	fmt.Println("Preference: ", a5Preference)
	fmt.Println("Allergy: ", a6Allergy)
	fmt.Println("Concerns: ", a7Concerns)
	fmt.Println("Benefits: ", a8Benefits)
	fmt.Println("********************  Answers  ******************** \n\n ")

	//func findBestProducts_RatingStrategy(dbClient database.Connector, ctx context.Context,
	//	q1SkinTypeAnswer string, q2SkinSensitivityAnswer string, q3AcneBreakoutsAnswer string, q4PreferencesAnswer []string,
	//	q5AllergiesAnswer []string, q6SkinConcernAnswer []string, q7AgeAnswer int, q8BenefitsAnswer []string) {

	findBestProducts_RatingStrategy(dbClient, ctx, a1SkinType, a2SkinSensitivity, a3Acne, []string{a5Preference}, []string{a6Allergy}, []string{a7Concerns}, a4Age, []string{a8Benefits})

	//storeIngredients(ctx, dbClient, "admin/ingredients_master.csv")
	//storeProducts(ctx, dbClient, "admin/products_master.csv")

	//q1SkinTypeAnswer := "dry"
	//q2SkinSensitivityAnswer := "often"
	//q3AcneBreakoutsAnswer := "occasionally"
	//q4PreferencesAnswer := []string{"vegan"}
	//q5AllergiesAnswer := []string{"nuts"}
	//q6SkinConcernAnswer := []string{"oiliness", "acne"}
	//q7AgeAnswer := 30
	//q8BenefitsAnswer := []string{"moisturizing", "soothing", "reduces_acne"}

	//findBestProducts_RatingStrategy(dbClient, ctx, q1SkinTypeAnswer, q2SkinSensitivityAnswer, q3AcneBreakoutsAnswer, q4PreferencesAnswer, q5AllergiesAnswer, q6SkinConcernAnswer, q7AgeAnswer, q8BenefitsAnswer)

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

func findIngredientsByQuestion(dbClient database.Connector, ctx context.Context,
	q1SkinTypeAnswer string, q2SkinSensitivityAnswer string, q3AcneBreakoutsAnswer string, q4PreferencesAnswer []string,
	q5AllergiesAnswer []string, q6SkinConcernAnswer []string, q7AgeAnswer int, q8BenefitsAnswer []string) (
	skintypeIng []model.Ingredient, skinSensIng []model.Ingredient, acneIng []model.Ingredient, prefIng []model.Ingredient,
	allergiesIng []model.Ingredient, skinConcernIng []model.Ingredient, ageIng []model.Ingredient, benefitsIng []model.Ingredient) {

	var err error

	skintypeIng, err = dbClient.GetIngredientsBySkintype(ctx, q1SkinTypeAnswer)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to get ingredients by skintype, error: %v", err))
	}

	skinSensIng, err = dbClient.GetIngredientsBySkinsensitivity(ctx, q2SkinSensitivityAnswer)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to get ingredients by skinsensitivity, error: %v", err))
	}

	acneIng, err = dbClient.GetIngredientsByAcneBreakouts(ctx, q3AcneBreakoutsAnswer)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to get ingredients by acnebreakouts, error: %v", err))
	}

	prefIng, err = dbClient.GetIngredientsByPreferences(ctx, q4PreferencesAnswer)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to get ingredients by preferences, error: %v", err))
	}

	allergiesIng, err = dbClient.GetIngredientsByAllergies(ctx, q5AllergiesAnswer)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to get ingredients by allergies, error: %v", err))
	}

	skinConcernIng, err = dbClient.GetIngredientsBySkinconcerns(ctx, q6SkinConcernAnswer)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to get ingredients by skinconcerns, error: %v", err))
	}

	ageIng, err = dbClient.GetIngredientsByAge(ctx, fmt.Sprintf("%d", q7AgeAnswer))
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to get ingredients by age, error: %v", err))
	}

	benefitsIng, err = dbClient.GetIngredientsByBenefits(ctx, q8BenefitsAnswer)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to get ingredients by benefits, error: %v", err))
	}

	fmt.Println(fmt.Sprintf("Skin type ingredients: %v", len(skintypeIng)))
	fmt.Println(fmt.Sprintf("Skin sensitivity ingredients: %v", len(skinSensIng)))
	fmt.Println(fmt.Sprintf("Acne breakout ingredients: %v", len(acneIng)))
	fmt.Println(fmt.Sprintf("Preference ingredients: %v", len(prefIng)))
	fmt.Println(fmt.Sprintf("Allergy ingredients: %v", len(allergiesIng)))
	fmt.Println(fmt.Sprintf("Skin concerns ingredients: %v", len(skinConcernIng)))
	fmt.Println(fmt.Sprintf("By Age ingredients: %v", len(ageIng)))
	fmt.Println(fmt.Sprintf("Benefits ingredients: %v", len(benefitsIng)))

	return
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
