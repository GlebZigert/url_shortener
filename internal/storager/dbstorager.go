package storager

import (
	"context"
	"fmt"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/db"
)

type DBStorager struct {
}

func (one *DBStorager) Load(shorten *[]*Shorten) error {

	rows, err := db.Get().Query("SELECT * FROM strazh")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var u Shorten
		err = rows.Scan(&u.ID, &u.UUID, &u.ShortURL, &u.OriginalURL, &u.DeletedFlag)
		if err != nil {
			panic(err)
		}
		*shorten = append(*shorten, &u)
	}

	return nil
}

func (one *DBStorager) StorageWrite(short, origin string, UUID int) error {

	return db.Insert(context.Background(), short, origin, UUID)

}

func (one *DBStorager) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return db.Ping(ctx)

}

func (one *DBStorager) Delete(str string) {
	fmt.Println("DBStorager delete ", str)
}
