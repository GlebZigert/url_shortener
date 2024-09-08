package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"go.uber.org/zap"
)

/*
Функция Delete принимает на входе массив шортов для удаления short
и айдишник пользователя который хочет удалить все эти шорты uid

Для каждого шорта из массива shorts я должен выполнить следующие действия:

найти этот шорт в локальной хранилке - это массив shorten
если шорт найден - посмотреть кто его создал
если его создал тот же пользователь который сейчас хочет его удалить


послылаю в бд команду выставить для этого шорта флаг deleted в true
если эта команда записи в бд выполняется успешно

повторяю действие по выставлению для этого шорта флага deleted в true уже в локальной хранилке
таким образом исключаю возможность расхождения информаии в бд и локальгой структуре-хранилке

Эту процедуру для нескольких шортов можно  выполнить параллельно

*/

func deleteShort(short string, uid int) {
	logger.Log.Info("deleteShort",
		zap.String("short:", short),
		zap.Int("uid:", uid))
	//ищу шорт в локальной хранилке
	for _, one := range shorten {
		var err error
		//если нащел этот шорт
		//проверяю кто его создал
		//если это тот же юзер который сейчас его удаляет - тогда надо удалять этотт шорт
		if one.ShortURL == short {
			if one.UUID == uid {

				//шорт считается удаленным если его akfu deleted выставлен в true

				//шорты хранятся в локальной хранилке и в бд
				//флаг надо выставить и там и там

				//сначала выставляю флаг в бд
				_, err = db.Get().Exec("UPDATE strazh SET deleted = true WHERE short = $1", short)

				//если запрос в бд был выполнен успешно
				//выставляю флаг и в хранилке
				if err == nil {
					one.DeletedFlag = true
					logger.Log.Info("удален",
						zap.String("short:", short),
						zap.Int("uid:", uid))
				}

				//

			} else {
				err = errors.New(fmt.Sprintln("шорт другого пользователя", one.UUID, uid))
			}
		}
		//если есть ошибки - их надо поднять вверх до обработчика в мидле errHandler
		if err != nil {

			logger.Log.Error("err",
				zap.String("short:", short),
				zap.String("err:", err.Error()))
		}
	}

}

func Delete(shorts []string, uid int) {
	logger.Log.Info("service.Delete ",
		zap.String("shorts:", strings.Join(shorts, "")),
		zap.Int("uid:", uid))

	//совершаю обход массива с шортами которые надо удалить
	for _, short := range shorts {
		//для каждого шорта в отдельной гоуртине запускаю функцию

		go deleteShort(short, uid)

	}

}
