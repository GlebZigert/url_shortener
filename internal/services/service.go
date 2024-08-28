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

var (
	mapa map[string]string
	id   int
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
	mapa = make(map[string]string)

	_ = storager.Load(&mapa)

}

func Short(oririn string) (string, error) {

	v, ok := mapa[oririn]

	if ok {
		return v, &ErrConflict409{config.Conflict409}
	}

	short := generateRandomString(8)

	mapa[oririn] = short
	storager.StorageWrite(short, oririn, len(mapa))

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

	for k, v := range mapa {
		if v == short {

			return k, nil
		}
	}

	return "", errors.New("отстуствует")

}

func GetAll() map[string]string {

	return mapa
}
