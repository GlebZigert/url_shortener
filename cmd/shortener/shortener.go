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

	shortURL := generateRandomString(8)

	Map()[url] = shortURL
	fmt.Println("Для ", url, " сгенерирован шорт: ", shortURL)
	return shortURL
}

func Origin(shortURL string) (string, error) {

	for k, v := range Map() {
		if v == shortURL {
			fmt.Println("Для шорта", shortURL, " найден url: ", v)
			return k, nil
		}
	}
	fmt.Println("Нет такого шорта как", shortURL)
	return "", errors.New("отстуствует")
}
