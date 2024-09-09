package services

import (
	"strings"

	"github.com/GlebZigert/url_shortener.git/internal/db"
)

func Delete(shorts []string, uid int) error {

	// сигнальный канал для завершения горутин
	doneCh := make(chan struct{})
	// закрываем его при завершении программы
	defer close(doneCh)

	// кладу в генератор слайс с шортами
	inputCh := generator(doneCh, shorts)

	// получаем слайс каналов из нескольких рабочих deleteShort
	channels := fanOut(doneCh, inputCh, len(shorts), uid)

	// а теперь объединяем эти каналы в один
	addResultCh := fanIn(doneCh, channels...)

	//получаю слайс айди тех шортов, которые прошли проверку на удаление
	res := multiply(doneCh, addResultCh)

	_, err := db.Get().Query("UPDATE strazh SET deleted = true WHERE id = ($1)", strings.Join(res, ","))
	return err

}
