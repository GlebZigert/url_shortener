package db

import (
	"context"
	"database/sql"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var table string = `
CREATE TABLE IF NOT EXISTS urls (
	id          SERIAL PRIMARY KEY,
  UUID          string,  
  ShortURL      string,  
  OriginslURL   string,  
  DeletedFlag   bool    
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

func Ping(ctx context.Context) error {

	err := db.PingContext(ctx)
	return err
}

func Insert(ctx context.Context, short, origin string, UUID int) error {

	_, err := db.ExecContext(ctx, "insert into urls (UUID,OriginslURL, ShortURL) values ($1, $2, $3)", UUID, origin, short)
	if err != nil {
		return err
	}

	return nil
}
