package storager

type IStorager interface {
	//Загружает хранилище в контейнер
	Load(AddToContainer func(short, origin string)) error
	StorageWrite(short, origin string) error
}

type Shorten struct {
	Id          int
	ShortURL    string
	OriginalURL string
}
