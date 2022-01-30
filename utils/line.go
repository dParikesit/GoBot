package utils

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"os"
)

var Bot *linebot.Client
var err error

func ConnectLine() {
	Bot, err = linebot.New(os.Getenv("CHANNELSECRET"), os.Getenv("CHANNELTOKEN"))
	if err != nil {
		log.Fatalln("Line bot connection error")
	}
}
