package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"go.uber.org/zap"
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

					logger.Log.Info("удаляю: ", zap.String("short", short))
					break
				}
			}

			ch <- id

		}(short, uid, ch)
		channels[i] = ch

	}
	listID := make([]string, len(shorts))
	for i, ch := range channels {
		listID[i] = strconv.Itoa(<-ch)
	}

	fmt.Println("arg: ", strings.Join(listID, ","))
	_, err := db.Get().Exec("UPDATE strazh SET deleted = true WHERE id IN (1,2)")
	//	_, err := db.Get().Exec("UPDATE strazh SET deleted = ? WHERE id IN (?)", true,
	//		strings.Join(listID, ","))
	if err != nil {
		logger.Log.Error(err.Error())
	}

	fmt.Println(listID)
}
