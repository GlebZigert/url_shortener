package services

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/GlebZigert/url_shortener.git/storager"
)

var (
	mapa map[string]string
	id   int
)

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

func Short(oririn string) string {

	v, ok := mapa[oririn]

	if ok {
		fmt.Println(oririn, " уже есть: ", v)
		return v
	}

	short := generateRandomString(8)

	//Map()[url] = shortURL
	mapa[oririn] = short

	err := storager.StorageWrite(short, oririn, len(mapa))
	if err != nil {
		fmt.Println("запись не прошла: ", err.Error())
	} else {
		fmt.Println("запись должна была пройти успешно")
	}

	fmt.Println("Для ", oririn, " сгенерирован шорт: ", short)
	return short
}

func Origin(short string) (string, error) {

	for k, v := range mapa {
		if v == short {
			fmt.Println("Для шорта", short, " найден url: ", k)
			return k, nil
		}
	}
	fmt.Println("Нет такого шорта как", short)
	return "", errors.New("отстуствует")

}

func GetAll() map[string]string {
	return mapa
}
