package storager

import "github.com/GlebZigert/url_shortener.git/internal/config"

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

func New(cfg *config.Config) (store Storager) {
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
