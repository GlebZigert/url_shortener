package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func Init() error {

	//host=localhost user=shortener password=userpassword dbname=shortener sslmode=disable

	var err error
	fmt.Println("config.DatabaseDSN: ", config.DatabaseDSN)
	db, err = sql.Open("pgx", config.DatabaseDSN)
	if err != nil {
	}

	fmt.Println(err)
	return err
}

func Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := db.PingContext(ctx)
	fmt.Println("ping:  ", err)
	return err
}
