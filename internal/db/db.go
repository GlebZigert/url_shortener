package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var table string = `
CREATE TABLE IF NOT EXISTS strazh (
	id          INTEGER PRIMARY KEY,
	origin        TEXT,
	short       TEXT
)`

var db *sql.DB

func Init() error {

	//host=localhost user=shortener password=userpassword dbname=shortener sslmode=disable

	var err error
	fmt.Println("config.DatabaseDSN: ", config.DatabaseDSN)
	db, err = sql.Open("pgx", config.DatabaseDSN)

	if err != nil {
		fmt.Println("err1: ", err)
		return err
	}

	res, err := db.Exec(table)

	if err != nil {
		fmt.Println("err2: ", err)
		return err
	}
	fmt.Println("res: ", res)
	return err
}

func Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := db.PingContext(ctx)
	fmt.Println("ping:  ", err)
	return err
}

func Insert(ctx context.Context, short, origin string, id int) error {
	// здесь будем вставлять записи в базу данных
	// ...

	_, err := db.ExecContext(ctx,
		"INSERT INTO strazh (id, origin, short)"+
			" VALUES(?,?,?,?,?)", id, origin, short)
	if err != nil {
		return err
	}

	return nil
}
