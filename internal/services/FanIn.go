package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"go.uber.org/zap"
)

// generator возвращает канал с данными
func generator(doneCh chan struct{}, input []string) chan string {
	// канал, в который будем отправлять данные из слайса
	inputCh := make(chan string)

	// горутина, в которой отправляем в канал  inputCh данные
	go func() {
		// как отправители закрываем канал, когда всё отправим
		defer close(inputCh)

		// перебираем все данные в слайсе
		for _, data := range input {
			select {
			// если doneCh закрыт, сразу выходим из горутины
			case <-doneCh:
				return
			// если doneCh не закрыт, кидаем в канал inputCh данные data
			case inputCh <- data:
			}
		}
	}()

	// возвращаем канал для данных
	return inputCh
}

// fanOut принимает канал данных, порождает несколько горутин
func fanOut(doneCh chan struct{}, inputCh chan string, numWorkers int, uid int) []chan int {
	// количество горутин add

	// каналы, в которые отправляются результаты
	channels := make([]chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		// получаем канал из горутины deleteShort
		addResultCh := deleteShortWorker(doneCh, inputCh, uid)
		// отправляем его в слайс каналов
		channels[i] = addResultCh
	}

	// возвращаем слайс каналов
	return channels
}

// fanIn объединяет несколько каналов resultChs в один.
func fanIn(doneCh chan struct{}, resultChs ...chan int) chan int {
	// конечный выходной канал в который отправляем данные из всех каналов из слайса, назовём его результирующим
	finalCh := make(chan int)

	// понадобится для ожидания всех горутин
	var wg sync.WaitGroup

	// перебираем все входящие каналы
	for _, ch := range resultChs {
		// в горутину передавать переменную цикла нельзя, поэтому делаем так
		chClosure := ch

		// инкрементируем счётчик горутин, которые нужно подождать
		wg.Add(1)

		go func() {
			// откладываем сообщение о том, что горутина завершилась
			defer wg.Done()

			// получаем данные из канала
			for data := range chClosure {
				select {
				// выходим из горутины, если канал закрылся
				case <-doneCh:
					return
				// если не закрылся, отправляем данные в конечный выходной канал
				case finalCh <- data:
				}
			}
		}()
	}

	go func() {
		// ждём завершения всех горутин
		wg.Wait()
		// когда все горутины завершились, закрываем результирующий канал
		close(finalCh)
	}()

	// возвращаем результирующий канал
	return finalCh
}

// multiply принимает на вход сигнальный канал для прекращения работы и канал с входными данными для работы

func multiply(doneCh chan struct{}, inputCh chan int) chan int {
	// канал с результатом

	multiplyRes := make(chan int)
	defer close(multiplyRes)
	go func() {
		// берем из канала inputCh значения, которые надо сложить в слайс
		for data := range inputCh {
			// изменяем данные

			select {
			// если канал doneCh закрылся, выходим из горутины
			case <-doneCh:
				return

			case multiplyRes <- data:

			}
		}
	}()

	// возвращаем канал для результатов вычислений
	return multiplyRes
}

func deleteShortWorker(doneCh chan struct{}, inputCh chan string, uid int) chan int {
	addRes := make(chan int)

	go func() {
		defer close(addRes)

		for data := range inputCh {
			id, err := deleteShort(data, uid)

			if err == nil {
				addRes <- id
			}

			select {
			case <-doneCh:
				return
			}

		}
	}()
	return addRes
}

// выполняет проверку шорта по пользователю
func deleteShort(short string, uid int) (id int, err error) {
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
					id = one.ID
					one.DeletedFlag = true
					logger.Log.Info("удален",
						zap.String("short:", short),
						zap.Int("uid:", uid))
					return id, err
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
	return

}

func final(Ch chan int) error {

	var res []string
	for data := range Ch {
		res = append(res, strconv.Itoa(data))

	}

	_, err := db.Get().Query("UPDATE strazh SET deleted = true WHERE id = ($1)", strings.Join(res, ","))
	return err

}
