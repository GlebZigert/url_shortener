package db

import (
	"context"
	"database/sql"
	"strings"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
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

type DBer struct {
	db *sql.DB
}

var dber DBer

func Get() *DBer {
	return &dber
}

// запуск бд
func (dber *DBer) Init(DatabaseDSN string) error {

	var err error
	dber.db, err = sql.Open("pgx", DatabaseDSN)

	if err != nil {
		return err
	}

	_, err = dber.db.Exec(table)

	if err != nil {
		return err
	}
	return err
}

// пинг дб
func (dber *DBer) Ping(ctx context.Context) error {

	err := dber.db.PingContext(ctx)
	return err
}

// вставка в  бд
func (dber *DBer) Insert(ctx context.Context, short, origin string, UUID int) error {

	_, err := dber.db.ExecContext(ctx, "insert into strazh (uid,origin, short) values ($1, $2, $3)", UUID, origin, short)
	if err != nil {
		return err
	}

	return nil
}

func (dber *DBer) DBLoad() (*sql.Rows, error) {
	return dber.db.Query("SELECT * FROM strazh")
}

// удалить из бд
func (dber *DBer) Delete(short interface{}) error {

	switch short := short.(type) {
	case string:
		_, err := dber.db.Exec("UPDATE strazh SET deleted = true WHERE short = $1", short)
		return err

	case []string:
		_, err := dber.db.Query("UPDATE strazh SET deleted = true WHERE id = ($1)", strings.Join(short, ","))

		return err
	default:
		return &packerr.WrongType
	}

}
