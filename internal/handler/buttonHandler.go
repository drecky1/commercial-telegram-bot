package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg_contour_bot/internal/cache"
	"tg_contour_bot/internal/telegram"
	"tg_contour_bot/utils"
)

func handleButton(query *tgbotapi.CallbackQuery, service *telegram.Service) error {
	var text string

	markup := tgbotapi.NewInlineKeyboardMarkup()
	message := query.Message

	switch query.Data {
	case nextButton:
		text = fmt.Sprintf(participateMenu, service.Settings.MainPrize)
		go service.SendMessage(message.Chat.ID, "ТУТ БУДЕТ СООБЩЕНИЕ ПО ПОВОДУ ПОЛЬЗОВАТЕЛЬСКОГО СОГЛАШЕНИЯ.")
		markup = participateMenuMarkup
	case backButton:
		text = startMenu
		markup = startMenuMarkup
	case participateButton:
		if int64(len(service.Cache.ParticipantsIDs)-1) >= service.Settings.MaxParticipants {
			return service.SendMessage(message.Chat.ID, "Простите, но уже слишком много участников.")
		}
		u, ok := service.Cache.Get(message.Chat.ID)
		if ok {
			text = fmt.Sprintf(utils.AlreadyRegistered, u.Title, u.Phone)
			return service.SendMenu(message.Chat.ID, text, updateParticipationDataMarkup)
		}
		if service.Settings.Registration {
			service.Cache.ParticipateStates[query.Message.Chat.ID] = cache.WaitingForTitle
			text = utils.PleaseIntroduceYourself
			markup = startMenuMarkup
		} else {
			text = utils.SorryRegisterClosed
			markup = startMenuMarkup
		}
		return service.SendMessage(message.Chat.ID, text)
	case changeTitle:
		if service.Settings.Registration {
			service.Cache.ParticipateStates[query.Message.Chat.ID] = cache.UpdatingTitle
			text = utils.PleaseIntroduceYourself
			markup = startMenuMarkup
		} else {
			text = utils.SorryRegisterClosed
			markup = startMenuMarkup
		}
		return service.SendMessage(message.Chat.ID, text)
	case changePhone:
		if service.Settings.Registration {
			service.Cache.ParticipateStates[query.Message.Chat.ID] = cache.UpdatingPhone
			text = utils.PleaseEnterYpuPhoneNumber
			markup = updateParticipationDataMarkup
		} else {
			text = utils.SorryRegisterClosed
			markup = updateParticipationDataMarkup
		}
		return service.SendMessage(message.Chat.ID, text)
	case giftButton:
		text = fmt.Sprintf(utils.PROMOCODE, service.Settings.Url, service.Settings.GiftCode)
		return service.SendMenu(message.Chat.ID, text, webUrlDataMarkup)
	}

	callbackCfg := tgbotapi.NewCallback(query.ID, "42")
	service.Bot.Send(callbackCfg)

	// Replace menu text and keyboard
	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, markup)
	msg.ParseMode = tgbotapi.ModeHTML
	service.Bot.Send(msg)
	return nil
}
