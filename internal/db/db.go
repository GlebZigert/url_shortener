package db

import (
	"context"
	"database/sql"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var table string = `
CREATE TABLE IF NOT EXISTS strazh (
	id          SERIAL PRIMARY KEY,
		uid 		INT ,
	origin        TEXT,
	short       TEXT,
deleted		BOOLEAN
)`

/*
var table string = `
CREATE TABLE IF NOT EXISTS strazh (
	id          SERIAL PRIMARY KEY,
	uid 		INT ,
	origin        TEXT,
	short       TEXT,
	deleted		BOOLEAN
)`
*/

var db *sql.DB

func Get() *sql.DB {
	return db
}

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

func Ping(ctx context.Context) error {

	err := db.PingContext(ctx)
	return err
}

func Insert(ctx context.Context, short, origin string, UUID int) error {

	_, err := db.ExecContext(ctx, "insert into strazh (uid,origin, short) values ($1, $2, $3)", UUID, origin, short)
	if err != nil {
		return err
	}

	return nil
}

func Del(ctx context.Context, short string) error {

	_, err := db.ExecContext(ctx, "UPDATE strazh SET deleted = true WHERE short = $1;", short)
	if err != nil {
		return err
	}

	return nil
}
