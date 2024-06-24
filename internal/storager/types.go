package storager

type IStorager interface {
	//Загружает хранилище в контейнер
	Load(Add_to_container func(short, origin string)) error
	StorageWrite(short, origin string) error
}

type Shorten struct {
	Id           int
	Short_url    string
	Original_url string
}
