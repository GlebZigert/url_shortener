package container

import (
	"errors"
	"fmt"

	"github.com/GlebZigert/url_shortener.git/internal/storager"
)

// Контейнер на основе мапы
type Map_container struct {
	mapa map[string]string
	storager.IStorager
}

func New_map_container(store storager.IStorager) *Map_container {
	fmt.Println("обьявляю мапу")
	ctn := Map_container{}
	ctn.mapa = make(map[string]string)

	if store != nil {
		ctn.IStorager = store
		ctn.Load(ctn.Set_short_without_db)

	}
	return &ctn
}

func (one *Map_container) Get_short(origin string) (v string, ok bool) {

	v, ok = one.mapa[origin]
	return

}
func (one *Map_container) Set_short_without_db(short, origin string) {
	fmt.Println("Set_short_without_db ", short, " ", origin)
	one.mapa[origin] = short

}

func (one *Map_container) Set_short(short, origin string) {
	fmt.Println("Set_short ", short, " ", origin)
	one.Set_short_without_db(short, origin)
	err := one.StorageWrite(short, origin)
	if err != nil {
		fmt.Println("запись не прошла: ", err.Error())
	} else {
		fmt.Println("запись должна была пройти успешно")
	}

}

func (one *Map_container) Get_origin(short string) (string, error) {

	for k, v := range one.mapa {
		if v == short {
			fmt.Println("Для шорта", short, " найден url: ", k)
			return k, nil
		}
	}
	fmt.Println("Нет такого шорта как", short)
	return "", errors.New("отстуствует")

}
