package services

import (
	"log"
	"net/http"
	"strings"
	feature "tiamat/m/v0/services/features"
	"tiamat/m/v0/services/magics"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

const linebotSecret = "8273b1607e7aad58ea027dda7dbcc57c"
const linebotToken = "uzNr4czd17pPxHKZ2aJ3erVkBU0XK7NcAYItFwqSEMMgR85BawivpWKMo4cFkSJTSjRERQv5XEMTXfzeKB5T1GTEGJwju80ZDQqJQzCBQTUTIr4860hOAyCeJFb1597sRb58kxD6HbcS+Vw1Y39AhwdB04t89/1O/w1cDnyilFU="

var (
	bot *linebot.Client
)

func Init() {
	var err error
	bot, err = linebot.New(linebotSecret, linebotToken)
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
			replyMsg = magicWord(text)
		}
		tokens := strings.Split(text, " ")
		switch tokens[0] {
		case feature.RollMagicWord:
			replyMsg = feature.Roll(text)
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

func magicWord(msg string) string {
	tokens := strings.Split(msg, " ")

	switch tokens[0] {
	case magics.MagicWordRoll:
		return feature.Roll(msg)
	default:
		return ""
	}
}
