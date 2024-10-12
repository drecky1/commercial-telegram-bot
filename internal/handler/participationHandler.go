package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"tg_contour_bot/internal/cache"
	"tg_contour_bot/internal/telegram"
	"tg_contour_bot/internal/validate"
	"tg_contour_bot/utils"
)

func handleParticipation(message *tgbotapi.Message, service *telegram.Service) error {
	user := message.Chat.ID
	if !service.Settings.Registration {
		return service.SendMessage(user, utils.SorryRegisterClosed)
	}
	text := message.Text
	state, ok := service.Cache.ParticipateStates[user]
	if !ok {
		return service.SendMessage(user, utils.DontKnowThatCommand)
	}
	if state == cache.WaitingForTitle {
		// need to validate name
		var u cache.UserInfo
		u.Title = text
		service.Cache.Participants[user] = u
		service.Cache.ParticipateStates[user] = cache.WaitingForPhone
		text = fmt.Sprintf(utils.CongratsNamePleaseEnterPhone, text)
		return service.SendMessage(user, text)
	}
	if state == cache.WaitingForPhone {
		u := service.Cache.Participants[user]
		// need to validate name
		text = strings.Replace(text, "+", "", 1)
		if !validate.IsValidRussianPhoneNumber(text) {
			return service.SendMessage(user, utils.NeedToSendValidPhone)
		}
		var uu cache.UserInfo
		var uname string
		if message.Chat.UserName != "" {
			uname = fmt.Sprintf("@%s", message.Chat.UserName)
		} else {
			uname = message.Chat.UserName
		}
		uu.TelegramUsername = uname
		uu.Title = u.Title
		if strings.HasPrefix(text, "7") {
			text = "+" + text
		} else if strings.HasPrefix(text, "8") {
			text = strings.Replace(text, "8", "+7", 1)
		}
		uu.Phone = text
		service.Mutex.Lock()
		delete(service.Cache.ParticipateStates, user)
		service.Cache.ParticipantsIDs = append(service.Cache.ParticipantsIDs, user)
		uu.ParticipantNumber = len(service.Cache.ParticipantsIDs)
		service.Cache.Participants[user] = uu
		service.Mutex.Unlock()
		text = fmt.Sprintf(utils.CongratsYouRegistered, uu.ParticipantNumber, uu.Title, uu.Phone)
		return service.SendMenu(user, text, updateParticipationDataMarkup)
	}
	if state == cache.UpdatingTitle {
		u := service.Cache.Participants[user]
		// need to validate name
		var uu cache.UserInfo
		uu.Title = text
		uu.Phone = u.Phone
		uu.ParticipantNumber = u.ParticipantNumber
		uu.TelegramUsername = u.TelegramUsername
		service.Mutex.Lock()
		service.Cache.Participants[user] = uu
		delete(service.Cache.ParticipateStates, user)
		service.Mutex.Unlock()
		text = fmt.Sprintf(utils.NameHasBeenUpdated, uu.Title)
		return service.SendMessage(user, text)
	}
	if state == cache.UpdatingPhone {
		u := service.Cache.Participants[user]
		// try to validate phone number
		phone := strings.Replace(text, "+", "", 1)
		if !validate.IsValidRussianPhoneNumber(phone) {
			return service.SendMessage(user, utils.NeedToSendValidPhone)
		}
		var uu cache.UserInfo
		uu.Title = u.Title
		uu.ParticipantNumber = u.ParticipantNumber
		uu.TelegramUsername = u.TelegramUsername
		if strings.HasPrefix(text, "7") {
			text = "+" + text
		} else if strings.HasPrefix(text, "8") {
			text = strings.Replace(text, "8", "+7", 1)
		}
		uu.Phone = text
		service.Mutex.Lock()
		service.Cache.Participants[user] = uu
		delete(service.Cache.ParticipateStates, user)
		service.Mutex.Unlock()
		text = fmt.Sprintf(utils.PhoneHasBeenUpdated, uu.Title, uu.Phone)
		return service.SendMessage(user, text)
	}
	return nil
}
