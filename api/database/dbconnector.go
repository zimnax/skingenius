package database

import (
	"context"
	"skingenius/database/model"
)

const dbname = "skingenius_new"

type Connector interface {
	SaveIngredient(ctx context.Context, ingredient *model.Ingredient) error
	FindIngredientByName(ctx context.Context, name string) (*model.Ingredient, error)
	FindIngredientByAlias(ctx context.Context, alias string) (*model.Ingredient, error)
	FindIngredientByINCIName(ctx context.Context, inci string) (*model.Ingredient, error)
	GetSkinConcernDescriptionByIngredients(ctx context.Context, ingredients []string, concern string) ([]model.SkinconcernToIngredientDescription, error)
	DeleteIngredientByName(ctx context.Context, name string) error

	SaveProduct(ctx context.Context, product *model.Product) error
	FindAllProducts(ctx context.Context) ([]model.Product, error)
	FindProductByName(ctx context.Context, name string) (*model.Product, error)
	FindAllProductsWithIngredients(ctx context.Context, ingredients []string, accuracy uint) ([]model.Product, error)
	FindProductsByIds(ctx context.Context, ids []int32) ([]model.Product, error)
	DeleteProductByName(ctx context.Context, name string) error

	GetAllIngredients(context.Context) ([]model.Ingredient, error)
	GetAllSkintypes(context.Context) ([]model.Skintype, error)
	GetAllSkinsensetivity(context.Context) ([]model.Skinsensitivity, error)
	GetAllAcneBreakouts(context.Context) ([]model.Acnebreakout, error)
	GetAllPreferences(context.Context) ([]model.Preference, error)
	GetAllAllergies(context.Context) ([]model.Allergy, error)
	GetAllSkinconcerns(context.Context) ([]model.Skinconcern, error)
	GetAllAge(context.Context) ([]model.Age, error)
	GetAllBenefits(context.Context) ([]model.Benefit, error)

	GetIngredientsBySkintype(context.Context, string) ([]model.Ingredient, error)
	GetIngredientsBySkinsensitivity(context.Context, string) ([]model.Ingredient, error)
	GetIngredientsByAcneBreakouts(context.Context, string) ([]model.Ingredient, error)
	GetIngredientsByPreferences(context.Context, []string) ([]model.Ingredient, error)
	GetIngredientsByAllergies(context.Context, []string) ([]model.Ingredient, error)
	GetIngredientsBySkinconcerns(context.Context, []string) ([]model.Ingredient, error)
	GetIngredientsByAge(context.Context, string) ([]model.Ingredient, error)
	GetIngredientsByBenefits(context.Context, []string) ([]model.Ingredient, error)

	SetupJoinTables() error

	SaveRecommendations(context.Context, []model.UserRecommendations) error
	GetRecommendations(context.Context, string) ([]model.UserRecommendations, error)
	DeleteRecommendations(context.Context, string) error

	SaveQuiz(ctx context.Context, quiz model.UserQuiz) error
	GetQuiz(ctx context.Context, userId string) (model.UserQuiz, error)

	LiveSearch(ctx context.Context, search string) ([]model.Product, error)

	SaveUserRoutine(ctx context.Context, routine model.UserRoutine) error
	GetUserRoutine(ctx context.Context, userId string) (model.UserRoutine, error)
	DeleteUserRoutine(ctx context.Context, userId string) error
}
