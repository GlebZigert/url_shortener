package storager

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"reflect"
	"time"
)

// Interface for DB
type DB interface {
	Ping(ctx context.Context) error
	Insert(ctx context.Context, short, origin string, UUID int) error
	DBLoad() (*sql.Rows, error)
	DBLoadUsers() (*sql.Rows, error)
	DBInsertUser(context.Context, int) error
	Delete(short interface{}) error
}

type dbstoreConfig interface {
	GetDatabaseDSN() string
}

// хранение в базе
type DBStorager struct {
	cfg dbstoreConfig
	DB
}

// загрузить из базы
func (one *DBStorager) Load(shorten *[]*Shorten) (*[]*Shorten, error) {

	rows, err := one.DBLoad()

	if err != nil {
		log.Println("db load err: ", err.Error())
		return nil, err
	}

	if rows.Err() != nil {
		log.Println("db load err: ", rows.Err().Error())
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

// загрузить из базы
func (one *DBStorager) LoadUsers(users *[]*User) (*[]*User, error) {

	rows, err := one.DBLoadUsers()

	if err != nil {
		log.Println("db load err: ", err.Error())
		return nil, err
	}

	if rows.Err() != nil {
		log.Println("db load err: ", rows.Err().Error())
		return nil, err
	}

	for rows.Next() {
		var u User
		err = rows.Scan(&u.Uid)
		if err != nil {
			return nil, err
		}
		*users = append(*users, &u)
	}

	return users, nil
}

// записать в базу
func (one *DBStorager) StorageWrite(short, origin string, UUID int) error {

	return one.Insert(context.Background(), short, origin, UUID)

}

// конструктор
func NewDBStorager(cfg dbstoreConfig, db DB) (*DBStorager, error) {

	if db == nil || reflect.ValueOf(db).IsNil() {
		log.Println("no db")
		return nil, errors.New("no db")
	}

	store := &DBStorager{cfg, db}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	return store, db.Ping(ctx)

}

// запись пользователя
func (one *DBStorager) WriteUID(uid int) error {
	//пока заглушка
	return one.DBInsertUser(context.Background(), uid)
}
