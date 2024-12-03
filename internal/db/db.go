package db

import (
	"context"
	"database/sql"
	"errors"
	"log"
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

// DBer is struct for db
type DBer struct {
	db *sql.DB
}

var ptr *DBer

// Get DBer struct
func Get(db *sql.DB) *DBer {

	if db == nil {
		log.Println("db 3 err: mo db")
		return nil
	}
	ptr = &DBer{db}

	return ptr
}

var db *sql.DB

// запуск бд
func Init(DatabaseDSN string) *sql.DB {

	var err error
	db, err = sql.Open("pgx", DatabaseDSN)

	if err != nil {
		log.Println("db 1 err: ", err.Error())
		return nil
	}

	_, err = db.Exec(table)

	if err != nil {
		log.Println("db 2 err: ", err.Error())
		return nil
	}
	return db
}

// пинг дб
func (dber *DBer) Ping(ctx context.Context) error {

	if dber.db == nil {
		return errors.New("no db")
	}
	err := dber.db.PingContext(ctx)
	return err
}

// вставка в  бд
func (dber *DBer) Insert(ctx context.Context, short, origin string, UUID int) error {

	stmt, err := dber.db.PrepareContext(ctx, "INSERT INTO strazh (uid,origin, short) VALUES ($1, $2, $3)")

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, UUID, origin, short)
	if err != nil {
		return err
	}

	return nil
}

// Load from db
func (dber *DBer) DBLoad() (*sql.Rows, error) {
	return dber.db.Query("SELECT * FROM strazh")
}

// Load Users from db
func (dber *DBer) DBLoadUsers() (*sql.Rows, error) {
	return dber.db.Query("SELECT uid FROM users")
}

// удалить из бд
func (dber *DBer) Delete(short interface{}) error {

	switch short := short.(type) {
	case string:
		_, err := dber.db.Exec("UPDATE strazh SET deleted = true WHERE short = $1", short)
		return err

	case []string:
		rows, err := dber.db.Query("UPDATE strazh SET deleted = true WHERE id = ($1)", strings.Join(short, ","))
		if err != nil {
			return err
		}
		if rows.Err() != nil {
			return rows.Err()
		}
		return err
	default:
		return &packerr.WrongType
	}

}

// DBInsertUser insert user to db
func (dber *DBer) DBInsertUser(ctx context.Context, uid int) error {

	stmt, err := dber.db.PrepareContext(ctx, "INSERT INTO users (uid) VALUES ($1)")

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, uid)
	if err != nil {
		return err
	}

	return nil
}
