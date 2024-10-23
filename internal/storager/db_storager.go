package storager

import (
	"context"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/db"
)

type DBStorager struct {
}

func (one *DBStorager) Load(shorten *[]*Shorten) error {

	rows, err := db.Get().Query("SELECT * FROM strazh")

	if err != nil {
		return err
	}

	for rows.Next() {
		var u Shorten
		err = rows.Scan(&u.ID, &u.UUID, &u.ShortURL, &u.OriginalURL, &u.DeletedFlag)
		if err != nil {
			return err
		}
		*shorten = append(*shorten, &u)
	}

	return nil
}

func (one *DBStorager) StorageWrite(short, origin string, UUID int) error {

	return db.Insert(context.Background(), short, origin, UUID)

}

func NewDBStorager() (*DBStorager, error) {

	store := &DBStorager{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return store, db.Ping(ctx)

}

func (one *DBStorager) Delete(listID []int) {

}