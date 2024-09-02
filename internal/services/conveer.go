package services

import "fmt"

func Delete(shorts []string, uid int) {
	channels := make([]chan int, len(shorts))
	for i, short := range shorts {

		ch := make(chan int)
		go func(short string, uid int, ch chan int) {
			id := -1
			for _, sh := range shorten {
				if sh.ShortURL == short && sh.UUID == uid {

					sh.DeletedFlag = true
					id = sh.ID
					break
				}
			}

			ch <- id

		}(short, uid, ch)
		channels[i] = ch

	}
	listID := make([]int, len(shorts))
	for i, ch := range channels {
		listID[i] = <-ch
	}
	fmt.Println(listID)
}
