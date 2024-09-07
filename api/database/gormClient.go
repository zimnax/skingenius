package database

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"skingenius/database/model"
	"skingenius/logger"
	"strings"
)

/*
cd /etc/postgresql/16/main/
nano pg_hba.conf
sudo service postgresql restart
*/

type GormConnector struct {
	db *gorm.DB
}

/*
returns the description of the skin concern for the given ingredients and concern
*/
func (g GormConnector) GetSkinConcernDescriptionByIngredients(ctx context.Context, ingredients []string, concern string) ([]model.SkinconcernToIngredientDescription, error) {
	/*
		SELECT ingredients.name as ingredient, skinconcerns.name as concern, ingredient_skinconcerns.description
		FROM ingredients
		INNER JOIN ingredient_skinconcerns ON ingredients.id = ingredient_skinconcerns.ingredient_id
		INNER JOIN skinconcerns ON skinconcerns.id = ingredient_skinconcerns.skinconcern_id
		where ingredients.name in ('algae extract','glycerin') AND skinconcerns.name in ('oiliness_shine', 'dark_circles','fine_lines_wrinkles','loss_of_elasticity_firmness')
	*/

	var cd []model.SkinconcernToIngredientDescription

	err := g.db.Select("ingredients.name as ingredientname, skinconcerns.name as concern, ingredient_skinconcerns.description").
		Table("ingredients").
		Joins("INNER JOIN ingredient_skinconcerns ON ingredients.id = ingredient_skinconcerns.ingredient_id").
		Joins("INNER JOIN skinconcerns ON skinconcerns.id = ingredient_skinconcerns.skinconcern_id").
		Where("ingredients.name IN (?) AND skinconcerns.name = ?", ingredients, concern).Find(&cd).Error

	return cd, err
}

func (g GormConnector) FindIngredientByAlias(ctx context.Context, alias string) (*model.Ingredient, error) {
	var ingredient model.Ingredient
	err := g.db.Where("? = any (synonyms)", alias).First(&ingredient).Error
	return &ingredient, err
}

func (g GormConnector) GetQuiz(ctx context.Context, userId string) (model.UserQuiz, error) {
	var uq model.UserQuiz
	//err := g.db.WithContext(ctx).Find(&uq, userId).Error
	err := g.db.WithContext(ctx).Where("user_id = ?", userId).First(&uq).Error

	return uq, err
}

func (g GormConnector) SaveQuiz(ctx context.Context, quiz model.UserQuiz) error {

	fmt.Println(fmt.Sprintf("quiz to save in db: %#v", quiz))
	return g.db.WithContext(ctx).Save(quiz).Error
}

func (g GormConnector) FindProductsByIds(ctx context.Context, ids []int32) ([]model.Product, error) {
	var products []model.Product
	// .Preload("Ingredients.Skinconcerns").
	err := g.db.WithContext(ctx).Preload("Ingredients").Where("products.id IN (?)", ids).Find(&products).Error

	return products, err
}

func (g GormConnector) GetRecommendations(ctx context.Context, s string) ([]model.UserRecommendations, error) {
	var ur []model.UserRecommendations

	//err := g.db.WithContext(ctx).Select("user_recommendations.user_id, user_recommendations.recommended_products").
	//	Table("user_recommendations").
	//	Where("user_recommendations.user_id = ?", s).Find(&ur).Error

	err := g.db.Where("user_recommendations.user_id = ?", s).Find(&ur).Error

	return ur, err
}

func (g GormConnector) SaveRecommendations(ctx context.Context, ur []model.UserRecommendations) error {
	//return g.db.WithContext(ctx).Create(ur).Error
	err := g.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "product_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"score"}),
	}).Create(&ur).Error

	return err
}

func (g GormConnector) FindAllProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	if err := g.db.Preload("Ingredients").Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

/* FIND ingredients count for each product

SELECT products.id, products.name,  count(ingredients.id) as i_count  FROM public.products
INNER JOIN product_ingredient ON products.id =product_ingredient.product_id
INNER JOIN ingredients ON ingredients.id =product_ingredient.ingredient_id
GROUP BY products.id, products.name
ORDER BY i_count DESC;
*/

func (g GormConnector) FindAllProductsWithIngredients(ctx context.Context, ingredients []string, accuracy uint) ([]model.Product, error) {

	/*

		SELECT products.id, products.name
		FROM public.products
		INNER JOIN product_ingredient ON products.id =product_ingredient.product_id
		INNER JOIN ingredients ON ingredients.id =product_ingredient.ingredient_id
		WHERE ingredients.name in (5)
		GROUP BY products.id, products.name
		HAVING COUNT(DISTINCT ingredients.name) = 1;

	*/

	var products []model.Product

	err := g.db.Select("products.id, products.name").
		Table("products").
		Joins("INNER JOIN product_ingredient ON products.id =product_ingredient.product_id").
		Joins("INNER JOIN ingredients ON ingredients.id =product_ingredient.ingredient_id").
		Where("ingredients.name IN (?)", ingredients).
		Group("products.id, products.name").
		Having("COUNT(DISTINCT ingredients.name) = ?", accuracy).
		Find(&products).Error

	return products, err
}

func (g GormConnector) SaveProduct(ctx context.Context, product *model.Product) error {
	return g.db.WithContext(ctx).Create(product).Error
}

func (g GormConnector) FindProductByName(ctx context.Context, name string) (*model.Product, error) {
	//clause.Associations
	var product model.Product
	err := g.db.Preload("Ingredients").Where("name = ?", name).First(&product).Error
	return &product, err
}

func (g GormConnector) FindIngredientByName(ctx context.Context, name string) (*model.Ingredient, error) {
	var ingredient model.Ingredient
	err := g.db.Where("name = ?", name).First(&ingredient).Error
	return &ingredient, err
}

func (g GormConnector) GetIngredientsByBenefits(ctx context.Context, benefits []string) ([]model.Ingredient, error) {
	var ingredients []model.Ingredient

	err := g.db.Select("ingredients.id, ingredients.name, sum(ingredient_benefits.score) as score").
		Table("ingredients").
		Joins("INNER JOIN ingredient_benefits ON ingredients.id = ingredient_benefits.ingredient_id").
		Joins("INNER JOIN benefits ON benefits.id = ingredient_benefits.benefit_id").
		Where("benefits.name IN (?)", benefits).
		Group("ingredients.id").
		Find(&ingredients).Error

	return ingredients, err
}

func (g GormConnector) GetIngredientsBySkinconcerns(ctx context.Context, concerns []string) ([]model.Ingredient, error) {
	var ingredients []model.Ingredient

	/*
		SELECT ingredients.id, ingredients.name, skinconcerns.name as skinconcerns_name, ingredient_skinconcerns.score
		FROM public.ingredients
		INNER JOIN ingredient_skinconcerns ON ingredients.id = ingredient_skinconcerns.ingredient_id
		INNER JOIN skinconcerns ON skinconcerns.id = ingredient_skinconcerns.skinconcern_id
		WHERE skinconcerns.name IN ('acne') AND ingredient_skinconcerns.score > 0
		GROUP BY ingredients.id
	*/

	/*
		SELECT ingredients.id, ingredients.name, SUM(ingredient_skinconcerns.score)
		FROM public.ingredients
		INNER JOIN ingredient_skinconcerns ON ingredients.id = ingredient_skinconcerns.ingredient_id
		INNER JOIN skinconcerns ON skinconcerns.id = ingredient_skinconcerns.skinconcern_id
		WHERE skinconcerns.name IN ('rosacea', 'hyperpigmentation_unevenskintone') AND ingredient_skinconcerns.score > 0
		GROUP BY ingredients.id
	*/

	err := g.db.Select("ingredients.id, ingredients.name, sum(ingredient_skinconcerns.score) as score").
		Table("ingredients").
		Joins("INNER JOIN ingredient_skinconcerns ON ingredients.id = ingredient_skinconcerns.ingredient_id").
		Joins("INNER JOIN skinconcerns ON skinconcerns.id = ingredient_skinconcerns.skinconcern_id").
		Where("skinconcerns.name IN (?)", concerns).
		Group("ingredients.id").
		Find(&ingredients).Error

	return ingredients, err
}

func (g GormConnector) GetIngredientsByAllergies(ctx context.Context, allergies []string) ([]model.Ingredient, error) {
	var ingredients []model.Ingredient

	/*
		SELECT ingredients.id, ingredients.name, allergies.name as allergies_name, ingredient_allergies.score
		FROM public.ingredients
		INNER JOIN ingredient_allergies ON ingredients.id = ingredient_allergies.ingredient_id
		INNER JOIN allergies ON allergies.id = ingredient_allergies.allergy_id
		WHERE allergies.name IN ('soy','nuts','latex') AND ingredient_allergies.score > 0
		GROUP BY ingredients.id
	*/

	//err := g.db.Select("ingredients.id, ingredients.name, ingredient_allergies.score").
	//	Table("ingredients").
	//	Joins("INNER JOIN ingredient_allergies ON ingredients.id = ingredient_allergies.ingredient_id").
	//	Joins("INNER JOIN allergies ON allergies.id = ingredient_allergies.allergy_id").
	//	Where("allergies.name IN (?)", allergies).
	//	Find(&ingredients).Error

	err := g.db.Select("ingredients.id, ingredients.name").
		Table("ingredients").
		Joins("INNER JOIN ingredient_allergies ON ingredients.id = ingredient_allergies.ingredient_id").
		Joins("INNER JOIN allergies ON allergies.id = ingredient_allergies.allergy_id").
		Where("allergies.name IN (?) AND ingredient_allergies.score = true", allergies).
		Group("ingredients.id").
		Find(&ingredients).Error

	return ingredients, err
}

func (g GormConnector) GetIngredientsByPreferences(ctx context.Context, preferences []string) ([]model.Ingredient, error) {
	var ingredients []model.Ingredient

	/*
		SELECT ingredients.id, ingredients.name, preferences.name as preferences_name, ingredient_preferences.score
		FROM public.ingredients
		INNER JOIN ingredient_preferences ON ingredients.id = ingredient_preferences.ingredient_id
		INNER JOIN preferences ON preferences.id = ingredient_preferences.preference_id
		WHERE preferences.name IN ('vegetarian','paleo') AND ingredient_preferences.score > 0
		GROUP BY ingredients.id
	*/

	//err := g.db.Select("ingredients.id, ingredients.name, ingredient_preferences.score").
	//	Table("ingredients").
	//	Joins("INNER JOIN ingredient_preferences ON ingredients.id = ingredient_preferences.ingredient_id").
	//	Joins("INNER JOIN preferences ON preferences.id = ingredient_preferences.preference_id").
	//	Where("preferences.name IN (?)", preferences).
	//	Find(&ingredients).Error

	err := g.db.Select("ingredients.id, ingredients.name").
		Table("ingredients").
		Joins("INNER JOIN ingredient_preferences ON ingredients.id = ingredient_preferences.ingredient_id").
		Joins("INNER JOIN preferences ON preferences.id = ingredient_preferences.preference_id").
		Where("preferences.name IN (?) AND ingredient_preferences.score = true", preferences).
		Group("ingredients.id").
		Find(&ingredients).Error

	return ingredients, err
}

func (g GormConnector) GetIngredientsByAge(ctx context.Context, age string) ([]model.Ingredient, error) {
	var ingredients []model.Ingredient

	/*
		SELECT ingredients.id, ingredients.name, ages.value, ingredient_ages.score
		FROM public.ingredients
		INNER JOIN ingredient_ages ON ingredients.id = ingredient_ages.ingredient_id
		INNER JOIN ages ON ages.id = ingredient_ages.age_id
		WHERE ages.value = 30 AND ingredient_ages.score > 0
	*/

	err := g.db.Select("ingredients.id, ingredients.name, ages.value").
		Table("ingredients").
		Joins("INNER JOIN ingredient_ages ON ingredients.id = ingredient_ages.ingredient_id").
		Joins("INNER JOIN ages ON ages.id = ingredient_ages.age_id").
		Where("ages.value = (?)", age).
		Find(&ingredients).Error

	return ingredients, err
}

func (g GormConnector) GetIngredientsByAcneBreakouts(ctx context.Context, frequency string) ([]model.Ingredient, error) {
	var ingredients []model.Ingredient

	//err := g.db.Select("ingredients.id, ingredients.name, acnebreakouts.frequency, ingredient_acnebreakouts.score").
	//	Table("ingredients").
	//	Joins("INNER JOIN ingredient_acnebreakouts ON ingredients.id = ingredient_acnebreakouts.ingredient_id").
	//	Joins("INNER JOIN acnebreakouts ON acnebreakouts.id = ingredient_acnebreakouts.acnebreakout_id").
	//	Where("acnebreakouts.frequency = (?)", frequency).
	//	Find(&ingredients).Error

	err := g.db.Select("ingredients.id, ingredients.name, acnebreakouts.frequency").
		Table("ingredients").
		Joins("INNER JOIN ingredient_acnebreakouts ON ingredients.id = ingredient_acnebreakouts.ingredient_id").
		Joins("INNER JOIN acnebreakouts ON acnebreakouts.id = ingredient_acnebreakouts.acnebreakout_id").
		Where("acnebreakouts.frequency = (?) AND ingredient_acnebreakouts.score = true", frequency).
		Find(&ingredients).Error

	return ingredients, err
}

func (g GormConnector) GetIngredientsBySkinsensitivity(ctx context.Context, skinsensitivity string) ([]model.Ingredient, error) {
	var ingredients []model.Ingredient

	/*
		SELECT ingredients.id, ingredients.name, skinsensitivities.sensitivity, ingredient_skinsensitivities.score
		FROM public.ingredients
		INNER JOIN ingredient_skinsensitivities ON ingredients.id = ingredient_skinsensitivities.ingredient_id
		INNER JOIN skinsensitivities ON Skinsensitivities.id = ingredient_skinsensitivities.skinsensitivity_id
		WHERE skinsensitivities.sensitivity = 'never' AND ingredient_skinsensitivities.score > 0
		ORDER BY id ASC
	*/

	//err := g.db.Select("ingredients.id, ingredients.name, skinsensitivities.sensitivity, ingredient_skinsensitivities.score").
	//	Table("ingredients").
	//	Joins("INNER JOIN ingredient_skinsensitivities ON ingredients.id = ingredient_skinsensitivities.ingredient_id").
	//	Joins("INNER JOIN skinsensitivities ON Skinsensitivities.id = ingredient_skinsensitivities.skinsensitivity_id").
	//	Where("skinsensitivities.sensitivity = (?)", skinsensitivity).
	//	Find(&ingredients).Error

	err := g.db.Select("ingredients.id, ingredients.name, skinsensitivities.sensitivity").
		Table("ingredients").
		Joins("INNER JOIN ingredient_skinsensitivities ON ingredients.id = ingredient_skinsensitivities.ingredient_id").
		Joins("INNER JOIN skinsensitivities ON Skinsensitivities.id = ingredient_skinsensitivities.skinsensitivity_id").
		Where("skinsensitivities.sensitivity = (?) AND ingredient_skinsensitivities.score = true ", skinsensitivity).
		Find(&ingredients).Error

	return ingredients, err
}

func (g GormConnector) GetIngredientsBySkintype(ctx context.Context, skintype string) ([]model.Ingredient, error) {
	var ingredients []model.Ingredient

	//err := g.db.Preload(clause.Associations).Find(&ingredients).Error // WORKS

	/*
		SELECT ingredients.id, ingredients.name, ingredient_skintypes.skintype_id, skintypes.type, ingredient_skintypes.score
		FROM public.ingredients
		INNER JOIN ingredient_skintypes ON ingredients.id = ingredient_skintypes.ingredient_id
		INNER JOIN skintypes ON skintypes.id = ingredient_skintypes.skintype_id
		WHERE skintypes.type = 'dry' AND ingredient_skintypes.score > 0
		ORDER BY id ASC
	*/

	err := g.db.Select("ingredients.id, ingredients.name, ingredient_skintypes.skintype_id,ingredient_skintypes.score").
		Table("ingredients").
		Joins("INNER JOIN ingredient_skintypes ON ingredients.id = ingredient_skintypes.ingredient_id").
		Joins("INNER JOIN skintypes ON skintypes.id = ingredient_skintypes.skintype_id").
		Where("skintypes.type = (?)", skintype).
		Find(&ingredients).Error

	return ingredients, err
}

func (g GormConnector) SetupJoinTables() error {
	var err error

	if err = g.db.SetupJoinTable(&model.Ingredient{}, "Preferences", &model.IngredientPreference{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientPreference], error: %v", err))
	}

	if err = g.db.SetupJoinTable(&model.Ingredient{}, "Skintypes", &model.IngredientSkintype{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientSkintype], error: %v", err))
	}

	if err = g.db.SetupJoinTable(&model.Ingredient{}, "Skinsensitivities", &model.IngredientSkinsensitivity{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientSkinsensitivity], error: %v", err))
	}

	if err = g.db.SetupJoinTable(&model.Ingredient{}, "Acnebreakouts", &model.IngredientAcnebreakout{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientAcnebreakout], error: %v", err))
	}

	if err = g.db.SetupJoinTable(&model.Ingredient{}, "Allergies", &model.IngredientAllergy{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientAllergy], error: %v", err))
	}

	if err = g.db.SetupJoinTable(&model.Ingredient{}, "Skinconcerns", &model.IngredientSkinconcern{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientSkinconcern], error: %v", err))
	}

	if err = g.db.SetupJoinTable(&model.Ingredient{}, "Ages", &model.IngredientAge{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientAge], error: %v", err))
	}

	if err = g.db.SetupJoinTable(&model.Ingredient{}, "Benefits", &model.IngredientBenefit{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientBenefit], error: %v", err))
	}

	return err
}

//func (g GormConnector) SavePreference(ctx context.Context, preference *model.Preference) error {
//
//	//g.db.WithContext(ctx).Model(&user).Association("Roles").Append(&role)
//
//	return nil
//}

func (g GormConnector) SaveIngredient(ctx context.Context, ingredient *model.Ingredient) error {
	db := g.db.WithContext(ctx).Save(ingredient)
	return db.Error
}

func (g GormConnector) GetAllIngredients(ctx context.Context) ([]model.Ingredient, error) {
	var ingredients []model.Ingredient
	if err := g.db.Find(&ingredients).Error; err != nil {
		return ingredients, err
	}

	return ingredients, nil
}

func (g GormConnector) GetAllPreferences(ctx context.Context) ([]model.Preference, error) {
	var preferences []model.Preference
	if err := g.db.Find(&preferences).Error; err != nil {
		return preferences, err
	}

	return preferences, nil
}

func (g GormConnector) GetAllAllergies(ctx context.Context) ([]model.Allergy, error) {
	var allergies []model.Allergy
	if err := g.db.Find(&allergies).Error; err != nil {
		return allergies, err
	}

	return allergies, nil
}

func (g GormConnector) GetAllSkinconcerns(ctx context.Context) ([]model.Skinconcern, error) {
	var skinconcerns []model.Skinconcern
	if err := g.db.Find(&skinconcerns).Error; err != nil {
		return skinconcerns, err
	}

	return skinconcerns, nil
}

func (g GormConnector) GetAllAge(ctx context.Context) ([]model.Age, error) {
	var ages []model.Age
	if err := g.db.Find(&ages).Error; err != nil {
		return ages, err
	}

	return ages, nil
}

func (g GormConnector) GetAllBenefits(ctx context.Context) ([]model.Benefit, error) {
	var benefits []model.Benefit
	if err := g.db.Find(&benefits).Error; err != nil {
		return benefits, err
	}

	return benefits, nil
}

func (g GormConnector) GetAllAcneBreakouts(ctx context.Context) ([]model.Acnebreakout, error) {
	var acnebreakout []model.Acnebreakout
	if err := g.db.Find(&acnebreakout).Error; err != nil {
		return acnebreakout, err
	}

	return acnebreakout, nil
}

func (g GormConnector) GetAllSkinsensetivity(ctx context.Context) ([]model.Skinsensitivity, error) {
	var skinsensitivity []model.Skinsensitivity
	if err := g.db.Find(&skinsensitivity).Error; err != nil {
		return skinsensitivity, err
	}

	return skinsensitivity, nil
}

func (g GormConnector) GetAllSkintypes(ctx context.Context) ([]model.Skintype, error) {
	var skintypes []model.Skintype
	if err := g.db.Find(&skintypes).Error; err != nil {
		return skintypes, err
	}

	return skintypes, nil
}

func NewGormClient(host string, port int, user, password string, migrate bool) (Connector, error) {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	logger.New().Info(context.Background(), fmt.Sprintf("DB config: %v", dbInfo))

	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		logger.New().Error(context.Background(), "error connecting to the database: ", err)
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}
	logger.New().Info(context.Background(), "Connected to the database!")

	if migrate {
		if migrationErr := automigrate(db); migrationErr != nil {
			return nil, migrationErr
		}

	}

	return &GormConnector{db: db}, nil
}

// <<---------------------- Migrations ----------------------->>

func automigrate(db *gorm.DB) error {
	logger.New().Info(context.Background(), "Auto-migration started")

	var err error

	if err = migrateSkinType(db); err != nil {
		return err
	}

	if err = migrateSkinSensitivity(db); err != nil {
		return err
	}

	if err = migrateAcnebreakout(db); err != nil {
		return err
	}

	if err = migratePreference(db); err != nil {
		return err
	}

	if err = migrateAllergy(db); err != nil {
		return err
	}

	if err = migrateSkinconcern(db); err != nil {
		return err
	}

	if err = migrateAge(db); err != nil {
		return err
	}

	if err = migrateBenefit(db); err != nil {
		return err
	}

	if err = db.AutoMigrate(&model.Ingredient{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [Ingredient], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [Ingredient], error: %v", err))
	}

	if err = db.AutoMigrate(&model.Product{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [Product], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [Product], error: %v", err))
	}

	db.Migrator().DropTable(&model.UserRecommendations{})
	if err = db.AutoMigrate(&model.UserRecommendations{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [UserRecommendations], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [UserRecommendations], error: %v", err))
	}

	if err = db.AutoMigrate(&model.UserQuiz{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [UserQuiz], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [UserQuiz], error: %v", err))
	}

	logger.New().Info(context.Background(), "Auto-migration finished successfully")
	return nil
}

func migrateBenefit(db *gorm.DB) error {
	var err error

	if err = db.SetupJoinTable(&model.Ingredient{}, "Benefits", &model.IngredientBenefit{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientBenefit], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("SetupJoinTable failed for table [IngredientBenefit], error: %v", err))
	}

	if err = db.AutoMigrate(&model.Benefit{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [Benefit], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [Benefit], error: %v", err))
	}

	benefits := []model.Benefit{
		{ID: 1, Name: model.BenefitMoisturizing},
		{ID: 2, Name: model.BenefitNourishing},
		{ID: 3, Name: model.BenefitHydrating},
		{ID: 4, Name: model.BenefitExfoliating},
		{ID: 5, Name: model.BenefitCalming},
		{ID: 6, Name: model.BenefitSoothing},
		{ID: 7, Name: model.BenefitUVBarrier},
		{ID: 8, Name: model.BenefitHealing},
		{ID: 9, Name: model.BenefitSmoothing},
		{ID: 10, Name: model.BenefitReducesAcne},
		{ID: 11, Name: model.BenefitReducesBlemishes},
		{ID: 12, Name: model.BenefitReducesWrinkles},
		{ID: 13, Name: model.BenefitImprovesSymptomsOfEczema},
		{ID: 14, Name: model.BenefitImprovesSymptomsOfPsoriasis},
		{ID: 15, Name: model.BenefitImprovesSymptomsOfDermatitis},
		{ID: 16, Name: model.BenefitBrightening},
		{ID: 17, Name: model.BenefitImprovesSkinTone},
		{ID: 18, Name: model.BenefitReducesInflammation},
		{ID: 19, Name: model.BenefitMinimizesPores},
		{ID: 20, Name: model.BenefitAntiAging},
		{ID: 21, Name: model.BenefitFirming},
		{ID: 22, Name: model.BenefitDetoxifying},
		{ID: 23, Name: model.BenefitBalancing},
		{ID: 24, Name: model.BenefitReducesRedness},
		{ID: 25, Name: model.BenefitClarifying},
		{ID: 26, Name: model.BenefitAntiBacterial},
		{ID: 27, Name: model.BenefitStimulatesCollagenProduction},
		{ID: 28, Name: model.BenefitReducesFineLines},
		{ID: 29, Name: model.BenefitAntioxidantProtection},
		{ID: 30, Name: model.BenefitSkinBarrierProtection},
	}

	if res := db.Create(benefits); res.Error != nil {
		if !strings.Contains(res.Error.Error(), "duplicate key") {
			return err
		}
	}

	return nil
}

func migrateAge(db *gorm.DB) error {
	var err error

	if err = db.SetupJoinTable(&model.Ingredient{}, "Ages", &model.IngredientAge{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientAge], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("SetupJoinTable failed for table [IngredientAge], error: %v", err))
	}

	if err = db.AutoMigrate(&model.Age{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [Age], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [Age], error: %v", err))
	}

	ages := []model.Age{
		{ID: 1, Value: model.Age10},
		{ID: 2, Value: model.Age20},
		{ID: 3, Value: model.Age30},
		{ID: 4, Value: model.Age40},
		{ID: 5, Value: model.Age50},
		{ID: 6, Value: model.Age60},
	}

	if res := db.Create(ages); res.Error != nil {
		if !strings.Contains(res.Error.Error(), "duplicate key") {
			return err
		}
	}

	return nil
}

func migrateSkinconcern(db *gorm.DB) error {
	var err error

	if err = db.SetupJoinTable(&model.Ingredient{}, "Skinconcerns", &model.IngredientSkinconcern{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientSkinconcern], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("SetupJoinTable failed for table [IngredientSkinconcern], error: %v", err))
	}

	db.Migrator().DropTable(&model.Skinconcern{})

	if err = db.AutoMigrate(&model.Skinconcern{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [Skinconcern], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [Skinconcern], error: %v", err))
	}

	//concerns := []model.Skinconcern{
	//	{ID: 1, Name: model.ConcernRosacea},
	//	{ID: 2, Name: model.ConcernHyperpigmentation},
	//	{ID: 3, Name: model.ConcernMelasma},
	//	{ID: 4, Name: model.ConcernCysticAcne},
	//	{ID: 5, Name: model.ConcernAcne},
	//	{ID: 6, Name: model.ConcernXerosis},
	//	{ID: 7, Name: model.ConcernDryness},
	//	{ID: 8, Name: model.ConcernOiliness},
	//	{ID: 9, Name: model.ConcernUnevenSkinTone},
	//	{ID: 10, Name: model.ConcernSignsOfAging},
	//	{ID: 11, Name: model.ConcernFineLines},
	//	{ID: 12, Name: model.ConcernWrinkles},
	//	{ID: 13, Name: model.ConcernDarkSpots},
	//	{ID: 14, Name: model.ConcernLostOfElasticityFirmness},
	//	{ID: 15, Name: model.ConcernVisiblePores},
	//	{ID: 16, Name: model.ConcernCloggedPoresBlackheads},
	//	{ID: 17, Name: model.ConcernRedness},
	//	{ID: 18, Name: model.ConcernDullness},
	//	{ID: 19, Name: model.ConcernDamagedSkin},
	//	{ID: 20, Name: model.ConcernUnevenTexture},
	//	{ID: 21, Name: model.ConcernEczema},
	//	{ID: 22, Name: model.ConcernPsoriasis},
	//	{ID: 23, Name: model.ConcernDermatitis},
	//	{ID: 24, Name: model.ConcernSunburnedSkin},
	//	{ID: 25, Name: model.ConcernDarkCircles},
	//	{ID: 26, Name: model.ConcernBlemishes},
	//	{ID: 27, Name: model.ConcernSensitiveSkin},
	//	{ID: 28, Name: model.ConcernNone},
	//}
	concerns := []model.Skinconcern{
		{ID: 1, Name: model.ConcernAcne},
		{ID: 2, Name: model.ConcernDryness_Dehydration},
		{ID: 3, Name: model.ConcernHyperpigmentation_UnevenSkinTone},
		{ID: 4, Name: model.ConcernOiliness_Shine},
		{ID: 5, Name: model.ConcernFine_lines_Wrinkles},
		{ID: 6, Name: model.ConcernLoss_of_Elasticity_firmness},
		{ID: 7, Name: model.ConcernVisible_pores_Uneven_texture},
		{ID: 8, Name: model.ConcernClogged_pores_blackheads},
		{ID: 9, Name: model.ConcernDullness},
		{ID: 10, Name: model.ConcernDark_circles},
		{ID: 11, Name: model.ConcernBlemishes},
		{ID: 12, Name: model.ConcernNone},
		{ID: 13, Name: model.ConcernRosacea},
	}

	if res := db.Create(concerns); res.Error != nil {
		if !strings.Contains(res.Error.Error(), "duplicate key") {
			return err
		}
	}

	return nil
}

func migrateAllergy(db *gorm.DB) error {
	var err error

	if err = db.SetupJoinTable(&model.Ingredient{}, "Allergies", &model.IngredientAllergy{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientAllergy], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("SetupJoinTable failed for table [IngredientAllergy], error: %v", err))
	}

	db.Migrator().DropTable(&model.Allergy{})

	if err = db.AutoMigrate(&model.Allergy{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [Allergy], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [Allergy], error: %v", err))
	}

	as := []model.Allergy{
		{ID: 1, Name: model.AllergyNuts},
		{ID: 2, Name: model.AllergySoy},
		{ID: 3, Name: model.AllergyLatex},
		{ID: 4, Name: model.AllergySesame},
		{ID: 5, Name: model.AllergyCitrus},
		{ID: 6, Name: model.AllergyDye},
		{ID: 7, Name: model.AllergyArtificialFragrance},
		{ID: 8, Name: model.AllergyScent},
		{ID: 9, Name: model.AllergyNone},
	}

	if res := db.Create(as); res.Error != nil {
		if !strings.Contains(res.Error.Error(), "duplicate key") {
			return err
		}
	}

	return nil
}

func migratePreference(db *gorm.DB) error {
	var err error

	db.Debug().Exec(`
	DO $$ BEGIN
		CREATE TYPE ingredient_preference AS ENUM ('paleo', 'vegetarian', 'vegan', 'glutenfree');
	END $$;`)

	if err = db.SetupJoinTable(&model.Ingredient{}, "Preferences", &model.IngredientPreference{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientPreference], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("SetupJoinTable failed for table [IngredientPreference], error: %v", err))
	}

	if err = db.AutoMigrate(&model.Preference{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [Preference], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [Preference], error: %v", err))
	}

	preferences := []model.Preference{
		{ID: 1, Name: model.Vegan},
		{ID: 2, Name: model.Vegetarian},
		{ID: 3, Name: model.GlutenFree},
		{ID: 4, Name: model.Paleo},
	}

	if res := db.Create(preferences); res.Error != nil {
		if !strings.Contains(res.Error.Error(), "duplicate key") {
			return err
		}
	}

	return nil
}

func migrateAcnebreakout(db *gorm.DB) error {
	var err error

	db.Debug().Exec(`
	DO $$ BEGIN
		CREATE TYPE acne_breakout AS ENUM ('never','rarely', 'occasionally', 'frequently', 'veryfrequently', 'always');
	END $$;`)

	if err = db.SetupJoinTable(&model.Ingredient{}, "Acnebreakouts", &model.IngredientAcnebreakout{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientAcnebreakout], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("SetupJoinTable failed for table [IngredientAcnebreakout], error: %v", err))
	}

	db.Migrator().DropTable(&model.Acnebreakout{})

	if err = db.AutoMigrate(&model.Acnebreakout{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [Acnebreakout], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [Acnebreakout], error: %v", err))
	}

	acneBreakouts := []model.Acnebreakout{
		{ID: 1, Frequency: model.NeverAcne},
		{ID: 2, Frequency: model.RarelyAcne},
		{ID: 3, Frequency: model.Occasionally},
		{ID: 4, Frequency: model.FrequentlyAcne},
		{ID: 5, Frequency: model.VeryFrequently},
		{ID: 6, Frequency: model.AlmostAlways},
	}

	if res := db.Create(acneBreakouts); res.Error != nil {
		if !strings.Contains(res.Error.Error(), "duplicate key") {
			return err
		}
	}

	return nil

}

func migrateSkinSensitivity(db *gorm.DB) error {
	var err error

	db.Debug().Exec(`
	DO $$ BEGIN
		CREATE TYPE skin_sensitivity AS ENUM ('never', 'rarely', 'sometimes', 'often', 'frequently', 'always');
	END $$;`)

	if err = db.SetupJoinTable(&model.Ingredient{}, "Skinsensitivities", &model.IngredientSkinsensitivity{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientSkinsensitivity], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("SetupJoinTable failed for table [IngredientSkinsensitivity], error: %v", err))
	}

	db.Migrator().DropTable(&model.Skinsensitivity{})

	if err = db.AutoMigrate(&model.Skinsensitivity{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [SkinSensitivity], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [SkinSensitivity], error: %v", err))
	}

	skinSens := []model.Skinsensitivity{
		{ID: 1, Sensitivity: model.Never},
		{ID: 2, Sensitivity: model.Rarely},
		{ID: 3, Sensitivity: model.Sometimes},
		{ID: 4, Sensitivity: model.Often},
		{ID: 5, Sensitivity: model.Frequently},
		{ID: 6, Sensitivity: model.Always},
	}

	if res := db.Create(skinSens); res.Error != nil {
		if !strings.Contains(res.Error.Error(), "duplicate key") {
			return err
		}
	}

	return nil

}

func migrateSkinType(db *gorm.DB) error {
	var err error

	db.Debug().Exec(`
	DO $$ BEGIN
		CREATE TYPE skin_type AS ENUM ('dry', 'normal', 'combination', 'oily');
	END $$;`)

	if err = db.SetupJoinTable(&model.Ingredient{}, "Skintypes", &model.IngredientSkintype{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientSkintype], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("SetupJoinTable failed for table [IngredientSkintype], error: %v", err))
	}

	if err = db.AutoMigrate(&model.Skintype{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [SkinType], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [SkinType], error: %v", err))
	}

	skinTypes := []*model.Skintype{
		{ID: 1, Type: model.Dry},
		{ID: 2, Type: model.Normal},
		{ID: 3, Type: model.Combination},
		{ID: 4, Type: model.Oily},
	}

	if res := db.Create(skinTypes); res.Error != nil {
		if !strings.Contains(res.Error.Error(), "duplicate key") {
			return err
		}
	}

	return nil
}
