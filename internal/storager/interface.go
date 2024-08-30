package storager

import (
	"github.com/GlebZigert/url_shortener.git/internal/logger"
)

type Shorten struct {
	ID          int
	UUID        int
	ShortURL    string
	OriginalURL string
	DeletedFlag bool
}

type Storager interface {
	Init() error
	Load(*[]Shorten) error
	StorageWrite(short, origin string, id, UUID int) error
}

var store Storager

func Load(shorten *[]Shorten) error {
	return store.Load(shorten)
}

func StorageWrite(sh Shorten) error {
	return store.StorageWrite(sh.ShortURL, sh.OriginalURL, sh.ID, sh.UUID)
}

func Init() {
	var err error
	store = &DBStorager{}
	if err = store.Init(); err == nil {
		logger.Log.Info("DB Storager")
		return
	}
	logger.Log.Error(err.Error())
	store = &FileStorager{}
	if err = store.Init(); err == nil {

		logger.Log.Info("File Storager")
		return
	}
	logger.Log.Error(err.Error())
	logger.Log.Info("No Storager")

	store = &EmptyStorager{}
}
