package utils

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"os"
)

var Bot *linebot.Client
var err error

type WebhookEvent struct {
	Destination string `json:"destination"`
	Events      Event  `json:"events"`
}

type Event struct {
	Type       string       `json:"type"`
	Timestamp  int          `json:"timestamp"`
	Source     EventSource  `json:"source"`
	ReplyToken string       `json:"reply_token,omitempty"`
	Mode       string       `json:"mode,omitempty"`
	Message    EventMessage `json:"message,omitempty"`
}

type EventSource struct {
	Type   string `json:"type"`
	UserId string `json:"user_id"`
}

type EventMessage struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Text string `json:"text"`
}

func ConnectLine() {
	Bot, err = linebot.New(os.Getenv("CHANNELSECRET"), os.Getenv("CHANNELTOKEN"))
	if err != nil {
		log.Fatalln("Line bot connection error")
	}
}
