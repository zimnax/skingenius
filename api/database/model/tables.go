package model

import (
	"database/sql/driver"
)

type Product struct {
	ID          uint
	Name        string
	Brand       string
	Link        string
	Ingredients []Ingredient `gorm:"many2many:product_ingredient;"`
}

type Ingredient struct {
	//<<<<<<< HEAD
	//	Id                   uint `gorm:"primaryKey;autoIncrement"`
	//	Name                 string
	//	PubchemId            int
	//	CasNumber            string
	//	ECNumber             string
	//	Synonyms             pq.StringArray          `gorm:"type:text[]"`
	//	SkinType             []SkinType              `gorm:"many2many:ingredient_skintype;"`
	//	SkinSensitivity      []SkinSensitivity       `gorm:"many2many:ingredient_skinsensitivity;"`
	//	AcneBreakout         []AcneBreakouts         `gorm:"many2many:ingredient_acnebreakout;"`
	//	IngredientPreference []IngredientPreferences `gorm:"many2many:ingredient_ingredientpreference;"`
	//	Allergies            []Allergy               `gorm:"many2many:ingredient_allergies;"`
	//	SkinConcerns         []Skinconcern           `gorm:"many2many:ingredient_skinconcern;"`
	//	Ages                 []Age                   `gorm:"many2many:ingredient_age;"`
	//	Benefits             []Benefit               `gorm:"many2many:ingredient_benefit;"`
	//=======
	ID                uint `gorm:"primaryKey;autoIncrement"`
	Name              string
	Skintypes         []Skintype        `gorm:"many2many:ingredient_skintypes;"`
	Skinsensitivities []Skinsensitivity `gorm:"many2many:ingredient_skinsensitivities;"`
	Acnebreakouts     []Acnebreakout    `gorm:"many2many:ingredient_acnebreakouts;"`
	Preferences       []Preference      `gorm:"many2many:ingredient_preferences;"`
	Allergies         []Allergy         `gorm:"many2many:ingredient_allergies;"`
	Skinconcerns      []Skinconcern     `gorm:"many2many:ingredient_skinconcerns;"`
	Ages              []Age             `gorm:"many2many:ingredient_ages;"`
	Benefits          []Benefit         `gorm:"many2many:ingredient_benefits;"`
	//>>>>>>> f64ca42e98e13531eefa029b4a555fbbece40a6b
}

type Allergy struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Skinconcern struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Age struct {
	ID    uint `gorm:"primaryKey"`
	Value int
}

type Benefit struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

// ---  Skin type
type Skintype struct {
	ID   uint          `gorm:"primaryKey"`
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
	*s = SkinTypeValue(value.(string))
	return nil
}

func (s SkinTypeValue) Value() (driver.Value, error) {
	return string(s), nil
}

// ---  Skin sensitivity
type Skinsensitivity struct {
	ID          uint                 `gorm:"primaryKey"`
	Sensitivity SkinSensitivityValue `gorm:"type:skin_sensitivity"`
}

type SkinSensitivityValue string

const (
	Never      SkinSensitivityValue = "never"
	Rarely     SkinSensitivityValue = "rarely"
	Sometimes  SkinSensitivityValue = "sometimes"
	Often      SkinSensitivityValue = "often"
	Frequently SkinSensitivityValue = "frequently"
	Always     SkinSensitivityValue = "always"
)

func (s *SkinSensitivityValue) Scan(value interface{}) error {
	*s = SkinSensitivityValue(value.(string))
	return nil
}

func (s SkinSensitivityValue) Value() (driver.Value, error) {
	return string(s), nil
}

// ---  Acne breakouts
type Acnebreakout struct {
	ID        uint               `gorm:"primaryKey"`
	Frequency AcneBreakoutsValue `gorm:"type:acne_breakout"`
}

type AcneBreakoutsValue string

const (
	NeverAcne      AcneBreakoutsValue = "never"
	RarelyAcne     AcneBreakoutsValue = "rarely"
	Occasionally   AcneBreakoutsValue = "occasionally"
	FrequentlyAcne AcneBreakoutsValue = "frequently"
	VeryFrequently AcneBreakoutsValue = "veryfrequently"
	AlmostAlways   AcneBreakoutsValue = "always"
)

func (s *AcneBreakoutsValue) Scan(value interface{}) error {
	*s = AcneBreakoutsValue(value.(string))
	return nil
}

func (s AcneBreakoutsValue) Value() (driver.Value, error) {
	return string(s), nil
}

// ---  Ingredient preferences
type Preference struct {
	ID   uint                       `gorm:"primaryKey"`
	Name IngredientPreferencesValue `gorm:"type:ingredient_preference"`
}

type IngredientPreferencesValue string

const (
	Paleo      IngredientPreferencesValue = "paleo"
	Vegetarian IngredientPreferencesValue = "vegetarian"
	Vegan      IngredientPreferencesValue = "vegan"
	GlutenFree IngredientPreferencesValue = "glutenfree"
)

func (s *IngredientPreferencesValue) Scan(value interface{}) error {
	*s = IngredientPreferencesValue(value.(string))
	return nil
}

func (s IngredientPreferencesValue) Value() (driver.Value, error) {
	return string(s), nil
}

// ------ Custom join tables

type IngredientSkintype struct {
	IngredientID uint `gorm:"primaryKey"`
	SkintypeID   uint `gorm:"primaryKey"` //  missing field skin_type_id for join table
	Score        uint
}

type IngredientSkinsensitivity struct {
	IngredientID      uint `gorm:"primaryKey"`
	SkinsensitivityID uint `gorm:"primaryKey"`
	Score             uint
}

type IngredientAcnebreakout struct {
	IngredientID   uint `gorm:"primaryKey"`
	AcnebreakoutID uint `gorm:"primaryKey"`
	Score          uint
}

type IngredientPreference struct {
	IngredientID uint `gorm:"primaryKey"`
	PreferenceID uint `gorm:"primaryKey"`
	Score        uint
}

type IngredientAllergy struct {
	IngredientID uint `gorm:"primaryKey"`
	AllergyID    uint `gorm:"primaryKey"`
	Score        uint
}

type IngredientSkinconcern struct {
	IngredientID  uint `gorm:"primaryKey"`
	SkinconcernID uint `gorm:"primaryKey"`
	Score         uint
}

type IngredientAge struct {
	IngredientID uint `gorm:"primaryKey"`
	AgeID        uint `gorm:"primaryKey"`
	Score        uint
}

type IngredientBenefit struct {
	IngredientID uint `gorm:"primaryKey"`
	BenefitID    uint `gorm:"primaryKey"`
	Score        uint
}
