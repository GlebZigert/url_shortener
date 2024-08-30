package storager

import (
	"context"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/db"
)

type DBStorager struct {
}

func (one *DBStorager) Load(shorten *[]*Shorten) error {
	return nil
}

func (one *DBStorager) StorageWrite(short, origin string, id, UUID int) error {

	return db.Insert(context.Background(), short, origin, UUID)

}

func (one *DBStorager) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return db.Ping(ctx)

}
