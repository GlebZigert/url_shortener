package storager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/GlebZigert/url_shortener.git/internal/config"
)

type Shorten struct {
	ID          int
	ShortURL    string
	OriginalURL string
}

func Load(mapa *map[string]string) error {
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

		(*mapa)[shorten.OriginalURL] = shorten.ShortURL
		//	id = shorten.ID + 1

	}

	//fmt.Println("id:", id)

	return nil
}

func StorageWrite(short, origin string, id int) error {
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
