package handler

import (
	"log"

	"github.com/Narushio/scarlet-bot/thirdparty/botgo/event"
)

func ErrorNotify() event.ErrorNotifyHandler {
	return func(err error) {
		log.Println("error notify receive: ", err)
	}
}
