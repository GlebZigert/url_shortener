package storager

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/GlebZigert/url_shortener.git/internal/config"
)

var id int

type FileStorager struct {
}

func (one *FileStorager) Init() error {

	file, err := os.OpenFile(config.FileStoragePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	return err
}

func (one *FileStorager) Load(shorten *[]*Shorten) error {

	file, err := os.OpenFile(config.FileStoragePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var data []byte
	err = nil
	for err == nil {
		data, err = reader.ReadBytes('\n')
		if err != nil {
			return err
		}

		var shorten Shorten
		err = json.Unmarshal(data, &shorten)
		if err != nil {
			return err
		}

	}

	return nil
}

func (one *FileStorager) Delete(str string) {

}

func (one *FileStorager) StorageWrite(short, origin string, UUID int) error {

	file, err := os.OpenFile(config.FileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	shorten := Shorten{id, UUID, short, origin, false}

	data, err := json.Marshal(&shorten)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	id++
	return nil
}
