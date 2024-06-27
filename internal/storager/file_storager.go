package storager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type FileStorager struct {
	path string
	id   int
}

func NewFileStorager(path string) *FileStorager {
	fmt.Println("file storager: ", path)
	fs := FileStorager{path, 0}

	//проверка на path

	return &fs
}

func (one *FileStorager) Load(f func(short, origin string)) error {
	fmt.Println("Load")
	if f == nil {
		fmt.Println("определи функцию ")
		return nil
	}

	file, err := os.OpenFile(one.path, os.O_RDONLY|os.O_CREATE, 0666)
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

		f(shorten.ShortURL, shorten.OriginalURL)
		one.id = shorten.ID + 1

	}

	fmt.Println("id:", one.id)

	return nil
}

func (one *FileStorager) StorageWrite(short, origin string) error {
	fmt.Println("write ", origin, " ", short, " to ", one.path)

	file, err := os.OpenFile(one.path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	shorten := Shorten{one.id, short, origin}

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
	one.id++
	return nil
}
