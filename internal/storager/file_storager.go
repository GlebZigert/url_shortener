package storager

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
)

var id int

type filestoreConfig interface {
	GetFileStoragePath() string
}

// хранение в файле
type FileStorager struct {
	cfg       filestoreConfig
	getreader GetFileReader
}

// GetFileReader to get reader
type GetFileReader interface {
	GetReader(path string) (*bufio.Reader, error)
}

// конструктор
func NewFileStorager(cfg filestoreConfig, getreader GetFileReader) (*FileStorager, error) {

	store := &FileStorager{cfg, getreader}
	file, err := os.OpenFile(cfg.GetFileStoragePath(), os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return store, err
	}
	err = file.Close()
	if err != nil {
		return store, err
	}
	return store, err
}

// загрузить из файла
func (one *FileStorager) Load(shorten *[]*Shorten) (err error) {

	reader, errCfgFile := one.getreader.GetReader(one.cfg.GetFileStoragePath())
	if errCfgFile != nil {
		return errCfgFile
	}

	var data []byte
	err = nil
	for err == nil {
		data, err = reader.ReadBytes('\n')
		if err != nil {
			return
		}

		var shorten Shorten
		err = json.Unmarshal(data, &shorten)
		if err != nil {
			return
		}

	}

	return
}

// удалить
func (one *FileStorager) Delete(short interface{}) error {
	return nil
}

// записать в файл
func (one *FileStorager) StorageWrite(short, origin string, UUID int) error {

	file, err := os.OpenFile(one.cfg.GetFileStoragePath(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer packerr.AddCloseErrToErr(&err, file)

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
