package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"tg_contour_bot/internal/telegram"
)

func handleMessage(message *tgbotapi.Message, service *telegram.Service) {
	user := message.From
	text := message.Text
	fmt.Println("Handled message", text, "from", user.UserName)

	if user == nil {
		return
	}

	if strings.HasPrefix(text, "/") {
		err := handleCommand(message, service)
		if err != nil {
			log.Printf("An error occured: %s", err.Error())
		}
	} else {
		err := handleParticipation(message, service)
		if err != nil {
			log.Printf("An error occured: %s", err.Error())
		}
	}
}
