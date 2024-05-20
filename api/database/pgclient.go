package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const dbname = "skingenius"

type Connector interface {
	FindMatch()
}

type PgConnector struct {
	db *sql.DB
}

func (c *PgConnector) FindMatch() {

}

func NewClient(host string, port int, user, password string) (Connector, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println(fmt.Sprintf("DB config: %v", psqlInfo))
	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, dbname))
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error pinging database: ", err)
	}
	fmt.Println("Connected to the database!")

	return &PgConnector{db: db}, nil
}
