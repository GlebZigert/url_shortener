package storager

import (
	"github.com/GlebZigert/url_shortener.git/internal/db"
)

type DBStorager struct {
}

func (one *DBStorager) Load(mapa *map[string]string) error {
	return nil
}

func (one *DBStorager) StorageWrite(short, origin string, id int) error {
	return nil
}

func (one *DBStorager) Init() error {

	return db.Ping()

}
