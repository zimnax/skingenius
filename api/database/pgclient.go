package database

import (
	"context"
	_ "github.com/lib/pq"
	"skingenius/database/model"
)

const dbname = "skingenius_new"

type Connector interface {
	//IngredientBySkinType(context.Context, string) ([]string, error)
	//IngredientBySkinSensitivity(context.Context, string) ([]string, error)
	//IngredientByAcne(context.Context, string) ([]string, error)
	//IngredientByPreferences(context.Context, []string) ([]string, error)
	//IngredientByAllergens(context.Context, []string) ([]string, error)
	//IngredientBySkinConcern(context.Context, string) ([]string, error)
	//IngredientByAge(context.Context, string) ([]string, error)
	//IngredientByProductBenefit(context.Context, string) ([]string, error)
	//
	//FilerHardParameters(context.Context, string, string, string) ([]string, error)

	GetAllSkintypes(context.Context) ([]model.Skintype, error)
	GetAllSkinsensetivity(context.Context) ([]model.Skinsensitivity, error)
	GetAllAcneBreakouts(context.Context) ([]model.Acnebreakout, error)
	GetAllPreferences(context.Context) ([]model.Preference, error)
	GetAllAllergies(context.Context) ([]model.Allergy, error)
	GetAllSkinconcerns(context.Context) ([]model.Skinconcern, error)
	GetAllAge(context.Context) ([]model.Age, error)
	GetAllBenefits(context.Context) ([]model.Benefit, error)
}

//type PgConnector struct {
//	db *sql.DB
//}
//
//func NewClient(host string, port int, user, password string) (Connector, error) {
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//
//	logger.New().Info(context.Background(), fmt.Sprintf("DB config: %v", psqlInfo))
//
//	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, dbname))
//	if err != nil {
//		logger.New().Error(context.Background(), "Error connecting to the database: ", err)
//	}
//	//defer db.Close()
//
//	err = db.Ping()
//	if err != nil {
//		return nil, fmt.Errorf("Error pinging database: ", err)
//	}
//	logger.New().Info(context.Background(), "Connected to the database!")
//
//	return &PgConnector{db: db}, nil
//}
//
//func (pg *PgConnector) FilerHardParameters(ctx context.Context, sensitivity, preferences, allergies string) ([]string, error) {
//	//query := fmt.Sprintf("SELECT ingredient FROM ingredient_skin_type WHERE %s = 'Yes'", skinType)
//	//logger.New().Info(ctx, fmt.Sprintf("FilerHardParameters query: %s", query))
//	return nil, nil
//}
//
//func (pg *PgConnector) IngredientByProductBenefit(ctx context.Context, benefits string) ([]string, error) {
//	val, ok := model.SkinBenefitsToDbValue[benefits]
//	if !ok {
//		logger.New().Error(context.Background(), fmt.Sprintf("failed to find db value for IngredientByProductBenefit value:'%s'", benefits))
//		return nil, errors.New(fmt.Sprintf("failed to find db value for IngredientByProductBenefit value:'%s'", benefits))
//	}
//
//	query := fmt.Sprintf("SELECT ingredient FROM benefit WHERE ")
//	var conditions []string
//	for _, p := range val {
//		conditions = append(conditions, fmt.Sprintf("%s=true", p))
//	}
//	query = query + strings.Join(conditions, " AND ")
//	logger.New().Info(ctx, fmt.Sprintf("IngredientByProductBenefit query: %s", query))
//
//	var res string
//	var ingredientsList []string
//
//	rows, err := pg.db.Query(query)
//	defer rows.Close()
//	if err != nil {
//		logger.New().Error(context.Background(), fmt.Sprintf("IngredientByProductBenefit err: %v", err))
//		return nil, err
//	}
//	for rows.Next() {
//		rows.Scan(&res)
//		ingredientsList = append(ingredientsList, res)
//	}
//
//	return ingredientsList, nil
//
//}
//
//func (pg *PgConnector) IngredientByAge(ctx context.Context, age string) ([]string, error) {
//	val, ok := model.AgeToDbValue[age]
//	if !ok {
//		logger.New().Error(context.Background(), fmt.Sprintf("failed to find db value for IngredientByAge value:'%s'", val))
//		return nil, errors.New(fmt.Sprintf("failed to find db value for IngredientByAge value:'%s'", val))
//	}
//
//	query := fmt.Sprintf("SELECT ingredient FROM age_range WHERE %s = true", val)
//	logger.New().Info(ctx, fmt.Sprintf("IngredientByAge query: %s", query))
//
//	var res string
//	var ingredientsList []string
//
//	rows, err := pg.db.Query(query)
//	defer rows.Close()
//	if err != nil {
//		logger.New().Error(context.Background(), fmt.Sprintf("IngredientByAge err: %v", err))
//		return nil, err
//	}
//	for rows.Next() {
//		rows.Scan(&res)
//		ingredientsList = append(ingredientsList, res)
//	}
//
//	return ingredientsList, nil
//}
//
//func (pg *PgConnector) IngredientBySkinConcern(ctx context.Context, concerns string) ([]string, error) {
//	val, ok := model.SkinConcernToDbValue[concerns]
//	if !ok {
//		logger.New().Error(context.Background(), fmt.Sprintf("failed to find db value for IngredientBySkinConcern value:'%s'", concerns))
//		return nil, errors.New(fmt.Sprintf("failed to find db value for IngredientBySkinConcern value:'%s'", concerns))
//	}
//
//	query := fmt.Sprintf("SELECT ingredient FROM skin_concern WHERE ")
//	var conditions []string
//	for _, p := range val {
//		conditions = append(conditions, fmt.Sprintf("%s=true", p))
//	}
//	query = query + strings.Join(conditions, " AND ")
//	logger.New().Info(ctx, fmt.Sprintf("IngredientBySkinConcern query: %s", query))
//
//	var res string
//	var ingredientsList []string
//
//	rows, err := pg.db.Query(query)
//	defer rows.Close()
//	if err != nil {
//		logger.New().Error(context.Background(), fmt.Sprintf("IngredientByAllergens err: %v", err))
//		return nil, err
//	}
//	for rows.Next() {
//		rows.Scan(&res)
//		ingredientsList = append(ingredientsList, res)
//	}
//
//	return ingredientsList, nil
//}
//
//func (pg *PgConnector) IngredientByAllergens(ctx context.Context, allergens []string) ([]string, error) {
//	if len(allergens) == 0 {
//		logger.New().Info(ctx, fmt.Sprintf("IngredientByAllergens - no allergens selected"))
//		return []string{}, nil
//	}
//
//	query := fmt.Sprintf("SELECT ingredient FROM ingredient_allergen WHERE ")
//	var conditions []string
//	for _, p := range allergens {
//		conditions = append(conditions, fmt.Sprintf("%s=true", p))
//	}
//	query = query + strings.Join(conditions, " AND ")
//	logger.New().Info(ctx, fmt.Sprintf("IngredientByAllergens query: %s", query))
//
//	var res string
//	var ingredientsList []string
//
//	rows, err := pg.db.Query(query)
//	defer rows.Close()
//	if err != nil {
//		logger.New().Error(context.Background(), fmt.Sprintf("IngredientByAllergens err: %v", err))
//		return nil, err
//	}
//	for rows.Next() {
//		rows.Scan(&res)
//		ingredientsList = append(ingredientsList, res)
//	}
//
//	return ingredientsList, nil
//}
//
//func (pg *PgConnector) IngredientByPreferences(ctx context.Context, pref []string) ([]string, error) {
//	if len(pref) == 0 {
//		logger.New().Info(ctx, fmt.Sprintf("IngredientByPreferences - no preferences selected"))
//		return []string{}, nil
//	}
//
//	query := fmt.Sprintf("SELECT ingredient FROM ingredient_preference WHERE ")
//	var conditions []string
//	for _, p := range pref {
//		conditions = append(conditions, fmt.Sprintf("%s=true", p))
//	}
//	query = query + strings.Join(conditions, " AND ")
//	logger.New().Info(ctx, fmt.Sprintf("IngredientByPreferences query: %s", query))
//
//	var res string
//	var ingredientsList []string
//
//	rows, err := pg.db.Query(query)
//	defer rows.Close()
//	if err != nil {
//		logger.New().Error(context.Background(), fmt.Sprintf("IngredientByPreferences err: %v", err))
//		return nil, err
//	}
//	for rows.Next() {
//		rows.Scan(&res)
//		ingredientsList = append(ingredientsList, res)
//	}
//
//	return ingredientsList, nil
//}
//
//func (pg *PgConnector) IngredientByAcne(ctx context.Context, acne string) ([]string, error) {
//	val, ok := model.SkinAcneToDbValue[acne]
//	if !ok {
//		logger.New().Error(context.Background(), fmt.Sprintf("failed to find db value for IngredientByAcne value:'%s'", acne))
//		return nil, errors.New(fmt.Sprintf("failed to find db value for IngredientByAcne value:'%s'", acne))
//	}
//
//	query := fmt.Sprintf("SELECT ingredient FROM ingredient_skin_type WHERE acne_prone = '%s'", val)
//	logger.New().Info(ctx, fmt.Sprintf("IngredientByAcne query: %s", query))
//
//	var res string
//	var ingredientsList []string
//
//	rows, err := pg.db.Query(query)
//	defer rows.Close()
//	if err != nil {
//		logger.New().Error(context.Background(), fmt.Sprintf("IngredientByAcne err: %v", err))
//		return nil, err
//	}
//	for rows.Next() {
//		rows.Scan(&res)
//		ingredientsList = append(ingredientsList, res)
//	}
//
//	return ingredientsList, nil
//}
//
//func (pg *PgConnector) IngredientBySkinSensitivity(ctx context.Context, sensitivity string) ([]string, error) {
//	val, ok := model.SkinSensitivityToDbValue[sensitivity]
//	if !ok {
//		logger.New().Error(context.Background(), fmt.Sprintf("failed to find db value for Skin sensitivity value:'%s'", sensitivity))
//		return nil, errors.New(fmt.Sprintf("failed to find db value for Skin sensitivity value:'%s'", sensitivity))
//	}
//
//	query := fmt.Sprintf("SELECT ingredient FROM ingredient_skin_type WHERE sensitive = '%s'", val)
//	logger.New().Info(ctx, fmt.Sprintf("IngredientBySkinSensitivity query: %s", query))
//
//	var res string
//	var ingredientsList []string
//
//	rows, err := pg.db.Query(query)
//	defer rows.Close()
//	if err != nil {
//		logger.New().Error(context.Background(), fmt.Sprintf("IngredientBySkinSensitivity err: %v", err))
//		return nil, err
//	}
//	for rows.Next() {
//		rows.Scan(&res)
//		ingredientsList = append(ingredientsList, res)
//	}
//
//	return ingredientsList, nil
//}
//
//func (pg *PgConnector) IngredientBySkinType(ctx context.Context, skinType string) ([]string, error) {
//	query := fmt.Sprintf("SELECT ingredient FROM ingredient_skin_type WHERE %s = 'Yes'", skinType)
//	logger.New().Info(ctx, fmt.Sprintf("IngredientBySkinType query: %s", query))
//
//	var res string
//	var ingredientsList []string
//
//	rows, err := pg.db.Query(query)
//	defer rows.Close()
//	if err != nil {
//		logger.New().Error(context.Background(), fmt.Sprintf("IngredientBySkinType err: %v", err))
//		return nil, err
//	}
//	for rows.Next() {
//		rows.Scan(&res)
//		ingredientsList = append(ingredientsList, res)
//	}
//
//	return ingredientsList, nil
//}
