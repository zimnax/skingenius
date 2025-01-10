package main

//const (
//	IngredientName = iota
//	PubChemCID
//	INCIName
//	INCINumber
//	CASNumber
//	Vegetarian
//	Vegan
//	GlutenFree
//	Paleo
//	NutFree
//	SoyFree
//	LatexFree
//	SesameFree
//	CitrusFree
//	DyeFree
//	FragranceFree
//	ScentFree
//	Normal
//	Dry
//	Oily
//	Combination
//	NotSensitive
//	ALittleSensitive
//	ModeratelySensitive
//	Sensitive
//	VerySensitive
//	ExtremelySensitive
//	NotAcneProne
//	ALittleAcneProne
//	ModeratelyAcneProne
//	AcneProne
//	VeryAcneProne
//	ExtremelyAcneProne
//	SuitableForMatureSkin
//	Acne
//	Rosacea
//	CysticAcne
//	Hyperpigmentation
//	Melasma
//	Xerosis
//	Dryness
//	Redness
//	Oiliness
//	SignsOfAging
//	DarkSpots
//	SensitiveSkin
//	Dullness
//	UnevenSkinTone
//	Wrinkles
//	FineLines
//	LossOfElasticityFirmness
//	DamagedSkin
//	VisiblePores
//	CloggedPoresBlackheads
//	UnevenTexture
//	Eczema
//	Psoriasis
//	Dermatitis
//	SunburnedSkin
//	DarkCircles
//	Blemishes
//	Moisturizing
//	Nourishing
//	Hydrating
//	Exfoliating
//	Calming
//	Soothing
//	UVBarrier
//	Healing
//	Smoothing
//	ReducesAcne
//	ReducesBlemishes
//	ReducesWrinkles
//	ImprovesSymptomsOfEczema
//	ImprovesSymptomsOfPsoriasis
//	ImprovesSymptomsOfDermatitis
//	Brightening
//	ImprovesSkinTone
//	ReducesInflammation
//	MinimizesPores
//	AntiAging
//	Firming
//	Detoxifying
//	Balancing
//	ReducesRedness
//	Clarifying
//	AntiBacterial
//	StimulatesCollagenProduction
//	ReducesFineLines
//	AntioxidantProtection
//	SkinBarrierProtection
//	Teen
//	Twenties
//	Thirties
//	Forties
//	Fifties
//	SixtiesPlus
//)

const (
	IngredientName = iota
	INCIName
	Aliases
	Allergen_Potential
	Risk_of_irritation
	Role_in_formulation
	Note
	Source
	Comedogenic_Rating

	Concentration_rinse_off
	Concentration_leave_on
	Concentration_sensitive

	SideEffects
	WorksWellWith

	Normal
	Dry
	Oily
	Combination

	NotSensitive
	ALittleSensitive
	ModeratelySensitive
	Sensitive
	VerySensitive
	ExtremelySensitive

	NotAcneProne
	ALittleAcneProne
	ModeratelyAcneProne
	AcneProne
	VeryAcneProne
	ExtremelyAcneProne

	Teen
	Twenties
	Thirties
	Forties
	Fifties
	SixtiesPlus

	Vegetarian
	Vegan
	GlutenFree
	Paleo
	Unscented
	ParabenFree
	SulphateFree
	SiliconFree

	NutFree
	SoyFree
	LatexFree
	SesameFree
	CitrusFree
	DyeFree
	FragranceFree
	ScentFree
	SeafoodFree
	DiaryFree

	Rosacea
	RosaceaDescription

	Hyperpigmentation_UnevenSkin_tone
	Hyperpigmentation_UnevenSkin_tone_Description

	Acne
	Acne_Description

	Dryness_Dehydration
	Dryness_Dehydration_Description

	Oiliness_Shine
	Oiliness_Shine_Description

	Fine_lines_Wrinkles
	Fine_lines_Wrinkles_Description

	Loss_of_Elasticity_firmness
	Loss_of_Elasticity_firmness_Description

	Visible_pores_Uneven_texture
	Visible_pores_Uneven_texture_Description

	Clogged_pores_blackheads
	Clogged_pores_blackheads_Description

	Dullness
	Dullness_Description

	Dark_circles
	Dark_circles_Description

	Blemishes
	Blemishes_Description

	Moisturizing
	Nourishing
	//Hydrating
	Exfoliating
	Soothing
	//Calming
	UVBarrier
	Healing
	Smoothing
	//ReducesAcne
	//ReducesBlemishes
	//ReducesWrinkles
	//ImprovesSymptomsOfEczema
	//ImprovesSymptomsOfPsoriasis
	//ImprovesSymptomsOfDermatitis
	Brightening
	//ImprovesSkinTone
	//ReducesInflammation
	MinimizesPores
	//AntiAging
	Firming
	//Detoxifying
	//Balancing
	//ReducesRedness
	//Clarifying
	//AntiBacterial
	//StimulatesCollagenProduction
	//ReducesFineLines
	//AntioxidantProtection
	//SkinBarrierProtection

	Concentrations
	Effective_at_low_concentrations
	ScoreCoefficientForConcentration
)

const (
	ProductIngredientName = iota
	ProductBrand
	ProductName
	ProductLink
	ProductType
	FormulationType
	Price
)
