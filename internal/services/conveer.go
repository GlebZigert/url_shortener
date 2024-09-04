package services

import (
	"fmt"
	"log"
	"strings"

	"github.com/GlebZigert/url_shortener.git/internal/db"
)

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
					fmt.Println("удаляю ", short, " ", id)
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
	var tags []string

	_, err := db.Get().Exec("UPDATE strazh SET deleted = ? WHERE tags = ?", true,
		strings.Join(tags, "|"), listID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(listID)
}
