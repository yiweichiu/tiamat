package services

import (
	"log"
	"net/http"
	"strings"
	"tiamat/m/v0/externals/line"
	"tiamat/m/v0/services/features"
	"tiamat/m/v0/services/magics"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var (
	bot *linebot.Client
)

func Init() {
	var err error
	bot, err = linebot.New(line.GetChannelSecret(), line.GetChannelAccessToken())
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx *gin.Context) {
	if bot == nil {
		Init()
	}

	events, err := bot.ParseRequest(ctx.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			ctx.JSON(http.StatusBadRequest, nil)
		} else {
			ctx.JSON(http.StatusInternalServerError, nil)
		}
		return
	}

	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			eventTypeMsg(event)
		default:
			log.Printf("event : %v", event)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": events,
	})
}

func eventTypeMsg(event *linebot.Event) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		text := message.Text
		replyMsg := ""

		if text[0] == '!' {
			replyMsg = magicWord(event, text)
		}

		if _, err := bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(replyMsg),
		).Do(); err != nil {
			log.Print(err)
		}
	default:
		log.Printf("message: %v", message)
	}
}

func magicWord(event *linebot.Event, msg string) string {
	tokens := strings.Split(msg, " ")

	switch tokens[0] {
	case magics.MagicRoll:
		return features.Roll(msg)
	default:
		return ""
	}
}
