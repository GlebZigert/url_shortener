package storager

import (
	"fmt"

	"github.com/GlebZigert/url_shortener.git/internal/config"
)

type Storager interface {
	Init() error
	Load(mapa *map[string]string) error
	StorageWrite(short, origin string, id int) error
}

var store Storager

func Load(mapa *map[string]string) error {
	return store.Load(mapa)
}

func StorageWrite(short, origin string, id int) error {
	fmt.Println("StorageWrite")
	return store.StorageWrite(short, origin, id)
}

func Init() {

	store = &DBStorager{}
	if err := store.Init(); err == nil {
		fmt.Println("Хранилище в БД ", config.DatabaseDSN)
		return
	}

	store = &FileStorager{}
	if err := store.Init(); err == nil {
		fmt.Println("Хранилище в файле ", config.FileStoragePath)
		return
	}
	fmt.Println("Без внешнего хранилища")
	store = &EmptyStorager{}
}
