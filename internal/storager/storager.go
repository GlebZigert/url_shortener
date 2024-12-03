package storager

// StoreConfig интерфейс конфига
type StoreConfig interface {
	GetFileStoragePath() string
	GetDatabaseDSN() string
}

// данные о шорте
type Shorten struct {
	ID          int    `db:"id"`
	UUID        int    `db:"user_id" json:"uuid"`
	ShortURL    string `db:"short_url" json:"short_url"`
	OriginalURL string `db:"original_url" json:"original_url"`
	DeletedFlag bool   `db:"is_deleted"`
}

// User struct for user
type User struct {
	Uid int `json:"uid"`
}

// операции с хранилищем
type Storager interface {
	Load(*[]*Shorten) (*[]*Shorten, error)
	LoadUsers(*[]*User) (*[]*User, error)
	StorageWrite(short, origin string, UUID int) error
	Delete(interface{}) error
	WriteUID(int) error
}

// конструктор
func New(cfg StoreConfig, rw FileReaderWriter, db DB) (store Storager) {
	var err error

	store, err = NewDBStorager(cfg, db)
	if err == nil {

		return
	}

	store, err = NewFileStorager(cfg, rw)
	if err == nil {

		return
	}
	store = &EmptyStorager{}

	return
}
