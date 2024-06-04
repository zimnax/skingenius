package model

import "database/sql/driver"

type Product struct {
	Id          uint `gorm:"primaryKey;autoIncrement"`
	Name        string
	Brand       string
	Link        string
	Ingredients []Ingredient `gorm:"many2many:product_ingredient;"`
}

type Ingredient struct {
	Id                   uint `gorm:"primaryKey;autoIncrement"`
	Name                 string
	SkinType             []SkinType              `gorm:"many2many:ingredient_skintype;"`
	SkinSensitivity      []SkinSensitivity       `gorm:"many2many:ingredient_skinsensitivity;"`
	AcneBreakout         []AcneBreakouts         `gorm:"many2many:ingredient_acnebreakout;"`
	IngredientPreference []IngredientPreferences `gorm:"many2many:ingredient_ingredientpreference;"`
	Allergies            []Allergy               `gorm:"many2many:ingredient_allergies;"`
	SkinConcerns         []Skinconcern           `gorm:"many2many:ingredient_skinconcern;"`
	Ages                 []Age                   `gorm:"many2many:ingredient_age;"`
	Benefits             []Benefit               `gorm:"many2many:ingredient_benefit;"`
}

type Allergy struct {
	Id   uint `gorm:"primaryKey;autoIncrement"`
	Name string
}

type Skinconcern struct {
	Id    uint `gorm:"primaryKey;autoIncrement"`
	Name  string
	Score int
}

type Age struct {
	Id    uint `gorm:"primaryKey;autoIncrement"`
	Value int
}

type Benefit struct {
	Id    uint `gorm:"primaryKey;autoIncrement"`
	Name  string
	Score int
}

// ---  Skin type
type SkinType struct {
	Id uint `gorm:"primaryKey;autoIncrement"`
	//SkinType SkinTypeValue `gorm:"column:skintype;type:enum('dry', 'normal', 'combination', 'oily')"`
	Type SkinTypeValue `gorm:"type:skin_type"`
}

type SkinTypeValue string

const (
	Dry         SkinTypeValue = "dry"
	Normal      SkinTypeValue = "normal"
	Combination SkinTypeValue = "combination"
	Oily        SkinTypeValue = "oily"
)

func (s *SkinTypeValue) Scan(value interface{}) error {
	*s = SkinTypeValue(value.([]byte))
	return nil
}

func (s SkinTypeValue) Value() (driver.Value, error) {
	return string(s), nil
}

// ---  Skin sensitivity
type SkinSensitivity struct {
	Id              uint                 `gorm:"primaryKey;autoIncrement"`
	SkinSensitivity SkinSensitivityValue `gorm:"type:skin_sensitivity"`
}

type SkinSensitivityValue string

const (
	Never      SkinSensitivityValue = "never"
	Rarely     SkinSensitivityValue = "rarely"
	Sometimes  SkinSensitivityValue = "sometimes"
	Often      SkinSensitivityValue = "often"
	Frequently SkinSensitivityValue = "frequently"
)

func (s *SkinSensitivityValue) Scan(value interface{}) error {
	*s = SkinSensitivityValue(value.([]byte))
	return nil
}

func (s SkinSensitivityValue) Value() (driver.Value, error) {
	return string(s), nil
}

// ---  Acne breakouts
type AcneBreakouts struct {
	Id        uint               `gorm:"primaryKey;autoIncrement"`
	Frequency AcneBreakoutsValue `gorm:"type:acne_breakout"`
}

type AcneBreakoutsValue string

const (
	RarelyAcne     AcneBreakoutsValue = "rarely"
	Occasionally   AcneBreakoutsValue = "occasionally"
	FrequentlyAcne AcneBreakoutsValue = "frequently"
	VeryFrequently AcneBreakoutsValue = "veryfrequently"
	AlmostAlways   AcneBreakoutsValue = "always"
)

func (s *AcneBreakoutsValue) Scan(value interface{}) error {
	*s = AcneBreakoutsValue(value.([]byte))
	return nil
}

func (s AcneBreakoutsValue) Value() (driver.Value, error) {
	return string(s), nil
}

// ---  Ingredient preferences
type IngredientPreferences struct {
	Id         uint                       `gorm:"primaryKey;autoIncrement"`
	Preference IngredientPreferencesValue `gorm:"type:ingredient_preference"`
}

type IngredientPreferencesValue string

const (
	Paleo      IngredientPreferencesValue = "paleo"
	Vegetarian IngredientPreferencesValue = "vegetarian"
	Vegan      IngredientPreferencesValue = "vegan"
	GlutenFree IngredientPreferencesValue = "glutenfree"
)

func (s *IngredientPreferencesValue) Scan(value interface{}) error {
	*s = IngredientPreferencesValue(value.([]byte))
	return nil
}

func (s IngredientPreferencesValue) Value() (driver.Value, error) {
	return string(s), nil
}
