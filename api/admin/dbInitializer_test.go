package main

import (
	"context"
	"fmt"
	"os"
	"skingenius/config"
	"skingenius/database"
	"testing"
)

func Test_SaveProducts(t *testing.T) {

	dbClient, err := database.NewGormClient(config.LocalHost, config.Port, config.User, config.Password, false, "test")
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	storeProducts(context.Background(), dbClient, "products-to-ingredients.csv")
}
