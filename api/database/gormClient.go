package database

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"skingenius/database/model"
	"skingenius/logger"
)

type GormConnector struct {
	db *gorm.DB
}

func (g *GormConnector) SaveIngredient(ingredient *model.Ingredient) error {
	db := g.db.Save(ingredient)
	return db.Error
}

func (g *GormConnector) IngredientBySkinType(ctx context.Context, s string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GormConnector) IngredientBySkinSensitivity(ctx context.Context, s string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GormConnector) IngredientByAcne(ctx context.Context, s string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GormConnector) IngredientByPreferences(ctx context.Context, strings []string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GormConnector) IngredientByAllergens(ctx context.Context, strings []string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GormConnector) IngredientBySkinConcern(ctx context.Context, s string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GormConnector) IngredientByAge(ctx context.Context, s string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GormConnector) IngredientByProductBenefit(ctx context.Context, s string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GormConnector) FilerHardParameters(ctx context.Context, s string, s2 string, s3 string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func NewGormClient(host string, port int, user, password string) (Connector, error) {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	logger.New().Info(context.Background(), fmt.Sprintf("DB config: %v", dbInfo))

	//db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, dbname))
	//if err != nil {
	//	logger.New().Error(context.Background(), "Error connecting to the database: ", err)
	//}
	////defer db.Close()
	//
	//err = db.Ping()
	//if err != nil {
	//	return nil, fmt.Errorf("Error pinging database: ", err)
	//}
	//logger.New().Info(context.Background(), "Connected to the database!")
	//
	//return &PgConnector{db: db}, nil

	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		logger.New().Error(context.Background(), "Error connecting to the database: ", err)
		return nil, fmt.Errorf("Error connecting to the database: ", err)
	}
	logger.New().Info(context.Background(), "Connected to the database!")

	if migrationErr := automigrate(db); migrationErr != nil {
		return nil, migrationErr
	}

	return &GormConnector{db: db}, nil
}

func automigrate(db *gorm.DB) error {
	logger.New().Info(context.Background(), "Auto-migration started")

	var err error

	// ----- skin type
	db.Debug().Exec(`
	DO $$ BEGIN
		CREATE TYPE skin_type AS ENUM ('dry', 'normal', 'combination', 'oily');
	END $$;`)

	if err = db.AutoMigrate(&model.SkinType{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [SkinType], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [SkinType], error: %v", err))
	}

	// ----- skin sensitivity
	db.Debug().Exec(`
	DO $$ BEGIN
		CREATE TYPE skin_sensitivity AS ENUM ('never', 'rarely', 'sometimes', 'often', 'frequently');
	END $$;`)

	if err = db.AutoMigrate(&model.SkinSensitivity{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [SkinSensitivity], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [SkinSensitivity], error: %v", err))
	}

	// ----- acne breakouts
	db.Debug().Exec(`
	DO $$ BEGIN
		CREATE TYPE acne_breakout AS ENUM ('rarely', 'occasionally', 'frequently', 'veryfrequently', 'always');
	END $$;`)

	if err = db.AutoMigrate(&model.AcneBreakouts{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [AcneBreakouts], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [AcneBreakouts], error: %v", err))
	}

	// ----- ingredient preferences
	db.Debug().Exec(`
	DO $$ BEGIN
		CREATE TYPE ingredient_preference AS ENUM ('paleo', 'vegetarian', 'vegan', 'glutenfree');
	END $$;`)

	if err = db.AutoMigrate(&model.IngredientPreferences{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [IngredientPreferences], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [IngredientPreferences], error: %v", err))
	}

	if err = db.AutoMigrate(&model.Product{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [Product], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [Product], error: %v", err))
	}

	if err = db.AutoMigrate(&model.Ingredient{}); err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("Automigration failed for table [Ingredient], error: %v", err))
		return fmt.Errorf(fmt.Sprintf("Automigration failed for table [Ingredient], error: %v", err))
	}

	logger.New().Info(context.Background(), "Auto-migration finished successfully")
	return nil
}
