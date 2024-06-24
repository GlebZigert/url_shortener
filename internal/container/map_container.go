package container

import (
	"errors"
	"fmt"

	"github.com/GlebZigert/url_shortener.git/internal/storager"
)

// Контейнер на основе мапы
type MapContainer struct {
	mapa map[string]string
	storager.IStorager
}

func NewMapContainer(store storager.IStorager) *MapContainer {
	fmt.Println("обьявляю мапу")
	ctn := MapContainer{}
	ctn.mapa = make(map[string]string)

	if store != nil {
		ctn.IStorager = store
		ctn.Load(ctn.SetShortWithoutDB)

	}
	return &ctn
}

func (one *MapContainer) GetShort(origin string) (v string, ok bool) {

	v, ok = one.mapa[origin]
	return

}
func (one *MapContainer) SetShortWithoutDB(short, origin string) {
	fmt.Println("SetShortWithoutDB ", short, " ", origin)
	one.mapa[origin] = short

}

func (one *MapContainer) SetShort(short, origin string) {
	fmt.Println("SetShort ", short, " ", origin)
	one.SetShortWithoutDB(short, origin)
	err := one.StorageWrite(short, origin)
	if err != nil {
		fmt.Println("запись не прошла: ", err.Error())
	} else {
		fmt.Println("запись должна была пройти успешно")
	}

}

func (one *MapContainer) GetOrigin(short string) (string, error) {

	for k, v := range one.mapa {
		if v == short {
			fmt.Println("Для шорта", short, " найден url: ", k)
			return k, nil
		}
	}
	fmt.Println("Нет такого шорта как", short)
	return "", errors.New("отстуствует")

}
