package storager

import (
	"context"
	"strings"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

type dbstoreConfig interface {
	GetDatabaseDSN() string
}

// хранение в базе
type DBStorager struct {
	cfg dbstoreConfig
}

// загрузить из базы
func (one *DBStorager) Load(shorten *[]*Shorten) (*[]*Shorten, error) {

	rows, err := db.Get().Query("SELECT * FROM strazh")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var u Shorten
		err = rows.Scan(&u.ID, &u.UUID, &u.ShortURL, &u.OriginalURL, &u.DeletedFlag)
		if err != nil {
			return nil, err
		}
		*shorten = append(*shorten, &u)
	}

	return shorten, nil
}

// записать в базу
func (one *DBStorager) StorageWrite(short, origin string, UUID int) error {

	return db.Insert(context.Background(), short, origin, UUID)

}

// конструктор
func NewDBStorager(cfg dbstoreConfig) (*DBStorager, error) {

	store := &DBStorager{cfg}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := db.Init(cfg.GetDatabaseDSN())
	if err != nil {
		return nil, err
	}

	return store, db.Ping(ctx)

}

// удалить из базы
func (one *DBStorager) Delete(short interface{}) error {

	switch short := short.(type) {
	case string:
		_, err := db.Get().Exec("UPDATE strazh SET deleted = true WHERE short = $1", short)
		return err

	case []string:
		_, err := db.Get().Query("UPDATE strazh SET deleted = true WHERE id = ($1)", strings.Join(short, ","))

		return err
	default:
		return &packerr.WrongType
	}

}
