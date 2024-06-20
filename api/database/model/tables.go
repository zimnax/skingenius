package model

import (
	"database/sql/driver"
	"fmt"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint
	Name        string
	Brand       string
	Link        string
	Ingredients []Ingredient `gorm:"many2many:product_ingredient;"`
}

type Ingredient struct {
	ID                uint `gorm:"primaryKey;autoIncrement"`
	Name              string
	PubchemId         string
	CasNumber         string
	ECNumber          string
	Synonyms          pq.StringArray    `gorm:"type:text[]"`
	Skintypes         []Skintype        `gorm:"many2many:ingredient_skintypes;"`
	Skinsensitivities []Skinsensitivity `gorm:"many2many:ingredient_skinsensitivities;"`
	Acnebreakouts     []Acnebreakout    `gorm:"many2many:ingredient_acnebreakouts;"`
	Preferences       []Preference      `gorm:"many2many:ingredient_preferences;"`
	Allergies         []Allergy         `gorm:"many2many:ingredient_allergies;"`
	Skinconcerns      []Skinconcern     `gorm:"many2many:ingredient_skinconcerns;"`
	Ages              []Age             `gorm:"many2many:ingredient_ages;"`
	Benefits          []Benefit         `gorm:"many2many:ingredient_benefits;"`
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
	Score        int
}

func (ip *IngredientSkintype) BeforeCreate(db *gorm.DB) error {
	fmt.Println("Before create IngredientSkintype")
	if customValue, ok := db.Statement.Context.Value(SkintypeCtxKey(ip.SkintypeID)).(int); ok {
		ip.Score = customValue
	}
	return nil
}

type IngredientSkinsensitivity struct {
	IngredientID      uint `gorm:"primaryKey"`
	SkinsensitivityID uint `gorm:"primaryKey"`
	Score             int
}

func (ip *IngredientSkinsensitivity) BeforeCreate(db *gorm.DB) error {
	fmt.Println("Before create IngredientSkinsensitivity")
	if customValue, ok := db.Statement.Context.Value(SkinsensetivityCtxKey(ip.SkinsensitivityID)).(int); ok {
		ip.Score = customValue
	}
	return nil
}

type IngredientAcnebreakout struct {
	IngredientID   uint `gorm:"primaryKey"`
	AcnebreakoutID uint `gorm:"primaryKey"`
	Score          int
}

func (ip *IngredientAcnebreakout) BeforeCreate(db *gorm.DB) error {
	fmt.Println("Before create IngredientAcnebreakout")
	if customValue, ok := db.Statement.Context.Value(AcnebreakoutsCtxKey(ip.AcnebreakoutID)).(int); ok {
		ip.Score = customValue
	}
	return nil
}

type IngredientPreference struct {
	IngredientID uint `gorm:"primaryKey"`
	PreferenceID uint `gorm:"primaryKey"`
	Score        int
}

func (ip *IngredientPreference) BeforeCreate(db *gorm.DB) error {
	fmt.Println("Before create IngredientPreference")
	if customValue, ok := db.Statement.Context.Value(PreferencesCtxKey(ip.PreferenceID)).(int); ok {
		ip.Score = customValue
	}
	return nil
}

type IngredientAllergy struct {
	IngredientID uint `gorm:"primaryKey"`
	AllergyID    uint `gorm:"primaryKey"`
	Score        int
}

func (ip *IngredientAllergy) BeforeCreate(db *gorm.DB) error {
	fmt.Println("Before create IngredientAllergy")
	if customValue, ok := db.Statement.Context.Value(AllergiesCtxKey(ip.AllergyID)).(int); ok {
		ip.Score = customValue
	}
	return nil
}

type IngredientSkinconcern struct {
	IngredientID  uint `gorm:"primaryKey"`
	SkinconcernID uint `gorm:"primaryKey"`
	Score         int
}

func (ip *IngredientSkinconcern) BeforeCreate(db *gorm.DB) error {
	fmt.Println("Before create IngredientSkinconcern")
	if customValue, ok := db.Statement.Context.Value(SkinconcernCtxKey(ip.SkinconcernID)).(int); ok {
		ip.Score = customValue
	}
	return nil
}

type IngredientAge struct {
	IngredientID uint `gorm:"primaryKey"`
	AgeID        uint `gorm:"primaryKey"`
	Score        int
}

func (ip *IngredientAge) BeforeCreate(db *gorm.DB) error {
	fmt.Println("Before create IngredientAge")
	if customValue, ok := db.Statement.Context.Value(AgeCtxKey).(int); ok {
		ip.Score = customValue
	}
	return nil
}

type IngredientBenefit struct {
	IngredientID uint `gorm:"primaryKey"`
	BenefitID    uint `gorm:"primaryKey"`
	Score        int
}

func (ip *IngredientBenefit) BeforeCreate(db *gorm.DB) error {
	fmt.Println("Before create IngredientBenefit")
	if customValue, ok := db.Statement.Context.Value(BenefitsCtxKey).(int); ok {
		ip.Score = customValue
	}
	return nil
}
