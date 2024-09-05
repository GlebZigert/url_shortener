package storager

import (
	"github.com/GlebZigert/url_shortener.git/internal/logger"
)

type Shorten struct {
	ID          int    `db:"id"`
	UUID        int    `db:"user_id"`
	ShortURL    string `db:"short_url"`
	OriginalURL string `db:"original_url"`
	DeletedFlag bool   `db:"is_deleted"`
}

type Storager interface {
	Load(*[]*Shorten) error
	StorageWrite(short, origin string, UUID int) error
	Delete([]int)
}

var store Storager

func Load(shorten *[]*Shorten) error {
	return store.Load(shorten)
}

func StorageWrite(sh Shorten) error {
	return store.StorageWrite(sh.ShortURL, sh.OriginalURL, sh.UUID)
}

func Init() {
	var err error
	store, err = NewDBStorager()
	if err == nil {
		logger.Log.Info("DB Storager")
		return
	}
	logger.Log.Error(err.Error())
	store, err = NewFileStorager()
	if err == nil {

		logger.Log.Info("File Storager")
		return
	}
	logger.Log.Error(err.Error())
	logger.Log.Info("No Storager")

	store = &EmptyStorager{}
}
