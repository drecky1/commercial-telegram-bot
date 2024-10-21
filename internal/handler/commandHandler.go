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
	var ok = false
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
		break
	case "/open":
		ok = true
		if msg.Chat.ID == service.Settings.Admin {
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
		ok = true
		if msg.Chat.ID == service.Settings.Admin {
			service.Settings.Registration = false
			commands.ShowParticipants(msg, service.Cache)
			go func() {
				if msg.Text != "" {
					err := service.SendMessage(service.Settings.Admin, msg.Text)
					if err != nil {
						return
					}
				}
			}()
			msg.Text = utils.CloseRegistrationMessage
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	case "/show":
		ok = true
		if msg.Chat.ID == service.Settings.Admin {
			commands.ShowParticipants(msg, service.Cache)
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	case "/rand":
		ok = true
		if msg.Chat.ID == service.Settings.Admin {
			commands.PickWinner(msg, service.Cache)
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	case "/delete":
		ok = true
		if msg.Chat.ID == service.Settings.Admin {
			commands.DeleteParticipants(msg, service.Settings, service.Cache)
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	case "/settings":
		ok = true
		if msg.Chat.ID == service.Settings.Admin {
			commands.CheckSettings(msg, service.Settings, service.Cache)
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	case "/promocode":
		if msg.Chat.ID == service.Settings.Admin {
			ok = true
			service.Settings.GiftCode = strings.ToUpper(strings.TrimSpace(strings.Replace(msg.Text, "/promocode", "", -1)))
			if service.Settings.GiftCode == "" {
				msg.Text = "введен пустой промокод"
				return service.SendMessage(msg.Chat.ID, msg.Text)
			}
			msg.Text = utils.PromocodeApplied
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	case "/prize":
		if msg.Chat.ID == service.Settings.Admin {
			ok = true
			service.Settings.MainPrize = strings.TrimSpace(strings.Replace(msg.Text, "/prize", "", -1))
			participateMessage := fmt.Sprintf(utils.ParticipateMessage, service.Settings.MainPrize)
			msg.Text = fmt.Sprintf(utils.MainPrizeApllied, participateMessage)
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	case "/change":
		ok = true
		if msg.Chat.ID == service.Settings.Admin {
			commands.ChangeParticipants(msg, service.Settings)
		} else {
			msg.Text = utils.YouAreNotAdministrator
		}
	}

	if !ok {
		return nil
	}
	if msg.Text == "" {
		return nil
	}
	return service.SendMessage(msg.Chat.ID, msg.Text)
}
