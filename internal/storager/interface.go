package storager

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
	return store.StorageWrite(short, origin, id)
}

func Init() {

	store = &DBStorager{}
	if err := store.Init(); err == nil {
		return
	}

	store = &FileStorager{}
	if err := store.Init(); err == nil {
		return
	}

	store = &EmptyStorager{}
}
