package handlers

import (
	"fmt"
	"github.com/dParikesit/dimsBot/controllers"
	"github.com/dParikesit/dimsBot/models"
	"github.com/dParikesit/dimsBot/utils"
	"github.com/gofiber/fiber/v2"
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

func GetAll(c *fiber.Ctx) error {
	var reminders []models.Reminder

	if c.Query("done") == "false" {
		reminders = controllers.GetBool(false)
	} else if c.Query("done") == "true" {
		reminders = controllers.GetBool(true)
	} else {
		reminders = controllers.GetAll()
	}

	return c.JSON(reminders)
}

func AddData(c *fiber.Ctx) error {
	dataTemp := new(models.ReminderTemp)
	if err := c.BodyParser(dataTemp); err != nil {
		return c.Status(500).JSON(err)
	}
	data, err := dataTemp.ConvTime()
	if err != nil {
		return c.Status(500).JSON(err)
	}

	log.Println(data)
	return c.JSON(data)
}
