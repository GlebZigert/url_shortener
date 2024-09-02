package storager

import (
	"github.com/GlebZigert/url_shortener.git/internal/logger"
)

type Storager interface {
	Init() error
	Load(mapa *map[string]string) error
	StorageWrite(short, origin string, id, UUID int) error
	Delete(string)
}

var store Storager

func Load(mapa *map[string]string) error {
	return store.Load(mapa)
}

func StorageWrite(short, origin string, id, UUID int) error {
	return store.StorageWrite(short, origin, id, UUID)
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
