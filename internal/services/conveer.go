package services

import "time"

//
func Delete(shorts []string, uid int) {

	// сигнальный канал для завершения горутин
	doneCh := make(chan struct{})
	inputCh := generator(shorts)
	channels := fanOut(doneCh, inputCh, len(shorts))
}

// generator — генератор, который создает канал и сразу возвращает его
func generator(input []string) chan string {
	inputCh := make(chan string)

	// через отдельную горутину генератор отправляет данные в канал
	go func() {
		// закрываем канал по завершению горутины — это отправитель
		defer close(inputCh)

		// перебираем данные в слайсе
		for _, data := range input {
			// отправляем данные в канал inputCh
			inputCh <- data
		}
	}()

	// возвращаем канал inputCh
	return inputCh
}

// fanOut принимает канал данных, порождает 10 горутин
func fanOut(doneCh chan struct{}, inputCh chan string, numWorkers int) []chan string {
	// количество горутин add

	// каналы, в которые отправляются результаты
	channels := make([]chan string, numWorkers)

	for i := 0; i < numWorkers; i++ {
		// получаем канал из горутины add
		addResultCh := add(doneCh, inputCh)
		// отправляем его в слайс каналов
		channels[i] = addResultCh
	}

	// возвращаем слайс каналов
	return channels
}

// ищет элемент структуры с short
//если uid совпадает
//выставляет флаг deleted
//возвращает id
//если нет
//возвращает -1
func DeleteOne(doneCh chan struct{}, inputCh chan string) chan int {
	addRes := make(chan int)

	go func() {
		defer close(addRes)

		for data := range inputCh {
			// замедлим вычисление, как будто функция add требует больше вычислительных ресурсов
			time.Sleep(time.Second)

			result := data + 1

			select {
			case <-doneCh:
				return
			case addRes <- result:
			}
		}
	}()
	return addRes
}
