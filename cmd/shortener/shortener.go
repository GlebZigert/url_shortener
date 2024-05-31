package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var mapa map[string]string

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func Map() map[string]string {

	if mapa == nil {
		fmt.Println("обьявляю мапу")
		mapa = make(map[string]string)
	}

	return mapa
}

func Short(url string) string {

	v, ok := Map()[url]
	if ok {
		fmt.Println(url, " уже есть в мапе: ", v)
		return v
	}

	short_url := generateRandomString(8)

	Map()[url] = short_url
	fmt.Println("Для ", url, " сгенерирован шорт: ", short_url)
	return short_url
}

func Origin(short_url string) (string, error) {

	for k, v := range Map() {
		if v == short_url {
			fmt.Println("Для шорта", short_url, " найден url: ", v)
			return k, nil
		}
	}
	fmt.Println("Нет такого шорта как", short_url)
	return "", errors.New("отстуствует")
}
