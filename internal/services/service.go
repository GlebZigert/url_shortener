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

var shorten []storager.Shorten

var (
	id int
)

var shortuser map[string]*list.List

type ErrConflict409 struct {
	s string
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
	shorten = []storager.Shorten{}

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
	AddUserToShort(int(uuid), short)
	sh := storager.Shorten{0, 0, short, oririn}
	shorten = append(shorten, storager.Shorten{0, 0, short, oririn})

	storager.StorageWrite(sh)

	return short, nil
}

func AddUserToShort(user int, short string) {
	l, ok := shortuser[short]
	if !ok {
		l = list.New()
		shortuser[short] = l

	}
	l.PushFront(user)
	for k, v := range shortuser {
		fmt.Println(k)
		for e := v.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}

	}
}

func CheckUserForShort(user int, short string) bool {
	l, ok := shortuser[short]
	if !ok {
		fmt.Println("-1")
		return false

	}
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == user {
			fmt.Println("+")
			return true
		}
	}
	fmt.Println("-2")
	return false
}

func Origin(short string) (string, error) {

	for _, sh := range shorten {
		if sh.ShortURL == short {

			return sh.OriginalURL, nil
		}
	}

	return "", errors.New("отстуствует")

}

func GetAll() *[]storager.Shorten {

	return &shorten
}
