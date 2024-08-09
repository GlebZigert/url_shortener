package storager

import (
	"github.com/GlebZigert/url_shortener.git/internal/db"
)

type DbStorager struct {
}

func (one *DbStorager) Load(mapa *map[string]string) error {
	return nil
}

func (one *DbStorager) StorageWrite(short, origin string, id int) error {
	return nil
}

func (one *DbStorager) Init() error {

	return db.Ping()

}
