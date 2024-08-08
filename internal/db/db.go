package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func Init() error {

	ps := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		`localhost`, `shortener`, `userpassword`, `shortener`)
	var err error
	db, err = sql.Open("pgx", ps)
	if err != nil {
	}

	fmt.Println(err)
	return err
}

func Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := db.PingContext(ctx)
	fmt.Println("ping: ", err)
	return err
}
