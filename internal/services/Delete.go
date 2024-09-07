package services

import (
	"github.com/GlebZigert/url_shortener.git/internal/db"
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

func delete_short(short string, uid int) {

	//ищу шорт в локальной хранилке
	for _, one := range shorten {

		//если нащел этот шорт
		//проверяю кто его создал
		//если это тот же юзер который сейчас его удаляет - тогда надо удалять этотт шорт
		if one.ShortURL == short && one.UUID == uid {

			//шорт считается удаленным если его akfu deleted выставлен в true

			//шорты хранятся в локальной хранилке и в бд
			//флаг надо выставить и там и там

			//сначала выставляю флаг в бд
			_, err := db.Get().Exec("UPDATE strazh SET deleted = true WHERE short = ?", short)

			//если запрос в бд был выполнен успешно
			//выставляю флаг и в хранилке
			if err == nil {
				one.DeletedFlag = true
			}

			//

		}
	}

}

func Delete(shorts []string, uid int) {

	//совершаю обход массива с шортами которые надо удалить
	for _, short := range shorts {
		//для каждого шорта в отдельной гоуртине запускаю функцию

		go delete_short(short, uid)

	}

}