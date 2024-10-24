package storager

type storeConfig interface {
	GetFileStoragePath() string
}

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

func New(cfg storeConfig) (store Storager) {
	var err error

	store, err = NewDBStorager()
	if err == nil {

		return
	}

	store, err = NewFileStorager(cfg)
	if err == nil {

		return
	}
	store = &EmptyStorager{}

	return
}
