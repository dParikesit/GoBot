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
	"time"
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
				var uid string
				if event.Source.Type == "room" {
					uid = event.Source.RoomID
				} else if event.Source.Type == "group" {
					uid = event.Source.GroupID
				} else {
					uid = event.Source.UserID
				}

				if kalimat[0] == "!remindme" {
					schedule, _ := time.Parse("02/01/2006", kalimat[1])
					reminder := models.Reminder{
						UserId:      uid,
						Schedule:    schedule,
						Description: kalimat[2],
						Done:        false,
					}

					if err = controllers.InsertOne(&reminder); err != nil {
						log.Println(err)
					}

					if _, err = utils.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(" OK! We will remind you")).Do(); err != nil {
						log.Println(err)
					}
				} else if kalimat[0] == "!todo" {
					reminders := controllers.GetBool(uid, false)
					if len(reminders) != 0 {
						answer := ""
						for _, reminder := range reminders {
							answer = answer + reminder.Schedule.Format("02/01") + " " + reminder.Description + "\n"
						}

						if _, err = utils.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(answer)).Do(); err != nil {
							log.Println(err)
						}
					} else {
						if _, err = utils.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Tidak ada todo list! Hore!!!")).Do(); err != nil {
							log.Println(err)
						}
					}

				} else if kalimat[0] == "!tagall" {
					log.Println(message)
				}
			}
		}
	}

	fmt.Fprint(w, "Hello World!")
}

func GetAll(c *fiber.Ctx) error {
	reminders := controllers.GetAll()
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
