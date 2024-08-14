package storager

import (
	"context"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/db"
)

type DBStorager struct {
}

func (one *DBStorager) Load(mapa *map[string]string) error {
	return nil
}

func (one *DBStorager) StorageWrite(short, origin string, id int) error {

	return db.Insert(context.Background(), short, origin, id)

}

func (one *DBStorager) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return db.Ping(ctx)

}
