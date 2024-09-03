package services

import (
	"container/list"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
)

var shorten []*storager.Shorten

var (
	id int
)

var shortuser map[string]*list.List

type ErrConflict409 struct {
	s string
}

type ErrDeleted struct {
	s string
}

func (e *ErrDeleted) Error() string {
	return e.s
}

func (e *ErrConflict409) Error() string {
	return e.s
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	//diff
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func Init() {

	l := list.New()
	fmt.Println(l)

	shortuser = make(map[string]*list.List)
	shorten = []*storager.Shorten{}

	_ = storager.Load(&shorten)

}

func Short(oririn string, uuid int) (string, error) {

	//v, ok := mapa[oririn]
	for _, sh := range shorten {
		if sh.OriginalURL == oririn {

			return sh.ShortURL, &ErrConflict409{config.Conflict409}
		}
	}

	short := generateRandomString(8)
	//AddUserToShort(int(uuid), short)
	sh := storager.Shorten{ID: shorten[len(shorten)-1].ID, UUID: uuid, ShortURL: short, OriginalURL: oririn, DeletedFlag: false}
	shorten = append(shorten, &storager.Shorten{ID: 0, UUID: 0, ShortURL: short, OriginalURL: oririn, DeletedFlag: false})

	storager.StorageWrite(sh)

	return short, nil
}

/*
func Delete(short string, uid int) {
	for _, sh := range shorten {
		if sh.ShortURL == short && sh.UUID == uid {
			fmt.Println("надо удалить ", short)
			sh.DeletedFlag = true
		}
	}
}
*/

func Origin(short string) (string, error) {

	for _, sh := range shorten {
		if sh.ShortURL == short {

			if sh.DeletedFlag {
				str := "шорт " + short + " удален"
				return sh.OriginalURL, &ErrDeleted{str}
			}

			return sh.OriginalURL, nil
		}
	}

	return "", errors.New("отстуствует")

}

func GetAll() *[]*storager.Shorten {

	return &shorten
}
