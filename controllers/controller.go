package controllers

import (
	"fmt"
	"github.com/dParikesit/dimsBot/utils"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"net/http"
	"strings"
)

func Line(w http.ResponseWriter, r *http.Request) {
	events, err := utils.Bot.ParseRequest(r)
	if err != nil {
		fmt.Fprint(w, "Parse failed")
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				kalimat := strings.SplitN(message.Text, " ", 3)
				log.Println(kalimat[1])
				if kalimat[0] == "!remindme" {
					if _, err = utils.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(" OK! We will remind you")).Do(); err != nil {
						log.Print(err)
					}
				}

			}
		}
	}

	fmt.Fprint(w, "Hello World!")
}
