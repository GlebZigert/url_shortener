// Package service is a package that what.
package services

import (
	"container/list"
	"errors"

	"math/rand"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
)

// это массив для хранения сокращенных url
var shorten []*storager.Shorten

// uid послежнего добавленного пользователя - он же и количество пользователей
var uid int

// хранилище
type Storager interface {
	Load(*[]*storager.Shorten) (*[]*storager.Shorten, error)
	StorageWrite(short, origin string, UUID int) error
	Delete(interface{}) error
	WriteUID(int) error
}

// логгер
type Logger interface {
	Info(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
}

// сервис
type Service struct {
	logger Logger
	store  Storager

	shortuser map[string]*list.List
	shorten   []*storager.Shorten
}

// конструткор
func NewService(logger Logger, store Storager) *Service {
	srv := Service{logger, store, make(map[string]*list.List), []*storager.Shorten{}}

	return &srv

}

var (
	id int
)

var shortuser map[string]*list.List

// генератор шорта
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

// сделать шорт на ориджин
func (s *Service) Short(oririn string, uuid int) (string, error) {

	s.logger.Info("Short: ", map[string]interface{}{

		"oririn": oririn,
		"uuid":   uuid,
	})

	//v, ok := mapa[oririn]
	for _, sh := range shorten {
		if sh.OriginalURL == oririn {

			return sh.ShortURL, &packerr.Conflict
		}
	}

	short := generateRandomString(8)
	//AddUserToShort(int(uuid), short)
	id := 0
	if len(shorten) > 0 {
		id = shorten[len(shorten)-1].ID + 1
	}

	err := s.store.StorageWrite(short, oririn, uuid)
	if err != nil {
		return "", err
	}

	sh := storager.Shorten{ID: id, UUID: uuid, ShortURL: short, OriginalURL: oririn, DeletedFlag: false}
	shorten = append(shorten, &sh)

	s.logger.Info("Сделан шорт: ", map[string]interface{}{

		"oririn": oririn,
		"short":  short,
		"uuid":   uuid,
	})
	return short, nil
}

// получить ориджин по шорту
func (s *Service) Origin(short string) (string, error) {

	for _, sh := range shorten {
		if sh.ShortURL == short {

			if sh.DeletedFlag {
				str := "шорт " + short + " удален"
				return sh.OriginalURL, &packerr.ErrDeleted{S: str}
			}

			return sh.OriginalURL, nil
		}
	}

	return "", errors.New("отстуствует")

}

// получить все шорты
func (s *Service) GetAll() *[]*storager.Shorten {

	return &shorten
}

func (s *Service) GetUsersCount() int {
	return uid
}
func (s *Service) GetUrlsCount() int {
	return len(shorten)
}

// ошибка в случае если не прошла запись свежего uid в бд или файл
type ErrCreateNextID error

var errCreateNextUID ErrCreateNextID

func (s *Service) CreateNextUID() (int, error) {
	next := uid + 1

	//если запись не прошла - возвращаем ошибку об этом
	if err := s.store.WriteUID(next); err != nil {
		return -1, errCreateNextUID
	}
	uid = next

	return uid, nil
}
