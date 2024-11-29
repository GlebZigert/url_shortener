package storager

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
)

var id int

// FilestoreConfig интерфейс конфига
type FilestoreConfig interface {
	GetFileStoragePath() string
}

// хранение в файле
type FileStorager struct {
	cfg FilestoreConfig
	rw  FileReaderWriter
}

// GetFileReader to get reader
type FileReaderWriter interface {
	Read(path string) ([]byte, error)
	Write(data []byte, filepath string) error
	CheckFile(filepath string) error
}

// конструктор
func NewFileStorager(cfg FilestoreConfig, rw FileReaderWriter) (*FileStorager, error) {

	store := &FileStorager{cfg, rw}
	if rw == nil {
		return nil, errors.New("no rw")
	}
	err := rw.CheckFile(cfg.GetFileStoragePath())
	return store, err
}

// загрузить из файла
func (one *FileStorager) Load(shorten *[]*Shorten) (res *[]*Shorten, users int, err error) {

	var data []byte
	err = nil
	for err == nil {

		data, err = one.rw.Read(one.cfg.GetFileStoragePath())
		if err != nil {
			log.Println(err.Error())
			continue
		}
		log.Println(string(data))

		var short Shorten
		err = json.Unmarshal(data, &short)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		*shorten = append(*shorten, &short)

	}
	//	log.Println("err: ", err.Error())

	res = shorten
	return
}

// удалить
func (one *FileStorager) Delete(short interface{}) error {
	return nil
}

// записать в файл
func (one *FileStorager) StorageWrite(short, origin string, UUID int) error {

	shorten := Shorten{id, UUID, short, origin, false}

	data, err := json.Marshal(&shorten)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	err = one.rw.Write(data, one.cfg.GetFileStoragePath())
	if err != nil {
		return err
	}

	id++
	return nil
}

func (one *FileStorager) WriteUID(uid int) error {
	//пока заглушка
	err := one.rw.Write([]byte("User: "+strconv.Itoa(uid)+"\n"), one.cfg.GetFileStoragePath())

	return err
}
