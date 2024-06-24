package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/container"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
)

type Icontainer interface {
	/*
		Беру из контейнера шорт для оригина:
		origin - оригинальный url
		ok - есть ли уже его шорт в контейнере
		short -значение шорта
	*/
	GetShort(origin string) (short string, ok bool)

	//Кладу в контейнер шорт short для оригина origin
	SetShort(short, origin string)
	/*
		Беру из контейнера оригин по шорту:
		short -значение шорта
		origin - оригинальный url
		err - ошибка

	*/

	GetOrigin(short string) (origin string, err error)
}

var pointer Icontainer

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
	ctn()
}

func ctn() Icontainer {

	if pointer == nil {

		// Задаю тип контейнера
		if config.FileStoragePath != "" {
			pointer = container.New_map_container(storager.New_file_storager(config.FileStoragePath))
		} else {
			pointer = container.New_map_container(nil)
		}

	}

	return pointer
}

func Short(oririn string) string {

	v, ok := ctn().GetShort(oririn)
	if ok {
		fmt.Println(oririn, " уже есть: ", v)
		return v
	}

	shortURL := generateRandomString(8)

	//Map()[url] = shortURL
	ctn().SetShort(shortURL, oririn)

	fmt.Println("Для ", oririn, " сгенерирован шорт: ", shortURL)
	return shortURL
}

func Origin(short string) (string, error) {

	return ctn().GetOrigin(short)

}
