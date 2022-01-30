package controllers

import (
	"fmt"
	"github.com/dParikesit/dimsBot/utils"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"net/http"
)

func Line(w http.ResponseWriter, r *http.Request) {
	events, err := utils.Bot.ParseRequest(r)
	if err != nil {
		fmt.Fprint(w, "Parse failed")
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			log.Println(event.Message)
			log.Println(event.Message.Type())
			log.Println(event.Message.Message)
		}
	}

	fmt.Fprint(w, "Hello World!")
}
