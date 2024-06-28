package model

import (
	"database/sql/driver"
	"fmt"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
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

	Score int
}

type Allergy struct {
	ID   uint `gorm:"primaryKey"`
	Name AllergyValue
}

type AllergyValue string

const (
	AllergyNuts                AllergyValue = "nuts"
	AllergySoy                 AllergyValue = "soy"
	AllergyLatex               AllergyValue = "latex"
	AllergySesame              AllergyValue = "sesame"
	AllergyCitrus              AllergyValue = "citrus"
	AllergyDye                 AllergyValue = "dye"
	AllergyArtificialFragrance AllergyValue = "artificial_fragrance"
	AllergyScent               AllergyValue = "scent"
)

func (s *AllergyValue) Scan(value interface{}) error {
	*s = AllergyValue(value.(string))
	return nil
}

func (s AllergyValue) Value() (driver.Value, error) {
	return string(s), nil
}

type Skinconcern struct {
	ID   uint `gorm:"primaryKey"`
	Name SkinconcernValue
}

type SkinconcernValue string

const (
	ConcernRosacea                  SkinconcernValue = "rosacea"
	ConcernHyperpigmentation        SkinconcernValue = "hyperpigmentation"
	ConcernMelasma                  SkinconcernValue = "melasma"
	ConcernCysticAcne               SkinconcernValue = "cystic_acne"
	ConcernAcne                     SkinconcernValue = "acne"
	ConcernXerosis                  SkinconcernValue = "xerosis"
	ConcernDryness                  SkinconcernValue = "dryness"
	ConcernOiliness                 SkinconcernValue = "oiliness"
	ConcernUnevenSkinTone           SkinconcernValue = "uneven_skin_tone"
	ConcernSignsOfAging             SkinconcernValue = "signs_of_aging"
	ConcernFineLines                SkinconcernValue = "fine_lines"
	ConcernWrinkles                 SkinconcernValue = "wrinkles"
	ConcernDarkSpots                SkinconcernValue = "dark_spots"
	ConcernLostOfElasticityFirmness SkinconcernValue = "lost_of_elasticity_firmness"
	ConcernVisiblePores             SkinconcernValue = "visible_pores"
	ConcernCloggedPoresBlackheads   SkinconcernValue = "clogged_pores_blackheads"
	ConcernRedness                  SkinconcernValue = "redness"
	ConcernDullness                 SkinconcernValue = "dullness"
	ConcernDamagedSkin              SkinconcernValue = "damaged_skin"
	ConcernUnevenTexture            SkinconcernValue = "uneven_texture"
	ConcernEczema                   SkinconcernValue = "eczema"
	ConcernPsoriasis                SkinconcernValue = "psoriasis"
	ConcernDermatitis               SkinconcernValue = "dermatitis"
	ConcernSunburnedSkin            SkinconcernValue = "sunburned_skin"
	ConcernDarkCircles              SkinconcernValue = "dark_circles"
	ConcernBlemishes                SkinconcernValue = "blemishes"
	ConcernSensitiveSkin            SkinconcernValue = "sensitive_skin"
)

func (s *SkinconcernValue) Scan(value interface{}) error {
	*s = SkinconcernValue(value.(string))
	return nil
}

func (s SkinconcernValue) Value() (driver.Value, error) {
	return string(s), nil
}

type Age struct {
	ID    uint `gorm:"primaryKey"`
	Value AgeValue
}

type AgeValue int

const (
	Age10 AgeValue = 10
	Age20 AgeValue = 20
	Age30 AgeValue = 30
	Age40 AgeValue = 40
	Age50 AgeValue = 50
	Age60 AgeValue = 60
)

func (s *AgeValue) Scan(value interface{}) error {
	*s = AgeValue(value.(int64))
	return nil
}

func (s AgeValue) Value() (driver.Value, error) {
	return int(s), nil
}

type Benefit struct {
	ID   uint `gorm:"primaryKey"`
	Name BenefitValue
}

type BenefitValue string

const (
	BenefitMoisturizing                 BenefitValue = "moisturizing"
	BenefitNourishing                   BenefitValue = "nourishing"
	BenefitHydrating                    BenefitValue = "hydrating"
	BenefitExfoliating                  BenefitValue = "exfoliating"
	BenefitCalming                      BenefitValue = "calming"
	BenefitSoothing                     BenefitValue = "soothing"
	BenefitUVBarrier                    BenefitValue = "uv_barrier"
	BenefitHealing                      BenefitValue = "healing"
	BenefitSmoothing                    BenefitValue = "smoothing"
	BenefitReducesAcne                  BenefitValue = "reduces_acne"
	BenefitReducesBlemishes             BenefitValue = "reduces_blemishes"
	BenefitReducesWrinkles              BenefitValue = "reduces_wrinkles"
	BenefitImprovesSymptomsOfEczema     BenefitValue = "improves_symptoms_of_eczema"
	BenefitImprovesSymptomsOfPsoriasis  BenefitValue = "improves_symptoms_of_psoriasis"
	BenefitImprovesSymptomsOfDermatitis BenefitValue = "improves_symptoms_of_dermatitis"
	BenefitBrightening                  BenefitValue = "brightening"
	BenefitImprovesSkinTone             BenefitValue = "improves_skin_tone"
	BenefitReducesInflammation          BenefitValue = "reduces_inflammation"
	BenefitMinimizesPores               BenefitValue = "minimizes_pores"
	BenefitAntiAging                    BenefitValue = "anti_aging"
	BenefitFirming                      BenefitValue = "firming"
	BenefitDetoxifying                  BenefitValue = "detoxifying"
	BenefitBalancing                    BenefitValue = "balancing"
	BenefitReducesRedness               BenefitValue = "reduces_redness"
	BenefitClarifying                   BenefitValue = "clarifying"
	BenefitAntiBacterial                BenefitValue = "anti_bacterial"
	BenefitStimulatesCollagenProduction BenefitValue = "stimulates_collagen_production"
	BenefitReducesFineLines             BenefitValue = "reduces_fine_lines"
	BenefitAntioxidantProtection        BenefitValue = "antioxidant_protection"
	BenefitSkinBarrierProtection        BenefitValue = "skin_barrier_protection"
)

func (s *BenefitValue) Scan(value interface{}) error {
	*s = BenefitValue(value.(string))
	return nil
}

func (s BenefitValue) Value() (driver.Value, error) {
	return string(s), nil
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
	if customValue, ok := db.Statement.Context.Value(AgeCtxKey(ip.AgeID)).(int); ok {
		ip.Score = customValue
	} else {
		fmt.Println(fmt.Sprintf("ERROR"))
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
	if customValue, ok := db.Statement.Context.Value(BenefitsCtxKey(ip.BenefitID)).(int); ok {
		ip.Score = customValue
	}
	return nil
}
