package database

import (
	"context"
	"skingenius/database/model"
)

const dbname = "skingenius_new"

type Connector interface {
	//FilerHardParameters(context.Context, string, string, string) ([]string, error)
	SaveIngredient(ctx context.Context, ingredient *model.Ingredient) error
	FindIngredientByName(ctx context.Context, name string) (*model.Ingredient, error)
	FindIngredientByAlias(ctx context.Context, alias string) (*model.Ingredient, error)
	//SavePreference(ctx context.Context, preference *model.Preference) error

	SaveProduct(ctx context.Context, product *model.Product) error
	FindAllProducts(ctx context.Context) ([]model.Product, error)
	FindProductByName(ctx context.Context, name string) (*model.Product, error)
	FindAllProductsWithIngredients(ctx context.Context, ingredients []string, accuracy uint) ([]model.Product, error)
	FindProductsByIds(ctx context.Context, ids []int32) ([]model.Product, error)

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

	SaveRecommendations(context.Context, string, []int32) error
	GetRecommendations(context.Context, string) ([]int32, error)

	SaveQuiz(ctx context.Context, quiz model.UserQuiz) error
	GetQuiz(ctx context.Context, userId string) (model.UserQuiz, error)
}
