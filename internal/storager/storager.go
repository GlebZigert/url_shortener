package storager

type StoreConfig interface {
	GetFileStoragePath() string
	GetDatabaseDSN() string
}

// данные о шорте
type Shorten struct {
	ID          int    `db:"id"`
	UUID        int    `db:"user_id"`
	ShortURL    string `db:"short_url"`
	OriginalURL string `db:"original_url"`
	DeletedFlag bool   `db:"is_deleted"`
}

// операции с хранилищем
type Storager interface {
	Load(*[]*Shorten) error
	StorageWrite(short, origin string, UUID int) error
	Delete(interface{}) error
}

// конструктор
func New(cfg StoreConfig, rw FileReaderWriter) (store Storager) {
	var err error

	store, err = NewDBStorager(cfg)
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
