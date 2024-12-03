package storager

import (
	"encoding/json"
	"errors"
	"log"
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
func (one *FileStorager) Load(shorten *[]*Shorten) (res *[]*Shorten, err error) {

	err = nil
	var flag bool
	for err == nil {

		flag = true
		data, err := one.rw.Read(one.cfg.GetFileStoragePath())
		if err != nil {
			flag = false
			log.Println(err.Error())
			break
		}

		log.Println(string(data))

		var short Shorten
		err = json.Unmarshal(data, &short)
		if err != nil || short.OriginalURL == "" {

			flag = false

		}

		if flag {
			log.Println(short.ID, short.UUID, short.OriginalURL, short.ShortURL)
			log.Println("...append... ", string(data))
			*shorten = append(*shorten, &short)
		}

	}
	//	log.Println("err: ", err.Error())

	res = shorten
	err = nil
	return
}

// загрузить из файла
func (one *FileStorager) LoadUsers(users *[]*User) (res *[]*User, err error) {

	err = nil

	for err == nil {

		data, err := one.rw.Read(one.cfg.GetFileStoragePath())
		if err != nil {

			log.Println(err.Error())
			break
		}

		log.Println(string(data))

		var user User
		err = json.Unmarshal(data, &user)

		if err != nil {
			log.Println(err.Error())
			continue
		}

		if user.Uid == 0 {

			continue
		}

		log.Println("...users append... ", string(data))
		*users = append(*users, &user)

	}
	//	log.Println("err: ", err.Error())

	res = users
	err = nil
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

	var id User
	id.Uid = uid

	data, err := json.Marshal(id)

	if err != nil {
		return err
	}

	err = one.rw.Write(data, one.cfg.GetFileStoragePath())

	return err
}
