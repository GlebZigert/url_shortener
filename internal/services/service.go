package services

import (
	"container/list"
	"errors"

	"math/rand"
	"time"

	"github.com/GlebZigert/url_shortener.git/internal/storager"
)

// это массив для хранения сокращенных url
var shorten []*storager.Shorten

type Storager interface {
	Load(*[]*storager.Shorten) error
	StorageWrite(short, origin string, UUID int) error
	Delete([]int)
}

type Logger interface {
	Info(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
}

type Service struct {
	logger Logger
	store  Storager

	shortuser map[string]*list.List
	shorten   []*storager.Shorten
}

/*

func Init() {

	shortuser = make(map[string]*list.List)
	shorten = []*storager.Shorten{}

	_ = storager.Load(&shorten)

}
*/

func NewService(logger Logger, store Storager) *Service {
	srv := Service{logger, store, make(map[string]*list.List), []*storager.Shorten{}}

	return &srv

}

var (
	id int
)

var shortuser map[string]*list.List

type ErrConflict409 struct {
	s string
}

type ErrDeleted struct {
	s string
}

func (e *ErrDeleted) Error() string {
	return e.s
}

func (e *ErrConflict409) Error() string {
	return e.s
}

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

func (s *Service) Short(oririn string, uuid int) (string, error) {

	s.logger.Info("try to login: ", map[string]interface{}{

		"oririn": oririn,
		"uuid":   uuid,
	})

	//v, ok := mapa[oririn]
	for _, sh := range shorten {
		if sh.OriginalURL == oririn {

			return sh.ShortURL, &ErrConflict409{"попытка сократить уже имеющийся в базе URL"}
		}
	}

	short := generateRandomString(8)
	//AddUserToShort(int(uuid), short)
	id := 0
	if len(shorten) > 0 {
		id = shorten[len(shorten)-1].ID + 1
	}

	sh := storager.Shorten{ID: id, UUID: uuid, ShortURL: short, OriginalURL: oririn, DeletedFlag: false}
	shorten = append(shorten, &sh)

	s.store.StorageWrite(short, oririn, uuid)

	s.logger.Info("try to login: ", map[string]interface{}{

		"oririn": oririn,
		"short":  short,
		"uuid":   uuid,
	})

	return short, nil
}

func (s *Service) Origin(short string) (string, error) {

	for _, sh := range shorten {
		if sh.ShortURL == short {

			if sh.DeletedFlag {
				str := "шорт " + short + " удален"
				return sh.OriginalURL, &ErrDeleted{str}
			}

			return sh.OriginalURL, nil
		}
	}

	return "", errors.New("отстуствует")

}

func (s *Service) GetAll() *[]*storager.Shorten {

	return &shorten
}
