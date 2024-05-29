package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"skingenius/logger"
)

const dbname = "skingenius"

type Connector interface {
	FindMatch()
	IngredientBySkinType(string) ([]string, error)
}

type PgConnector struct {
	db *sql.DB
}

func (c *PgConnector) FindMatch() {

}

func NewClient(host string, port int, user, password string) (Connector, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	logger.New().Info(context.Background(), fmt.Sprintf("DB config: %v", psqlInfo))

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, dbname))
	if err != nil {
		logger.New().Error(context.Background(), "Error connecting to the database: ", err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error pinging database: ", err)
	}
	logger.New().Info(context.Background(), "Connected to the database!")

	return &PgConnector{db: db}, nil
}

func (pg *PgConnector) IngredientBySkinType(skinType string) ([]string, error) {
	rows, err := pg.db.Query("SELECT ingredient FROM ingredient_skin_type WHERE $1 = Yes", skinType)

	var res string
	var ingredientsList []string

	defer rows.Close()
	if err != nil {
		logger.New().Error(context.Background(), fmt.Sprintf("IngredientBySkinType err: %v", err))
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&res)
		ingredientsList = append(ingredientsList, res)
	}

	return ingredientsList, nil
}
