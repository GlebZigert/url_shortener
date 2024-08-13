package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var table string = `
CREATE TABLE IF NOT EXISTS strazh (
	id          SERIAL PRIMARY KEY,
	origin        TEXT,
	short       TEXT
)`

var db *sql.DB

func Init() error {

	var err error
	db, err = sql.Open("pgx", config.DatabaseDSN)

	if err != nil {
		return err
	}

	_, err = db.Exec(table)

	if err != nil {
		return err
	}
	return err
}

func Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := db.PingContext(ctx)
	return err
}

func Insert(ctx context.Context, short, origin string, id int) error {

	_, err := db.ExecContext(ctx, "insert into strazh (origin, short) values ($1, $2)", origin, short)
	if err != nil {
		return err
	}

	return nil
}
