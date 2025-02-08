package database

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"skingenius/database/model"
	"skingenius/logger"
	"strings"
)

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

	if err = migrateRoleInFormulation(db); err != nil {
		return err
	}

	if err = db.SetupJoinTable(&model.Product{}, "Ingredients", &model.ProductIngredients{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [ProductIngredients], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("SetupJoinTable failed for table [ProductIngredients], error: %v", err))
	}

	if err = db.AutoMigrate(&model.Product{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [Product], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [Product], error: %v", err))
	}

	if err = db.AutoMigrate(&model.Ingredient{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [Ingredient], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [Ingredient], error: %v", err))
	}

	if err = db.AutoMigrate(&model.UserRecommendations{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [UserRecommendations], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [UserRecommendations], error: %v", err))
	}

	if err = db.AutoMigrate(&model.UserQuiz{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [UserQuiz], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [UserQuiz], error: %v", err))
	}

	if err = db.AutoMigrate(&model.UserRoutine{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [UserRoutine], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [UserRoutine], error: %v", err))
	}

	/*
			LIVE SEARCH

			alter table products ADD column ts tsvector
			generated always as (
			to_tsvector('english', products.name) || ' ' ||
			to_tsvector('english', products.brand)
			) STORED;

			create index ts_idx on products using GIN (ts);
			select name, brand from products where ts @@ websearch_to_tsquery('english', 'hydr');


		CREATE EXTENSION pg_trgm // check if needed


		select name, brand from products where name LIKE '%bra%' OR brand LIKE '%bra%'

	*/

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
		{ID: 9, Name: model.AllergySeafood},
		{ID: 10, Name: model.AllergyNone},
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
		CREATE TYPE ingredient_preference AS ENUM ('paleo', 'vegetarian', 'vegan', 'glutenfree','unscented', 'parabenfree', 'sulphatefree', 'siliconfree');
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
		{ID: 5, Name: model.ParabenFree},
		{ID: 6, Name: model.SulphateFree},
		{ID: 7, Name: model.SiliconFree},
		{ID: 8, Name: model.Unscented},
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

func migrateRoleInFormulation(db *gorm.DB) error {
	var err error

	if err = db.SetupJoinTable(&model.Ingredient{}, "Roleinformulations", &model.IngredientRoleinformulation{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("SetupJoinTable failed for table [IngredientRoleinformulation], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("SetupJoinTable failed for table [IngredientRoleinformulation], error: %v", err))
	}

	if err = db.AutoMigrate(&model.Roleinformulation{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [Roleinformulation], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [Roleinformulation], error: %v", err))
	}

	ages := []model.Roleinformulation{
		{ID: 1, Name: model.Active},
		{ID: 2, Name: model.Antioxidant},
		{ID: 3, Name: model.ChelatingAgent},
		{ID: 4, Name: model.Colorant},
		{ID: 5, Name: model.Emollient},
		{ID: 6, Name: model.Emulsifier},
		{ID: 7, Name: model.Exfoliant},
		{ID: 8, Name: model.Fragrance},
		{ID: 9, Name: model.Humectant},
		{ID: 10, Name: model.Occlusive},
		{ID: 11, Name: model.PenetrationEnhancer},
		{ID: 12, Name: model.Preservative},
		{ID: 13, Name: model.Solvent},
		{ID: 14, Name: model.Stabilizer},
		{ID: 15, Name: model.Sunscreen},
		{ID: 16, Name: model.TextureEnhancer},
		{ID: 17, Name: model.Thickener},
		{ID: 18, Name: model.PHAdjuster},
	}

	if res := db.Create(ages); res.Error != nil {
		if !strings.Contains(res.Error.Error(), "duplicate key") {
			return err
		}
	}

	return nil
}
