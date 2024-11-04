package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tg_contour_bot/internal/telegram"
	"tg_contour_bot/utils"
)

func HandleUpdate(update tgbotapi.Update, service *telegram.Service) {
	switch {
	case update.Message != nil:
		if update.Message.Chat.Type != utils.PrivateChat {
			return
		}
		handleMessage(update.Message, service)
		return

	case update.CallbackQuery != nil:
		if update.CallbackQuery.Message.Chat.Type != utils.PrivateChat {
			return
		}
		err := handleButton(update.CallbackQuery, service)
		if err != nil {
			log.Println(err)
		}
		return
	}
}
