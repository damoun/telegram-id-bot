package main

import (
	"os"
	"strconv"
	"context"
	"net/http"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	telegramBotApi *tgbotapi.BotAPI
	telegramApiToken = os.Getenv("TELEGRAM_API_TOKEN")
)

// Initialize Telegram API client
func init() {
	var err error
	telegramBotApi, err = tgbotapi.NewBotAPI(telegramApiToken)
	if err != nil {
		panic(err)
	}
}

// GetReply Create the string to reply.
func GetReply(update tgbotapi.Update) string {
	if update.Message.IsCommand() {
		switch update.Message.Command() {
			case "getid":
				return strconv.FormatInt(update.Message.From.ID, 10)
			case "getgroupid":
				if update.Message.Chat.Type != "private" {
					return strconv.FormatInt(update.Message.Chat.ID, 10)
				}
				return "This is not a group! Stop fooling me!"
		}
	}
	return `This bot can supply reply with your Telegram user ID.
Use /getid command to get your Telegram ID.`
}

// WebHookHandler Handler for Telegram Webhook.
func WebHookHandler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var update tgbotapi.Update
	err := json.Unmarshal([]byte(req.Body), &update)
	if err == nil && update.Message != nil {
		reply := GetReply(update)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		msg.ReplyToMessageID = update.Message.MessageID
		_, err = telegramBotApi.Send(msg)
	}
	return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusOK}, err
}

func main() {
    lambda.Start(WebHookHandler)
}