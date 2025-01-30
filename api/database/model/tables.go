package model

import (
	"context"
	"database/sql/driver"
	"fmt"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"skingenius/logger"
)

type UserRoutine struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	UserId      string
	ProductID   uint //`gorm:"foreignKey:ID"` // Product `gorm:"many2one:user_routine_products;"` //`gorm:"foreignKey:ID"`
	Product     Product
	TimeOfDay   string
	TimesPerDay string
	HowLong     string
	Note        string
}

//func (u *UserRoutine) BeforeDelete(tx *gorm.DB) (err error) {
//	fmt.Println(fmt.Sprintf("Before delete UserRoutine, userid:[%s] ", u.UserId))
//	return tx.Exec("DELETE FROM user_routine_products WHERE user_routine_user_id = ?", u.UserId).Error
//}

type QuizIngredients struct {
	SkinTypeIng           []Ingredient
	SkinSensitivityIng    []Ingredient
	AcneBreakoutsIng      []Ingredient
	ProductPreferencesIng []Ingredient
	FreeFromAllergensIng  []Ingredient
	SkinConcernIng        []Ingredient
	AgeIng                []Ingredient
	ProductBenefitIng     []Ingredient
}

type UserRecommendations struct {
	UserId    string  `gorm:"primaryKey"`
	ProductId int     `gorm:"primaryKey"`
	Score     float64 `gorm:"type:decimal(4,2);"`
	//RecommendedProducts pq.Int32Array `gorm:"type:integer[]"`
}

type UserQuiz struct {
	UserId             string `gorm:"primaryKey"`
	SkinType           string
	SkinSensitivity    string
	AcneBreakouts      string
	ProductPreferences pq.StringArray `gorm:"type:text[]"`
	FreeFromAllergens  pq.StringArray `gorm:"type:text[]"`
	SkinConcern        pq.StringArray `gorm:"type:text[]"`
	Age                int
	ProductBenefit     pq.StringArray `gorm:"type:text[]"`
}

type Product struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	Name            string
	Brand           string
	Ingredients     []Ingredient `gorm:"many2many:product_ingredients;"`
	Link            string
	Type            string
	FormulationType string
	FormulatedFor   string
	Price           float64
	Image           string
	Description     string

	Jsoningredients    string             `gorm:"type:jsonb"`
	Deleted            gorm.DeletedAt     // db.Unscoped().Where("age = 20").Find(&users)
	Score              float64            `sql:"-" gorm:"-"`
	Concentrations     map[string]float64 `sql:"-" gorm:"-"`
	ActiveIngredients  []string           `sql:"-" gorm:"-"` // ingredients in product which is same as concern ingredients
	PassiveIngredients []string           `sql:"-" gorm:"-"` // ingredients not accounted for in concern or benefit
	WASTotal           float64            `sql:"-" gorm:"-"` //  Sum up weighted average scores of ingredients for each benefit or concern
}

type Ingredient struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	PubchemId string
	CasNumber string
	ECNumber  string
	INCIName  string `json:"inci_name"`
	//Type               string              // Active, Inactive
	Roleinformulations []Roleinformulation `gorm:"many2many:ingredient_roleinformulations;"`
	Synonyms           pq.StringArray      `gorm:"type:text[]"`
	Skintypes          []Skintype          `gorm:"many2many:ingredient_skintypes;"`
	Skinsensitivities  []Skinsensitivity   `gorm:"many2many:ingredient_skinsensitivities;"`
	Acnebreakouts      []Acnebreakout      `gorm:"many2many:ingredient_acnebreakouts;"`
	Preferences        []Preference        `gorm:"many2many:ingredient_preferences;"`
	Allergies          []Allergy           `gorm:"many2many:ingredient_allergies;"`
	Skinconcerns       []Skinconcern       `gorm:"many2many:ingredient_skinconcerns;"`
	Ages               []Age               `gorm:"many2many:ingredient_ages;"`
	Benefits           []Benefit           `gorm:"many2many:ingredient_benefits;"`

	ConcentrationRinseOffMin float64 `json:"concentration_rinse_off_min"`
	ConcentrationRinseOffMax float64 `json:"concentration_rinse_off_max"`

	ConcentrationLeaveOnMin float64 `json:"concentration_leave_on_min"`
	ConcentrationLeaveOnMax float64 `json:"concentration_leave_on_max"`

	EffectiveAtLowConcentration ConcentrationEffectiveness

	Score float64 `json:"score"`
	Index int     `json:"index"`

	ConcernDescription string `gorm:"-"`
}

//type IngredientScoreEffectiveness struct {
//	Score         float64
//	Effectiveness ConcentrationEffectiveness
//}
//
//func (ise IngredientScoreEffectiveness) AddScoreAndEff(score float64, effectiveness ConcentrationEffectiveness) IngredientScoreEffectiveness {
//	ise.Score = ise.Score + score
//	ise.Effectiveness = effectiveness
//
//	return ise
//}

type ConcentrationEffectiveness string

const (
	EffectiveYes      ConcentrationEffectiveness = "yes"
	EffectiveNo       ConcentrationEffectiveness = "no"
	EffectiveModerate ConcentrationEffectiveness = "moderately effective"
	EffectiveUnknown  ConcentrationEffectiveness = "unknown"
)

type SkinconcernToIngredientDescription struct {
	Ingredientname string
	Concern        string
	Description    string
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
	AllergySeafood             AllergyValue = "seafood"
	AllergyDiary               AllergyValue = "diary"
	AllergyNone                AllergyValue = "no_allergy"
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
	//ConcernRosacea                  SkinconcernValue = "rosacea"
	//ConcernHyperpigmentation        SkinconcernValue = "hyperpigmentation"
	//ConcernMelasma                  SkinconcernValue = "melasma"
	//ConcernCysticAcne               SkinconcernValue = "cystic_acne"
	//ConcernAcne                     SkinconcernValue = "acne"
	//ConcernXerosis                  SkinconcernValue = "xerosis"
	//ConcernDryness                  SkinconcernValue = "dryness"
	//ConcernOiliness                 SkinconcernValue = "oiliness"
	//ConcernUnevenSkinTone           SkinconcernValue = "uneven_skin_tone"
	//ConcernSignsOfAging             SkinconcernValue = "signs_of_aging"
	//ConcernFineLines                SkinconcernValue = "fine_lines"
	//ConcernWrinkles                 SkinconcernValue = "wrinkles"
	//ConcernDarkSpots                SkinconcernValue = "dark_spots"
	//ConcernLostOfElasticityFirmness SkinconcernValue = "lost_of_elasticity_firmness"
	//ConcernVisiblePores             SkinconcernValue = "visible_pores"
	//ConcernCloggedPoresBlackheads   SkinconcernValue = "clogged_pores_blackheads"
	//ConcernRedness                  SkinconcernValue = "redness"
	//ConcernDullness                 SkinconcernValue = "dullness"
	//ConcernDamagedSkin              SkinconcernValue = "damaged_skin"
	//ConcernUnevenTexture            SkinconcernValue = "uneven_texture"
	//ConcernEczema                   SkinconcernValue = "eczema"
	//ConcernPsoriasis                SkinconcernValue = "psoriasis"
	//ConcernDermatitis               SkinconcernValue = "dermatitis"
	//ConcernSunburnedSkin            SkinconcernValue = "sunburned_skin"
	//ConcernDarkCircles              SkinconcernValue = "dark_circles"
	//ConcernBlemishes                SkinconcernValue = "blemishes"
	//ConcernSensitiveSkin            SkinconcernValue = "sensitive_skin"

	ConcernAcne                             SkinconcernValue = "acne"
	ConcernRosacea                          SkinconcernValue = "rosacea"
	ConcernDryness_Dehydration              SkinconcernValue = "dryness_dehydration"
	ConcernHyperpigmentation_UnevenSkinTone SkinconcernValue = "hyperpigmentation_unevenskintone"
	ConcernOiliness_Shine                   SkinconcernValue = "oiliness_shine"
	ConcernFine_lines_Wrinkles              SkinconcernValue = "fine_lines_wrinkles"
	ConcernLoss_of_Elasticity_firmness      SkinconcernValue = "loss_of_elasticity_firmness"
	ConcernVisible_pores_Uneven_texture     SkinconcernValue = "visible_pores_uneven_texture"
	ConcernClogged_pores_blackheads         SkinconcernValue = "clogged_pores_blackheads"
	ConcernDullness                         SkinconcernValue = "dullness"
	ConcernDark_circles                     SkinconcernValue = "dark_circles"
	ConcernBlemishes                        SkinconcernValue = "blemishes"
	ConcernNone                             SkinconcernValue = "no_concern"
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

// ---  Role in formulation
type Roleinformulation struct {
	ID   uint `gorm:"primaryKey"`
	Name RoleInFormulationValue
}

type RoleInFormulationValue string

const (
	Active              RoleInFormulationValue = "active"
	Antioxidant         RoleInFormulationValue = "antioxidant"
	ChelatingAgent      RoleInFormulationValue = "chelating_agent"
	Colorant            RoleInFormulationValue = "colorant"
	Emollient           RoleInFormulationValue = "emollient"
	Emulsifier          RoleInFormulationValue = "emulsifier"
	Exfoliant           RoleInFormulationValue = "exfoliant"
	Fragrance           RoleInFormulationValue = "fragrance"
	Humectant           RoleInFormulationValue = "humectant"
	Occlusive           RoleInFormulationValue = "occlusive"
	PenetrationEnhancer RoleInFormulationValue = "penetration_enhancer"
	Preservative        RoleInFormulationValue = "preservative"
	Solvent             RoleInFormulationValue = "solvent"
	Stabilizer          RoleInFormulationValue = "stabilizer"
	Sunscreen           RoleInFormulationValue = "sunscreen"
	TextureEnhancer     RoleInFormulationValue = "texture_enhancer"
	Thickener           RoleInFormulationValue = "thickener"
	PHAdjuster          RoleInFormulationValue = "pH_adjuster"
)

func (s *RoleInFormulationValue) Scan(value interface{}) error {
	*s = RoleInFormulationValue(value.(string))
	return nil
}

func (s RoleInFormulationValue) Value() (driver.Value, error) {
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
	Paleo        IngredientPreferencesValue = "paleo"
	Vegetarian   IngredientPreferencesValue = "vegetarian"
	Vegan        IngredientPreferencesValue = "vegan"
	GlutenFree   IngredientPreferencesValue = "glutenfree"
	ParabenFree  IngredientPreferencesValue = "parabenfree"
	SulphateFree IngredientPreferencesValue = "sulphatefree"
	SiliconFree  IngredientPreferencesValue = "siliconfree"
	Unscented    IngredientPreferencesValue = "unscented"
	NoPreference IngredientPreferencesValue = "no_preference"
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
	IngredientID uint    `gorm:"primaryKey"`
	SkintypeID   uint    `gorm:"primaryKey"` //  missing field skin_type_id for join table
	Score        float64 `gorm:"type:decimal(4,2);"`
}

func (ip *IngredientSkintype) BeforeCreate(db *gorm.DB) error {
	logger.New().Debug(context.Background(), "Before create IngredientSkintype")

	if customValue, ok := db.Statement.Context.Value(SkintypeCtxKey(ip.SkintypeID)).(float64); ok {
		ip.Score = customValue
	}
	return nil
}

type IngredientSkinsensitivity struct {
	IngredientID      uint `gorm:"primaryKey"`
	SkinsensitivityID uint `gorm:"primaryKey"`
	Score             bool
}

func (ip *IngredientSkinsensitivity) BeforeCreate(db *gorm.DB) error {
	logger.New().Debug(context.Background(), "Before create IngredientSkinsensitivity")
	if customValue, ok := db.Statement.Context.Value(SkinsensetivityCtxKey(ip.SkinsensitivityID)).(bool); ok {
		ip.Score = customValue
	}
	return nil
}

type IngredientAcnebreakout struct {
	IngredientID   uint `gorm:"primaryKey"`
	AcnebreakoutID uint `gorm:"primaryKey"`
	Score          bool
}

func (ip *IngredientAcnebreakout) BeforeCreate(db *gorm.DB) error {
	logger.New().Debug(context.Background(), "Before create IngredientAcnebreakout")

	if customValue, ok := db.Statement.Context.Value(AcnebreakoutsCtxKey(ip.AcnebreakoutID)).(bool); ok {
		ip.Score = customValue
	}
	return nil
}

type IngredientPreference struct {
	IngredientID uint `gorm:"primaryKey"`
	PreferenceID uint `gorm:"primaryKey"`
	Score        bool
}

func (ip *IngredientPreference) BeforeCreate(db *gorm.DB) error {
	logger.New().Debug(context.Background(), "Before create IngredientPreference")

	if customValue, ok := db.Statement.Context.Value(PreferencesCtxKey(ip.PreferenceID)).(bool); ok {
		ip.Score = customValue
	}
	return nil
}

type IngredientAllergy struct {
	IngredientID uint `gorm:"primaryKey"`
	AllergyID    uint `gorm:"primaryKey"`
	Score        bool
}

func (ip *IngredientAllergy) BeforeCreate(db *gorm.DB) error {
	logger.New().Debug(context.Background(), "Before create IngredientAllergy")

	if customValue, ok := db.Statement.Context.Value(AllergiesCtxKey(ip.AllergyID)).(bool); ok {
		ip.Score = customValue
	}
	return nil
}

type IngredientSkinconcern struct {
	IngredientID  uint    `gorm:"primaryKey"`
	SkinconcernID uint    `gorm:"primaryKey"`
	Score         float64 `gorm:"type:decimal(4,2);"`
	Description   string
}

func (ip *IngredientSkinconcern) BeforeCreate(db *gorm.DB) error {
	logger.New().Debug(context.Background(), "Before create IngredientSkinconcern")

	if customValue, ok := db.Statement.Context.Value(SkinconcernCtxKey(ip.SkinconcernID)).(float64); ok {
		ip.Score = customValue
	}

	if customValue, ok := db.Statement.Context.Value(SkinconcernDescCtxKey(ip.SkinconcernID)).(string); ok {
		ip.Description = customValue
	}

	return nil
}

type IngredientAge struct {
	IngredientID uint `gorm:"primaryKey"`
	AgeID        uint `gorm:"primaryKey"`
	Score        bool
}

func (ip *IngredientAge) BeforeCreate(db *gorm.DB) error {
	logger.New().Debug(context.Background(), "Before create IngredientAge")

	if customValue, ok := db.Statement.Context.Value(AgeCtxKey(ip.AgeID)).(bool); ok {
		ip.Score = customValue
	} else {
		fmt.Println(fmt.Sprintf("ERROR, failed to get Age value from context by key: %s", AgeCtxKey(ip.AgeID)))
	}
	return nil
}

type IngredientBenefit struct {
	IngredientID uint `gorm:"primaryKey"`
	BenefitID    uint `gorm:"primaryKey"`
	Score        float64
}

func (ip *IngredientBenefit) BeforeCreate(db *gorm.DB) error {
	logger.New().Debug(context.Background(), "Before create IngredientBenefit")

	if customValue, ok := db.Statement.Context.Value(BenefitsCtxKey(ip.BenefitID)).(float64); ok {
		ip.Score = customValue
	}
	return nil
}

type IngredientRoleinformulation struct {
	IngredientID        uint `gorm:"primaryKey"`
	RoleinformulationID uint `gorm:"primaryKey"`
}

type ProductIngredients struct {
	IngredientID uint `gorm:"primaryKey"`
	ProductID    uint `gorm:"primaryKey"`
	Index        int
}

func (pi *ProductIngredients) BeforeCreate(db *gorm.DB) error {
	logger.New().Debug(context.Background(), "Before create ProductIngredient")

	if indexVal, ok := db.Statement.Context.Value(IngredientIndexCtxKey(pi.IngredientID)).(int); ok {
		pi.Index = indexVal
	}
	return nil
}
