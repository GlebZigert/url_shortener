package services

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/config"
)

type Shorten struct {
	ID          int
	ShortURL    string
	OriginalURL string
}

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
	Load()

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

	err := StorageWrite(short, oririn)
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

func Load() error {
	fmt.Println("Load")

	file, err := os.OpenFile(config.FileStoragePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var data []byte
	err = nil
	for err == nil {
		data, err = reader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err.Error())

			return err
		}

		fmt.Println("storage data:", string(data))

		var shorten Shorten
		err = json.Unmarshal(data, &shorten)
		if err != nil {
			fmt.Println(err.Error())

			return err
		}

		mapa[shorten.OriginalURL] = shorten.ShortURL
		id = shorten.ID + 1

	}

	fmt.Println("id:", id)

	return nil
}

func StorageWrite(short, origin string) error {
	fmt.Println("write ", origin, " ", short, " to ", config.FileStoragePath)

	file, err := os.OpenFile(config.FileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	shorten := Shorten{id, short, origin}

	data, err := json.Marshal(&shorten)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	data = append(data, '\n')
	fmt.Println("data: ", len(data), string(data))
	// записываем событие в буфер
	nn, err := writer.Write(data)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Записалось :", nn)
	id++
	return nil
}
