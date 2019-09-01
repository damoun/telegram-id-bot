package main

import (
	"os"
	"strconv"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// GetReply Create the string to reply.
func GetReply(update tgbotapi.Update) string {
	if update.Message.Text == "/getid" {
		return strconv.Itoa(update.Message.From.ID)
	} else if update.Message.Text == "/getgroupid" {
		if update.Message.Chat.Type != "private" {
			return strconv.FormatInt(update.Message.Chat.ID, 10)
		}
		return "This is not a group! Stop fooling me!"
	}
	return `This bot can supply reply with your Telegram user ID.
Use /getid command to get your Telegram ID.`
}

//WebHookHandler Handler for Telegram Webhook.
func WebHookHandler(update tgbotapi.Update) error {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err == nil {
		reply := GetReply(update)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		msg.ReplyToMessageID = update.Message.MessageID
		_, err = bot.Send(msg)
	}
	return err
}

func main() {
    lambda.Start(WebHookHandler)
}