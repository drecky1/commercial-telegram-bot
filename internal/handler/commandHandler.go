package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"tg_contour_bot/internal/commands"
	"tg_contour_bot/internal/telegram"
	"tg_contour_bot/utils"
)

func handleCommand(msg *tgbotapi.Message, service *telegram.Service) error {
	var isAdmin bool
	switch msg.Chat.ID {
	case service.Settings.Admin:
		isAdmin = true
	case service.Settings.Moderator:
		isAdmin = true
	default:
		isAdmin = false
	}
	parts := strings.Split(msg.Text, " ")
	switch parts[0] {
	case "/get_id":
		_ = service.SendMessage(msg.Chat.ID, fmt.Sprintf("%d", msg.Chat.ID))
		break
	case "/start":
		err := service.SendMenu(msg.Chat.ID, startMenu, startMenuMarkup)
		if err != nil {
			return err
		}
		msg.Text = ""
	case "/open":
		if isAdmin {
			if service.Settings.GiftCode == "" {
				msg.Text = utils.WrongOpenRegistrationCheckPromocode
				return service.SendMessage(msg.Chat.ID, msg.Text)
			}
			if service.Settings.MainPrize == "" {
				msg.Text = utils.WrongOpenRegistrationMainPrize
				return service.SendMessage(msg.Chat.ID, msg.Text)
			}
			service.Settings.Registration = true
			msg.Text = utils.OpenRegistrationMessage
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	case "/close":
		if isAdmin {
			commands.ShowParticipants(msg.Chat.ID, service)
			service.Settings.Registration = false
			msg.Text = utils.CloseRegistrationMessage
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	case "/show":
		if isAdmin {
			commands.ShowParticipants(msg.Chat.ID, service)
			msg.Text = ""
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	case "/rand":
		if isAdmin {
			commands.PickWinner(msg, service.Cache)
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	case "/delete":
		if isAdmin {
			commands.DeleteParticipants(msg, service.Settings, service.Cache)
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	case "/settings":
		if isAdmin {
			commands.CheckSettings(msg, service.Settings, service.Cache)
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	case "/promocode":
		if isAdmin {
			service.Settings.GiftCode = strings.ToUpper(strings.TrimSpace(strings.Replace(msg.Text, "/promocode", "", -1)))
			if service.Settings.GiftCode == "" {
				msg.Text = "введен пустой промокод"
				return service.SendMessage(msg.Chat.ID, msg.Text)
			}
			msg.Text = utils.PromocodeApplied
		} else {
			msg.Text = utils.YouAreNotAdministrator
			return service.SendMessage(msg.Chat.ID, msg.Text)
		}
	case "/prize":
		if isAdmin {
			service.Settings.MainPrize = strings.TrimSpace(strings.Replace(msg.Text, "/prize", "", -1))
			participateMessage := fmt.Sprintf(utils.ParticipateMessage, service.Settings.MainPrize)
			msg.Text = fmt.Sprintf(utils.MainPrizeApllied, participateMessage)
		} else {
			msg.Text = utils.YouAreNotAdministrator
			return service.SendMessage(msg.Chat.ID, msg.Text)
		}
	case "/change":
		if isAdmin {
			commands.ChangeParticipants(msg, service.Settings)
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	default:
		break
	}
	if msg.Text == "" {
		return nil
	}
	return service.SendMessage(msg.Chat.ID, msg.Text)
}
