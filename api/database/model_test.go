package database

import (
	"context"
	"fmt"
	"os"
	"skingenius/config"
	"testing"
)

func Test_FindTop3ByIds(t *testing.T) {
	db, err := NewGormClient(config.RemoteHost, config.Port, config.User, config.Password, false)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	ps, err := db.FindProductsByIds(context.Background(), []int32{100, 89, 85})

	fmt.Println(err)
	fmt.Println(fmt.Sprintf("0: %d", len(ps[0].Ingredients)))
}
