package services

import (
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
